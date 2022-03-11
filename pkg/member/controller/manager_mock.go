// Code generated by MockGen. DO NOT EDIT.
// Source: ../../apis/member.go

// Package controller is a generated GoMock package.
package controller

import (
	reflect "reflect"

	apis "github.com/hwameistor/local-storage/pkg/apis"
	apisv1alpha1 "github.com/hwameistor/local-storage/pkg/apis/hwameistor/v1alpha1"
	storage "github.com/hwameistor/local-storage/pkg/member/node/storage"
	gomock "github.com/golang/mock/gomock"
	runtime "k8s.io/apimachinery/pkg/runtime"
	cache "sigs.k8s.io/controller-runtime/pkg/cache"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// MockLocalStorageMember is a mock of LocalStorageMember interface.
type MockLocalStorageMember struct {
	ctrl     *gomock.Controller
	recorder *MockLocalStorageMemberMockRecorder
}

// MockLocalStorageMemberMockRecorder is the mock recorder for MockLocalStorageMember.
type MockLocalStorageMemberMockRecorder struct {
	mock *MockLocalStorageMember
}

// NewMockLocalStorageMember creates a new mock instance.
func NewMockLocalStorageMember(ctrl *gomock.Controller) *MockLocalStorageMember {
	mock := &MockLocalStorageMember{ctrl: ctrl}
	mock.recorder = &MockLocalStorageMemberMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLocalStorageMember) EXPECT() *MockLocalStorageMemberMockRecorder {
	return m.recorder
}

// ConfigureBase mocks base method.
func (m *MockLocalStorageMember) ConfigureBase(name, namespace string, haSystemConfig apisv1alpha1.SystemConfig, cli client.Client, informersCache cache.Cache) apis.LocalStorageMember {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConfigureBase", name, namespace, haSystemConfig, cli, informersCache)
	ret0, _ := ret[0].(apis.LocalStorageMember)
	return ret0
}

// ConfigureBase indicates an expected call of ConfigureBase.
func (mr *MockLocalStorageMemberMockRecorder) ConfigureBase(name, namespace, haSystemConfig, cli, informersCache interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConfigureBase", reflect.TypeOf((*MockLocalStorageMember)(nil).ConfigureBase), name, namespace, haSystemConfig, cli, informersCache)
}

// ConfigureCSIDriver mocks base method.
func (m *MockLocalStorageMember) ConfigureCSIDriver(driverName, sockAddr string) apis.LocalStorageMember {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConfigureCSIDriver", driverName, sockAddr)
	ret0, _ := ret[0].(apis.LocalStorageMember)
	return ret0
}

// ConfigureCSIDriver indicates an expected call of ConfigureCSIDriver.
func (mr *MockLocalStorageMemberMockRecorder) ConfigureCSIDriver(driverName, sockAddr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConfigureCSIDriver", reflect.TypeOf((*MockLocalStorageMember)(nil).ConfigureCSIDriver), driverName, sockAddr)
}

// ConfigureController mocks base method.
func (m *MockLocalStorageMember) ConfigureController(scheme *runtime.Scheme) apis.LocalStorageMember {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConfigureController", scheme)
	ret0, _ := ret[0].(apis.LocalStorageMember)
	return ret0
}

// ConfigureController indicates an expected call of ConfigureController.
func (mr *MockLocalStorageMemberMockRecorder) ConfigureController(scheme interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConfigureController", reflect.TypeOf((*MockLocalStorageMember)(nil).ConfigureController), scheme)
}

// ConfigureNode mocks base method.
func (m *MockLocalStorageMember) ConfigureNode() apis.LocalStorageMember {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConfigureNode")
	ret0, _ := ret[0].(apis.LocalStorageMember)
	return ret0
}

// ConfigureNode indicates an expected call of ConfigureNode.
func (mr *MockLocalStorageMemberMockRecorder) ConfigureNode() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConfigureNode", reflect.TypeOf((*MockLocalStorageMember)(nil).ConfigureNode))
}

// ConfigureRESTServer mocks base method.
func (m *MockLocalStorageMember) ConfigureRESTServer(httpPort int) apis.LocalStorageMember {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConfigureRESTServer", httpPort)
	ret0, _ := ret[0].(apis.LocalStorageMember)
	return ret0
}

