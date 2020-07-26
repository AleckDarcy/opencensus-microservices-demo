// Copyright 2018 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"bytes"
	"fmt"
	"github.com/AleckDarcy/reload/core/tracer"
	"html/template"
	"io"
	"math/rand"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"time"
	"unsafe"

	pb "github.com/AleckDarcy/opencensus-microservices-demo/src/frontend/genproto"
	"github.com/AleckDarcy/opencensus-microservices-demo/src/frontend/money"
	rHtml "github.com/AleckDarcy/reload/runtime/html"
	rTemplate "github.com/AleckDarcy/reload/runtime/html/template"

	//_ "git.apache.org/thrift.git/lib/go/thrift"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

var (
	templates = rTemplate.Must(
		template.New("").Funcs(template.FuncMap{
			"renderMoney":    renderMoney,
			"marshalTracing": rTemplate.MarshalTracing,
		}).ParseGlob("templates/*.html"))
)

type productView struct {
	Item  *pb.Product
	Price *pb.Money
}

func (fe *frontendServer) homeHandler(w http.ResponseWriter, r *http.Request) {
	// unavoidable modification since we need to assign threadID by calling context.WithValue
	// we need to assign the new Context to replace the current one
	r = rHtml.Init(r)

	log := r.Context().Value(ctxKeyLog{}).(logrus.FieldLogger)
	//log.WithField("currency", currentCurrency(r)).Info("home")
	//log.Info("home")

	useDefaultCurrency := false
	currencies, err := fe.getCurrencies(r.Context())
	if err != nil {
		log.Errorf("[RELOAD] homeHandler getCurrencies fail: %s", err)
		// [RELOAD: FAULT TOLERANT] starts
		currencies = []string{"USD"}
		useDefaultCurrency = true
		//renderHTTPError(log, r, w, errors.Wrap(err, "could not retrieve currencies"), http.StatusInternalServerError)
		//return
		// [RELOAD: FAULT TOLERANT] ends
	}
	products, err := fe.getProducts(r.Context())
	if err != nil {
		log.Errorf("[RELOAD] homeHandler getProducts fail: %s", err)
		// [RELOAD: FAULT TOLERANT] starts
		//renderHTTPError(log, r, w, errors.Wrap(err, "could not retrieve products"), http.StatusInternalServerError)
		//return
		// [RELOAD: FAULT TOLERANT] ends
	}
	cart, err := fe.getCart(r.Context(), sessionID(r))
	if err != nil {
		log.Errorf("[RELOAD] homeHandler getCart fail: %s", err)
		// [RELOAD: FAULT TOLERANT] starts
		//renderHTTPError(log, r, w, errors.Wrap(err, "could not retrieve cart"), http.StatusInternalServerError)
		//return
		// [RELOAD: FAULT TOLERANT] ends
	}

	ps := make([]productView, len(products))

	rate := 1.0
	currencyCode := currentCurrency(r)

	for i, p := range products {
		usd := p.GetPriceUsd()

		// [RELOAD]
		if i == 0 && !useDefaultCurrency {
			numUSD := float64(usd.Units) + float64(usd.Nanos)/1e9
			price, _ := fe.convertCurrency(r.Context(), usd, currencyCode)
			rate = (float64(price.Units) + float64(price.Nanos)/1e9) / numUSD

			ps[i] = productView{p, price}
		} else if rate == 1.0 {
			ps[i] = productView{p, p.GetPriceUsd()}
		} else {
			num := float64(usd.Units) + float64(usd.Nanos)/1e9*rate
			units := int64(num)
			nanos := int32((num - float64(units)) * 1e9)

			ps[i] = productView{p, &pb.Money{
				CurrencyCode: currencyCode,
				Units:        units,
				Nanos:        nanos,
			}}
		}

		//if useDefault {
		//	ps[i] = productView{p, p.GetPriceUsd()}
		//} else {
		//	price, _ := fe.convertCurrency(r.Context(), p.GetPriceUsd(), currentCurrency(r))
		//	ps[i] = productView{p, price}
		//}
	}

	//ExecuteHomeTemplate(w, map[string]interface{}{
	//	"session_id":    sessionID(r),
	//	"request_id":    r.Context().Value(ctxKeyRequestID{}),
	//	"user_currency": currencyCode,
	//	"currencies":    currencies,
	//	"products":      ps,
	//	"cart_size":     len(cart),
	//	"banner_color":  os.Getenv("BANNER_COLOR"), // illustrates canary deployments
	//	"ad":            fe.chooseAd(r, log),
	//	"render":        r.FormValue("render"),
	//})

	if err := templates.ExecuteTemplateReload(r.Context(), w, "home", map[string]interface{}{
		"session_id":    sessionID(r),
		"request_id":    r.Context().Value(ctxKeyRequestID{}),
		"user_currency": currencyCode,
		"currencies":    currencies,
		"products":      ps,
		"cart_size":     len(cart),
		"banner_color":  os.Getenv("BANNER_COLOR"), // illustrates canary deployments
		"ad":            fe.chooseAd(r, log),
		"render":        r.FormValue("render"),
	}); err != nil {
		log.Error(err)
	}
}

