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
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	pb "github.com/AleckDarcy/opencensus-microservices-demo/src/frontend/genproto"
	"github.com/AleckDarcy/opencensus-microservices-demo/src/frontend/money"
	"github.com/AleckDarcy/reload/core/tracer"
	rHtml "github.com/AleckDarcy/reload/runtime/html"
	rTemplate "github.com/AleckDarcy/reload/runtime/html/template"

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

	rate := 0.0
	currencyCode := currentCurrency(r)

	for i, p := range products {
		usd := p.GetPriceUsd()

		// [RELOAD]
		if i == 0 && !useDefaultCurrency {
			numUSD := float64(usd.Units) + float64(usd.Nanos)/1e9
			price, _ := fe.convertCurrency(r.Context(), usd, currencyCode)
			rate = (float64(price.Units) + float64(price.Nanos)/1e9) / numUSD

			ps[i] = productView{p, price}
		} else if rate == 1.0 || useDefaultCurrency {
			ps[i] = productView{p, p.GetPriceUsd()}
		} else {
			price, _ := fe.convertCurrency(r.Context(), usd, currencyCode)
			ps[i] = productView{p, price}
		}

		//if useDefault {
		//	ps[i] = productView{p, p.GetPriceUsd()}
		//} else {
		//	price, _ := fe.convertCurrency(r.Context(), p.GetPriceUsd(), currentCurrency(r))
		//	ps[i] = productView{p, price}
		//}
	}

	ExecuteHomeTemplate(r.Context(), w, map[string]interface{}{
		"session_id":    sessionID(r),
		"request_id":    r.Context().Value(ctxKeyRequestID{}),
		"user_currency": currencyCode,
		"currencies":    currencies,
		"products":      ps,
		"cart_size":     len(cart),
		"banner_color":  os.Getenv("BANNER_COLOR"), // illustrates canary deployments
		"ad":            fe.chooseAd(r, log),
		"render":        r.FormValue("render"),
	})

	//if err := templates.ExecuteTemplateReload(r.Context(), w, "home", map[string]interface{}{
	//	"session_id":    sessionID(r),
	//	"request_id":    r.Context().Value(ctxKeyRequestID{}),
	//	"user_currency": currencyCode,
	//	"currencies":    currencies,
	//	"products":      ps,
	//	"cart_size":     len(cart),
	//	"banner_color":  os.Getenv("BANNER_COLOR"), // illustrates canary deployments
	//	"ad":            fe.chooseAd(r, log),
	//	"render":        r.FormValue("render"),
	//}); err != nil {
	//	log.Error(err)
	//}
}

type writer interface {
	io.Writer
	WriteString(s string) (int, error)
}

func ExecuteHomeTemplate(ctx context.Context, w http.ResponseWriter, data map[string]interface{}) {
	if metaVal := ctx.Value(tracer.ContextMetaKey{}); metaVal != nil {
		meta := metaVal.(*tracer.ContextMeta)
		if trace, ok := tracer.Store.GetByContextMeta(meta); ok {
			trace.Records = append(trace.Records, &tracer.Record{
				Type:        tracer.RecordType_RecordSend,
				Timestamp:   time.Now().UnixNano(),
				MessageName: meta.Url(),
				Uuid:        meta.UUID(),
				Service:     tracer.ServiceUUID,
			})

			trace.Rlfis = nil
			trace.Tfis = nil

			data["fi_trace"] = trace

			//delete trace from tracer.Store
			tracer.Store.DeleteByContextMeta(meta)
		}
	}

	if data["render"] == "json" {
		w.Header().Set(rHtml.ContentType, rHtml.ContentTypeJSON)

		json.NewEncoder(w).Encode(data)

		return
	}

	wr := bufio.NewWriter(w)

	HeaderTemplate(wr, data)
	HomeTemplate(wr, data)
	FooterTemplate(wr, data)

	wr.Flush()
}

