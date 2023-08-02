// Code generated by MockGen. DO NOT EDIT.
// Source: info.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	central "github.com/stackrox/rox/generated/internalapi/central"
	gomock "go.uber.org/mock/gomock"
)

// MockInfo is a mock of Info interface.
type MockInfo struct {
	ctrl     *gomock.Controller
	recorder *MockInfoMockRecorder
}

// MockInfoMockRecorder is the mock recorder for MockInfo.
type MockInfoMockRecorder struct {
	mock *MockInfo
}

// NewMockInfo creates a new mock instance.
func NewMockInfo(ctrl *gomock.Controller) *MockInfo {
	mock := &MockInfo{ctrl: ctrl}
	mock.recorder = &MockInfoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInfo) EXPECT() *MockInfoMockRecorder {
	return m.recorder
}

// DeleteClusterMetrics mocks base method.
func (m *MockInfo) DeleteClusterMetrics(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DeleteClusterMetrics", arg0)
}

// DeleteClusterMetrics indicates an expected call of DeleteClusterMetrics.
func (mr *MockInfoMockRecorder) DeleteClusterMetrics(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteClusterMetrics", reflect.TypeOf((*MockInfo)(nil).DeleteClusterMetrics), arg0)
}

// SetClusterMetrics mocks base method.
func (m *MockInfo) SetClusterMetrics(arg0 string, arg1 *central.ClusterMetrics) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetClusterMetrics", arg0, arg1)
}

// SetClusterMetrics indicates an expected call of SetClusterMetrics.
func (mr *MockInfoMockRecorder) SetClusterMetrics(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetClusterMetrics", reflect.TypeOf((*MockInfo)(nil).SetClusterMetrics), arg0, arg1)
}

// Start mocks base method.
func (m *MockInfo) Start() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Start")
}

// Start indicates an expected call of Start.
func (mr *MockInfoMockRecorder) Start() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockInfo)(nil).Start))
}