// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/openservicemesh/osm/pkg/catalog (interfaces: MeshCataloger)

// Package catalog is a generated GoMock package.
package catalog

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	endpoint "github.com/openservicemesh/osm/pkg/endpoint"
	identity "github.com/openservicemesh/osm/pkg/identity"
	kubernetes "github.com/openservicemesh/osm/pkg/kubernetes"
	service "github.com/openservicemesh/osm/pkg/service"
	trafficpolicy "github.com/openservicemesh/osm/pkg/trafficpolicy"
)

// MockMeshCataloger is a mock of MeshCataloger interface
type MockMeshCataloger struct {
	ctrl     *gomock.Controller
	recorder *MockMeshCatalogerMockRecorder
}

// MockMeshCatalogerMockRecorder is the mock recorder for MockMeshCataloger
type MockMeshCatalogerMockRecorder struct {
	mock *MockMeshCataloger
}

// NewMockMeshCataloger creates a new mock instance
func NewMockMeshCataloger(ctrl *gomock.Controller) *MockMeshCataloger {
	mock := &MockMeshCataloger{ctrl: ctrl}
	mock.recorder = &MockMeshCatalogerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMeshCataloger) EXPECT() *MockMeshCatalogerMockRecorder {
	return m.recorder
}

// GetEgressTrafficPolicy mocks base method
func (m *MockMeshCataloger) GetEgressTrafficPolicy(arg0 identity.ServiceIdentity) (*trafficpolicy.EgressTrafficPolicy, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEgressTrafficPolicy", arg0)
	ret0, _ := ret[0].(*trafficpolicy.EgressTrafficPolicy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEgressTrafficPolicy indicates an expected call of GetEgressTrafficPolicy
func (mr *MockMeshCatalogerMockRecorder) GetEgressTrafficPolicy(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEgressTrafficPolicy", reflect.TypeOf((*MockMeshCataloger)(nil).GetEgressTrafficPolicy), arg0)
}

// GetIngressPoliciesForService mocks base method
func (m *MockMeshCataloger) GetIngressPoliciesForService(arg0 service.MeshService) ([]*trafficpolicy.InboundTrafficPolicy, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIngressPoliciesForService", arg0)
	ret0, _ := ret[0].([]*trafficpolicy.InboundTrafficPolicy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIngressPoliciesForService indicates an expected call of GetIngressPoliciesForService
func (mr *MockMeshCatalogerMockRecorder) GetIngressPoliciesForService(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIngressPoliciesForService", reflect.TypeOf((*MockMeshCataloger)(nil).GetIngressPoliciesForService), arg0)
}

// GetKubecontroller mocks base method
func (m *MockMeshCataloger) GetKubecontroller() kubernetes.Controller {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetKubecontroller")
	ret0, _ := ret[0].(kubernetes.Controller)
	return ret0
}

// GetKubecontroller indicates an expected call of GetKubecontroller
func (mr *MockMeshCatalogerMockRecorder) GetKubecontroller() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetKubecontroller", reflect.TypeOf((*MockMeshCataloger)(nil).GetKubecontroller))
}

// GetPortToProtocolMappingForService mocks base method
func (m *MockMeshCataloger) GetPortToProtocolMappingForService(arg0 service.MeshService) (map[uint32]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPortToProtocolMappingForService", arg0)
	ret0, _ := ret[0].(map[uint32]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPortToProtocolMappingForService indicates an expected call of GetPortToProtocolMappingForService
func (mr *MockMeshCatalogerMockRecorder) GetPortToProtocolMappingForService(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPortToProtocolMappingForService", reflect.TypeOf((*MockMeshCataloger)(nil).GetPortToProtocolMappingForService), arg0)
}

// GetResolvableServiceEndpoints mocks base method
func (m *MockMeshCataloger) GetResolvableServiceEndpoints(arg0 service.MeshService) ([]endpoint.Endpoint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetResolvableServiceEndpoints", arg0)
	ret0, _ := ret[0].([]endpoint.Endpoint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetResolvableServiceEndpoints indicates an expected call of GetResolvableServiceEndpoints
func (mr *MockMeshCatalogerMockRecorder) GetResolvableServiceEndpoints(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetResolvableServiceEndpoints", reflect.TypeOf((*MockMeshCataloger)(nil).GetResolvableServiceEndpoints), arg0)
}

// GetTargetPortToProtocolMappingForService mocks base method
func (m *MockMeshCataloger) GetTargetPortToProtocolMappingForService(arg0 service.MeshService) (map[uint32]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTargetPortToProtocolMappingForService", arg0)
	ret0, _ := ret[0].(map[uint32]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTargetPortToProtocolMappingForService indicates an expected call of GetTargetPortToProtocolMappingForService
func (mr *MockMeshCatalogerMockRecorder) GetTargetPortToProtocolMappingForService(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTargetPortToProtocolMappingForService", reflect.TypeOf((*MockMeshCataloger)(nil).GetTargetPortToProtocolMappingForService), arg0)
}

// GetWeightedClustersForUpstream mocks base method
func (m *MockMeshCataloger) GetWeightedClustersForUpstream(arg0 service.MeshService) []service.WeightedCluster {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWeightedClustersForUpstream", arg0)
	ret0, _ := ret[0].([]service.WeightedCluster)
	return ret0
}

// GetWeightedClustersForUpstream indicates an expected call of GetWeightedClustersForUpstream
func (mr *MockMeshCatalogerMockRecorder) GetWeightedClustersForUpstream(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWeightedClustersForUpstream", reflect.TypeOf((*MockMeshCataloger)(nil).GetWeightedClustersForUpstream), arg0)
}

// ListAllowedEndpointsForService mocks base method
func (m *MockMeshCataloger) ListAllowedEndpointsForService(arg0 identity.ServiceIdentity, arg1 service.MeshService) ([]endpoint.Endpoint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAllowedEndpointsForService", arg0, arg1)
	ret0, _ := ret[0].([]endpoint.Endpoint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAllowedEndpointsForService indicates an expected call of ListAllowedEndpointsForService
func (mr *MockMeshCatalogerMockRecorder) ListAllowedEndpointsForService(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAllowedEndpointsForService", reflect.TypeOf((*MockMeshCataloger)(nil).ListAllowedEndpointsForService), arg0, arg1)
}

// ListAllowedInboundServiceIdentities mocks base method
func (m *MockMeshCataloger) ListAllowedInboundServiceIdentities(arg0 identity.ServiceIdentity) ([]identity.ServiceIdentity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAllowedInboundServiceIdentities", arg0)
	ret0, _ := ret[0].([]identity.ServiceIdentity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAllowedInboundServiceIdentities indicates an expected call of ListAllowedInboundServiceIdentities
func (mr *MockMeshCatalogerMockRecorder) ListAllowedInboundServiceIdentities(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAllowedInboundServiceIdentities", reflect.TypeOf((*MockMeshCataloger)(nil).ListAllowedInboundServiceIdentities), arg0)
}

// ListAllowedOutboundServiceIdentities mocks base method
func (m *MockMeshCataloger) ListAllowedOutboundServiceIdentities(arg0 identity.ServiceIdentity) ([]identity.ServiceIdentity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAllowedOutboundServiceIdentities", arg0)
	ret0, _ := ret[0].([]identity.ServiceIdentity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAllowedOutboundServiceIdentities indicates an expected call of ListAllowedOutboundServiceIdentities
func (mr *MockMeshCatalogerMockRecorder) ListAllowedOutboundServiceIdentities(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAllowedOutboundServiceIdentities", reflect.TypeOf((*MockMeshCataloger)(nil).ListAllowedOutboundServiceIdentities), arg0)
}

// ListAllowedOutboundServicesForIdentity mocks base method
func (m *MockMeshCataloger) ListAllowedOutboundServicesForIdentity(arg0 identity.ServiceIdentity) []service.MeshService {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAllowedOutboundServicesForIdentity", arg0)
	ret0, _ := ret[0].([]service.MeshService)
	return ret0
}

// ListAllowedOutboundServicesForIdentity indicates an expected call of ListAllowedOutboundServicesForIdentity
func (mr *MockMeshCatalogerMockRecorder) ListAllowedOutboundServicesForIdentity(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAllowedOutboundServicesForIdentity", reflect.TypeOf((*MockMeshCataloger)(nil).ListAllowedOutboundServicesForIdentity), arg0)
}

// ListInboundTrafficPolicies mocks base method
func (m *MockMeshCataloger) ListInboundTrafficPolicies(arg0 identity.ServiceIdentity, arg1 []service.MeshService) []*trafficpolicy.InboundTrafficPolicy {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListInboundTrafficPolicies", arg0, arg1)
	ret0, _ := ret[0].([]*trafficpolicy.InboundTrafficPolicy)
	return ret0
}

// ListInboundTrafficPolicies indicates an expected call of ListInboundTrafficPolicies
func (mr *MockMeshCatalogerMockRecorder) ListInboundTrafficPolicies(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListInboundTrafficPolicies", reflect.TypeOf((*MockMeshCataloger)(nil).ListInboundTrafficPolicies), arg0, arg1)
}

// ListInboundTrafficTargetsWithRoutes mocks base method
func (m *MockMeshCataloger) ListInboundTrafficTargetsWithRoutes(arg0 identity.ServiceIdentity) ([]trafficpolicy.TrafficTargetWithRoutes, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListInboundTrafficTargetsWithRoutes", arg0)
	ret0, _ := ret[0].([]trafficpolicy.TrafficTargetWithRoutes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListInboundTrafficTargetsWithRoutes indicates an expected call of ListInboundTrafficTargetsWithRoutes
func (mr *MockMeshCatalogerMockRecorder) ListInboundTrafficTargetsWithRoutes(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListInboundTrafficTargetsWithRoutes", reflect.TypeOf((*MockMeshCataloger)(nil).ListInboundTrafficTargetsWithRoutes), arg0)
}

// ListMeshServicesForIdentity mocks base method
func (m *MockMeshCataloger) ListMeshServicesForIdentity(arg0 identity.ServiceIdentity) []service.MeshService {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListMeshServicesForIdentity", arg0)
	ret0, _ := ret[0].([]service.MeshService)
	return ret0
}

// ListMeshServicesForIdentity indicates an expected call of ListMeshServicesForIdentity
func (mr *MockMeshCatalogerMockRecorder) ListMeshServicesForIdentity(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListMeshServicesForIdentity", reflect.TypeOf((*MockMeshCataloger)(nil).ListMeshServicesForIdentity), arg0)
}

// ListOutboundTrafficPolicies mocks base method
func (m *MockMeshCataloger) ListOutboundTrafficPolicies(arg0 identity.ServiceIdentity) []*trafficpolicy.OutboundTrafficPolicy {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListOutboundTrafficPolicies", arg0)
	ret0, _ := ret[0].([]*trafficpolicy.OutboundTrafficPolicy)
	return ret0
}

// ListOutboundTrafficPolicies indicates an expected call of ListOutboundTrafficPolicies
func (mr *MockMeshCatalogerMockRecorder) ListOutboundTrafficPolicies(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListOutboundTrafficPolicies", reflect.TypeOf((*MockMeshCataloger)(nil).ListOutboundTrafficPolicies), arg0)
}

// ListServiceIdentitiesForService mocks base method
func (m *MockMeshCataloger) ListServiceIdentitiesForService(arg0 service.MeshService) ([]identity.ServiceIdentity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListServiceIdentitiesForService", arg0)
	ret0, _ := ret[0].([]identity.ServiceIdentity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListServiceIdentitiesForService indicates an expected call of ListServiceIdentitiesForService
func (mr *MockMeshCatalogerMockRecorder) ListServiceIdentitiesForService(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListServiceIdentitiesForService", reflect.TypeOf((*MockMeshCataloger)(nil).ListServiceIdentitiesForService), arg0)
}