func HomeTemplate(wr writer, data map[string]interface{}) {
	wr.WriteString(`
	<main role="main">
        <section class="jumbotron text-center mb-0" style="background-color: `)
	wr.WriteString(data["banner_color"].(string))
	wr.WriteString(`;">
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
`)

	if products, ok := data["products"].([]productView); ok {
		for _, product := range products {
			wr.WriteString(`
                	<div class="col-md-4">
                 	   <div class="card mb-4 box-shadow">
                 	       <a href="/product/`)
			wr.WriteString(product.Item.Id)
			wr.WriteString(`">
                    	        <img class="card-img-top" alt =""
                     	           style="width: 100%%; height: auto;"
                     	           src="`)
			wr.WriteString(product.Item.Picture)
			wr.WriteString(`">
                        	</a>
                        	<div class="card-body">
                            	<h5 class="card-title">
                                `)
			wr.WriteString(product.Item.Name)
			wr.WriteString(`
                            	</h5>
                            	<div class="d-flex justify-content-between align-items-center">
                                	<div class="btn-group">
                                    	<a href="/product/`)
			wr.WriteString(product.Item.Id)
			wr.WriteString(`">
                                        	<button type="button" class="btn btn-sm btn-outline-secondary">Buy</button>
                                    	</a>
                                	</div>
                                	<small class="text-muted">
`)

			money := product.Price
			fmt.Fprintf(wr, "%s %d.%02d", money.CurrencyCode, money.Units, money.Nanos/1e7)

			wr.WriteString(`
                                	</small>
                            	</div>
                        	</div>
                    	</div>
                	</div>
`)
		}
	}

	wr.WriteString(`
            	</div>
            	<div class="row">
`)

	AdTemplate(wr, data)

	if fi_trace, ok := data["fi_trace"]; ok {
		wr.WriteString(`
				</div>
                <div class="trace">
                    `)
		json.NewEncoder(wr).Encode(fi_trace)
		//wr.WriteString(rTemplate.MarshalTracing(fi_trace.(*tracer.Trace)))
		wr.WriteString(`
                </div>
			</div>
        </div>
    </main>
`)
	} else {
		wr.WriteString(`
				</div>
            </div>
        </div>
    </main>
`)
	}
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

func HeaderTemplate(wr writer, data map[string]interface{}) {
	if currencies, ok := data["currencies"]; ok {
		user_currency := data["user_currency"].(string)

		wr.WriteString(`
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
					`)
		for _, currency := range currencies.([]string) {
			if currency == user_currency {
				wr.WriteString(`
						<option value="`)
				wr.WriteString(currency)
				wr.WriteString(`" selected="selected">`)
				wr.WriteString(currency)
				wr.WriteString(`</option>
`)
			} else {
				wr.WriteString(`
						<option value="`)
				wr.WriteString(currency)
				wr.WriteString(`">`)
				wr.WriteString(currency)
				wr.WriteString(`</option>
`)
			}
		}

		wr.WriteString(`
					</select>
					<a class="btn btn-primary btn-light ml-2" href="/cart" role="button">View Cart (`)
		fmt.Fprintf(wr, "%v", data["cart_size"])
		wr.WriteString(`</a>
				</form>
            </div>
        </div>
    </header>
`)
	} else {
		wr.WriteString(headerStr)
	}
}

func AdTemplate(wr writer, data map[string]interface{}) {
	if ad := data["ad"].(*pb.Ad); ad != nil {
		wr.WriteString(`
				<div class="container">
					<div class="alert alert-dark" role="alert">
						<strong>Advertisement:</strong>
						<a href="`)
		wr.WriteString(ad.RedirectUrl)
		wr.WriteString(`" rel="nofollow" target="_blank" class="alert-link">
							`)
		wr.WriteString(ad.Text)
		wr.WriteString(`
						</a>
					</div>
				</div>
`)
	}
}

func FooterTemplate(wr writer, data map[string]interface{}) {
	wr.WriteString(`
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
                session-id: `)
	wr.WriteString(data["session_id"].(string))
	wr.WriteString(
		`</br>
                request-id: `)
	wr.WriteString(data["request_id"].(string))
	wr.WriteString(`</br>
            </small>
        </div>
    </footer>
    <script src="https://stackpath.bootstrapcdn.com/bootstrap/4.1.1/js/bootstrap.min.js" integrity="sha384-smHYKdLADwkXOn1EmN1qk/HfnUcbVRZyYmZ4qpPea6sjB/pTJ0euyQp0Mk8ck+5T" crossorigin="anonymous"></script>
</body>
</html>
`)
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
		log.Errorf("failed to retrieve ads: %s", err.Error())

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
