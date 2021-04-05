// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/openservicemesh/osm/pkg/ingress (interfaces: Monitor)

// Package ingress is a generated GoMock package.
package ingress

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	service "github.com/openservicemesh/osm/pkg/service"
	v1 "k8s.io/api/networking/v1"
	v1beta1 "k8s.io/api/networking/v1beta1"
)

// MockMonitor is a mock of Monitor interface
type MockMonitor struct {
	ctrl     *gomock.Controller
	recorder *MockMonitorMockRecorder
}

// MockMonitorMockRecorder is the mock recorder for MockMonitor
type MockMonitorMockRecorder struct {
	mock *MockMonitor
}

// NewMockMonitor creates a new mock instance
func NewMockMonitor(ctrl *gomock.Controller) *MockMonitor {
	mock := &MockMonitor{ctrl: ctrl}
	mock.recorder = &MockMonitorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMonitor) EXPECT() *MockMonitorMockRecorder {
	return m.recorder
}

// GetAPIVersion mocks base method
func (m *MockMonitor) GetAPIVersion() APIVersion {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAPIVersion")
	ret0, _ := ret[0].(APIVersion)
	return ret0
}

// GetAPIVersion indicates an expected call of GetAPIVersion
func (mr *MockMonitorMockRecorder) GetAPIVersion() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAPIVersion", reflect.TypeOf((*MockMonitor)(nil).GetAPIVersion))
}

// GetIngressNetworkingV1 mocks base method
func (m *MockMonitor) GetIngressNetworkingV1(arg0 service.MeshService) ([]*v1.Ingress, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIngressNetworkingV1", arg0)
	ret0, _ := ret[0].([]*v1.Ingress)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIngressNetworkingV1 indicates an expected call of GetIngressNetworkingV1
func (mr *MockMonitorMockRecorder) GetIngressNetworkingV1(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIngressNetworkingV1", reflect.TypeOf((*MockMonitor)(nil).GetIngressNetworkingV1), arg0)
}

// GetIngressNetworkingV1beta1 mocks base method
func (m *MockMonitor) GetIngressNetworkingV1beta1(arg0 service.MeshService) ([]*v1beta1.Ingress, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIngressNetworkingV1beta1", arg0)
	ret0, _ := ret[0].([]*v1beta1.Ingress)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIngressNetworkingV1beta1 indicates an expected call of GetIngressNetworkingV1beta1
func (mr *MockMonitorMockRecorder) GetIngressNetworkingV1beta1(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIngressNetworkingV1beta1", reflect.TypeOf((*MockMonitor)(nil).GetIngressNetworkingV1beta1), arg0)
}
