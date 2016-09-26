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
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/m3db/m3db/generated/thrift/rpc"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/uber/tchannel-go/thrift"
)

var (
	testWriteBatchPool writeBatchRequestPool
	testWriteArrayPool idDatapointArrayPool
)

func init() {
	testWriteBatchPool = newWriteBatchRequestPool(nil)
	testWriteBatchPool.Init()
	testWriteArrayPool = newIDDatapointArrayPool(nil, 0)
	testWriteArrayPool.Init()
}

type hostQueueResult struct {
	result interface{}
	err    error
}

func newHostQueueTestOptions() Options {
	return NewOptions().
		SetHostQueueOpsFlushSize(4).
		SetHostQueueOpsArrayPoolSize(4).
		SetWriteBatchSize(4).
		SetFetchBatchSize(4).
		SetHostQueueOpsFlushInterval(0)
}

func TestHostQueueWriteErrorBeforeOpen(t *testing.T) {
	opts := newHostQueueTestOptions()
	queue := newHostQueue(h, testWriteBatchPool, testWriteArrayPool, opts)

	err := queue.Enqueue(&writeOp{})
	assert.Error(t, err)
	assert.Equal(t, err, errQueueNotOpen)
}

func TestHostQueueWriteErrorAfterClose(t *testing.T) {
	opts := newHostQueueTestOptions()
	queue := newHostQueue(h, testWriteBatchPool, testWriteArrayPool, opts)

	queue.Open()
	queue.Close()

	err := queue.Enqueue(&writeOp{})
	assert.Error(t, err)
	assert.Equal(t, err, errQueueNotOpen)
}

func TestHostQueueWriteBatches(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockConnPool := NewMockconnectionPool(ctrl)

	opts := newHostQueueTestOptions()
	queue := newHostQueue(h, testWriteBatchPool, testWriteArrayPool, opts).(*queue)
	queue.connPool = mockConnPool

	// Open
	mockConnPool.EXPECT().Open()
	queue.Open()
	assert.Equal(t, stateOpen, queue.state)

	// Prepare callback for writes
	var (
		results []hostQueueResult
		wg      sync.WaitGroup
	)
	callback := func(r interface{}, err error) {
		results = append(results, hostQueueResult{r, err})
		wg.Done()
	}

	// Prepare writes
	writes := []*writeOp{
		testWriteOp("testNs", "foo", 1.0, 1000, rpc.TimeType_UNIX_SECONDS, callback),
		testWriteOp("testNs", "bar", 2.0, 2000, rpc.TimeType_UNIX_SECONDS, callback),
		testWriteOp("testNs", "baz", 3.0, 3000, rpc.TimeType_UNIX_SECONDS, callback),
		testWriteOp("testNs", "qux", 4.0, 4000, rpc.TimeType_UNIX_SECONDS, callback),
	}
	wg.Add(len(writes))

	for i, write := range writes[:3] {
		assert.NoError(t, queue.Enqueue(write))
		assert.Equal(t, i+1, queue.Len())

		// Sleep some so that we can ensure flushing is not happening until queue is full
		time.Sleep(20 * time.Millisecond)
	}

	// Prepare mocks for flush
	mockClient := rpc.NewMockTChanNode(ctrl)
	writeBatch := func(ctx thrift.Context, req *rpc.WriteBatchRequest) {
		for i, write := range writes {
			assert.Equal(t, *req.Elements[i], *write.request.IdDatapoint)
		}
	}
	mockClient.EXPECT().WriteBatch(gomock.Any(), gomock.Any()).Do(writeBatch).Return(nil)

	mockConnPool.EXPECT().NextClient().Return(mockClient, nil)

	// Final write will flush
	assert.NoError(t, queue.Enqueue(writes[3]))
	assert.Equal(t, 0, queue.Len())

	// Wait for all writes
	wg.Wait()

	// Assert writes successful
	success := []hostQueueResult{{nil, nil}, {nil, nil}, {nil, nil}, {nil, nil}}
	assert.Equal(t, success, results)

	// Close
	var closeWg sync.WaitGroup
	closeWg.Add(1)
	mockConnPool.EXPECT().Close().Do(func() {
		closeWg.Done()
	})
	queue.Close()
	closeWg.Wait()
}

