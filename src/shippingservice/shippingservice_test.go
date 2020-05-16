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
	"fmt"
	"log"
	"net"
	"testing"
	"time"

	"github.com/AleckDarcy/reload/core/tracer"
	"go.opencensus.io/plugin/ocgrpc"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"

	pb "github.com/AleckDarcy/opencensus-microservices-demo/src/shippingservice/genproto"
)

// TestGetQuote is a basic check on the GetQuote RPC service.
func TestGetQuote(t *testing.T) {
	s := server{}

	// A basic test case to test logic and protobuf interactions.
	req := &pb.GetQuoteRequest{
		Address: &pb.Address{
			StreetAddress: "Muffin Man",
			City:          "London",
			State:         "",
			Country:       "England",
		},
		Items: []*pb.CartItem{
			{
				ProductId: "23",
				Quantity:  1,
			},
			{
				ProductId: "46",
				Quantity:  3,
			},
		},
	}

	res, err := s.GetQuote(context.Background(), req)
	if err != nil {
		t.Errorf("TestGetQuote (%v) failed", err)
	}
	if res.CostUsd.GetUnits() != 11 || res.CostUsd.GetNanos() != 220000000 {
		t.Errorf("TestGetQuote: Quote value '%d.%d' does not match expected '%s'", res.CostUsd.GetUnits(), res.CostUsd.GetNanos(), "11.220000000")
	}
}

func run(port int) string {
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal(err)
	}
	srv := grpc.NewServer(grpc.StatsHandler(&ocgrpc.ServerHandler{}))
	svc := &server{}
	pb.RegisterShippingServiceServer(srv, svc)
	healthpb.RegisterHealthServer(srv, svc)
	go srv.Serve(l)
	return l.Addr().String()
}

func TestGetQuote_Tracing(t *testing.T) {
	ctx := context.Background()

	addr := run(0)
	conn, err := grpc.Dial(addr,
		grpc.WithInsecure(),
		grpc.WithStatsHandler(&ocgrpc.ClientHandler{}))
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()
	client := pb.NewShippingServiceClient(conn)

	trace := &tracer.Trace{Id: 1}
	req := &pb.GetQuoteRequest{
		Address:  &pb.Address{StreetAddress: "Muffin Man", City: "London", State: "", Country: "England"},
		Items:    []*pb.CartItem{{ProductId: "23", Quantity: 1}, {ProductId: "46", Quantity: 3}},
		FI_Trace: &tracer.Trace{Id: 1},
	}

	// set fault-injection
	req.FI_Trace.Tfi = &tracer.TFI{
		Type:  tracer.FaultType_FaultCrash,
		Name:  "GetQuoteRequest",
		After: []*tracer.TFIMeta{{Name: "GetQuoteRequest", Times: 2}},
	}
	testSuccess(t, client, ctx, trace, req)
	testSuccess(t, client, ctx, trace, req)
	testFault(t, client, ctx, trace, req, "crash", 0)
	testFault(t, client, ctx, trace, req, "crash", 0)

	// set fault-injection
	req.FI_Trace.Tfi = &tracer.TFI{
		Type:  tracer.FaultType_FaultDelay,
		Name:  "GetQuoteRequest",
		Delay: 1000,
		After: []*tracer.TFIMeta{{Name: "GetQuoteRequest", Times: 2, Already: 2}},
	}
	testFault(t, client, ctx, trace, req, "delay", 1000)
	testFault(t, client, ctx, trace, req, "delay", 1000)

	// unset fault-injection
	req.FI_Trace.Tfi = nil
	testSuccess(t, client, ctx, trace, req)

	for i, record := range trace.Records {
		t.Logf("%d type: %v time: %v name: %s", i, record.Type, record.Timestamp, record.MessageName)
	}
}

func testSuccess(t *testing.T, client pb.ShippingServiceClient, ctx context.Context, trace *tracer.Trace, msg tracer.Tracer) bool {
	record := &tracer.Record{
		Type:        tracer.RecordType_RecordSend,
		Timestamp:   time.Now().UnixNano(),
		MessageName: msg.GetFI_Name(),
		Uuid:        tracer.NewUUID(),
	}
	trace.Records = append(trace.Records, record)

	msgT := msg.GetFI_Trace()
	msgT.Records = []*tracer.Record{record}

	var res tracer.Tracer
	switch req := msg.(type) {
	case *pb.GetQuoteRequest:
		res, _ = client.GetQuote(ctx, req)
	case *pb.ShipOrderRequest:
		res, _ = client.ShipOrder(ctx, req)
	default:
		panic("unreachable code")
	}

	trace.Records = append(trace.Records, res.GetFI_Trace().Records...)
	record = &tracer.Record{
		Type:        tracer.RecordType_RecordReceive,
		Timestamp:   time.Now().UnixNano(),
		MessageName: res.GetFI_Name(),
		Uuid:        trace.Records[len(trace.Records)-1].Uuid,
	}
	trace.Records = append(trace.Records, record)

	msg.GetFI_Trace().CalFI(append(res.GetFI_Trace().Records, record))

	return true
}

func testFault(t *testing.T, client pb.ShippingServiceClient, ctx context.Context, trace *tracer.Trace, msg tracer.Tracer, fault string, delay int64) bool {
	signal := make(chan bool, 1)
	go func(signal chan bool) {
		signal <- testSuccess(t, client, ctx, trace, msg)
	}(signal)

	if fault == "crash" {
		select {
		case <-signal:
			t.Fatal("should have triggered crash")
		case <-time.After(1 * time.Second):
			t.Log("crash has been triggered")
		}
	} else if fault == "delay" {
		select {
		case res := <-signal:
			t.Log("delay has been triggered")

			return res
		case <-time.After(time.Duration(100+delay) * time.Millisecond):
			t.Fatal("should have triggered delay")
		}
	} else {
		panic("unreachable code")
	}

	return true
}

// TestShipOrder is a basic check on the ShipOrder RPC service.
func TestShipOrder(t *testing.T) {
	s := server{}

	// A basic test case to test logic and protobuf interactions.
	req := &pb.ShipOrderRequest{
		Address: &pb.Address{
			StreetAddress: "Muffin Man",
			City:          "London",
			State:         "",
			Country:       "England",
		},
		Items: []*pb.CartItem{
			{
				ProductId: "23",
				Quantity:  1,
			},
			{
				ProductId: "46",
				Quantity:  3,
			},
		},
	}

	res, err := s.ShipOrder(context.Background(), req)
	if err != nil {
		t.Errorf("TestShipOrder (%v) failed", err)
	}
	// @todo improve quality of this test to check for a pattern such as '[A-Z]{2}-\d+-\d+'.
	if len(res.TrackingId) != 18 {
		t.Errorf("TestShipOrder: Tracking ID is malformed - has %d characters, %d expected", len(res.TrackingId), 18)
	}
}
