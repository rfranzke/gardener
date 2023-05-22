// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/gardener/gardener/extensions/pkg/webhook/controlplane/genericmutator (interfaces: Ensurer)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	semver "github.com/Masterminds/semver"
	unit "github.com/coreos/go-systemd/v22/unit"
	v1alpha1 "github.com/gardener/etcd-druid/api/v1alpha1"
	context0 "github.com/gardener/gardener/extensions/pkg/webhook/context"
	v1alpha10 "github.com/gardener/gardener/pkg/apis/extensions/v1alpha1"
	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/api/apps/v1"
	v10 "k8s.io/api/core/v1"
	v11 "k8s.io/autoscaler/vertical-pod-autoscaler/pkg/apis/autoscaling.k8s.io/v1"
	v1beta1 "k8s.io/kubelet/config/v1beta1"
	reflect "reflect"
)

// MockEnsurer is a mock of Ensurer interface
type MockEnsurer struct {
	ctrl     *gomock.Controller
	recorder *MockEnsurerMockRecorder
}

// MockEnsurerMockRecorder is the mock recorder for MockEnsurer
type MockEnsurerMockRecorder struct {
	mock *MockEnsurer
}

// NewMockEnsurer creates a new mock instance
func NewMockEnsurer(ctrl *gomock.Controller) *MockEnsurer {
	mock := &MockEnsurer{ctrl: ctrl}
	mock.recorder = &MockEnsurerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockEnsurer) EXPECT() *MockEnsurerMockRecorder {
	return m.recorder
}

// EnsureAdditionalFiles mocks base method
func (m *MockEnsurer) EnsureAdditionalFiles(arg0 context.Context, arg1 context0.GardenContext, arg2, arg3 *[]v1alpha10.File) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnsureAdditionalFiles", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnsureAdditionalFiles indicates an expected call of EnsureAdditionalFiles
func (mr *MockEnsurerMockRecorder) EnsureAdditionalFiles(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureAdditionalFiles", reflect.TypeOf((*MockEnsurer)(nil).EnsureAdditionalFiles), arg0, arg1, arg2, arg3)
}

// EnsureAdditionalUnits mocks base method
func (m *MockEnsurer) EnsureAdditionalUnits(arg0 context.Context, arg1 context0.GardenContext, arg2, arg3 *[]v1alpha10.Unit) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnsureAdditionalUnits", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnsureAdditionalUnits indicates an expected call of EnsureAdditionalUnits
func (mr *MockEnsurerMockRecorder) EnsureAdditionalUnits(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureAdditionalUnits", reflect.TypeOf((*MockEnsurer)(nil).EnsureAdditionalUnits), arg0, arg1, arg2, arg3)
}

// EnsureClusterAutoscalerDeployment mocks base method
func (m *MockEnsurer) EnsureClusterAutoscalerDeployment(arg0 context.Context, arg1 context0.GardenContext, arg2, arg3 *v1.Deployment) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnsureClusterAutoscalerDeployment", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnsureClusterAutoscalerDeployment indicates an expected call of EnsureClusterAutoscalerDeployment
func (mr *MockEnsurerMockRecorder) EnsureClusterAutoscalerDeployment(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureClusterAutoscalerDeployment", reflect.TypeOf((*MockEnsurer)(nil).EnsureClusterAutoscalerDeployment), arg0, arg1, arg2, arg3)
}

// EnsureETCD mocks base method
func (m *MockEnsurer) EnsureETCD(arg0 context.Context, arg1 context0.GardenContext, arg2, arg3 *v1alpha1.Etcd) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnsureETCD", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnsureETCD indicates an expected call of EnsureETCD
func (mr *MockEnsurerMockRecorder) EnsureETCD(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureETCD", reflect.TypeOf((*MockEnsurer)(nil).EnsureETCD), arg0, arg1, arg2, arg3)
}