func ExecuteHomeTemplate(w io.Writer, data map[string]interface{}) {
	wr := bytes.NewBufferString("")

	HeaderTemplate(wr, data)

	HomeTemplate(wr, data)

	FooterTemplate(wr, data)

	w.Write(*(*[]byte)(unsafe.Pointer(&home)))
	runtime.KeepAlive(home)
}

var homeStr = `
	<main role="main">
        <section class="jumbotron text-center mb-0" style="background-color: %s;">
            <div class="container">
                <h1 class="jumbotron-heading">
                    One-stop for Hipster Fashion &amp; Style Online
                </h1>
                <p class="lead text-muted">
                    Tired of mainstream fashion ideas, popular trends and
                    societal norms? This line of lifestyle products will help
                    you catch up with the hipster trend and express your
                    personal style. Start shopping hip and vintage items now!
                </p>
            </div>
        </section>

        <div class="py-5 bg-light">
            <div class="container">
            <div class="row">
                %s
            </div>
            <div class="row">
                %s
            </div>

            %s
            </div>
        </div>
    </main>
`

func HomeTemplate(data map[string]interface{}) string {
	productsStr := ""
	if products, ok := data["products"].([]productView); ok {
		for _, product := range products {
			productsStr += fmt.Sprintf(`
                <div class="col-md-4">
                    <div class="card mb-4 box-shadow">
                        <a href="/product/{{.Item.Id}}">
                            <img class="card-img-top" alt =""
                                style="width: 100%%; height: auto;"
                                src="%s">
                        </a>
                        <div class="card-body">
                            <h5 class="card-title">
                                %s
                            </h5>
                            <div class="d-flex justify-content-between align-items-center">
                                <div class="btn-group">
                                    <a href="/product/%s">
                                        <button type="button" class="btn btn-sm btn-outline-secondary">Buy</button>
                                    </a>
                                </div>
                                <small class="text-muted">
                                    %s
                                </strong>
                                </small>
                            </div>
                        </div>
                    </div>
                </div>
`, product.Item.Picture, product.Item.Name, product.Item.Id, renderMoney(product.Price))
		}
	}

	adStr := AdTemplate(data)

	traceStr := ""

	if fi_trace, ok := data["fi_trace"]; ok {
		traceStr = fmt.Sprintf(`
                <div class="trace">
                    %s
                </div>
`, rTemplate.MarshalTracing(fi_trace.(*tracer.Trace)))
	}

	return fmt.Sprintf(homeStr, data["banner_color"], productsStr, adStr, traceStr)
}

var headerStr = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0, shrink-to-fit=no">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Hipster Shop</title>
    <link href="https://stackpath.bootstrapcdn.com/bootstrap/4.1.1/css/bootstrap.min.css" rel="stylesheet" integrity="sha384-WskhaSGFgHYWDcbwN70/dfYBj47jz9qbsMId/iRN3ewGhXQFZCSftd1LZCfmhktB" crossorigin="anonymous">
</head>
<body>

    <header>
        <div class="navbar navbar-dark bg-dark box-shadow">
            <div class="container d-flex justify-content-between">
                <a href="/" class="navbar-brand d-flex align-items-center">
                    Hipster Shop
                </a>
            </div>
        </div>
    </header>
