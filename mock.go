// Code generated by MockGen. DO NOT EDIT.
// Source: contracts.go

// Package rabbitmq is a generated GoMock package.
package rabbitmq

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	amqp "github.com/streadway/amqp"
	proto "google.golang.org/protobuf/proto"
)

// MockIProducer is a mock of IProducer interface.
type MockIProducer struct {
	ctrl     *gomock.Controller
	recorder *MockIProducerMockRecorder
}

// MockIProducerMockRecorder is the mock recorder for MockIProducer.
type MockIProducerMockRecorder struct {
	mock *MockIProducer
}

// NewMockIProducer creates a new mock instance.
func NewMockIProducer(ctrl *gomock.Controller) *MockIProducer {
	mock := &MockIProducer{ctrl: ctrl}
	mock.recorder = &MockIProducerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIProducer) EXPECT() *MockIProducerMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockIProducer) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockIProducerMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockIProducer)(nil).Close))
}

// ProduceJson mocks base method.
func (m *MockIProducer) ProduceJson(ctx context.Context, routerKey string, data interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProduceJson", ctx, routerKey, data)
	ret0, _ := ret[0].(error)
	return ret0
}

// ProduceJson indicates an expected call of ProduceJson.
func (mr *MockIProducerMockRecorder) ProduceJson(ctx, routerKey, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProduceJson", reflect.TypeOf((*MockIProducer)(nil).ProduceJson), ctx, routerKey, data)
}

// ProduceJsonWithoutLogs mocks base method.
func (m *MockIProducer) ProduceJsonWithoutLogs(ctx context.Context, routerKey string, data interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProduceJsonWithoutLogs", ctx, routerKey, data)
	ret0, _ := ret[0].(error)
	return ret0
}

// ProduceJsonWithoutLogs indicates an expected call of ProduceJsonWithoutLogs.
func (mr *MockIProducerMockRecorder) ProduceJsonWithoutLogs(ctx, routerKey, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProduceJsonWithoutLogs", reflect.TypeOf((*MockIProducer)(nil).ProduceJsonWithoutLogs), ctx, routerKey, data)
}

// ProduceProto mocks base method.
func (m *MockIProducer) ProduceProto(ctx context.Context, routerKey string, data proto.Message) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProduceProto", ctx, routerKey, data)
	ret0, _ := ret[0].(error)
	return ret0
}

// ProduceProto indicates an expected call of ProduceProto.
func (mr *MockIProducerMockRecorder) ProduceProto(ctx, routerKey, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProduceProto", reflect.TypeOf((*MockIProducer)(nil).ProduceProto), ctx, routerKey, data)
}

// MockIConsumer is a mock of IConsumer interface.
type MockIConsumer struct {
	ctrl     *gomock.Controller
	recorder *MockIConsumerMockRecorder
}

// MockIConsumerMockRecorder is the mock recorder for MockIConsumer.
type MockIConsumerMockRecorder struct {
	mock *MockIConsumer
}

// NewMockIConsumer creates a new mock instance.
func NewMockIConsumer(ctrl *gomock.Controller) *MockIConsumer {
	mock := &MockIConsumer{ctrl: ctrl}
	mock.recorder = &MockIConsumerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIConsumer) EXPECT() *MockIConsumerMockRecorder {
	return m.recorder
}

// Consume mocks base method.
func (m *MockIConsumer) Consume(handler func(context.Context, amqp.Delivery) bool) (*Channel, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Consume", handler)
	ret0, _ := ret[0].(*Channel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Consume indicates an expected call of Consume.
func (mr *MockIConsumerMockRecorder) Consume(handler interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Consume", reflect.TypeOf((*MockIConsumer)(nil).Consume), handler)
}
