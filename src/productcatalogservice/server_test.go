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
	"context"
	"testing"
	"time"

	"github.com/AleckDarcy/reload/core/tracer"
	"github.com/golang/protobuf/proto"
	"github.com/google/go-cmp/cmp"
	"go.opencensus.io/plugin/ocgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/census-ecosystem/opencensus-microservices-demo/src/productcatalogservice/genproto"
)

func TestServer(t *testing.T) {
	ctx := context.Background()
	addr := run(0)
	conn, err := grpc.Dial(addr,
		grpc.WithInsecure(),
		grpc.WithStatsHandler(&ocgrpc.ClientHandler{}))
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()
	client := pb.NewProductCatalogServiceClient(conn)
	res, err := client.ListProducts(ctx, &pb.ListProductsRequest{})
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(res.Products, parseCatalog(), cmp.Comparer(proto.Equal)); diff != "" {
		t.Error(diff)
	}

	got, err := client.GetProduct(ctx, &pb.GetProductRequest{Id: "OLJCESPC7Z"})
	if err != nil {
		t.Fatal(err)
	}
	if want := parseCatalog()[0]; !proto.Equal(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
	_, err = client.GetProduct(ctx, &pb.GetProductRequest{Id: "N/A"})
	if got, want := status.Code(err), codes.NotFound; got != want {
		t.Errorf("got %s, want %s", got, want)
	}

	sres, err := client.SearchProducts(ctx, &pb.SearchProductsRequest{Query: "typewriter"})
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(sres.Results, []*pb.Product{parseCatalog()[0]}, cmp.Comparer(proto.Equal)); diff != "" {
		t.Error(diff)
	}
}

func TestServer_TFI1(t *testing.T) {
	ctx := context.Background()

	addr := run(0)
	conn, err := grpc.Dial(addr,
		grpc.WithInsecure(),
		grpc.WithStatsHandler(&ocgrpc.ClientHandler{}))
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()
	client := pb.NewProductCatalogServiceClient(conn)

	trace := &tracer.Trace{Id: 1}
	req := &pb.ListProductsRequest{FI_Trace: &tracer.Trace{Id: 1}}

	// set fault-injection
	req.FI_Trace.Tfi = &tracer.TFI{
		Type:  tracer.FaultType_FaultCrash,
		Name:  "ListProductsRequest",
		After: []*tracer.TFIMeta{{Name: "ListProductsRequest", Times: 2}},
	}
	testSuccess(t, client, ctx, trace, req)
	testSuccess(t, client, ctx, trace, req)
	testFault(t, client, ctx, trace, req, "crash", 0)
	testFault(t, client, ctx, trace, req, "crash", 0)

	// set fault-injection
	req.FI_Trace.Tfi = &tracer.TFI{
		Type:  tracer.FaultType_FaultDelay,
		Name:  "ListProductsRequest",
		Delay: 1000,
		After: []*tracer.TFIMeta{{Name: "ListProductsRequest", Times: 1, Already: 2}},
	}
	testFault(t, client, ctx, trace, req, "delay", 1000)
	testFault(t, client, ctx, trace, req, "delay", 1000)

	// unset fault-injection
	req.FI_Trace.Tfi = nil
	testSuccess(t, client, ctx, trace, req)

	for i, record := range trace.Records {
		t.Logf("%d type: %v time: %v name: %s", i, record.Type, time.Unix(record.Timestamp/1e9, record.Timestamp%1e9), record.MessageName)
	}
}

func TestServer_TFI2(t *testing.T) {
	ctx := context.Background()

	addr := run(0)
	conn, err := grpc.Dial(addr,
		grpc.WithInsecure(),
		grpc.WithStatsHandler(&ocgrpc.ClientHandler{}))
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()
	client := pb.NewProductCatalogServiceClient(conn)

	trace := &tracer.Trace{Id: 1}
	req1 := &pb.ListProductsRequest{FI_Trace: &tracer.Trace{Id: 1}}

	// set fault-injection
	req1.FI_Trace.Tfi = &tracer.TFI{
		Type: tracer.FaultType_FaultCrash,
		Name: "ListProductsRequest",
		After: []*tracer.TFIMeta{
			{Name: "GetProductRequest", Times: 1},
			{Name: "SearchProductsRequest", Times: 1},
		},
	}
	testSuccess(t, client, ctx, trace, req1)

	req2 := &pb.GetProductRequest{Id: "OLJCESPC7Z"}
	req2.FI_Trace = req1.FI_Trace
	testSuccess(t, client, ctx, trace, req2)
	testSuccess(t, client, ctx, trace, req1)

	req3 := &pb.SearchProductsRequest{Query: "typewriter"}
	req3.FI_Trace = req2.FI_Trace
	testSuccess(t, client, ctx, trace, req3)

	req1.FI_Trace = req3.FI_Trace
	testFault(t, client, ctx, trace, req1, "crash", 0)

	// unset fault-injection
	req1.FI_Trace.Tfi = nil
	testSuccess(t, client, ctx, trace, req1)

	for i, record := range trace.Records {
		t.Logf("%d type: %v time: %v name: %s", i, record.Type, time.Unix(record.Timestamp/1e9, record.Timestamp%1e9), record.MessageName)
	}
}

func testSuccess(t *testing.T, client pb.ProductCatalogServiceClient, ctx context.Context, trace *tracer.Trace, msg tracer.Tracer) bool {
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
	case *pb.ListProductsRequest:
		res, _ = client.ListProducts(ctx, req)
	case *pb.GetProductRequest:
		res, _ = client.GetProduct(ctx, req)
	case *pb.SearchProductsRequest:
		res, _ = client.SearchProducts(ctx, req)
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

func testFault(t *testing.T, client pb.ProductCatalogServiceClient, ctx context.Context, trace *tracer.Trace, msg tracer.Tracer, fault string, delay int64) bool {
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
