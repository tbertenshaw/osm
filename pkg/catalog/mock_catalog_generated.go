// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/openservicemesh/osm/pkg/catalog (interfaces: MeshCataloger)

// Package catalog is a generated GoMock package.
package catalog

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1alpha2 "github.com/openservicemesh/osm/pkg/apis/config/v1alpha2"
	v1alpha1 "github.com/openservicemesh/osm/pkg/apis/policy/v1alpha1"
	endpoint "github.com/openservicemesh/osm/pkg/endpoint"
	envoy "github.com/openservicemesh/osm/pkg/envoy"
	identity "github.com/openservicemesh/osm/pkg/identity"
	service "github.com/openservicemesh/osm/pkg/service"
	trafficpolicy "github.com/openservicemesh/osm/pkg/trafficpolicy"
	types "k8s.io/apimachinery/pkg/types"
)

// MockMeshCataloger is a mock of MeshCataloger interface.
type MockMeshCataloger struct {
	ctrl     *gomock.Controller
	recorder *MockMeshCatalogerMockRecorder
}

// MockMeshCatalogerMockRecorder is the mock recorder for MockMeshCataloger.
type MockMeshCatalogerMockRecorder struct {
	mock *MockMeshCataloger
}

// NewMockMeshCataloger creates a new mock instance.
func NewMockMeshCataloger(ctrl *gomock.Controller) *MockMeshCataloger {
	mock := &MockMeshCataloger{ctrl: ctrl}
	mock.recorder = &MockMeshCatalogerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMeshCataloger) EXPECT() *MockMeshCatalogerMockRecorder {
	return m.recorder
}

