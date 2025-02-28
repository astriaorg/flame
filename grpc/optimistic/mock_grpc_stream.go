package optimistic

import (
	"context"
	"google.golang.org/grpc/metadata"
	"io"
	"time"
)

type MockBidirectionalStreaming[K any, V any] struct {
	requestStream        []*K
	accumulatedResponses []*V
	requestCounter       uint64
}

func (ms *MockBidirectionalStreaming[K, V]) Recv() (*K, error) {
	// add a delay to make it look like an async stream
	time.Sleep(500 * time.Millisecond)
	if ms.requestCounter > uint64(len(ms.requestStream)-1) {
		// end the stream after all the packets have been sent
		return nil, io.EOF
	}

	req := ms.requestStream[ms.requestCounter]
	ms.requestCounter += 1

	return req, nil
}

func (ms *MockBidirectionalStreaming[K, V]) Send(res *V) error {
	ms.accumulatedResponses = append(ms.accumulatedResponses, res)
	return nil
}

func (ms *MockBidirectionalStreaming[K, V]) SetHeader(md metadata.MD) error {
	panic("implement SetHeader in the mock bi-directional grpc stream if used!")
}

func (ms *MockBidirectionalStreaming[K, V]) SendHeader(md metadata.MD) error {
	panic("implement SendHeader in the mock bi-directional grpc stream if used!")
}

func (ms *MockBidirectionalStreaming[K, V]) SetTrailer(md metadata.MD) {
	panic("implement SetTrailer in the mock bi-directional grpc stream if used!")
}

func (ms *MockBidirectionalStreaming[K, V]) Context() context.Context {
	return context.Background()
}

func (ms *MockBidirectionalStreaming[K, V]) SendMsg(m any) error {
	panic("implement SendMsg in the mock bi-directional grpc stream if used!")
}

func (ms *MockBidirectionalStreaming[K, V]) RecvMsg(m any) error {
	panic("implement RecvMsg in the mock bi-directional grpc stream if used!")
}

type MockServerSideStreaming[K any] struct {
	sentResponses []*K
}

func (ms *MockServerSideStreaming[K]) SendMsg(m any) error {
	//TODO implement me
	panic("implement SendMsg in the mock server-side grpc stream if used!")
}

func (ms *MockServerSideStreaming[K]) Send(res *K) error {
	ms.sentResponses = append(ms.sentResponses, res)
	return nil
}

func (ms *MockServerSideStreaming[K]) SetHeader(md metadata.MD) error {
	panic("implement SetHeader in the mock server-side grpc stream if used!")
}

func (ms *MockServerSideStreaming[K]) SendHeader(md metadata.MD) error {
	panic("implement SendHeader in the mock server-side grpc stream if used!")
}

func (ms *MockServerSideStreaming[K]) SetTrailer(md metadata.MD) {
	panic("implement SetTrailer in the mock server-side grpc stream if used!")
}

func (ms *MockServerSideStreaming[K]) Context() context.Context {
	return context.Background()
}

func (ms *MockServerSideStreaming[K]) RecvMsg(m any) error {
	panic("implement RecvMsg in the mock server-side grpc stream if used!")
}
