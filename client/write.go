// Copyright (c) 2016 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package client

import (
	"github.com/m3db/m3db/generated/thrift/rpc"
	"github.com/m3db/m3db/pool"
)

var (
	writeOpZeroed writeOp
)

type writeOp struct {
	request      rpc.WriteRequest
	idDatapoint  rpc.IDDatapoint
	datapoint    rpc.Datapoint
	completionFn completionFn
}

func (w *writeOp) reset() {
	*w = writeOpZeroed
	w.idDatapoint.Datapoint = &w.datapoint
	w.request.IdDatapoint = &w.idDatapoint
}

func (w *writeOp) Size() int {
	// Writes always represent a single write
	return 1
}

func (w *writeOp) CompletionFn() completionFn {
	return w.completionFn
}

type writeOpPool interface {
	// Init pool
	Init()

	// Get a write op
	Get() *writeOp

	// Put a write op
	Put(w *writeOp)
}

type poolOfWriteOp struct {
	pool pool.ObjectPool
}

func newWriteOpPool(opts pool.ObjectPoolOptions) writeOpPool {
	p := pool.NewObjectPool(opts)
	return &poolOfWriteOp{p}
}

func (p *poolOfWriteOp) Init() {
	p.pool.Init(func() interface{} {
		w := &writeOp{}
		w.reset()
		return w
	})
}

func (p *poolOfWriteOp) Get() *writeOp {
	w := p.pool.Get().(*writeOp)
	return w
}

func (p *poolOfWriteOp) Put(w *writeOp) {
	w.reset()
	p.pool.Put(w)
}
