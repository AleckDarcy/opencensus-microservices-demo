package main

import (
	"github.com/AleckDarcy/reload/core/client/data"
	"github.com/AleckDarcy/reload/core/tracer"
	"math"

	"sync"
)

type LatencyID uint64

const (
	GetSupportedCurrenciesRequestLatencyID LatencyID = iota
	ListProductsRequestID
	CurrencyConversionRequestID
	AdRequestID
	LatencyIDMax
)

type requestLatencyStore struct {
	latencies map[int64]*RequestLatencies

	lock sync.RWMutex
}

var RequestLatencyStore = &requestLatencyStore{
	latencies: map[int64]*RequestLatencies{},
}

type RequestLatencies struct {
	latencies [LatencyIDMax]RequestLatency
}

type RequestLatency struct {
	count     float64
	latencies []int64
}

type RequestLatenciesAvg struct {
	Latencies [LatencyIDMax]RequestLatencyAvg
}

type RequestLatencyAvg struct {
	Max, Min int64
	Mean     float64
	StdErr   float64
	Total    float64
}

func (s *requestLatencyStore) NClientInit() {
	s.latencies = map[int64]*RequestLatencies{}
}

func (s *requestLatencyStore) RoundInit() {}

func (s *requestLatencyStore) RspFunc(req *data.Request, rsp *data.Response) {
	trace := rsp.Trace

	reqLatencies := &RequestLatencies{}
	sendTimes := map[tracer.UUID]int64{}

	for _, record := range trace.Records {
		if record.Type == tracer.RecordType_RecordSend {
			switch record.MessageName {
			case "GetSupportedCurrenciesRequest", "ListProductsRequest", "CurrencyConversionRequest", "AdRequest":
				sendTimes[record.Uuid] = record.Timestamp
			}
		} else if record.Type == tracer.RecordType_RecordReceive {
			latencyID := LatencyIDMax

			switch record.MessageName {
			case "GetSupportedCurrenciesResponse":
				latencyID = GetSupportedCurrenciesRequestLatencyID
			case "ListProductsResponse":
				latencyID = ListProductsRequestID
			case "Money":
				latencyID = CurrencyConversionRequestID
			case "AdResponse":
				latencyID = AdRequestID
			default:
				continue
			}

			latency := getLatency(sendTimes, record)
			reqLatency := &reqLatencies.latencies[latencyID]

			reqLatency.latencies = append(reqLatency.latencies, latency)
			reqLatency.count++
		}
	}

	s.lock.Lock()
	s.latencies[rsp.Trace.Id] = reqLatencies
	s.lock.Unlock()
}

func getLatency(sendTimes map[tracer.UUID]int64, record *tracer.Record) int64 {
	sendTime := sendTimes[record.Uuid]

	return (record.Timestamp - sendTime)/1e3
}

func (s *requestLatencyStore) NClientFinish() interface{} {
	reqLatenciesAvg := &RequestLatenciesAvg{}

	totalCounts := [LatencyIDMax]float64{}
	for _, reqLatencies := range RequestLatencyStore.latencies {
		for i, reqLatency := range reqLatencies.latencies {
			totalCounts[i] += reqLatency.count
		}
	}

	for _, reqLatencies := range RequestLatencyStore.latencies {
		for i, reqLatency := range reqLatencies.latencies {
			totalCount := totalCounts[i]
			avg := &reqLatenciesAvg.Latencies[i]
			for _, latency := range reqLatency.latencies {
				if avg.Min == 0 || avg.Min > latency {
					avg.Min = latency
				}

				if avg.Max < latency {
					avg.Max = latency
				}

				avg.Mean += float64(latency) / totalCount
			}
		}
	}

	for _, reqLatencies := range RequestLatencyStore.latencies {
		for i, reqLatency := range reqLatencies.latencies {
			avg := &reqLatenciesAvg.Latencies[i]
			for _, latency := range reqLatency.latencies {
				avg.StdErr += math.Pow(float64(latency)-avg.Mean, 2)
			}

			avg.StdErr = math.Sqrt(avg.StdErr)
			avg.StdErr /= totalCounts[i]
			avg.Total = totalCounts[i]
		}
	}

	return reqLatenciesAvg
}

func (s *requestLatencyStore) RoundFinish() interface{} {return nil}

func (s *requestLatencyStore) TestFinish() interface{} {return nil}
