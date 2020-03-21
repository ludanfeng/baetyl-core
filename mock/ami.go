// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/baetyl/baetyl-core/ami (interfaces: AMI)

// Package mock is a generated GoMock package.
package mock

import (
	v1 "github.com/baetyl/baetyl-go/spec/v1"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockAMI is a mock of AMI interface
type MockAMI struct {
	ctrl     *gomock.Controller
	recorder *MockAMIMockRecorder
}

// MockAMIMockRecorder is the mock recorder for MockAMI
type MockAMIMockRecorder struct {
	mock *MockAMI
}

// NewMockAMI creates a new mock instance
func NewMockAMI(ctrl *gomock.Controller) *MockAMI {
	mock := &MockAMI{ctrl: ctrl}
	mock.recorder = &MockAMIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAMI) EXPECT() *MockAMIMockRecorder {
	return m.recorder
}

// ApplyApplications mocks base method
func (m *MockAMI) ApplyApplications(arg0 v1.Desire) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApplyApplications", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// ApplyApplications indicates an expected call of ApplyApplications
func (mr *MockAMIMockRecorder) ApplyApplications(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplyApplications", reflect.TypeOf((*MockAMI)(nil).ApplyApplications), arg0)
}

// CollectInfo mocks base method
func (m *MockAMI) CollectInfo() (v1.Report, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CollectInfo")
	ret0, _ := ret[0].(v1.Report)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CollectInfo indicates an expected call of CollectInfo
func (mr *MockAMIMockRecorder) CollectInfo() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CollectInfo", reflect.TypeOf((*MockAMI)(nil).CollectInfo))
}