func TestHostQueueWriteBatchesDifferentNamespaces(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockConnPool := NewMockconnectionPool(ctrl)

	opts := newHostQueueTestOptions()
	queue := newHostQueue(h, testWriteBatchPool, testWriteArrayPool, opts).(*queue)
	queue.connPool = mockConnPool

	// Open
	mockConnPool.EXPECT().Open()
	queue.Open()
	assert.Equal(t, stateOpen, queue.state)

	// Prepare callback for writes
	var (
		results     []hostQueueResult
		resultsLock sync.Mutex
		wg          sync.WaitGroup
	)
	callback := func(r interface{}, err error) {
		resultsLock.Lock()
		results = append(results, hostQueueResult{r, err})
		resultsLock.Unlock()
		wg.Done()
	}

	// Prepare writes
	writes := []*writeOp{
		testWriteOp("testNs1", "foo", 1.0, 1000, rpc.TimeType_UNIX_SECONDS, callback),
		testWriteOp("testNs1", "bar", 2.0, 2000, rpc.TimeType_UNIX_SECONDS, callback),
		testWriteOp("testNs1", "baz", 3.0, 3000, rpc.TimeType_UNIX_SECONDS, callback),
		testWriteOp("testNs2", "qux", 4.0, 4000, rpc.TimeType_UNIX_SECONDS, callback),
	}
	wg.Add(len(writes))

	// Prepare mocks for flush
	mockClient := rpc.NewMockTChanNode(ctrl)
	writeBatch := func(ctx thrift.Context, req *rpc.WriteBatchRequest) {
		var writesForNamespace []*writeOp
		if req.NameSpace == "testNs1" {
			writesForNamespace = writes[:3]
		} else {
			writesForNamespace = writes[3:]
		}
		assert.Equal(t, len(writesForNamespace), len(req.Elements))
		for i, write := range writesForNamespace {
			assert.Equal(t, *req.Elements[i], *write.request.IdDatapoint)
		}
	}

	// Assert the writes will be handled in two batches
	mockClient.EXPECT().WriteBatch(gomock.Any(), gomock.Any()).Do(writeBatch).Return(nil).MinTimes(2).MaxTimes(2)
	mockConnPool.EXPECT().NextClient().Return(mockClient, nil).MinTimes(2).MaxTimes(2)

	for _, write := range writes {
		assert.NoError(t, queue.Enqueue(write))
	}

	// Wait for all writes
	wg.Wait()

	// Assert writes successful
	success := []hostQueueResult{{nil, nil}, {nil, nil}, {nil, nil}, {nil, nil}}
	assert.Equal(t, success, results)

	// Close
	var closeWg sync.WaitGroup
	closeWg.Add(1)
	mockConnPool.EXPECT().Close().Do(func() {
		closeWg.Done()
	})
	queue.Close()
	closeWg.Wait()
}

func TestHostQueueWriteBatchesNoClientAvailable(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockConnPool := NewMockconnectionPool(ctrl)

	opts := newHostQueueTestOptions()
	opts = opts.SetHostQueueOpsFlushInterval(time.Millisecond)
	queue := newHostQueue(h, testWriteBatchPool, testWriteArrayPool, opts).(*queue)
	queue.connPool = mockConnPool

	// Open
	mockConnPool.EXPECT().Open()
	queue.Open()
	assert.Equal(t, stateOpen, queue.state)

	// Prepare mocks for flush
	nextClientErr := fmt.Errorf("an error")
	mockConnPool.EXPECT().NextClient().Return(nil, nextClientErr)

	// Write
	var wg sync.WaitGroup
	wg.Add(1)
	callback := func(r interface{}, err error) {
		assert.Error(t, err)
		assert.Equal(t, nextClientErr, err)
		wg.Done()
	}
	assert.NoError(t, queue.Enqueue(testWriteOp("testNs", "foo", 1.0, 1000, rpc.TimeType_UNIX_SECONDS, callback)))

	// Wait for background flush
	wg.Wait()

	// Close
	var closeWg sync.WaitGroup
	closeWg.Add(1)
	mockConnPool.EXPECT().Close().Do(func() {
		closeWg.Done()
	})
	queue.Close()
	closeWg.Wait()
}