`

var headerStr1 = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0, shrink-to-fit=no">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Hipster Shop</title>
    <link href="https://stackpath.bootstrapcdn.com/bootstrap/4.1.1/css/bootstrap.min.css" rel="stylesheet" integrity="sha384-WskhaSGFgHYWDcbwN70/dfYBj47jz9qbsMId/iRN3ewGhXQFZCSftd1LZCfmhktB" crossorigin="anonymous">
</head>
<body>

    <header>
        <div class="navbar navbar-dark bg-dark box-shadow">
            <div class="container d-flex justify-content-between">
                <a href="/" class="navbar-brand d-flex align-items-center">
                    Hipster Shop
                </a>
                <form class="form-inline ml-auto" method="POST" action="/setCurrency" id="currency_form">
				<select name="currency_code" class="form-control"
					onchange="document.getElementById('currency_form').submit();" style="width:auto;">
					%s
					</select>
					<a class="btn btn-primary btn-light ml-2" href="/cart" role="button">View Cart (%v)</a>
				</form>
            </div>
        </div>
    </header>
`

func HeaderTemplate(wr *bytes.Buffer, data map[string]interface{}) string {
	if currencies, ok := data["currencies"]; ok {
		user_currency := data["user_currency"].(string)
		currenciesStr := ""
		for _, currency := range currencies.([]string) {
			if currency == user_currency {
				currenciesStr += fmt.Sprintf(`
						<option value="%s" selected="selected"">%s</option>
`, currency, currency)
			} else {
				currenciesStr += fmt.Sprintf(`
						<option value="%s">%s</option>
`, currency, currency)
			}
		}

		return fmt.Sprintf(headerStr1, currenciesStr, data["cart_size"])
	}

	return headerStr
}

var footerStr = `
	<footer class="py-5 px-5">
        <div class="container">
            <p>
                &copy; 2018 Google Inc
                <span class="text-muted">
                    <a href="https://github.com/AleckDarcy/opencensus-microservices-demo/">(Source Code)</a>
                </span>
            </p>
            <p>
                <small class="text-muted">
                    This website is hosted for demo purposes only. It is not an
                    actual shop. This is not an official Google project.
                </small>
            </p>
            <small class="text-muted">
                session-id: %s</br>
                request-id: %s</br>
            </small>
        </div>
    </footer>
    <script src="https://stackpath.bootstrapcdn.com/bootstrap/4.1.1/js/bootstrap.min.js" integrity="sha384-smHYKdLADwkXOn1EmN1qk/HfnUcbVRZyYmZ4qpPea6sjB/pTJ0euyQp0Mk8ck+5T" crossorigin="anonymous"></script>
</body>
</html>
`

func AdTemplate(data map[string]interface{}) string {
	if ad := data["ad"].(*pb.Ad); ad != nil {
		return fmt.Sprintf(`
				<div class="container">
					<div class="alert alert-dark" role="alert">
						<strong>Advertisement:</strong>
						<a href="%s" rel="nofollow" target="_blank" class="alert-link">
							%s
						</a>
					</div>
				</div>
`, ad.RedirectUrl, ad.Text)
	}

	return ""
}

func FooterTemplate(data map[string]interface{}) string {
	return fmt.Sprintf(footerStr, data["session_id"], data["request_id"])
}

