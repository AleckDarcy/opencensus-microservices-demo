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
	"context"
	"flag"
	"fmt"
	"github.com/golang/protobuf/jsonpb"
	"io/ioutil"
	"net"
	"os"
	"strings"
	"time"

	pb "github.com/AleckDarcy/opencensus-microservices-demo/src/productcatalogservice/genproto"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"

	"cloud.google.com/go/profiler"
	"contrib.go.opencensus.io/exporter/stackdriver"
	"github.com/sirupsen/logrus"
	"go.opencensus.io/exporter/jaeger"
	"go.opencensus.io/plugin/ocgrpc"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	_ "github.com/AleckDarcy/reload"
)

var (
	catalogJSON []byte

	port = flag.Int("port", 3550, "port to listen at")

	log = logrus.New()
)

func init() {
	c, err := ioutil.ReadFile("products.json")
	if err != nil {
		log.Errorf("failed to open product catalog json file: %v", err)
	}
	catalogJSON = c

	products = func() []*pb.Product {
		var cat pb.ListProductsResponse

		if err := jsonpb.Unmarshal(bytes.NewReader(catalogJSON), &cat); err != nil {
			log.Errorf("warning: failed to parse the catalog JSON: %v", err)
			return nil
		}

		return cat.Products
	}()

	log.Level = getLogLevel()
	log.Infof("successfully parsed product catalog json")
}

var jaegerOn string

func getLogLevel() logrus.Level {
	v := os.Getenv("LOG_LEVEL")

	switch v {
	case "panic":
		return logrus.PanicLevel
	case "fatal":
		return logrus.FatalLevel
	case "error":
		return logrus.ErrorLevel
	case "warn":
		return logrus.WarnLevel
	case "info":
		return logrus.InfoLevel
	default:
		return logrus.DebugLevel
	}
}

func main() {
	mustMapEnv(&jaegerOn, "JAEGER_ON")

	go initTracing()
	go initProfiling("productcatalogservice", "1.0.0")
	flag.Parse()

	log.Infof("starting grpc server at :%d", *port)
	run(*port)
	select {}
}

func run(port int) string {
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal(err)
	}
	srv := grpc.NewServer(grpc.StatsHandler(&ocgrpc.ServerHandler{}))
	svc := &productCatalog{}
	pb.RegisterProductCatalogServiceServer(srv, svc)
	healthpb.RegisterHealthServer(srv, svc)
	go srv.Serve(l)
	return l.Addr().String()
}

func initJaegerTracing() {

	// Register the Jaeger exporter to be able to retrieve
	// the collected spans.
	exporter, err := jaeger.NewExporter(jaeger.Options{
		Endpoint: "http://jaeger:14268",
		Process: jaeger.Process{
			ServiceName: "productcatalogservice",
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	trace.RegisterExporter(exporter)
}

func initStats(exporter *stackdriver.Exporter) {
	view.RegisterExporter(exporter)
	if err := view.Register(ocgrpc.DefaultServerViews...); err != nil {
		log.Info("Error registering default server views")
	} else {
		log.Info("Registered default server views")
	}
}

func initTracing() {
	if jaegerOn == "false" {
		return
	}

	// This is a demo app with low QPS. trace.AlwaysSample() is used here
	// to make sure traces are available for observation and analysis.
	// In a production environment or high QPS setup please use
	// trace.ProbabilitySampler set at the desired probability.
	trace.ApplyConfig(trace.Config{DefaultSampler: trace.AlwaysSample()})

	initJaegerTracing()

	// TODO(ahmetb) this method is duplicated in other microservices using Go
	// since they are not sharing packages.
	for i := 1; i <= 3; i++ {
		exporter, err := stackdriver.NewExporter(stackdriver.Options{})
		if err != nil {
			log.Errorf("info: failed to initialize stackdriver exporter: %+v", err)
		} else {
			trace.RegisterExporter(exporter)
			trace.ApplyConfig(trace.Config{DefaultSampler: trace.AlwaysSample()})
			log.Warn("registered stackdriver tracing")

			// Register the views to collect server stats.
			initStats(exporter)
			return
		}
		d := time.Second * 10 * time.Duration(i)
		log.Warn("sleeping %v to retry initializing stackdriver exporter", d)
		time.Sleep(d)
	}
	log.Printf("warning: could not initialize stackdriver exporter after retrying, giving up")
}

func initProfiling(service, version string) {
	// TODO(ahmetb) this method is duplicated in other microservices using Go
	// since they are not sharing packages.
	for i := 1; i <= 3; i++ {
		if err := profiler.Start(profiler.Config{
			Service:        service,
			ServiceVersion: version,
			// ProjectID must be set if not running on GCP.
			// ProjectID: "my-project",
		}); err != nil {
			log.Errorf("warn: failed to start profiler: %+v", err)
		} else {
			log.Warn("started stackdriver profiler")
			return
		}
		d := time.Second * 10 * time.Duration(i)
		log.Warnf("sleeping %v to retry initializing stackdriver profiler", d)
		time.Sleep(d)
	}
	log.Warnf("warning: could not initialize stackdriver profiler after retrying, giving up")
}

func mustMapEnv(target *string, envKey string) {
	v := os.Getenv(envKey)
	if v == "" {
		panic(fmt.Sprintf("environment variable %q not set", envKey))
	}
	*target = v
}

type productCatalog struct{}

var products []*pb.Product

func (p *productCatalog) Check(ctx context.Context, req *healthpb.HealthCheckRequest) (*healthpb.HealthCheckResponse, error) {
	return &healthpb.HealthCheckResponse{Status: healthpb.HealthCheckResponse_SERVING}, nil
}

func (p *productCatalog) ListProducts(context.Context, *pb.ListProductsRequest) (*pb.ListProductsResponse, error) {
	return &pb.ListProductsResponse{Products: products}, nil
}

func (p *productCatalog) GetProduct(ctx context.Context, req *pb.GetProductRequest) (*pb.Product, error) {
	for _, product := range products {
		if req.Id == product.Id {
			return product, nil
		}
	}

	return nil, status.Errorf(codes.NotFound, "no product with ID %s", req.Id)
}

func (p *productCatalog) SearchProducts(ctx context.Context, req *pb.SearchProductsRequest) (*pb.SearchProductsResponse, error) {
	// Intepret query as a substring match in name or description.
	var ps []*pb.Product
	for _, p := range products {
		if strings.Contains(strings.ToLower(p.Name), strings.ToLower(req.Query)) ||
			strings.Contains(strings.ToLower(p.Description), strings.ToLower(req.Query)) {
			ps = append(ps, p)
		}
	}
	return &pb.SearchProductsResponse{Results: ps}, nil
}
