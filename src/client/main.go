package main

import (
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strconv"
	"sync/atomic"
	"time"

	"github.com/AleckDarcy/reload/core/client/core"
	"github.com/AleckDarcy/reload/core/client/data"
	"github.com/AleckDarcy/reload/core/tracer"
	rHtml "github.com/AleckDarcy/reload/runtime/html"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"go.opencensus.io/plugin/ocgrpc"
	"google.golang.org/grpc"
)

const (
	port = "8081"
)

func main() {
	log := logrus.New()
	log.Level = logrus.DebugLevel
	log.Formatter = &logrus.TextFormatter{}

	addr := os.Getenv("LISTEN_ADDR")
	srvPort := port
	svc := new(clientSvc)

	r := mux.NewRouter()
	r.HandleFunc("/", svc.homeHandler).Methods(http.MethodGet)
	r.HandleFunc("/perf", svc.perfHandler).Methods(http.MethodPost, http.MethodGet)

	handler := &logHandler{log: log, next: r}

	log.Infof("starting server on " + addr + ":" + srvPort)
	log.Fatal(http.ListenAndServe(addr+":"+srvPort, handler))
}

func mustMapEnv(target *string, envKey string) {
	v := os.Getenv(envKey)
	if v == "" {
		panic(fmt.Sprintf("environment variable %q not set", envKey))
	}
	*target = v
}

func mustConnGRPC(ctx context.Context, addr string) *grpc.ClientConn {
	conn, err := grpc.DialContext(ctx, addr,
		grpc.WithInsecure(),
		grpc.WithStatsHandler(&ocgrpc.ClientHandler{}))
	if err != nil {
		panic(errors.Wrapf(err, "grpc: failed to connect %s", addr))
	}
	return conn
}

var (
	templates = template.Must(
		template.New("").Funcs(template.FuncMap{
			"renderJSON": renderJSON,
			//"marshalTracing": rTemplate.MarshalTracing,
		}).ParseGlob("templates/*.html"))
)

func renderJSON(j interface{}) string {
	bytes, err := json.Marshal(j)
	if err != nil {
		return fmt.Sprintf("rendor JSON err: %v", err)
	}

	return string(bytes)
}

type clientSvc struct {
	frontendSvcConn *grpc.ClientConn

	status core.Status
	result *core.Perf
}

func (s *clientSvc) homeHandler(w http.ResponseWriter, r *http.Request) {
	log := r.Context().Value(ctxKeyLog{}).(logrus.FieldLogger)

	status := atomic.LoadInt64(&s.status.Status)
	if status == core.Idle {
		s.renderResult(0, w, log)

		return
	}

	s.renderStatus(0, w, log)
}

func (s *clientSvc) perfHandler(w http.ResponseWriter, r *http.Request) {
	log := r.Context().Value(ctxKeyLog{}).(logrus.FieldLogger)

	task := time.Now().UnixNano()
	if sTask := r.FormValue("task"); sTask != "" {
		task, _ = strconv.ParseInt(sTask, 10, 64)
	}

	if task == s.status.TaskID {
		s.renderStatus(task, w, log)

		return
	} else if !atomic.CompareAndSwapInt64(&s.status.Status, core.Idle, core.Running) {
		s.renderStatus(task, w, log)

		return
	}

	s.status.TaskID = task

	id, _ := strconv.ParseUint(r.FormValue("id"), 10, 64)
	s.status.ID = int(id)

	nTest := int64(100)
	if sTest := r.FormValue("test"); sTest != "" {
		nTest, _ = strconv.ParseInt(sTest, 10, 64)
	}

	nRound := int64(1)
	if sRound := r.FormValue("round"); sRound != "" {
		nRound, _ = strconv.ParseInt(sRound, 10, 64)
	}

	nClients := []int{1, 2, 4, 8, 16, 32, 64, 128}
	if sClient := r.FormValue("client"); sClient != "" {
		nClient, _ := strconv.ParseInt(sClient, 10, 64)
		nClients = []int{int(nClient)}
	}

	url := r.FormValue("url")

	var caseConf []core.CaseConf
	var customizer core.Customizer

	switch id {
	case 1: // trace off & jaeger on
		caseConf = []core.CaseConf{
			{
				Request: &data.Request{
					Method:      data.HTTPGet,
					URL:         url,
					MessageName: "home",
					Trace:       nil,
					Expect: &data.ExpectedResponse{
						ContentType: rHtml.ContentTypeHTML,
					},
				},
			},
		}
	case 2: // trace on
		caseConf = []core.CaseConf{
			{
				Request: &data.Request{
					Method:      data.HTTPGet,
					URL:         url,
					MessageName: "home",
					Trace:       &tracer.Trace{},
					Expect: &data.ExpectedResponse{
						ContentType: rHtml.ContentTypeHTML,
					},
				},
			},
		}
	case 3: // trace on with frontend latency
		caseConf = []core.CaseConf{
			{
				Request: &data.Request{
					Method:      data.HTTPGet,
					URL:         url,
					MessageName: "home",
					Trace:       &tracer.Trace{},
					Expect: &data.ExpectedResponse{
						ContentType: rHtml.ContentTypeHTML,
						Action:      data.DeserializeTrace,
					},
				},
			},
		}
	case 4: // trace on with service latency
		caseConf = []core.CaseConf{
			{
				Request: &data.Request{
					Method:      data.HTTPGet,
					URL:         url,
					MessageName: "home",
					Trace:       &tracer.Trace{},
					Expect: &data.ExpectedResponse{
						ContentType: rHtml.ContentTypeHTML,
						Action:      data.DeserializeTrace | data.CustomizedRspFunc,
					},
				},
			},
		}

		customizer = RequestLatencyStore
	case 5: // trace on with trace sampling
		caseConf = []core.CaseConf{
			{
				Request: &data.Request{
					Method:      data.HTTPGet,
					URL:         url,
					MessageName: "home",
					Trace:       &tracer.Trace{},
					Expect: &data.ExpectedResponse{
						ContentType: rHtml.ContentTypeHTML,
						Action:      data.CustomizedRspFunc,
					},
				},
			},
		}

		mask := int64(100)
		if sMask := r.FormValue("mask"); sMask != "" {
			if mask, _ = strconv.ParseInt(sMask, 10, 64); mask <= 0 {
				mask = 100
			}
		}

		customizer = NewTraceSample(nTest, nRound, mask, nClients)
	default:
		s.renderStatus(task, w, log)
		atomic.StoreInt64(&s.status.Status, core.Idle)

		return
	}

	go func() {
		log.Infof("starting task, test: %d, round: %d, url: %s", nTest, nRound, url)
		s.result = core.RunPerf(nTest, nRound, nClients, caseConf, &s.status, customizer)
		s.status.CaseID = 0
		s.status.NClient = 0
		s.status.Round = 0

		atomic.StoreInt64(&s.status.Status, core.Idle)
	}()

	s.renderStatus(task, w, log)
}

func (s *clientSvc) renderResult(task int64, w http.ResponseWriter, log logrus.FieldLogger) {
	report := &core.Report{}
	if s.result != nil {
		report = core.GetReport(s.result)
	}

	if err := templates.ExecuteTemplate(w, "result", map[string]interface{}{
		"task":   task,
		"report": report,
		"result": s.result,
	}); err != nil {
		log.Error(err)
	}
}

func (s *clientSvc) renderStatus(task int64, w http.ResponseWriter, log logrus.FieldLogger) {
	if err := templates.ExecuteTemplate(w, "status", map[string]interface{}{
		"task":   task,
		"status": &s.status,
	}); err != nil {
		log.Error(err)
	}
}