func (fe *frontendServer) productHandler(w http.ResponseWriter, r *http.Request) {
	// unavoidable modification since we need to assign threadID by calling context.WithValue
	// we need to assign the new Context to replace the current one
	r = rHtml.Init(r)

	log := r.Context().Value(ctxKeyLog{}).(logrus.FieldLogger)
	id := mux.Vars(r)["id"]
	if id == "" {
		renderHTTPError(log, r, w, errors.New("product id not specified"), http.StatusBadRequest)
		return
	}
	log.WithField("id", id).WithField("currency", currentCurrency(r)).
		Debug("serving product page")

	p, err := fe.getProduct(r.Context(), id)
	if err != nil {
		renderHTTPError(log, r, w, errors.Wrap(err, "could not retrieve product"), http.StatusInternalServerError)
		return
	}
	currencies, err := fe.getCurrencies(r.Context())
	if err != nil {
		renderHTTPError(log, r, w, errors.Wrap(err, "could not retrieve currencies"), http.StatusInternalServerError)
		return
	}

	cart, err := fe.getCart(r.Context(), sessionID(r))
	if err != nil {
		renderHTTPError(log, r, w, errors.Wrap(err, "could not retrieve cart"), http.StatusInternalServerError)
		return
	}

	price, err := fe.convertCurrency(r.Context(), p.GetPriceUsd(), currentCurrency(r))
	if err != nil {
		renderHTTPError(log, r, w, errors.Wrap(err, "failed to convert currency"), http.StatusInternalServerError)
		return
	}

	recommendations, err := fe.getRecommendations(r.Context(), sessionID(r), []string{id})
	if err != nil {
		renderHTTPError(log, r, w, errors.Wrap(err, "failed to get product recommendations"), http.StatusInternalServerError)
		return
	}

	product := struct {
		Item  *pb.Product
		Price *pb.Money
	}{p, price}

	// unavoidable modification since we cannot hack into runtime of Go-lang
	if err := templates.ExecuteTemplateReload(r.Context(), w, "product", map[string]interface{}{
		"session_id":      sessionID(r),
		"request_id":      r.Context().Value(ctxKeyRequestID{}),
		"ad":              fe.chooseAd(r, log),
		"user_currency":   currentCurrency(r),
		"currencies":      currencies,
		"product":         product,
		"recommendations": recommendations,
		"cart_size":       len(cart),
	}); err != nil {
		log.Println(err)
	}
}

func (fe *frontendServer) addToCartHandler(w http.ResponseWriter, r *http.Request) {
	// unavoidable modification since we need to assign threadID by calling context.WithValue
	// we need to assign the new Context to replace the current one
	r = rHtml.Init(r)

	log := r.Context().Value(ctxKeyLog{}).(logrus.FieldLogger)
	quantity, _ := strconv.ParseUint(r.FormValue("quantity"), 10, 32)
	productID := r.FormValue("product_id")
	if productID == "" || quantity == 0 {
		renderHTTPError(log, r, w, errors.New("invalid form input"), http.StatusBadRequest)
		return
	}
	log.WithField("product", productID).WithField("quantity", quantity).Debug("adding to cart")

	p, err := fe.getProduct(r.Context(), productID)
	if err != nil {
		renderHTTPError(log, r, w, errors.Wrap(err, "could not retrieve product"), http.StatusInternalServerError)
		return
	}

	if err := fe.insertCart(r.Context(), sessionID(r), p.GetId(), int32(quantity)); err != nil {
		renderHTTPError(log, r, w, errors.Wrap(err, "failed to add to cart"), http.StatusInternalServerError)
		return
	}
	w.Header().Set("location", "/cart")
	w.WriteHeader(http.StatusFound)
}

func (fe *frontendServer) emptyCartHandler(w http.ResponseWriter, r *http.Request) {
	// unavoidable modification since we need to assign threadID by calling context.WithValue
	// we need to assign the new Context to replace the current one
	r = rHtml.Init(r)

	log := r.Context().Value(ctxKeyLog{}).(logrus.FieldLogger)
	log.Debug("emptying cart")

	if err := fe.emptyCart(r.Context(), sessionID(r)); err != nil {
		renderHTTPError(log, r, w, errors.Wrap(err, "failed to empty cart"), http.StatusInternalServerError)
		return
	}
	w.Header().Set("location", "/")
	w.WriteHeader(http.StatusFound)
}

