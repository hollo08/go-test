// Code generated by MockGen. DO NOT EDIT.
// Source: myfunc.go

// Package mock_testmock is a generated GoMock package.
package testmock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockMyFunc is a mock of MyFunc interface.
type MockMyFunc struct {
	ctrl     *gomock.Controller
	recorder *MockMyFuncMockRecorder
}

// MockMyFuncMockRecorder is the mock recorder for MockMyFunc.
type MockMyFuncMockRecorder struct {
	mock *MockMyFunc
}

// NewMockMyFunc creates a new mock instance.
func NewMockMyFunc(ctrl *gomock.Controller) *MockMyFunc {
	mock := &MockMyFunc{ctrl: ctrl}
	mock.recorder = &MockMyFuncMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMyFunc) EXPECT() *MockMyFuncMockRecorder {
	return m.recorder
}

// GetInfo mocks base method.
func (m *MockMyFunc) GetInfo() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInfo")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetInfo indicates an expected call of GetInfo.
func (mr *MockMyFuncMockRecorder) GetInfo() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInfo", reflect.TypeOf((*MockMyFunc)(nil).GetInfo))
}