// ConfigureRESTServer indicates an expected call of ConfigureRESTServer.
func (mr *MockLocalStorageMemberMockRecorder) ConfigureRESTServer(httpPort interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConfigureRESTServer", reflect.TypeOf((*MockLocalStorageMember)(nil).ConfigureRESTServer), httpPort)
}

// Controller mocks base method.
func (m *MockLocalStorageMember) Controller() apis.ControllerManager {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Controller")
	ret0, _ := ret[0].(apis.ControllerManager)
	return ret0
}

// Controller indicates an expected call of Controller.
func (mr *MockLocalStorageMemberMockRecorder) Controller() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Controller", reflect.TypeOf((*MockLocalStorageMember)(nil).Controller))
}

// DriverName mocks base method.
func (m *MockLocalStorageMember) DriverName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DriverName")
	ret0, _ := ret[0].(string)
	return ret0
}

// DriverName indicates an expected call of DriverName.
func (mr *MockLocalStorageMemberMockRecorder) DriverName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DriverName", reflect.TypeOf((*MockLocalStorageMember)(nil).DriverName))
}

// Name mocks base method.
func (m *MockLocalStorageMember) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockLocalStorageMemberMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockLocalStorageMember)(nil).Name))
}

// Node mocks base method.
func (m *MockLocalStorageMember) Node() apis.NodeManager {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Node")
	ret0, _ := ret[0].(apis.NodeManager)
	return ret0
}

// Node indicates an expected call of Node.
func (mr *MockLocalStorageMemberMockRecorder) Node() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Node", reflect.TypeOf((*MockLocalStorageMember)(nil).Node))
}

// Run mocks base method.
func (m *MockLocalStorageMember) Run(stopCh <-chan struct{}) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Run", stopCh)
}

// Run indicates an expected call of Run.
func (mr *MockLocalStorageMemberMockRecorder) Run(stopCh interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockLocalStorageMember)(nil).Run), stopCh)
}

// Version mocks base method.
func (m *MockLocalStorageMember) Version() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Version")
	ret0, _ := ret[0].(string)
	return ret0
}

// Version indicates an expected call of Version.
func (mr *MockLocalStorageMemberMockRecorder) Version() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Version", reflect.TypeOf((*MockLocalStorageMember)(nil).Version))
}

// MockControllerManager is a mock of ControllerManager interface.
type MockControllerManager struct {
	ctrl     *gomock.Controller
	recorder *MockControllerManagerMockRecorder
}

// MockControllerManagerMockRecorder is the mock recorder for MockControllerManager.
type MockControllerManagerMockRecorder struct {
	mock *MockControllerManager
}

// NewMockControllerManager creates a new mock instance.
func NewMockControllerManager(ctrl *gomock.Controller) *MockControllerManager {
	mock := &MockControllerManager{ctrl: ctrl}
	mock.recorder = &MockControllerManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockControllerManager) EXPECT() *MockControllerManagerMockRecorder {
	return m.recorder
}

// ReconcileNode mocks base method.
func (m *MockControllerManager) ReconcileNode(node *apisv1alpha1.LocalStorageNode) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ReconcileNode", node)
}

// ReconcileNode indicates an expected call of ReconcileNode.
func (mr *MockControllerManagerMockRecorder) ReconcileNode(node interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileNode", reflect.TypeOf((*MockControllerManager)(nil).ReconcileNode), node)
}

// ReconcileVolume mocks base method.
func (m *MockControllerManager) ReconcileVolume(vol *apisv1alpha1.LocalVolume) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ReconcileVolume", vol)
}

// ReconcileVolume indicates an expected call of ReconcileVolume.
func (mr *MockControllerManagerMockRecorder) ReconcileVolume(vol interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileVolume", reflect.TypeOf((*MockControllerManager)(nil).ReconcileVolume), vol)
}

// ReconcileVolumeConvert mocks base method.
func (m *MockControllerManager) ReconcileVolumeConvert(vol *apisv1alpha1.LocalVolumeConvert) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ReconcileVolumeConvert", vol)
}