// EnsureKubeAPIServerDeployment mocks base method
func (m *MockEnsurer) EnsureKubeAPIServerDeployment(arg0 context.Context, arg1 context0.GardenContext, arg2, arg3 *v1.Deployment) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnsureKubeAPIServerDeployment", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnsureKubeAPIServerDeployment indicates an expected call of EnsureKubeAPIServerDeployment
func (mr *MockEnsurerMockRecorder) EnsureKubeAPIServerDeployment(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureKubeAPIServerDeployment", reflect.TypeOf((*MockEnsurer)(nil).EnsureKubeAPIServerDeployment), arg0, arg1, arg2, arg3)
}

// EnsureKubeAPIServerService mocks base method
func (m *MockEnsurer) EnsureKubeAPIServerService(arg0 context.Context, arg1 context0.GardenContext, arg2, arg3 *v10.Service) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnsureKubeAPIServerService", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnsureKubeAPIServerService indicates an expected call of EnsureKubeAPIServerService
func (mr *MockEnsurerMockRecorder) EnsureKubeAPIServerService(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureKubeAPIServerService", reflect.TypeOf((*MockEnsurer)(nil).EnsureKubeAPIServerService), arg0, arg1, arg2, arg3)
}

// EnsureKubeControllerManagerDeployment mocks base method
func (m *MockEnsurer) EnsureKubeControllerManagerDeployment(arg0 context.Context, arg1 context0.GardenContext, arg2, arg3 *v1.Deployment) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnsureKubeControllerManagerDeployment", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnsureKubeControllerManagerDeployment indicates an expected call of EnsureKubeControllerManagerDeployment
func (mr *MockEnsurerMockRecorder) EnsureKubeControllerManagerDeployment(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureKubeControllerManagerDeployment", reflect.TypeOf((*MockEnsurer)(nil).EnsureKubeControllerManagerDeployment), arg0, arg1, arg2, arg3)
}

// EnsureKubeSchedulerDeployment mocks base method
func (m *MockEnsurer) EnsureKubeSchedulerDeployment(arg0 context.Context, arg1 context0.GardenContext, arg2, arg3 *v1.Deployment) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnsureKubeSchedulerDeployment", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnsureKubeSchedulerDeployment indicates an expected call of EnsureKubeSchedulerDeployment
func (mr *MockEnsurerMockRecorder) EnsureKubeSchedulerDeployment(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureKubeSchedulerDeployment", reflect.TypeOf((*MockEnsurer)(nil).EnsureKubeSchedulerDeployment), arg0, arg1, arg2, arg3)
}