// GetEgressTrafficPolicy mocks base method.
func (m *MockMeshCataloger) GetEgressTrafficPolicy(arg0 identity.ServiceIdentity) (*trafficpolicy.EgressTrafficPolicy, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEgressTrafficPolicy", arg0)
	ret0, _ := ret[0].(*trafficpolicy.EgressTrafficPolicy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEgressTrafficPolicy indicates an expected call of GetEgressTrafficPolicy.
func (mr *MockMeshCatalogerMockRecorder) GetEgressTrafficPolicy(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEgressTrafficPolicy", reflect.TypeOf((*MockMeshCataloger)(nil).GetEgressTrafficPolicy), arg0)
}

// GetHostnamesForService mocks base method.
func (m *MockMeshCataloger) GetHostnamesForService(arg0 service.MeshService, arg1 bool) []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHostnamesForService", arg0, arg1)
	ret0, _ := ret[0].([]string)
	return ret0
}

// GetHostnamesForService indicates an expected call of GetHostnamesForService.
func (mr *MockMeshCatalogerMockRecorder) GetHostnamesForService(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHostnamesForService", reflect.TypeOf((*MockMeshCataloger)(nil).GetHostnamesForService), arg0, arg1)
}

// GetInboundMeshTrafficPolicy mocks base method.
func (m *MockMeshCataloger) GetInboundMeshTrafficPolicy(arg0 identity.ServiceIdentity, arg1 []service.MeshService) *trafficpolicy.InboundMeshTrafficPolicy {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInboundMeshTrafficPolicy", arg0, arg1)
	ret0, _ := ret[0].(*trafficpolicy.InboundMeshTrafficPolicy)
	return ret0
}

// GetInboundMeshTrafficPolicy indicates an expected call of GetInboundMeshTrafficPolicy.
func (mr *MockMeshCatalogerMockRecorder) GetInboundMeshTrafficPolicy(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInboundMeshTrafficPolicy", reflect.TypeOf((*MockMeshCataloger)(nil).GetInboundMeshTrafficPolicy), arg0, arg1)
}

// GetIngressBackendPolicyForService mocks base method.
func (m *MockMeshCataloger) GetIngressBackendPolicyForService(arg0 service.MeshService) *v1alpha1.IngressBackend {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIngressBackendPolicyForService", arg0)
	ret0, _ := ret[0].(*v1alpha1.IngressBackend)
	return ret0
}

// GetIngressBackendPolicyForService indicates an expected call of GetIngressBackendPolicyForService.
func (mr *MockMeshCatalogerMockRecorder) GetIngressBackendPolicyForService(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIngressBackendPolicyForService", reflect.TypeOf((*MockMeshCataloger)(nil).GetIngressBackendPolicyForService), arg0)
}

// GetIngressTrafficPolicies mocks base method.
func (m *MockMeshCataloger) GetIngressTrafficPolicies(arg0 []service.MeshService) []*trafficpolicy.IngressTrafficPolicy {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIngressTrafficPolicies", arg0)
	ret0, _ := ret[0].([]*trafficpolicy.IngressTrafficPolicy)
	return ret0
}

// GetIngressTrafficPolicies indicates an expected call of GetIngressTrafficPolicies.
func (mr *MockMeshCatalogerMockRecorder) GetIngressTrafficPolicies(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIngressTrafficPolicies", reflect.TypeOf((*MockMeshCataloger)(nil).GetIngressTrafficPolicies), arg0)
}

// GetIngressTrafficPolicy mocks base method.
func (m *MockMeshCataloger) GetIngressTrafficPolicy(arg0 service.MeshService) (*trafficpolicy.IngressTrafficPolicy, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIngressTrafficPolicy", arg0)
	ret0, _ := ret[0].(*trafficpolicy.IngressTrafficPolicy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIngressTrafficPolicy indicates an expected call of GetIngressTrafficPolicy.
func (mr *MockMeshCatalogerMockRecorder) GetIngressTrafficPolicy(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIngressTrafficPolicy", reflect.TypeOf((*MockMeshCataloger)(nil).GetIngressTrafficPolicy), arg0)
}

// GetMeshConfig mocks base method.
func (m *MockMeshCataloger) GetMeshConfig() v1alpha2.MeshConfig {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMeshConfig")
	ret0, _ := ret[0].(v1alpha2.MeshConfig)
	return ret0
}

// GetMeshConfig indicates an expected call of GetMeshConfig.
func (mr *MockMeshCatalogerMockRecorder) GetMeshConfig() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMeshConfig", reflect.TypeOf((*MockMeshCataloger)(nil).GetMeshConfig))
}

// GetOSMNamespace mocks base method.
func (m *MockMeshCataloger) GetOSMNamespace() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOSMNamespace")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetOSMNamespace indicates an expected call of GetOSMNamespace.
func (mr *MockMeshCatalogerMockRecorder) GetOSMNamespace() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOSMNamespace", reflect.TypeOf((*MockMeshCataloger)(nil).GetOSMNamespace))
}

// GetOutboundMeshTrafficPolicy mocks base method.
func (m *MockMeshCataloger) GetOutboundMeshTrafficPolicy(arg0 identity.ServiceIdentity) *trafficpolicy.OutboundMeshTrafficPolicy {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOutboundMeshTrafficPolicy", arg0)
	ret0, _ := ret[0].(*trafficpolicy.OutboundMeshTrafficPolicy)
	return ret0
}

// GetOutboundMeshTrafficPolicy indicates an expected call of GetOutboundMeshTrafficPolicy.
func (mr *MockMeshCatalogerMockRecorder) GetOutboundMeshTrafficPolicy(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOutboundMeshTrafficPolicy", reflect.TypeOf((*MockMeshCataloger)(nil).GetOutboundMeshTrafficPolicy), arg0)
}

// GetProxyStatsHeaders mocks base method.
func (m *MockMeshCataloger) GetProxyStatsHeaders(arg0 *envoy.Proxy) (map[string]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProxyStatsHeaders", arg0)
	ret0, _ := ret[0].(map[string]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProxyStatsHeaders indicates an expected call of GetProxyStatsHeaders.
func (mr *MockMeshCatalogerMockRecorder) GetProxyStatsHeaders(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProxyStatsHeaders", reflect.TypeOf((*MockMeshCataloger)(nil).GetProxyStatsHeaders), arg0)
}

// GetResolvableEndpointsForService mocks base method.
func (m *MockMeshCataloger) GetResolvableEndpointsForService(arg0 service.MeshService) []endpoint.Endpoint {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetResolvableEndpointsForService", arg0)
	ret0, _ := ret[0].([]endpoint.Endpoint)
	return ret0
}

// GetResolvableEndpointsForService indicates an expected call of GetResolvableEndpointsForService.
func (mr *MockMeshCatalogerMockRecorder) GetResolvableEndpointsForService(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetResolvableEndpointsForService", reflect.TypeOf((*MockMeshCataloger)(nil).GetResolvableEndpointsForService), arg0)
}

// GetServicesForServiceIdentity mocks base method.
func (m *MockMeshCataloger) GetServicesForServiceIdentity(arg0 identity.ServiceIdentity) []service.MeshService {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetServicesForServiceIdentity", arg0)
	ret0, _ := ret[0].([]service.MeshService)
	return ret0
}

// GetServicesForServiceIdentity indicates an expected call of GetServicesForServiceIdentity.
func (mr *MockMeshCatalogerMockRecorder) GetServicesForServiceIdentity(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetServicesForServiceIdentity", reflect.TypeOf((*MockMeshCataloger)(nil).GetServicesForServiceIdentity), arg0)
}

// GetTargetPortForServicePort mocks base method.
func (m *MockMeshCataloger) GetTargetPortForServicePort(arg0 types.NamespacedName, arg1 uint16) (uint16, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTargetPortForServicePort", arg0, arg1)
	ret0, _ := ret[0].(uint16)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTargetPortForServicePort indicates an expected call of GetTargetPortForServicePort.
func (mr *MockMeshCatalogerMockRecorder) GetTargetPortForServicePort(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTargetPortForServicePort", reflect.TypeOf((*MockMeshCataloger)(nil).GetTargetPortForServicePort), arg0, arg1)
}

// GetUpstreamTrafficSetting mocks base method.
func (m *MockMeshCataloger) GetUpstreamTrafficSetting(arg0 *types.NamespacedName) *v1alpha1.UpstreamTrafficSetting {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUpstreamTrafficSetting", arg0)
	ret0, _ := ret[0].(*v1alpha1.UpstreamTrafficSetting)
	return ret0
}

// GetUpstreamTrafficSetting indicates an expected call of GetUpstreamTrafficSetting.
func (mr *MockMeshCatalogerMockRecorder) GetUpstreamTrafficSetting(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUpstreamTrafficSetting", reflect.TypeOf((*MockMeshCataloger)(nil).GetUpstreamTrafficSetting), arg0)
}

// GetUpstreamTrafficSettingByHost mocks base method.
func (m *MockMeshCataloger) GetUpstreamTrafficSettingByHost(arg0 string) *v1alpha1.UpstreamTrafficSetting {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUpstreamTrafficSettingByHost", arg0)
	ret0, _ := ret[0].(*v1alpha1.UpstreamTrafficSetting)
	return ret0
}

// GetUpstreamTrafficSettingByHost indicates an expected call of GetUpstreamTrafficSettingByHost.
func (mr *MockMeshCatalogerMockRecorder) GetUpstreamTrafficSettingByHost(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUpstreamTrafficSettingByHost", reflect.TypeOf((*MockMeshCataloger)(nil).GetUpstreamTrafficSettingByHost), arg0)
}

// GetUpstreamTrafficSettingByNamespace mocks base method.
func (m *MockMeshCataloger) GetUpstreamTrafficSettingByNamespace(arg0 *types.NamespacedName) *v1alpha1.UpstreamTrafficSetting {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUpstreamTrafficSettingByNamespace", arg0)
	ret0, _ := ret[0].(*v1alpha1.UpstreamTrafficSetting)
	return ret0
}

// GetUpstreamTrafficSettingByNamespace indicates an expected call of GetUpstreamTrafficSettingByNamespace.
func (mr *MockMeshCatalogerMockRecorder) GetUpstreamTrafficSettingByNamespace(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUpstreamTrafficSettingByNamespace", reflect.TypeOf((*MockMeshCataloger)(nil).GetUpstreamTrafficSettingByNamespace), arg0)
}

// GetUpstreamTrafficSettingByService mocks base method.
func (m *MockMeshCataloger) GetUpstreamTrafficSettingByService(arg0 *service.MeshService) *v1alpha1.UpstreamTrafficSetting {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUpstreamTrafficSettingByService", arg0)
	ret0, _ := ret[0].(*v1alpha1.UpstreamTrafficSetting)
	return ret0
}

// GetUpstreamTrafficSettingByService indicates an expected call of GetUpstreamTrafficSettingByService.
func (mr *MockMeshCatalogerMockRecorder) GetUpstreamTrafficSettingByService(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUpstreamTrafficSettingByService", reflect.TypeOf((*MockMeshCataloger)(nil).GetUpstreamTrafficSettingByService), arg0)
}

// IsMetricsEnabled mocks base method.
func (m *MockMeshCataloger) IsMetricsEnabled(arg0 *envoy.Proxy) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsMetricsEnabled", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsMetricsEnabled indicates an expected call of IsMetricsEnabled.
func (mr *MockMeshCatalogerMockRecorder) IsMetricsEnabled(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsMetricsEnabled", reflect.TypeOf((*MockMeshCataloger)(nil).IsMetricsEnabled), arg0)
}

// ListAllowedUpstreamEndpointsForService mocks base method.
func (m *MockMeshCataloger) ListAllowedUpstreamEndpointsForService(arg0 identity.ServiceIdentity, arg1 service.MeshService) []endpoint.Endpoint {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAllowedUpstreamEndpointsForService", arg0, arg1)
	ret0, _ := ret[0].([]endpoint.Endpoint)
	return ret0
}

// ListAllowedUpstreamEndpointsForService indicates an expected call of ListAllowedUpstreamEndpointsForService.
func (mr *MockMeshCatalogerMockRecorder) ListAllowedUpstreamEndpointsForService(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAllowedUpstreamEndpointsForService", reflect.TypeOf((*MockMeshCataloger)(nil).ListAllowedUpstreamEndpointsForService), arg0, arg1)
}

// ListEgressPolicies mocks base method.
func (m *MockMeshCataloger) ListEgressPolicies() []*v1alpha1.Egress {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListEgressPolicies")
	ret0, _ := ret[0].([]*v1alpha1.Egress)
	return ret0
}

// ListEgressPolicies indicates an expected call of ListEgressPolicies.
func (mr *MockMeshCatalogerMockRecorder) ListEgressPolicies() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListEgressPolicies", reflect.TypeOf((*MockMeshCataloger)(nil).ListEgressPolicies))
}

// ListEgressPoliciesForServiceAccount mocks base method.
func (m *MockMeshCataloger) ListEgressPoliciesForServiceAccount(arg0 identity.K8sServiceAccount) []*v1alpha1.Egress {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListEgressPoliciesForServiceAccount", arg0)
	ret0, _ := ret[0].([]*v1alpha1.Egress)
	return ret0
}

// ListEgressPoliciesForServiceAccount indicates an expected call of ListEgressPoliciesForServiceAccount.
func (mr *MockMeshCatalogerMockRecorder) ListEgressPoliciesForServiceAccount(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListEgressPoliciesForServiceAccount", reflect.TypeOf((*MockMeshCataloger)(nil).ListEgressPoliciesForServiceAccount), arg0)
}

// ListEndpointsForIdentity mocks base method.
func (m *MockMeshCataloger) ListEndpointsForIdentity(arg0 identity.ServiceIdentity) []endpoint.Endpoint {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListEndpointsForIdentity", arg0)
	ret0, _ := ret[0].([]endpoint.Endpoint)
	return ret0
}

// ListEndpointsForIdentity indicates an expected call of ListEndpointsForIdentity.
func (mr *MockMeshCatalogerMockRecorder) ListEndpointsForIdentity(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListEndpointsForIdentity", reflect.TypeOf((*MockMeshCataloger)(nil).ListEndpointsForIdentity), arg0)
}

// ListEndpointsForService mocks base method.
func (m *MockMeshCataloger) ListEndpointsForService(arg0 service.MeshService) []endpoint.Endpoint {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListEndpointsForService", arg0)
	ret0, _ := ret[0].([]endpoint.Endpoint)
	return ret0
}

// ListEndpointsForService indicates an expected call of ListEndpointsForService.
func (mr *MockMeshCatalogerMockRecorder) ListEndpointsForService(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListEndpointsForService", reflect.TypeOf((*MockMeshCataloger)(nil).ListEndpointsForService), arg0)
}

// ListInboundServiceIdentities mocks base method.
func (m *MockMeshCataloger) ListInboundServiceIdentities(arg0 identity.ServiceIdentity) []identity.ServiceIdentity {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListInboundServiceIdentities", arg0)
	ret0, _ := ret[0].([]identity.ServiceIdentity)
	return ret0
}

// ListInboundServiceIdentities indicates an expected call of ListInboundServiceIdentities.
func (mr *MockMeshCatalogerMockRecorder) ListInboundServiceIdentities(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListInboundServiceIdentities", reflect.TypeOf((*MockMeshCataloger)(nil).ListInboundServiceIdentities), arg0)
}

// ListInboundTrafficTargetsWithRoutes mocks base method.
func (m *MockMeshCataloger) ListInboundTrafficTargetsWithRoutes(arg0 identity.ServiceIdentity) ([]trafficpolicy.TrafficTargetWithRoutes, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListInboundTrafficTargetsWithRoutes", arg0)
	ret0, _ := ret[0].([]trafficpolicy.TrafficTargetWithRoutes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListInboundTrafficTargetsWithRoutes indicates an expected call of ListInboundTrafficTargetsWithRoutes.
func (mr *MockMeshCatalogerMockRecorder) ListInboundTrafficTargetsWithRoutes(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListInboundTrafficTargetsWithRoutes", reflect.TypeOf((*MockMeshCataloger)(nil).ListInboundTrafficTargetsWithRoutes), arg0)
}

// ListIngressBackendPolicies mocks base method.
func (m *MockMeshCataloger) ListIngressBackendPolicies() []*v1alpha1.IngressBackend {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListIngressBackendPolicies")
	ret0, _ := ret[0].([]*v1alpha1.IngressBackend)
	return ret0
}

// ListIngressBackendPolicies indicates an expected call of ListIngressBackendPolicies.
func (mr *MockMeshCatalogerMockRecorder) ListIngressBackendPolicies() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListIngressBackendPolicies", reflect.TypeOf((*MockMeshCataloger)(nil).ListIngressBackendPolicies))
}

// ListOutboundServiceIdentities mocks base method.
func (m *MockMeshCataloger) ListOutboundServiceIdentities(arg0 identity.ServiceIdentity) []identity.ServiceIdentity {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListOutboundServiceIdentities", arg0)
	ret0, _ := ret[0].([]identity.ServiceIdentity)
	return ret0
}

// ListOutboundServiceIdentities indicates an expected call of ListOutboundServiceIdentities.
func (mr *MockMeshCatalogerMockRecorder) ListOutboundServiceIdentities(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListOutboundServiceIdentities", reflect.TypeOf((*MockMeshCataloger)(nil).ListOutboundServiceIdentities), arg0)
}

// ListOutboundServicesForIdentity mocks base method.
func (m *MockMeshCataloger) ListOutboundServicesForIdentity(arg0 identity.ServiceIdentity) []service.MeshService {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListOutboundServicesForIdentity", arg0)
	ret0, _ := ret[0].([]service.MeshService)
	return ret0
}

// ListOutboundServicesForIdentity indicates an expected call of ListOutboundServicesForIdentity.
func (mr *MockMeshCatalogerMockRecorder) ListOutboundServicesForIdentity(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListOutboundServicesForIdentity", reflect.TypeOf((*MockMeshCataloger)(nil).ListOutboundServicesForIdentity), arg0)
}

// ListRetryPolicies mocks base method.
func (m *MockMeshCataloger) ListRetryPolicies() []*v1alpha1.Retry {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListRetryPolicies")
	ret0, _ := ret[0].([]*v1alpha1.Retry)
	return ret0
}

// ListRetryPolicies indicates an expected call of ListRetryPolicies.
func (mr *MockMeshCatalogerMockRecorder) ListRetryPolicies() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRetryPolicies", reflect.TypeOf((*MockMeshCataloger)(nil).ListRetryPolicies))
}

// ListRetryPoliciesForServiceAccount mocks base method.
func (m *MockMeshCataloger) ListRetryPoliciesForServiceAccount(arg0 identity.K8sServiceAccount) []*v1alpha1.Retry {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListRetryPoliciesForServiceAccount", arg0)
	ret0, _ := ret[0].([]*v1alpha1.Retry)
	return ret0
}

// ListRetryPoliciesForServiceAccount indicates an expected call of ListRetryPoliciesForServiceAccount.
func (mr *MockMeshCatalogerMockRecorder) ListRetryPoliciesForServiceAccount(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRetryPoliciesForServiceAccount", reflect.TypeOf((*MockMeshCataloger)(nil).ListRetryPoliciesForServiceAccount), arg0)
}

// ListServiceIdentitiesForService mocks base method.
func (m *MockMeshCataloger) ListServiceIdentitiesForService(arg0 service.MeshService) []identity.ServiceIdentity {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListServiceIdentitiesForService", arg0)
	ret0, _ := ret[0].([]identity.ServiceIdentity)
	return ret0
}

// ListServiceIdentitiesForService indicates an expected call of ListServiceIdentitiesForService.
func (mr *MockMeshCatalogerMockRecorder) ListServiceIdentitiesForService(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListServiceIdentitiesForService", reflect.TypeOf((*MockMeshCataloger)(nil).ListServiceIdentitiesForService), arg0)
}

// ListServices mocks base method.
func (m *MockMeshCataloger) ListServices() []service.MeshService {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListServices")
	ret0, _ := ret[0].([]service.MeshService)
	return ret0
}

// ListServices indicates an expected call of ListServices.
func (mr *MockMeshCatalogerMockRecorder) ListServices() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListServices", reflect.TypeOf((*MockMeshCataloger)(nil).ListServices))
}

// ListServicesForProxy mocks base method.
func (m *MockMeshCataloger) ListServicesForProxy(arg0 *envoy.Proxy) ([]service.MeshService, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListServicesForProxy", arg0)
	ret0, _ := ret[0].([]service.MeshService)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListServicesForProxy indicates an expected call of ListServicesForProxy.
func (mr *MockMeshCatalogerMockRecorder) ListServicesForProxy(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListServicesForProxy", reflect.TypeOf((*MockMeshCataloger)(nil).ListServicesForProxy), arg0)
}

// ListUpstreamTrafficSettings mocks base method.
func (m *MockMeshCataloger) ListUpstreamTrafficSettings() []*v1alpha1.UpstreamTrafficSetting {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListUpstreamTrafficSettings")
	ret0, _ := ret[0].([]*v1alpha1.UpstreamTrafficSetting)
	return ret0
}

// ListUpstreamTrafficSettings indicates an expected call of ListUpstreamTrafficSettings.
func (mr *MockMeshCatalogerMockRecorder) ListUpstreamTrafficSettings() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUpstreamTrafficSettings", reflect.TypeOf((*MockMeshCataloger)(nil).ListUpstreamTrafficSettings))
}

// UpdateIngressBackendStatus mocks base method.
func (m *MockMeshCataloger) UpdateIngressBackendStatus(arg0 *v1alpha1.IngressBackend) (*v1alpha1.IngressBackend, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateIngressBackendStatus", arg0)
	ret0, _ := ret[0].(*v1alpha1.IngressBackend)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateIngressBackendStatus indicates an expected call of UpdateIngressBackendStatus.
func (mr *MockMeshCatalogerMockRecorder) UpdateIngressBackendStatus(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateIngressBackendStatus", reflect.TypeOf((*MockMeshCataloger)(nil).UpdateIngressBackendStatus), arg0)
}

// UpdateUpstreamTrafficSettingStatus mocks base method.
func (m *MockMeshCataloger) UpdateUpstreamTrafficSettingStatus(arg0 *v1alpha1.UpstreamTrafficSetting) (*v1alpha1.UpstreamTrafficSetting, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUpstreamTrafficSettingStatus", arg0)
	ret0, _ := ret[0].(*v1alpha1.UpstreamTrafficSetting)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUpstreamTrafficSettingStatus indicates an expected call of UpdateUpstreamTrafficSettingStatus.
func (mr *MockMeshCatalogerMockRecorder) UpdateUpstreamTrafficSettingStatus(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUpstreamTrafficSettingStatus", reflect.TypeOf((*MockMeshCataloger)(nil).UpdateUpstreamTrafficSettingStatus), arg0)
}

// VerifyProxy mocks base method.
func (m *MockMeshCataloger) VerifyProxy(arg0 *envoy.Proxy) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyProxy", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// VerifyProxy indicates an expected call of VerifyProxy.
func (mr *MockMeshCatalogerMockRecorder) VerifyProxy(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyProxy", reflect.TypeOf((*MockMeshCataloger)(nil).VerifyProxy), arg0)
}