func TestHostQueueWriteBatchesPartialBatchErrs(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockConnPool := NewMockconnectionPool(ctrl)

	opts := newHostQueueTestOptions()
	opts = opts.SetHostQueueOpsFlushSize(2)
	queue := newHostQueue(h, testWriteBatchPool, testWriteArrayPool, opts).(*queue)
	queue.connPool = mockConnPool

	// Open
	mockConnPool.EXPECT().Open()
	queue.Open()
	assert.Equal(t, stateOpen, queue.state)

	// Prepare writes
	var wg sync.WaitGroup
	writeErr := "a write error"
	writes := []*writeOp{
		testWriteOp("testNs", "foo", 1.0, 1000, rpc.TimeType_UNIX_SECONDS, func(r interface{}, err error) {
			assert.Error(t, err)
			rpcErr, ok := err.(*rpc.Error)
			assert.True(t, ok)
			assert.Equal(t, rpc.ErrorType_INTERNAL_ERROR, rpcErr.Type)
			assert.Equal(t, writeErr, rpcErr.Message)
			wg.Done()
		}),
		testWriteOp("testNs", "bar", 2.0, 2000, rpc.TimeType_UNIX_SECONDS, func(r interface{}, err error) {
			assert.NoError(t, err)
			wg.Done()
		}),
	}
	wg.Add(len(writes))

	// Prepare mocks for flush
	mockClient := rpc.NewMockTChanNode(ctrl)
	writeBatch := func(ctx thrift.Context, req *rpc.WriteBatchRequest) {
		for i, write := range writes {
			assert.Equal(t, *req.Elements[i], *write.request.IdDatapoint)
		}
	}
	batchErrs := &rpc.WriteBatchErrors{Errors: []*rpc.WriteBatchError{
		&rpc.WriteBatchError{Index: 0, Err: &rpc.Error{
			Type:    rpc.ErrorType_INTERNAL_ERROR,
			Message: writeErr,
		}},
	}}
	mockClient.EXPECT().WriteBatch(gomock.Any(), gomock.Any()).Do(writeBatch).Return(batchErrs)
	mockConnPool.EXPECT().NextClient().Return(mockClient, nil)

	// Perform writes
	for _, write := range writes {
		assert.NoError(t, queue.Enqueue(write))
	}

	// Wait for flush
	wg.Wait()

	// Close
	var closeWg sync.WaitGroup
	closeWg.Add(1)
	mockConnPool.EXPECT().Close().Do(func() {
		closeWg.Done()
	})
	queue.Close()
	closeWg.Wait()
}

func TestHostQueueWriteBatchesEntireBatchErr(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockConnPool := NewMockconnectionPool(ctrl)

	opts := newHostQueueTestOptions()
	opts = opts.SetHostQueueOpsFlushSize(2)
	queue := newHostQueue(h, testWriteBatchPool, testWriteArrayPool, opts).(*queue)
	queue.connPool = mockConnPool

	// Open
	mockConnPool.EXPECT().Open()
	queue.Open()
	assert.Equal(t, stateOpen, queue.state)

	// Prepare writes
	var wg sync.WaitGroup
	writeErr := fmt.Errorf("an error")
	callback := func(r interface{}, err error) {
		assert.Error(t, err)
		assert.Equal(t, writeErr, err)
		wg.Done()
	}
	writes := []*writeOp{
		testWriteOp("testNs", "foo", 1.0, 1000, rpc.TimeType_UNIX_SECONDS, callback),
		testWriteOp("testNs", "bar", 2.0, 2000, rpc.TimeType_UNIX_SECONDS, callback),
	}
	wg.Add(len(writes))

	// Prepare mocks for flush
	mockClient := rpc.NewMockTChanNode(ctrl)
	writeBatch := func(ctx thrift.Context, req *rpc.WriteBatchRequest) {
		for i, write := range writes {
			assert.Equal(t, *req.Elements[i], *write.request.IdDatapoint)
		}
	}
	mockClient.EXPECT().WriteBatch(gomock.Any(), gomock.Any()).Do(writeBatch).Return(writeErr)
	mockConnPool.EXPECT().NextClient().Return(mockClient, nil)

	// Perform writes
	for _, write := range writes {
		assert.NoError(t, queue.Enqueue(write))
	}

	// Wait for flush
	wg.Wait()

	// Close
	var closeWg sync.WaitGroup
	closeWg.Add(1)
	mockConnPool.EXPECT().Close().Do(func() {
		closeWg.Done()
	})
	queue.Close()
	closeWg.Wait()
}

