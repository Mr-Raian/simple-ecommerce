// Code generated by MockGen. DO NOT EDIT.
// Source: api/internal/handler (interfaces: CardPaymentProccesor)

// Package mocks is a generated GoMock package.
package mocks

import (
	data "api/internal/data"
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockCardPaymentProccesor is a mock of CardPaymentProccesor interface.
type MockCardPaymentProccesor struct {
	ctrl     *gomock.Controller
	recorder *MockCardPaymentProccesorMockRecorder
}

// MockCardPaymentProccesorMockRecorder is the mock recorder for MockCardPaymentProccesor.
type MockCardPaymentProccesorMockRecorder struct {
	mock *MockCardPaymentProccesor
}

// NewMockCardPaymentProccesor creates a new mock instance.
func NewMockCardPaymentProccesor(ctrl *gomock.Controller) *MockCardPaymentProccesor {
	mock := &MockCardPaymentProccesor{ctrl: ctrl}
	mock.recorder = &MockCardPaymentProccesorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCardPaymentProccesor) EXPECT() *MockCardPaymentProccesorMockRecorder {
	return m.recorder
}

// CreateCheckout mocks base method.
func (m *MockCardPaymentProccesor) CreateCheckout(arg0 context.Context, arg1 uint) (data.Checkout, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCheckout", arg0, arg1)
	ret0, _ := ret[0].(data.Checkout)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCheckout indicates an expected call of CreateCheckout.
func (mr *MockCardPaymentProccesorMockRecorder) CreateCheckout(arg0, arg1 interface{}) *CardPaymentProccesorCreateCheckoutCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCheckout", reflect.TypeOf((*MockCardPaymentProccesor)(nil).CreateCheckout), arg0, arg1)
	return &CardPaymentProccesorCreateCheckoutCall{Call: call}
}

// CardPaymentProccesorCreateCheckoutCall wrap *gomock.Call
type CardPaymentProccesorCreateCheckoutCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *CardPaymentProccesorCreateCheckoutCall) Return(arg0 data.Checkout, arg1 error) *CardPaymentProccesorCreateCheckoutCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *CardPaymentProccesorCreateCheckoutCall) Do(f func(context.Context, uint) (data.Checkout, error)) *CardPaymentProccesorCreateCheckoutCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *CardPaymentProccesorCreateCheckoutCall) DoAndReturn(f func(context.Context, uint) (data.Checkout, error)) *CardPaymentProccesorCreateCheckoutCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ParseWebhook mocks base method.
func (m *MockCardPaymentProccesor) ParseWebhook(arg0 []byte, arg1 string) (string, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParseWebhook", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ParseWebhook indicates an expected call of ParseWebhook.
func (mr *MockCardPaymentProccesorMockRecorder) ParseWebhook(arg0, arg1 interface{}) *CardPaymentProccesorParseWebhookCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseWebhook", reflect.TypeOf((*MockCardPaymentProccesor)(nil).ParseWebhook), arg0, arg1)
	return &CardPaymentProccesorParseWebhookCall{Call: call}
}

// CardPaymentProccesorParseWebhookCall wrap *gomock.Call
type CardPaymentProccesorParseWebhookCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *CardPaymentProccesorParseWebhookCall) Return(arg0, arg1 string, arg2 error) *CardPaymentProccesorParseWebhookCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *CardPaymentProccesorParseWebhookCall) Do(f func([]byte, string) (string, string, error)) *CardPaymentProccesorParseWebhookCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *CardPaymentProccesorParseWebhookCall) DoAndReturn(f func([]byte, string) (string, string, error)) *CardPaymentProccesorParseWebhookCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
