// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/gardener/gardener/pkg/client/kubernetes (interfaces: Interface)

// Package kubernetes is a generated GoMock package.
package kubernetes

import (
	versioned "github.com/gardener/gardener/pkg/client/core/clientset/versioned"
	versioned0 "github.com/gardener/gardener/pkg/client/garden/clientset/versioned"
	kubernetes "github.com/gardener/gardener/pkg/client/kubernetes"
	versioned1 "github.com/gardener/gardener/pkg/client/machine/clientset/versioned"
	gomock "github.com/golang/mock/gomock"
	logrus "github.com/sirupsen/logrus"
	v1 "k8s.io/api/apps/v1"
	v10 "k8s.io/api/batch/v1"
	v11 "k8s.io/api/core/v1"
	v12 "k8s.io/api/rbac/v1"
	v1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	clientset "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	meta "k8s.io/apimachinery/pkg/api/meta"
	v13 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kubernetes0 "k8s.io/client-go/kubernetes"
	rest "k8s.io/client-go/rest"
	v1beta10 "k8s.io/kube-aggregator/pkg/apis/apiregistration/v1beta1"
	clientset0 "k8s.io/kube-aggregator/pkg/client/clientset_generated/clientset"
	reflect "reflect"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// MockInterface is a mock of Interface interface
type MockInterface struct {
	ctrl     *gomock.Controller
	recorder *MockInterfaceMockRecorder
}

// MockInterfaceMockRecorder is the mock recorder for MockInterface
type MockInterfaceMockRecorder struct {
	mock *MockInterface
}

// NewMockInterface creates a new mock instance
func NewMockInterface(ctrl *gomock.Controller) *MockInterface {
	mock := &MockInterface{ctrl: ctrl}
	mock.recorder = &MockInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockInterface) EXPECT() *MockInterfaceMockRecorder {
	return m.recorder
}

// APIExtension mocks base method
func (m *MockInterface) APIExtension() clientset.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "APIExtension")
	ret0, _ := ret[0].(clientset.Interface)
	return ret0
}

// APIExtension indicates an expected call of APIExtension
func (mr *MockInterfaceMockRecorder) APIExtension() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "APIExtension", reflect.TypeOf((*MockInterface)(nil).APIExtension))
}

// APIRegistration mocks base method
func (m *MockInterface) APIRegistration() clientset0.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "APIRegistration")
	ret0, _ := ret[0].(clientset0.Interface)
	return ret0
}

// APIRegistration indicates an expected call of APIRegistration
func (mr *MockInterfaceMockRecorder) APIRegistration() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "APIRegistration", reflect.TypeOf((*MockInterface)(nil).APIRegistration))
}

// Applier mocks base method
func (m *MockInterface) Applier() kubernetes.ApplierInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Applier")
	ret0, _ := ret[0].(kubernetes.ApplierInterface)
	return ret0
}

// Applier indicates an expected call of Applier
func (mr *MockInterfaceMockRecorder) Applier() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Applier", reflect.TypeOf((*MockInterface)(nil).Applier))
}

// CheckForwardPodPort mocks base method
func (m *MockInterface) CheckForwardPodPort(arg0, arg1 string, arg2, arg3 int) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckForwardPodPort", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckForwardPodPort indicates an expected call of CheckForwardPodPort
func (mr *MockInterfaceMockRecorder) CheckForwardPodPort(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckForwardPodPort", reflect.TypeOf((*MockInterface)(nil).CheckForwardPodPort), arg0, arg1, arg2, arg3)
}

