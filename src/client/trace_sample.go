package main

import (
	"encoding/json"
	"strings"
	"sync/atomic"
	"time"

	"github.com/AleckDarcy/reload/core/client/core"
	"github.com/AleckDarcy/reload/core/client/data"
	"github.com/AleckDarcy/reload/core/client/parser"
	"github.com/AleckDarcy/reload/core/tracer"

	"golang.org/x/net/html"
)

func NewTraceSample(nTest, nRound, mask int64, nClients []int) core.Customizer {
	t := &TraceSample{
		NClients: make([]NClientTraceSample, len(nClients)),

		mask: mask,
	}

	for i := range t.NClients {
		nClient := &t.NClients[i]
		nClient.NClient = nClients[i]
		nClient.Rounds = make([]RoundTraceSample, nRound)

		for j := range nClient.Rounds {
			round := &nClient.Rounds[j]
			round.Traces = make([]*tracer.Trace, nTest/mask)
		}
	}

	return t
}

type TraceSample struct {
	NClients []NClientTraceSample

	mask     int64
	iNClient int
	nClient  *NClientTraceSample // current NClientTraceSample
	round    *RoundTraceSample   // current RoundTraceSample
}

type NClientTraceSample struct {
	NClient int
	Rounds  []RoundTraceSample

	iRound int
}

type RoundTraceSample struct {
	Traces []*tracer.Trace

	traceID int64
}

func (t *TraceSample) NClientInit() {
	t.nClient = &t.NClients[t.iNClient]
	t.iNClient++
}

func (t *TraceSample) RoundInit() {
	t.round = &t.nClient.Rounds[t.nClient.iRound]
	t.nClient.iRound++
}

func (t *TraceSample) RspFunc(rsp *data.Response) {
	if traceID := atomic.AddInt64(&t.round.traceID, 1); traceID%t.mask == 0 {
		node, _ := html.Parse(strings.NewReader(string(rsp.Body)))
		if traceNode := parser.GetElementByClass(node, "trace"); traceNode != nil {
			traceString := parser.GetJSON(traceNode)

			rsp.Trace = &tracer.Trace{}
			if err := json.Unmarshal([]byte(traceString), rsp.Trace); err != nil {
				t.round.Traces[traceID/t.mask-1] = nil

				return
			}

			rsp.Trace.Records = append(rsp.Trace.Records, &tracer.Record{
				Type:        tracer.RecordType_RecordReceive,
				Timestamp:   time.Now().UnixNano(),
				MessageName: "/", // hard code
				Uuid:        "generated by client",
				Service:     "client",
			})

			t.round.Traces[traceID/t.mask-1] = rsp.Trace
		}
	}
}

func (t *TraceSample) NClientFinish() interface{} {
	return nil
}

func (t *TraceSample) RoundFinish() interface{} {
	return nil
}

func (t *TraceSample) TestFinish() interface{} {
	return t
}