// ReconcileVolumeConvert indicates an expected call of ReconcileVolumeConvert.
func (mr *MockControllerManagerMockRecorder) ReconcileVolumeConvert(vol interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileVolumeConvert", reflect.TypeOf((*MockControllerManager)(nil).ReconcileVolumeConvert), vol)
}

// ReconcileVolumeExpand mocks base method.
func (m *MockControllerManager) ReconcileVolumeExpand(vol *apisv1alpha1.LocalVolumeExpand) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ReconcileVolumeExpand", vol)
}

// ReconcileVolumeExpand indicates an expected call of ReconcileVolumeExpand.
func (mr *MockControllerManagerMockRecorder) ReconcileVolumeExpand(vol interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileVolumeExpand", reflect.TypeOf((*MockControllerManager)(nil).ReconcileVolumeExpand), vol)
}

// ReconcileVolumeMigrate mocks base method.
func (m *MockControllerManager) ReconcileVolumeMigrate(vol *apisv1alpha1.LocalVolumeMigrate) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ReconcileVolumeMigrate", vol)
}

// ReconcileVolumeMigrate indicates an expected call of ReconcileVolumeMigrate.
func (mr *MockControllerManagerMockRecorder) ReconcileVolumeMigrate(vol interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileVolumeMigrate", reflect.TypeOf((*MockControllerManager)(nil).ReconcileVolumeMigrate), vol)
}

// Run mocks base method.
func (m *MockControllerManager) Run(stopCh <-chan struct{}) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Run", stopCh)
}

// Run indicates an expected call of Run.
func (mr *MockControllerManagerMockRecorder) Run(stopCh interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockControllerManager)(nil).Run), stopCh)
}

// MockNodeManager is a mock of NodeManager interface.
type MockNodeManager struct {
	ctrl     *gomock.Controller
	recorder *MockNodeManagerMockRecorder
}

// MockNodeManagerMockRecorder is the mock recorder for MockNodeManager.
type MockNodeManagerMockRecorder struct {
	mock *MockNodeManager
}

// NewMockNodeManager creates a new mock instance.
func NewMockNodeManager(ctrl *gomock.Controller) *MockNodeManager {
	mock := &MockNodeManager{ctrl: ctrl}
	mock.recorder = &MockNodeManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNodeManager) EXPECT() *MockNodeManagerMockRecorder {
	return m.recorder
}

// ReconcileVolumeReplica mocks base method.
func (m *MockNodeManager) ReconcileVolumeReplica(replica *apisv1alpha1.LocalVolumeReplica) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ReconcileVolumeReplica", replica)
}

// ReconcileVolumeReplica indicates an expected call of ReconcileVolumeReplica.
func (mr *MockNodeManagerMockRecorder) ReconcileVolumeReplica(replica interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileVolumeReplica", reflect.TypeOf((*MockNodeManager)(nil).ReconcileVolumeReplica), replica)
}

// Run mocks base method.
func (m *MockNodeManager) Run(stopCh <-chan struct{}) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Run", stopCh)
}

// Run indicates an expected call of Run.
func (mr *MockNodeManagerMockRecorder) Run(stopCh interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockNodeManager)(nil).Run), stopCh)
}

// Storage mocks base method.
func (m *MockNodeManager) Storage() *storage.LocalManager {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Storage")
	ret0, _ := ret[0].(*storage.LocalManager)
	return ret0
}

// Storage indicates an expected call of Storage.
func (mr *MockNodeManagerMockRecorder) Storage() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Storage", reflect.TypeOf((*MockNodeManager)(nil).Storage))
}

// TakeVolumeReplicaTaskAssignment mocks base method.
func (m *MockNodeManager) TakeVolumeReplicaTaskAssignment(vol *apisv1alpha1.LocalVolume) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "TakeVolumeReplicaTaskAssignment", vol)
}

// TakeVolumeReplicaTaskAssignment indicates an expected call of TakeVolumeReplicaTaskAssignment.
func (mr *MockNodeManagerMockRecorder) TakeVolumeReplicaTaskAssignment(vol interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TakeVolumeReplicaTaskAssignment", reflect.TypeOf((*MockNodeManager)(nil).TakeVolumeReplicaTaskAssignment), vol)
}