// CheckResourceCleanup mocks base method
func (m *MockInterface) CheckResourceCleanup(arg0 *logrus.Entry, arg1 map[string]map[string]bool, arg2 string, arg3 []string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckResourceCleanup", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckResourceCleanup indicates an expected call of CheckResourceCleanup
func (mr *MockInterfaceMockRecorder) CheckResourceCleanup(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckResourceCleanup", reflect.TypeOf((*MockInterface)(nil).CheckResourceCleanup), arg0, arg1, arg2, arg3)
}

// CleanupAPIGroupResources mocks base method
func (m *MockInterface) CleanupAPIGroupResources(arg0 map[string]map[string]bool, arg1 string, arg2 []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CleanupAPIGroupResources", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// CleanupAPIGroupResources indicates an expected call of CleanupAPIGroupResources
func (mr *MockInterfaceMockRecorder) CleanupAPIGroupResources(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CleanupAPIGroupResources", reflect.TypeOf((*MockInterface)(nil).CleanupAPIGroupResources), arg0, arg1, arg2)
}

// CleanupResources mocks base method
func (m *MockInterface) CleanupResources(arg0 map[string]map[string]bool, arg1 map[string][]string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CleanupResources", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CleanupResources indicates an expected call of CleanupResources
func (mr *MockInterfaceMockRecorder) CleanupResources(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CleanupResources", reflect.TypeOf((*MockInterface)(nil).CleanupResources), arg0, arg1)
}

// Client mocks base method
func (m *MockInterface) Client() client.Client {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Client")
	ret0, _ := ret[0].(client.Client)
	return ret0
}

// Client indicates an expected call of Client
func (mr *MockInterfaceMockRecorder) Client() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Client", reflect.TypeOf((*MockInterface)(nil).Client))
}

// CreateConfigMap mocks base method
func (m *MockInterface) CreateConfigMap(arg0, arg1 string, arg2 map[string]string, arg3 bool) (*v11.ConfigMap, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateConfigMap", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*v11.ConfigMap)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateConfigMap indicates an expected call of CreateConfigMap
func (mr *MockInterfaceMockRecorder) CreateConfigMap(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateConfigMap", reflect.TypeOf((*MockInterface)(nil).CreateConfigMap), arg0, arg1, arg2, arg3)
}

// CreateNamespace mocks base method
func (m *MockInterface) CreateNamespace(arg0 *v11.Namespace, arg1 bool) (*v11.Namespace, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateNamespace", arg0, arg1)
	ret0, _ := ret[0].(*v11.Namespace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateNamespace indicates an expected call of CreateNamespace
func (mr *MockInterfaceMockRecorder) CreateNamespace(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateNamespace", reflect.TypeOf((*MockInterface)(nil).CreateNamespace), arg0, arg1)
}

// CreateSecret mocks base method
func (m *MockInterface) CreateSecret(arg0, arg1 string, arg2 v11.SecretType, arg3 map[string][]byte, arg4 bool) (*v11.Secret, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSecret", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(*v11.Secret)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSecret indicates an expected call of CreateSecret
func (mr *MockInterfaceMockRecorder) CreateSecret(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSecret", reflect.TypeOf((*MockInterface)(nil).CreateSecret), arg0, arg1, arg2, arg3, arg4)
}

// CreateSecretObject mocks base method
func (m *MockInterface) CreateSecretObject(arg0 *v11.Secret, arg1 bool) (*v11.Secret, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSecretObject", arg0, arg1)
	ret0, _ := ret[0].(*v11.Secret)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSecretObject indicates an expected call of CreateSecretObject
func (mr *MockInterfaceMockRecorder) CreateSecretObject(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSecretObject", reflect.TypeOf((*MockInterface)(nil).CreateSecretObject), arg0, arg1)
}

// DeleteAPIService mocks base method
func (m *MockInterface) DeleteAPIService(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAPIService", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAPIService indicates an expected call of DeleteAPIService
func (mr *MockInterfaceMockRecorder) DeleteAPIService(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAPIService", reflect.TypeOf((*MockInterface)(nil).DeleteAPIService), arg0)
}

// DeleteAPIServiceForcefully mocks base method
func (m *MockInterface) DeleteAPIServiceForcefully(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAPIServiceForcefully", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAPIServiceForcefully indicates an expected call of DeleteAPIServiceForcefully
func (mr *MockInterfaceMockRecorder) DeleteAPIServiceForcefully(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAPIServiceForcefully", reflect.TypeOf((*MockInterface)(nil).DeleteAPIServiceForcefully), arg0)
}

// DeleteCRDForcefully mocks base method
func (m *MockInterface) DeleteCRDForcefully(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCRDForcefully", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCRDForcefully indicates an expected call of DeleteCRDForcefully
func (mr *MockInterfaceMockRecorder) DeleteCRDForcefully(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCRDForcefully", reflect.TypeOf((*MockInterface)(nil).DeleteCRDForcefully), arg0)
}

// DeleteClusterRole mocks base method
func (m *MockInterface) DeleteClusterRole(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteClusterRole", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteClusterRole indicates an expected call of DeleteClusterRole
func (mr *MockInterfaceMockRecorder) DeleteClusterRole(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteClusterRole", reflect.TypeOf((*MockInterface)(nil).DeleteClusterRole), arg0)
}

// DeleteClusterRoleBinding mocks base method
func (m *MockInterface) DeleteClusterRoleBinding(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteClusterRoleBinding", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteClusterRoleBinding indicates an expected call of DeleteClusterRoleBinding
func (mr *MockInterfaceMockRecorder) DeleteClusterRoleBinding(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteClusterRoleBinding", reflect.TypeOf((*MockInterface)(nil).DeleteClusterRoleBinding), arg0)
}

// DeleteConfigMap mocks base method
func (m *MockInterface) DeleteConfigMap(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteConfigMap", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteConfigMap indicates an expected call of DeleteConfigMap
func (mr *MockInterfaceMockRecorder) DeleteConfigMap(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteConfigMap", reflect.TypeOf((*MockInterface)(nil).DeleteConfigMap), arg0, arg1)
}

// DeleteCronJob mocks base method
func (m *MockInterface) DeleteCronJob(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCronJob", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCronJob indicates an expected call of DeleteCronJob
func (mr *MockInterfaceMockRecorder) DeleteCronJob(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCronJob", reflect.TypeOf((*MockInterface)(nil).DeleteCronJob), arg0, arg1)
}

// DeleteDaemonSet mocks base method
func (m *MockInterface) DeleteDaemonSet(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteDaemonSet", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteDaemonSet indicates an expected call of DeleteDaemonSet
func (mr *MockInterfaceMockRecorder) DeleteDaemonSet(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteDaemonSet", reflect.TypeOf((*MockInterface)(nil).DeleteDaemonSet), arg0, arg1)
}

// DeleteDeployment mocks base method
func (m *MockInterface) DeleteDeployment(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteDeployment", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteDeployment indicates an expected call of DeleteDeployment
func (mr *MockInterfaceMockRecorder) DeleteDeployment(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteDeployment", reflect.TypeOf((*MockInterface)(nil).DeleteDeployment), arg0, arg1)
}

// DeleteHorizontalPodAutoscaler mocks base method
func (m *MockInterface) DeleteHorizontalPodAutoscaler(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteHorizontalPodAutoscaler", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteHorizontalPodAutoscaler indicates an expected call of DeleteHorizontalPodAutoscaler
func (mr *MockInterfaceMockRecorder) DeleteHorizontalPodAutoscaler(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteHorizontalPodAutoscaler", reflect.TypeOf((*MockInterface)(nil).DeleteHorizontalPodAutoscaler), arg0, arg1)
}

// DeleteIngress mocks base method
func (m *MockInterface) DeleteIngress(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteIngress", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteIngress indicates an expected call of DeleteIngress
func (mr *MockInterfaceMockRecorder) DeleteIngress(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteIngress", reflect.TypeOf((*MockInterface)(nil).DeleteIngress), arg0, arg1)
}

// DeleteJob mocks base method
func (m *MockInterface) DeleteJob(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteJob", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteJob indicates an expected call of DeleteJob
func (mr *MockInterfaceMockRecorder) DeleteJob(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteJob", reflect.TypeOf((*MockInterface)(nil).DeleteJob), arg0, arg1)
}

// DeleteNamespace mocks base method
func (m *MockInterface) DeleteNamespace(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteNamespace", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteNamespace indicates an expected call of DeleteNamespace
func (mr *MockInterfaceMockRecorder) DeleteNamespace(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteNamespace", reflect.TypeOf((*MockInterface)(nil).DeleteNamespace), arg0)
}

// DeleteNetworkPolicy mocks base method
func (m *MockInterface) DeleteNetworkPolicy(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteNetworkPolicy", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteNetworkPolicy indicates an expected call of DeleteNetworkPolicy
func (mr *MockInterfaceMockRecorder) DeleteNetworkPolicy(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteNetworkPolicy", reflect.TypeOf((*MockInterface)(nil).DeleteNetworkPolicy), arg0, arg1)
}

// DeletePod mocks base method
func (m *MockInterface) DeletePod(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePod", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePod indicates an expected call of DeletePod
func (mr *MockInterfaceMockRecorder) DeletePod(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePod", reflect.TypeOf((*MockInterface)(nil).DeletePod), arg0, arg1)
}

// DeletePodForcefully mocks base method
func (m *MockInterface) DeletePodForcefully(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePodForcefully", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePodForcefully indicates an expected call of DeletePodForcefully
func (mr *MockInterfaceMockRecorder) DeletePodForcefully(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePodForcefully", reflect.TypeOf((*MockInterface)(nil).DeletePodForcefully), arg0, arg1)
}

// DeleteRoleBinding mocks base method
func (m *MockInterface) DeleteRoleBinding(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteRoleBinding", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteRoleBinding indicates an expected call of DeleteRoleBinding
func (mr *MockInterfaceMockRecorder) DeleteRoleBinding(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRoleBinding", reflect.TypeOf((*MockInterface)(nil).DeleteRoleBinding), arg0, arg1)
}

// DeleteSecret mocks base method
func (m *MockInterface) DeleteSecret(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSecret", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSecret indicates an expected call of DeleteSecret
func (mr *MockInterfaceMockRecorder) DeleteSecret(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSecret", reflect.TypeOf((*MockInterface)(nil).DeleteSecret), arg0, arg1)
}

// DeleteService mocks base method
func (m *MockInterface) DeleteService(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteService", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteService indicates an expected call of DeleteService
func (mr *MockInterfaceMockRecorder) DeleteService(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteService", reflect.TypeOf((*MockInterface)(nil).DeleteService), arg0, arg1)
}

// DeleteServiceAccount mocks base method
func (m *MockInterface) DeleteServiceAccount(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteServiceAccount", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteServiceAccount indicates an expected call of DeleteServiceAccount
func (mr *MockInterfaceMockRecorder) DeleteServiceAccount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteServiceAccount", reflect.TypeOf((*MockInterface)(nil).DeleteServiceAccount), arg0, arg1)
}

// DeleteStatefulSet mocks base method
func (m *MockInterface) DeleteStatefulSet(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteStatefulSet", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteStatefulSet indicates an expected call of DeleteStatefulSet
func (mr *MockInterfaceMockRecorder) DeleteStatefulSet(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteStatefulSet", reflect.TypeOf((*MockInterface)(nil).DeleteStatefulSet), arg0, arg1)
}

// ForwardPodPort mocks base method
func (m *MockInterface) ForwardPodPort(arg0, arg1 string, arg2, arg3 int) (chan struct{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ForwardPodPort", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(chan struct{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ForwardPodPort indicates an expected call of ForwardPodPort
func (mr *MockInterfaceMockRecorder) ForwardPodPort(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ForwardPodPort", reflect.TypeOf((*MockInterface)(nil).ForwardPodPort), arg0, arg1, arg2, arg3)
}

// Garden mocks base method
func (m *MockInterface) Garden() versioned0.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Garden")
	ret0, _ := ret[0].(versioned0.Interface)
	return ret0
}

// Garden indicates an expected call of Garden
func (mr *MockInterfaceMockRecorder) Garden() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Garden", reflect.TypeOf((*MockInterface)(nil).Garden))
}

// GardenCore mocks base method
func (m *MockInterface) GardenCore() versioned.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GardenCore")
	ret0, _ := ret[0].(versioned.Interface)
	return ret0
}

// GardenCore indicates an expected call of GardenCore
func (mr *MockInterfaceMockRecorder) GardenCore() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GardenCore", reflect.TypeOf((*MockInterface)(nil).GardenCore))
}

// GetConfigMap mocks base method
func (m *MockInterface) GetConfigMap(arg0, arg1 string) (*v11.ConfigMap, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetConfigMap", arg0, arg1)
	ret0, _ := ret[0].(*v11.ConfigMap)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetConfigMap indicates an expected call of GetConfigMap
func (mr *MockInterfaceMockRecorder) GetConfigMap(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConfigMap", reflect.TypeOf((*MockInterface)(nil).GetConfigMap), arg0, arg1)
}

// GetDeployment mocks base method
func (m *MockInterface) GetDeployment(arg0, arg1 string) (*v1.Deployment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDeployment", arg0, arg1)
	ret0, _ := ret[0].(*v1.Deployment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDeployment indicates an expected call of GetDeployment
func (mr *MockInterfaceMockRecorder) GetDeployment(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeployment", reflect.TypeOf((*MockInterface)(nil).GetDeployment), arg0, arg1)
}

// GetJob mocks base method
func (m *MockInterface) GetJob(arg0, arg1 string) (*v10.Job, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetJob", arg0, arg1)
	ret0, _ := ret[0].(*v10.Job)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetJob indicates an expected call of GetJob
func (mr *MockInterfaceMockRecorder) GetJob(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetJob", reflect.TypeOf((*MockInterface)(nil).GetJob), arg0, arg1)
}

// GetNamespace mocks base method
func (m *MockInterface) GetNamespace(arg0 string) (*v11.Namespace, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNamespace", arg0)
	ret0, _ := ret[0].(*v11.Namespace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNamespace indicates an expected call of GetNamespace
func (mr *MockInterfaceMockRecorder) GetNamespace(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNamespace", reflect.TypeOf((*MockInterface)(nil).GetNamespace), arg0)
}

// GetPod mocks base method
func (m *MockInterface) GetPod(arg0, arg1 string) (*v11.Pod, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPod", arg0, arg1)
	ret0, _ := ret[0].(*v11.Pod)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPod indicates an expected call of GetPod
func (mr *MockInterfaceMockRecorder) GetPod(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPod", reflect.TypeOf((*MockInterface)(nil).GetPod), arg0, arg1)
}

// GetResourceAPIGroups mocks base method
func (m *MockInterface) GetResourceAPIGroups() map[string][]string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetResourceAPIGroups")
	ret0, _ := ret[0].(map[string][]string)
	return ret0
}

// GetResourceAPIGroups indicates an expected call of GetResourceAPIGroups
func (mr *MockInterfaceMockRecorder) GetResourceAPIGroups() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetResourceAPIGroups", reflect.TypeOf((*MockInterface)(nil).GetResourceAPIGroups))
}

// GetSecret mocks base method
func (m *MockInterface) GetSecret(arg0, arg1 string) (*v11.Secret, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSecret", arg0, arg1)
	ret0, _ := ret[0].(*v11.Secret)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSecret indicates an expected call of GetSecret
func (mr *MockInterfaceMockRecorder) GetSecret(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSecret", reflect.TypeOf((*MockInterface)(nil).GetSecret), arg0, arg1)
}

// GetService mocks base method
func (m *MockInterface) GetService(arg0, arg1 string) (*v11.Service, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetService", arg0, arg1)
	ret0, _ := ret[0].(*v11.Service)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetService indicates an expected call of GetService
func (mr *MockInterfaceMockRecorder) GetService(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetService", reflect.TypeOf((*MockInterface)(nil).GetService), arg0, arg1)
}

// Kubernetes mocks base method
func (m *MockInterface) Kubernetes() kubernetes0.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Kubernetes")
	ret0, _ := ret[0].(kubernetes0.Interface)
	return ret0
}

// Kubernetes indicates an expected call of Kubernetes
func (mr *MockInterfaceMockRecorder) Kubernetes() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Kubernetes", reflect.TypeOf((*MockInterface)(nil).Kubernetes))
}

// ListAPIServices mocks base method
func (m *MockInterface) ListAPIServices(arg0 v13.ListOptions) (*v1beta10.APIServiceList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAPIServices", arg0)
	ret0, _ := ret[0].(*v1beta10.APIServiceList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAPIServices indicates an expected call of ListAPIServices
func (mr *MockInterfaceMockRecorder) ListAPIServices(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAPIServices", reflect.TypeOf((*MockInterface)(nil).ListAPIServices), arg0)
}

// ListCRDs mocks base method
func (m *MockInterface) ListCRDs(arg0 v13.ListOptions) (*v1beta1.CustomResourceDefinitionList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListCRDs", arg0)
	ret0, _ := ret[0].(*v1beta1.CustomResourceDefinitionList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListCRDs indicates an expected call of ListCRDs
func (mr *MockInterfaceMockRecorder) ListCRDs(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListCRDs", reflect.TypeOf((*MockInterface)(nil).ListCRDs), arg0)
}

// ListDeployments mocks base method
func (m *MockInterface) ListDeployments(arg0 string, arg1 v13.ListOptions) (*v1.DeploymentList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListDeployments", arg0, arg1)
	ret0, _ := ret[0].(*v1.DeploymentList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListDeployments indicates an expected call of ListDeployments
func (mr *MockInterfaceMockRecorder) ListDeployments(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListDeployments", reflect.TypeOf((*MockInterface)(nil).ListDeployments), arg0, arg1)
}

// ListNamespaces mocks base method
func (m *MockInterface) ListNamespaces(arg0 v13.ListOptions) (*v11.NamespaceList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListNamespaces", arg0)
	ret0, _ := ret[0].(*v11.NamespaceList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListNamespaces indicates an expected call of ListNamespaces
func (mr *MockInterfaceMockRecorder) ListNamespaces(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListNamespaces", reflect.TypeOf((*MockInterface)(nil).ListNamespaces), arg0)
}

// ListNodes mocks base method
func (m *MockInterface) ListNodes(arg0 v13.ListOptions) (*v11.NodeList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListNodes", arg0)
	ret0, _ := ret[0].(*v11.NodeList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListNodes indicates an expected call of ListNodes
func (mr *MockInterfaceMockRecorder) ListNodes(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListNodes", reflect.TypeOf((*MockInterface)(nil).ListNodes), arg0)
}

// ListPods mocks base method
func (m *MockInterface) ListPods(arg0 string, arg1 v13.ListOptions) (*v11.PodList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPods", arg0, arg1)
	ret0, _ := ret[0].(*v11.PodList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPods indicates an expected call of ListPods
func (mr *MockInterfaceMockRecorder) ListPods(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPods", reflect.TypeOf((*MockInterface)(nil).ListPods), arg0, arg1)
}

// ListRoleBindings mocks base method
func (m *MockInterface) ListRoleBindings(arg0 string, arg1 v13.ListOptions) (*v12.RoleBindingList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListRoleBindings", arg0, arg1)
	ret0, _ := ret[0].(*v12.RoleBindingList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListRoleBindings indicates an expected call of ListRoleBindings
func (mr *MockInterfaceMockRecorder) ListRoleBindings(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRoleBindings", reflect.TypeOf((*MockInterface)(nil).ListRoleBindings), arg0, arg1)
}

// ListSecrets mocks base method
func (m *MockInterface) ListSecrets(arg0 string, arg1 v13.ListOptions) (*v11.SecretList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListSecrets", arg0, arg1)
	ret0, _ := ret[0].(*v11.SecretList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSecrets indicates an expected call of ListSecrets
func (mr *MockInterfaceMockRecorder) ListSecrets(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSecrets", reflect.TypeOf((*MockInterface)(nil).ListSecrets), arg0, arg1)
}

// ListStatefulSets mocks base method
func (m *MockInterface) ListStatefulSets(arg0 string, arg1 v13.ListOptions) (*v1.StatefulSetList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListStatefulSets", arg0, arg1)
	ret0, _ := ret[0].(*v1.StatefulSetList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListStatefulSets indicates an expected call of ListStatefulSets
func (mr *MockInterfaceMockRecorder) ListStatefulSets(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListStatefulSets", reflect.TypeOf((*MockInterface)(nil).ListStatefulSets), arg0, arg1)
}

// Machine mocks base method
func (m *MockInterface) Machine() versioned1.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Machine")
	ret0, _ := ret[0].(versioned1.Interface)
	return ret0
}

// Machine indicates an expected call of Machine
func (mr *MockInterfaceMockRecorder) Machine() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Machine", reflect.TypeOf((*MockInterface)(nil).Machine))
}

// PatchDeployment mocks base method
func (m *MockInterface) PatchDeployment(arg0, arg1 string, arg2 []byte) (*v1.Deployment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PatchDeployment", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.Deployment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PatchDeployment indicates an expected call of PatchDeployment
func (mr *MockInterfaceMockRecorder) PatchDeployment(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchDeployment", reflect.TypeOf((*MockInterface)(nil).PatchDeployment), arg0, arg1, arg2)
}

// PatchNamespace mocks base method
func (m *MockInterface) PatchNamespace(arg0 string, arg1 []byte) (*v11.Namespace, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PatchNamespace", arg0, arg1)
	ret0, _ := ret[0].(*v11.Namespace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PatchNamespace indicates an expected call of PatchNamespace
func (mr *MockInterfaceMockRecorder) PatchNamespace(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchNamespace", reflect.TypeOf((*MockInterface)(nil).PatchNamespace), arg0, arg1)
}

// RESTClient mocks base method
func (m *MockInterface) RESTClient() rest.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RESTClient")
	ret0, _ := ret[0].(rest.Interface)
	return ret0
}

// RESTClient indicates an expected call of RESTClient
func (mr *MockInterfaceMockRecorder) RESTClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RESTClient", reflect.TypeOf((*MockInterface)(nil).RESTClient))
}

// RESTConfig mocks base method
func (m *MockInterface) RESTConfig() *rest.Config {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RESTConfig")
	ret0, _ := ret[0].(*rest.Config)
	return ret0
}

// RESTConfig indicates an expected call of RESTConfig
func (mr *MockInterfaceMockRecorder) RESTConfig() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RESTConfig", reflect.TypeOf((*MockInterface)(nil).RESTConfig))
}

// RESTMapper mocks base method
func (m *MockInterface) RESTMapper() meta.RESTMapper {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RESTMapper")
	ret0, _ := ret[0].(meta.RESTMapper)
	return ret0
}

// RESTMapper indicates an expected call of RESTMapper
func (mr *MockInterfaceMockRecorder) RESTMapper() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RESTMapper", reflect.TypeOf((*MockInterface)(nil).RESTMapper))
}

// UpdateConfigMap mocks base method
func (m *MockInterface) UpdateConfigMap(arg0, arg1 string, arg2 map[string]string) (*v11.ConfigMap, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateConfigMap", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v11.ConfigMap)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateConfigMap indicates an expected call of UpdateConfigMap
func (mr *MockInterfaceMockRecorder) UpdateConfigMap(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateConfigMap", reflect.TypeOf((*MockInterface)(nil).UpdateConfigMap), arg0, arg1, arg2)
}

// UpdateSecretObject mocks base method
func (m *MockInterface) UpdateSecretObject(arg0 *v11.Secret) (*v11.Secret, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateSecretObject", arg0)
	ret0, _ := ret[0].(*v11.Secret)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateSecretObject indicates an expected call of UpdateSecretObject
func (mr *MockInterfaceMockRecorder) UpdateSecretObject(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSecretObject", reflect.TypeOf((*MockInterface)(nil).UpdateSecretObject), arg0)
}

// Version mocks base method
func (m *MockInterface) Version() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Version")
	ret0, _ := ret[0].(string)
	return ret0
}

// Version indicates an expected call of Version
func (mr *MockInterfaceMockRecorder) Version() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Version", reflect.TypeOf((*MockInterface)(nil).Version))
}
