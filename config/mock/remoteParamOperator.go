// Code generated by MockGen. DO NOT EDIT.
// Source: ./remoteParam/base.go

// Package mock_remoteParam is a generated GoMock package.
package mock_remoteParam

import (
	reflect "reflect"

	remoteParam "github.com/ChainSafe/sygma-fee-oracle/remoteParam"
	gomock "github.com/golang/mock/gomock"
)

// MockRemoteParamOperator is a mock of RemoteParamOperator interface.
type MockRemoteParamOperator struct {
	ctrl     *gomock.Controller
	recorder *MockRemoteParamOperatorMockRecorder
}

// MockRemoteParamOperatorMockRecorder is the mock recorder for MockRemoteParamOperator.
type MockRemoteParamOperatorMockRecorder struct {
	mock *MockRemoteParamOperator
}

// NewMockRemoteParamOperator creates a new mock instance.
func NewMockRemoteParamOperator(ctrl *gomock.Controller) *MockRemoteParamOperator {
	mock := &MockRemoteParamOperator{ctrl: ctrl}
	mock.recorder = &MockRemoteParamOperatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRemoteParamOperator) EXPECT() *MockRemoteParamOperatorMockRecorder {
	return m.recorder
}

// LoadParameter mocks base method.
func (m *MockRemoteParamOperator) LoadParameter(paramName string) (*remoteParam.RemoteParamResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadParameter", paramName)
	ret0, _ := ret[0].(*remoteParam.RemoteParamResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoadParameter indicates an expected call of LoadParameter.
func (mr *MockRemoteParamOperatorMockRecorder) LoadParameter(paramName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadParameter", reflect.TypeOf((*MockRemoteParamOperator)(nil).LoadParameter), paramName)
}