func TestHostQueueFetchBatches(t *testing.T) {
	namespace := "testNs"
	ids := []string{"foo", "bar", "baz", "qux"}
	result := &rpc.FetchRawBatchResult_{}
	for _ = range ids {
		result.Elements = append(result.Elements, &rpc.FetchRawResult_{Segments: []*rpc.Segments{}})
	}
	var expected []hostQueueResult
	for i := range ids {
		expected = append(expected, hostQueueResult{result.Elements[i].Segments, nil})
	}
	testHostQueueFetchBatches(t, namespace, ids, result, expected, nil, func(results []hostQueueResult) {
		assert.Equal(t, expected, results)
	})
}

func TestHostQueueFetchBatchesErrorOnNextClientUnavailable(t *testing.T) {
	namespace := "testNs"
	ids := []string{"foo", "bar", "baz", "qux"}
	expectedErr := fmt.Errorf("an error")
	var expected []hostQueueResult
	for _ = range ids {
		expected = append(expected, hostQueueResult{nil, expectedErr})
	}
	opts := &testHostQueueFetchBatchesOptions{
		nextClientErr: expectedErr,
	}
	testHostQueueFetchBatches(t, namespace, ids, nil, expected, opts, func(results []hostQueueResult) {
		assert.Equal(t, expected, results)
	})
}

func TestHostQueueFetchBatchesErrorOnFetchRawBatchError(t *testing.T) {
	namespace := "testNs"
	ids := []string{"foo", "bar", "baz", "qux"}
	expectedErr := fmt.Errorf("an error")
	var expected []hostQueueResult
	for _ = range ids {
		expected = append(expected, hostQueueResult{nil, expectedErr})
	}
	opts := &testHostQueueFetchBatchesOptions{
		fetchRawBatchErr: expectedErr,
	}
	testHostQueueFetchBatches(t, namespace, ids, nil, expected, opts, func(results []hostQueueResult) {
		assert.Equal(t, expected, results)
	})
}

func TestHostQueueFetchBatchesErrorOnFetchNoResponse(t *testing.T) {
	namespace := "testNs"
	ids := []string{"foo", "bar", "baz", "qux"}
	result := &rpc.FetchRawBatchResult_{}
	for _ = range ids[:len(ids)-1] {
		result.Elements = append(result.Elements, &rpc.FetchRawResult_{Segments: []*rpc.Segments{}})
	}
	var expected []hostQueueResult
	for i := range ids[:len(ids)-1] {
		expected = append(expected, hostQueueResult{result.Elements[i].Segments, nil})
	}
	expected = append(expected, hostQueueResult{nil, errQueueFetchNoResponse})
	testHostQueueFetchBatches(t, namespace, ids, result, expected, nil, func(results []hostQueueResult) {
		assert.Equal(t, expected, results)
	})
}

func TestHostQueueFetchBatchesErrorOnResultError(t *testing.T) {
	namespace := "testNs"
	ids := []string{"foo", "bar", "baz", "qux"}
	anError := &rpc.Error{Type: rpc.ErrorType_INTERNAL_ERROR, Message: "an error"}
	result := &rpc.FetchRawBatchResult_{}
	for _ = range ids[:len(ids)-1] {
		result.Elements = append(result.Elements, &rpc.FetchRawResult_{Segments: []*rpc.Segments{}})
	}
	result.Elements = append(result.Elements, &rpc.FetchRawResult_{Err: anError})
	var expected []hostQueueResult
	for i := range ids[:len(ids)-1] {
		expected = append(expected, hostQueueResult{result.Elements[i].Segments, nil})
	}
	testHostQueueFetchBatches(t, namespace, ids, result, expected, nil, func(results []hostQueueResult) {
		assert.Equal(t, expected, results[:len(results)-1])
		rpcErr, ok := results[len(results)-1].err.(*rpc.Error)
		assert.True(t, ok)
		assert.Equal(t, anError.Type, rpcErr.Type)
		assert.Equal(t, anError.Message, rpcErr.Message)
	})
}

