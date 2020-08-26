package main

import (
	"github.com/AleckDarcy/reload/core/client/data"
	"time"
)

type TraceSampleRate struct {}

func (t *TraceSampleRate) NClientInit() {}

func (t *TraceSampleRate) RoundInit() {}

func (t *TraceSampleRate) RspFunc(req *data.Request, rsp *data.Response) {}

func (t *TraceSampleRate) NClientFinish() interface{} {
	time.Sleep(time.Second)

	return nil
}

func (t *TraceSampleRate) RoundFinish() interface{} {
	return nil
}

func (t *TraceSampleRate) TestFinish() interface{} {
	return t
}
