package main

import (
	"bytes"
	pb "github.com/AleckDarcy/opencensus-microservices-demo/src/frontend/genproto"
	"testing"
)

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
		"cart_size": 1,
		"banner_color": "black",
		"ad": (*pb.Ad)(nil),
	}

	wr := bytes.NewBufferString("")

	ExecuteHomeTemplate(wr, data)

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
		"cart_size": 1,
		"banner_color": "black",
		"ad": (*pb.Ad)(nil),
	}

	for i := 0; i < b.N; i++ {
		wr := bytes.NewBufferString("")

		ExecuteHomeTemplate(wr, data)
	}
}