// EnsureKubeletCloudProviderConfig mocks base method
func (m *MockEnsurer) EnsureKubeletCloudProviderConfig(arg0 context.Context, arg1 context0.GardenContext, arg2 *semver.Version, arg3 *string, arg4 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnsureKubeletCloudProviderConfig", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnsureKubeletCloudProviderConfig indicates an expected call of EnsureKubeletCloudProviderConfig
func (mr *MockEnsurerMockRecorder) EnsureKubeletCloudProviderConfig(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureKubeletCloudProviderConfig", reflect.TypeOf((*MockEnsurer)(nil).EnsureKubeletCloudProviderConfig), arg0, arg1, arg2, arg3, arg4)
}

// EnsureKubeletConfiguration mocks base method
func (m *MockEnsurer) EnsureKubeletConfiguration(arg0 context.Context, arg1 context0.GardenContext, arg2 *semver.Version, arg3, arg4 *v1beta1.KubeletConfiguration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnsureKubeletConfiguration", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnsureKubeletConfiguration indicates an expected call of EnsureKubeletConfiguration
func (mr *MockEnsurerMockRecorder) EnsureKubeletConfiguration(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureKubeletConfiguration", reflect.TypeOf((*MockEnsurer)(nil).EnsureKubeletConfiguration), arg0, arg1, arg2, arg3, arg4)
}

// EnsureKubeletServiceUnitOptions mocks base method
func (m *MockEnsurer) EnsureKubeletServiceUnitOptions(arg0 context.Context, arg1 context0.GardenContext, arg2 *semver.Version, arg3, arg4 []*unit.UnitOption) ([]*unit.UnitOption, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnsureKubeletServiceUnitOptions", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].([]*unit.UnitOption)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EnsureKubeletServiceUnitOptions indicates an expected call of EnsureKubeletServiceUnitOptions
func (mr *MockEnsurerMockRecorder) EnsureKubeletServiceUnitOptions(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureKubeletServiceUnitOptions", reflect.TypeOf((*MockEnsurer)(nil).EnsureKubeletServiceUnitOptions), arg0, arg1, arg2, arg3, arg4)
}

// EnsureKubernetesGeneralConfiguration mocks base method
func (m *MockEnsurer) EnsureKubernetesGeneralConfiguration(arg0 context.Context, arg1 context0.GardenContext, arg2, arg3 *string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnsureKubernetesGeneralConfiguration", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnsureKubernetesGeneralConfiguration indicates an expected call of EnsureKubernetesGeneralConfiguration
func (mr *MockEnsurerMockRecorder) EnsureKubernetesGeneralConfiguration(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureKubernetesGeneralConfiguration", reflect.TypeOf((*MockEnsurer)(nil).EnsureKubernetesGeneralConfiguration), arg0, arg1, arg2, arg3)
}

// EnsureMachineControllerManagerDeployment mocks base method
func (m *MockEnsurer) EnsureMachineControllerManagerDeployment(arg0 context.Context, arg1 context0.GardenContext, arg2, arg3 *v1.Deployment) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnsureMachineControllerManagerDeployment", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnsureMachineControllerManagerDeployment indicates an expected call of EnsureMachineControllerManagerDeployment
func (mr *MockEnsurerMockRecorder) EnsureMachineControllerManagerDeployment(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureMachineControllerManagerDeployment", reflect.TypeOf((*MockEnsurer)(nil).EnsureMachineControllerManagerDeployment), arg0, arg1, arg2, arg3)
}

// EnsureMachineControllerManagerVPA mocks base method
func (m *MockEnsurer) EnsureMachineControllerManagerVPA(arg0 context.Context, arg1 context0.GardenContext, arg2, arg3 *v11.VerticalPodAutoscaler) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnsureMachineControllerManagerVPA", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnsureMachineControllerManagerVPA indicates an expected call of EnsureMachineControllerManagerVPA
func (mr *MockEnsurerMockRecorder) EnsureMachineControllerManagerVPA(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureMachineControllerManagerVPA", reflect.TypeOf((*MockEnsurer)(nil).EnsureMachineControllerManagerVPA), arg0, arg1, arg2, arg3)
}

// EnsureVPNSeedServerDeployment mocks base method
func (m *MockEnsurer) EnsureVPNSeedServerDeployment(arg0 context.Context, arg1 context0.GardenContext, arg2, arg3 *v1.Deployment) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnsureVPNSeedServerDeployment", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnsureVPNSeedServerDeployment indicates an expected call of EnsureVPNSeedServerDeployment
func (mr *MockEnsurerMockRecorder) EnsureVPNSeedServerDeployment(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureVPNSeedServerDeployment", reflect.TypeOf((*MockEnsurer)(nil).EnsureVPNSeedServerDeployment), arg0, arg1, arg2, arg3)
}

// ShouldProvisionKubeletCloudProviderConfig mocks base method
func (m *MockEnsurer) ShouldProvisionKubeletCloudProviderConfig(arg0 context.Context, arg1 context0.GardenContext, arg2 *semver.Version) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ShouldProvisionKubeletCloudProviderConfig", arg0, arg1, arg2)
	ret0, _ := ret[0].(bool)
	return ret0
}

// ShouldProvisionKubeletCloudProviderConfig indicates an expected call of ShouldProvisionKubeletCloudProviderConfig
func (mr *MockEnsurerMockRecorder) ShouldProvisionKubeletCloudProviderConfig(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShouldProvisionKubeletCloudProviderConfig", reflect.TypeOf((*MockEnsurer)(nil).ShouldProvisionKubeletCloudProviderConfig), arg0, arg1, arg2)
}