type testHostQueueFetchBatchesOptions struct {
	nextClientErr    error
	fetchRawBatchErr error
}

func testHostQueueFetchBatches(
	t *testing.T,
	namespace string,
	ids []string,
	result *rpc.FetchRawBatchResult_,
	expected []hostQueueResult,
	testOpts *testHostQueueFetchBatchesOptions,
	assertion func(results []hostQueueResult),
) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockConnPool := NewMockconnectionPool(ctrl)

	opts := newHostQueueTestOptions()
	queue := newHostQueue(h, testWriteBatchPool, testWriteArrayPool, opts).(*queue)
	queue.connPool = mockConnPool

	// Open
	mockConnPool.EXPECT().Open()
	queue.Open()
	assert.Equal(t, stateOpen, queue.state)

	// Prepare callback for fetches
	var (
		results []hostQueueResult
		wg      sync.WaitGroup
	)
	callback := func(r interface{}, err error) {
		results = append(results, hostQueueResult{r, err})
		wg.Done()
	}

	// Prepare fetch batch op
	fetchBatch := &fetchBatchOp{
		request: rpc.FetchRawBatchRequest{
			RangeStart: 0,
			RangeEnd:   1,
			NameSpace:  namespace,
			Ids:        ids,
		},
	}
	for _ = range fetchBatch.request.Ids {
		fetchBatch.completionFns = append(fetchBatch.completionFns, callback)
	}
	wg.Add(len(fetchBatch.request.Ids))

	// Prepare mocks for flush
	mockClient := rpc.NewMockTChanNode(ctrl)
	if testOpts != nil && testOpts.nextClientErr != nil {
		mockConnPool.EXPECT().NextClient().Return(nil, testOpts.nextClientErr)
	} else if testOpts != nil && testOpts.fetchRawBatchErr != nil {
		fetchRawBatch := func(ctx thrift.Context, req *rpc.FetchRawBatchRequest) {
			assert.Equal(t, &fetchBatch.request, req)
		}
		mockClient.EXPECT().
			FetchRawBatch(gomock.Any(), gomock.Any()).
			Do(fetchRawBatch).
			Return(nil, testOpts.fetchRawBatchErr)

		mockConnPool.EXPECT().NextClient().Return(mockClient, nil)
	} else {
		fetchRawBatch := func(ctx thrift.Context, req *rpc.FetchRawBatchRequest) {
			assert.Equal(t, &fetchBatch.request, req)
		}
		mockClient.EXPECT().
			FetchRawBatch(gomock.Any(), gomock.Any()).
			Do(fetchRawBatch).
			Return(result, nil)

		mockConnPool.EXPECT().NextClient().Return(mockClient, nil)
	}

	// Fetch
	assert.NoError(t, queue.Enqueue(fetchBatch))

	// Wait for fetch to complete
	wg.Wait()

	// Assert results match expected
	assertion(results)

	// Close
	var closeWg sync.WaitGroup
	closeWg.Add(1)
	mockConnPool.EXPECT().Close().Do(func() {
		closeWg.Done()
	})
	queue.Close()
	closeWg.Wait()
}

func testWriteOp(
	namespace string,
	id string,
	value float64,
	timestamp int64,
	timeType rpc.TimeType,
	completionFn completionFn,
) *writeOp {
	w := &writeOp{}
	w.reset()
	w.request.NameSpace = namespace
	w.idDatapoint.ID = id
	w.datapoint = rpc.Datapoint{
		Value:         value,
		Timestamp:     timestamp,
		TimestampType: timeType,
	}
	w.completionFn = completionFn
	return w
}