func (fe *frontendServer) viewCartHandler(w http.ResponseWriter, r *http.Request) {
	// unavoidable modification since we need to assign threadID by calling context.WithValue
	// we need to assign the new Context to replace the current one
	r = rHtml.Init(r)

	log := r.Context().Value(ctxKeyLog{}).(logrus.FieldLogger)
	log.Debug("view user cart")
	currencies, err := fe.getCurrencies(r.Context())
	if err != nil {
		renderHTTPError(log, r, w, errors.Wrap(err, "could not retrieve currencies"), http.StatusInternalServerError)
		return
	}
	cart, err := fe.getCart(r.Context(), sessionID(r))
	if err != nil {
		renderHTTPError(log, r, w, errors.Wrap(err, "could not retrieve cart"), http.StatusInternalServerError)
		return
	}

	recommendations, err := fe.getRecommendations(r.Context(), sessionID(r), cartIDs(cart))
	if err != nil {
		renderHTTPError(log, r, w, errors.Wrap(err, "failed to get product recommendations"), http.StatusInternalServerError)
		return
	}

	shippingCost, err := fe.getShippingQuote(r.Context(), cart, currentCurrency(r))
	if err != nil {
		renderHTTPError(log, r, w, errors.Wrap(err, "failed to get shipping quote"), http.StatusInternalServerError)
		return
	}

	type cartItemView struct {
		Item     *pb.Product
		Quantity int32
		Price    *pb.Money
	}
	items := make([]cartItemView, len(cart))
	totalPrice := pb.Money{CurrencyCode: currentCurrency(r)}
	for i, item := range cart {
		p, err := fe.getProduct(r.Context(), item.GetProductId())
		if err != nil {
			renderHTTPError(log, r, w, errors.Wrapf(err, "could not retrieve product #%s", item.GetProductId()), http.StatusInternalServerError)
			return
		}
		price, err := fe.convertCurrency(r.Context(), p.GetPriceUsd(), currentCurrency(r))
		if err != nil {
			renderHTTPError(log, r, w, errors.Wrapf(err, "could not convert currency for product #%s", item.GetProductId()), http.StatusInternalServerError)
			return
		}

		multPrice := money.MultiplySlow(*price, uint32(item.GetQuantity()))
		items[i] = cartItemView{
			Item:     p,
			Quantity: item.GetQuantity(),
			Price:    &multPrice}
		totalPrice = money.Must(money.Sum(totalPrice, multPrice))
	}
	totalPrice = money.Must(money.Sum(totalPrice, *shippingCost))

	year := time.Now().Year()
	if err := templates.ExecuteTemplateReload(r.Context(), w, "cart", map[string]interface{}{
		"session_id":       sessionID(r),
		"request_id":       r.Context().Value(ctxKeyRequestID{}),
		"user_currency":    currentCurrency(r),
		"currencies":       currencies,
		"recommendations":  recommendations,
		"cart_size":        len(cart),
		"shipping_cost":    shippingCost,
		"total_cost":       totalPrice,
		"items":            items,
		"expiration_years": []int{year, year + 1, year + 2, year + 3, year + 4},
	}); err != nil {
		log.Println(err)
	}
}

func (fe *frontendServer) placeOrderHandler(w http.ResponseWriter, r *http.Request) {
	// unavoidable modification since we need to assign threadID by calling context.WithValue
	// we need to assign the new Context to replace the current one
	r = rHtml.Init(r)

	log := r.Context().Value(ctxKeyLog{}).(logrus.FieldLogger)
	log.Debug("placing order")

	var (
		email         = r.FormValue("email")
		streetAddress = r.FormValue("street_address")
		zipCode, _    = strconv.ParseInt(r.FormValue("zip_code"), 10, 32)
		city          = r.FormValue("city")
		state         = r.FormValue("state")
		country       = r.FormValue("country")
		ccNumber      = r.FormValue("credit_card_number")
		ccMonth, _    = strconv.ParseInt(r.FormValue("credit_card_expiration_month"), 10, 32)
		ccYear, _     = strconv.ParseInt(r.FormValue("credit_card_expiration_year"), 10, 32)
		ccCVV, _      = strconv.ParseInt(r.FormValue("credit_card_cvv"), 10, 32)
	)

	order, err := pb.NewCheckoutServiceClient(fe.checkoutSvcConn).
		PlaceOrder(r.Context(), &pb.PlaceOrderRequest{
			Email: email,
			CreditCard: &pb.CreditCardInfo{
				CreditCardNumber:          ccNumber,
				CreditCardExpirationMonth: int32(ccMonth),
				CreditCardExpirationYear:  int32(ccYear),
				CreditCardCvv:             int32(ccCVV)},
			UserId:       sessionID(r),
			UserCurrency: currentCurrency(r),
			Address: &pb.Address{
				StreetAddress: streetAddress,
				City:          city,
				State:         state,
				ZipCode:       int32(zipCode),
				Country:       country},
		})
	if err != nil {
		renderHTTPError(log, r, w, errors.Wrap(err, "failed to complete the order"), http.StatusInternalServerError)
		return
	}
	log.WithField("order", order.GetOrder().GetOrderId()).Info("order placed")

	order.GetOrder().GetItems()
	recommendations, _ := fe.getRecommendations(r.Context(), sessionID(r), nil)

	totalPaid := *order.GetOrder().GetShippingCost()
	for _, v := range order.GetOrder().GetItems() {
		totalPaid = money.Must(money.Sum(totalPaid, *v.GetCost()))
	}

	if err := templates.ExecuteTemplateReload(r.Context(), w, "order", map[string]interface{}{
		"session_id":      sessionID(r),
		"request_id":      r.Context().Value(ctxKeyRequestID{}),
		"user_currency":   currentCurrency(r),
		"order":           order.GetOrder(),
		"total_paid":      &totalPaid,
		"recommendations": recommendations,
	}); err != nil {
		log.Println(err)
	}
}

