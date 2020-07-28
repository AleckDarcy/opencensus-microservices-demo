package main

import (
	"bytes"
	"context"
	pb "github.com/AleckDarcy/opencensus-microservices-demo/src/frontend/genproto"
	"net/http"
	"testing"
	"time"
)

type mockResponseWriter struct {
	*bytes.Buffer
}

func (wr *mockResponseWriter) Header() http.Header {
	return nil
}

func (wr *mockResponseWriter) WriteHeader(statusCode int) {

}

type emptyCtx int

var mockCtx = emptyCtx(0)

func (*emptyCtx) Deadline() (deadline time.Time, ok bool) {
	return
}

func (*emptyCtx) Done() <-chan struct{} {
	return nil
}

func (*emptyCtx) Err() error {
	return nil
}

func (*emptyCtx) Value(key interface{}) interface{} {
	return nil
}

func TestHomeTemplate(t *testing.T) {
	data := map[string]interface{}{
		"session_id":    "session_id",
		"request_id":    "request_id",
		"user_currency": "USD",
		"currencies":    []string{"USD", "CAD"},
		"products": []productView{
			{
				Item: &pb.Product{
					Id:          "11111",
					Name:        "AAAAA",
					Description: "aaaaa",
					Picture:     ",,,,,",
					PriceUsd: &pb.Money{
						CurrencyCode: "USD",
						Units:        1,
						Nanos:        1e9,
					},
				},
				Price: &pb.Money{
					CurrencyCode: "USD",
					Units:        1,
					Nanos:        1e9,
				},
			},
		},
		"cart_size":    1,
		"banner_color": "black",
		"ad":           (*pb.Ad)(nil),
	}

	wr := &mockResponseWriter{bytes.NewBufferString("")}
	t.Log(context.WithValue(nil, "test", nil))
	ExecuteHomeTemplate(context.WithValue(&mockCtx, "test", nil), wr, data)

	t.Log(wr.String())
}

func BenchmarkTemplate(b *testing.B) {
	data := map[string]interface{}{
		"session_id":    "session_id",
		"request_id":    "request_id",
		"user_currency": "USD",
		"currencies":    []string{"USD", "CAD"},
		"products": []productView{
			{
				Item: &pb.Product{
					Id:          "11111",
					Name:        "AAAAA",
					Description: "aaaaa",
					Picture:     ",,,,,",
					PriceUsd: &pb.Money{
						CurrencyCode: "USD",
						Units:        1,
						Nanos:        1e9,
					},
				},
				Price: &pb.Money{
					CurrencyCode: "USD",
					Units:        1,
					Nanos:        1e9,
				},
			},
		},
		"cart_size":    1,
		"banner_color": "black",
		"ad":           (*pb.Ad)(nil),
	}

	for i := 0; i < b.N; i++ {
		wr := &mockResponseWriter{bytes.NewBufferString("")}

		ExecuteHomeTemplate(context.WithValue(&mockCtx, "test", nil), wr, data)
	}
}
