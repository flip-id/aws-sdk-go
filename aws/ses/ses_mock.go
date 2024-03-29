// Code generated by MockGen. DO NOT EDIT.
// Source: ses.go

// Package ses is a generated GoMock package.
package ses

import (
	context "context"
	reflect "reflect"

	ses "github.com/aws/aws-sdk-go-v2/service/ses"
	gomock "github.com/golang/mock/gomock"
)

// MockSESServiceInterface is a mock of SESServiceInterface interface.
type MockSESServiceInterface struct {
	ctrl     *gomock.Controller
	recorder *MockSESServiceInterfaceMockRecorder
}

// MockSESServiceInterfaceMockRecorder is the mock recorder for MockSESServiceInterface.
type MockSESServiceInterfaceMockRecorder struct {
	mock *MockSESServiceInterface
}

// NewMockSESServiceInterface creates a new mock instance.
func NewMockSESServiceInterface(ctrl *gomock.Controller) *MockSESServiceInterface {
	mock := &MockSESServiceInterface{ctrl: ctrl}
	mock.recorder = &MockSESServiceInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSESServiceInterface) EXPECT() *MockSESServiceInterfaceMockRecorder {
	return m.recorder
}

// SendEmail mocks base method.
func (m *MockSESServiceInterface) SendEmail(ctx context.Context, params *ses.SendEmailInput, optFns ...func(*ses.Options)) (*ses.SendEmailOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, params}
	for _, a := range optFns {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SendEmail", varargs...)
	ret0, _ := ret[0].(*ses.SendEmailOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SendEmail indicates an expected call of SendEmail.
func (mr *MockSESServiceInterfaceMockRecorder) SendEmail(ctx, params interface{}, optFns ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, params}, optFns...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendEmail", reflect.TypeOf((*MockSESServiceInterface)(nil).SendEmail), varargs...)
}

// SendRawEmail mocks base method.
func (m *MockSESServiceInterface) SendRawEmail(ctx context.Context, params *ses.SendRawEmailInput, optFns ...func(*ses.Options)) (*ses.SendRawEmailOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, params}
	for _, a := range optFns {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SendRawEmail", varargs...)
	ret0, _ := ret[0].(*ses.SendRawEmailOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SendRawEmail indicates an expected call of SendRawEmail.
func (mr *MockSESServiceInterfaceMockRecorder) SendRawEmail(ctx, params interface{}, optFns ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, params}, optFns...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendRawEmail", reflect.TypeOf((*MockSESServiceInterface)(nil).SendRawEmail), varargs...)
}