func (fe *frontendServer) logoutHandler(w http.ResponseWriter, r *http.Request) {
	log := r.Context().Value(ctxKeyLog{}).(logrus.FieldLogger)
	log.Debug("logging out")
	for _, c := range r.Cookies() {
		c.Expires = time.Now().Add(-time.Hour * 24 * 365)
		c.MaxAge = -1
		http.SetCookie(w, c)
	}
	w.Header().Set("Location", "/")
	w.WriteHeader(http.StatusFound)
}

func (fe *frontendServer) setCurrencyHandler(w http.ResponseWriter, r *http.Request) {
	log := r.Context().Value(ctxKeyLog{}).(logrus.FieldLogger)
	cur := r.FormValue("currency_code")
	log.WithField("curr.new", cur).WithField("curr.old", currentCurrency(r)).
		Debug("setting currency")

	if cur != "" {
		http.SetCookie(w, &http.Cookie{
			Name:   cookieCurrency,
			Value:  cur,
			MaxAge: cookieMaxAge,
		})
	}
	referer := r.Header.Get("referer")
	if referer == "" {
		referer = "/"
	}
	w.Header().Set("Location", referer)
	w.WriteHeader(http.StatusFound)
}

// chooseAd queries for advertisements available and randomly chooses one, if
// available. It ignores the error retrieving the ad since it is not critical.
func (fe *frontendServer) chooseAd(r *http.Request, log logrus.FieldLogger) *pb.Ad {
	ctx := r.Context()

	ads, err := fe.getAd(ctx)
	if err != nil {
		log.Warnf("failed to retrieve ads: %s", err.Error())

		// [RELOAD]
		price := &pb.Money{
			CurrencyCode: "USD",
			Units:        1,
			Nanos:        0,
		}

		//if !useDefault {
		price, _ = fe.convertCurrency(ctx, price, currentCurrency(r))
		//}

		defaultAd := &pb.Ad{
			RedirectUrl: "https://www.google.com",
			Text:        fmt.Sprintf("default product, price: %d %s", price.Units, currentCurrency(r)),
		}

		return defaultAd
	}

	return ads[rand.Intn(len(ads))]
}

func renderHTTPError(log logrus.FieldLogger, r *http.Request, w http.ResponseWriter, err error, code int) {
	log.WithField("error", err).Error("request error")
	errMsg := fmt.Sprintf("%+v", err)

	w.WriteHeader(code)
	templates.ExecuteTemplateReload(r.Context(), w, "error", map[string]interface{}{
		"session_id":  sessionID(r),
		"request_id":  r.Context().Value(ctxKeyRequestID{}),
		"error":       errMsg,
		"status_code": code,
		"status":      http.StatusText(code)})
}

func currentCurrency(r *http.Request) string {
	c, _ := r.Cookie(cookieCurrency)
	if c != nil {
		return c.Value
	}
	return defaultCurrency
}

func sessionID(r *http.Request) string {
	v := r.Context().Value(ctxKeySessionID{})
	if v != nil {
		return v.(string)
	}
	return ""
}

func cartIDs(c []*pb.CartItem) []string {
	out := make([]string, len(c))
	for i, v := range c {
		out[i] = v.GetProductId()
	}
	return out
}

func renderMoney(money *pb.Money) string {
	return fmt.Sprintf("%s %d.%02d", money.CurrencyCode, money.Units, money.Nanos/10000000)
}
