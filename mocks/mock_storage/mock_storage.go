// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/netapp/trident/storage (interfaces: Backend,Pool)

// Package mock_storage is a generated GoMock package.
package mock_storage

import (
	context "context"
	reflect "reflect"

	roaring "github.com/RoaringBitmap/roaring"
	gomock "github.com/golang/mock/gomock"
	config "github.com/netapp/trident/config"
	storage "github.com/netapp/trident/storage"
	storageattribute "github.com/netapp/trident/storage_attribute"
	utils "github.com/netapp/trident/utils"
)

// MockBackend is a mock of Backend interface.
type MockBackend struct {
	ctrl     *gomock.Controller
	recorder *MockBackendMockRecorder
}

// MockBackendMockRecorder is the mock recorder for MockBackend.
type MockBackendMockRecorder struct {
	mock *MockBackend
}

// NewMockBackend creates a new mock instance.
func NewMockBackend(ctrl *gomock.Controller) *MockBackend {
	mock := &MockBackend{ctrl: ctrl}
	mock.recorder = &MockBackendMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBackend) EXPECT() *MockBackendMockRecorder {
	return m.recorder
}

// AddStoragePool mocks base method.
func (m *MockBackend) AddStoragePool(arg0 storage.Pool) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddStoragePool", arg0)
}

// AddStoragePool indicates an expected call of AddStoragePool.
func (mr *MockBackendMockRecorder) AddStoragePool(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddStoragePool", reflect.TypeOf((*MockBackend)(nil).AddStoragePool), arg0)
}

// AddVolume mocks base method.
func (m *MockBackend) AddVolume(arg0 context.Context, arg1 *storage.VolumeConfig, arg2 storage.Pool, arg3 map[string]storageattribute.Request, arg4 bool) (*storage.Volume, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddVolume", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(*storage.Volume)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddVolume indicates an expected call of AddVolume.
func (mr *MockBackendMockRecorder) AddVolume(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddVolume", reflect.TypeOf((*MockBackend)(nil).AddVolume), arg0, arg1, arg2, arg3, arg4)
}

// BackendUUID mocks base method.
func (m *MockBackend) BackendUUID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BackendUUID")
	ret0, _ := ret[0].(string)
	return ret0
}

// BackendUUID indicates an expected call of BackendUUID.
func (mr *MockBackendMockRecorder) BackendUUID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BackendUUID", reflect.TypeOf((*MockBackend)(nil).BackendUUID))
}

// CanEnablePublishEnforcement mocks base method.
func (m *MockBackend) CanEnablePublishEnforcement() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CanEnablePublishEnforcement")
	ret0, _ := ret[0].(bool)
	return ret0
}

// CanEnablePublishEnforcement indicates an expected call of CanEnablePublishEnforcement.
func (mr *MockBackendMockRecorder) CanEnablePublishEnforcement() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CanEnablePublishEnforcement", reflect.TypeOf((*MockBackend)(nil).CanEnablePublishEnforcement))
}

// CanGetState mocks base method.
func (m *MockBackend) CanGetState() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CanGetState")
	ret0, _ := ret[0].(bool)
	return ret0
}

// CanGetState indicates an expected call of CanGetState.
func (mr *MockBackendMockRecorder) CanGetState() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CanGetState", reflect.TypeOf((*MockBackend)(nil).CanGetState))
}

// CanMirror mocks base method.
func (m *MockBackend) CanMirror() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CanMirror")
	ret0, _ := ret[0].(bool)
	return ret0
}

// CanMirror indicates an expected call of CanMirror.
func (mr *MockBackendMockRecorder) CanMirror() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CanMirror", reflect.TypeOf((*MockBackend)(nil).CanMirror))
}

// CanSnapshot mocks base method.
func (m *MockBackend) CanSnapshot(arg0 context.Context, arg1 *storage.SnapshotConfig, arg2 *storage.VolumeConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CanSnapshot", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// CanSnapshot indicates an expected call of CanSnapshot.
func (mr *MockBackendMockRecorder) CanSnapshot(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CanSnapshot", reflect.TypeOf((*MockBackend)(nil).CanSnapshot), arg0, arg1, arg2)
}

// CloneVolume mocks base method.
func (m *MockBackend) CloneVolume(arg0 context.Context, arg1, arg2 *storage.VolumeConfig, arg3 storage.Pool, arg4 bool) (*storage.Volume, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloneVolume", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(*storage.Volume)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CloneVolume indicates an expected call of CloneVolume.
func (mr *MockBackendMockRecorder) CloneVolume(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloneVolume", reflect.TypeOf((*MockBackend)(nil).CloneVolume), arg0, arg1, arg2, arg3, arg4)
}

// ConfigRef mocks base method.
func (m *MockBackend) ConfigRef() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConfigRef")
	ret0, _ := ret[0].(string)
	return ret0
}

// ConfigRef indicates an expected call of ConfigRef.
func (mr *MockBackendMockRecorder) ConfigRef() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConfigRef", reflect.TypeOf((*MockBackend)(nil).ConfigRef))
}

// ConstructExternal mocks base method.
func (m *MockBackend) ConstructExternal(arg0 context.Context) *storage.BackendExternal {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConstructExternal", arg0)
	ret0, _ := ret[0].(*storage.BackendExternal)
	return ret0
}

// ConstructExternal indicates an expected call of ConstructExternal.
func (mr *MockBackendMockRecorder) ConstructExternal(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConstructExternal", reflect.TypeOf((*MockBackend)(nil).ConstructExternal), arg0)
}

// ConstructPersistent mocks base method.
func (m *MockBackend) ConstructPersistent(arg0 context.Context) *storage.BackendPersistent {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConstructPersistent", arg0)
	ret0, _ := ret[0].(*storage.BackendPersistent)
	return ret0
}

// ConstructPersistent indicates an expected call of ConstructPersistent.
func (mr *MockBackendMockRecorder) ConstructPersistent(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConstructPersistent", reflect.TypeOf((*MockBackend)(nil).ConstructPersistent), arg0)
}

// CreateSnapshot mocks base method.
func (m *MockBackend) CreateSnapshot(arg0 context.Context, arg1 *storage.SnapshotConfig, arg2 *storage.VolumeConfig) (*storage.Snapshot, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSnapshot", arg0, arg1, arg2)
	ret0, _ := ret[0].(*storage.Snapshot)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSnapshot indicates an expected call of CreateSnapshot.
func (mr *MockBackendMockRecorder) CreateSnapshot(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSnapshot", reflect.TypeOf((*MockBackend)(nil).CreateSnapshot), arg0, arg1, arg2)
}

// DeleteSnapshot mocks base method.
func (m *MockBackend) DeleteSnapshot(arg0 context.Context, arg1 *storage.SnapshotConfig, arg2 *storage.VolumeConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSnapshot", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSnapshot indicates an expected call of DeleteSnapshot.
func (mr *MockBackendMockRecorder) DeleteSnapshot(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSnapshot", reflect.TypeOf((*MockBackend)(nil).DeleteSnapshot), arg0, arg1, arg2)
}

// Driver mocks base method.
func (m *MockBackend) Driver() storage.Driver {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Driver")
	ret0, _ := ret[0].(storage.Driver)
	return ret0
}

// Driver indicates an expected call of Driver.
func (mr *MockBackendMockRecorder) Driver() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Driver", reflect.TypeOf((*MockBackend)(nil).Driver))
}

// EnablePublishEnforcement mocks base method.
func (m *MockBackend) EnablePublishEnforcement(arg0 context.Context, arg1 *storage.Volume) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnablePublishEnforcement", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnablePublishEnforcement indicates an expected call of EnablePublishEnforcement.
func (mr *MockBackendMockRecorder) EnablePublishEnforcement(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnablePublishEnforcement", reflect.TypeOf((*MockBackend)(nil).EnablePublishEnforcement), arg0, arg1)
}

// GetBackendState mocks base method.
func (m *MockBackend) GetBackendState(arg0 context.Context) (string, *roaring.Bitmap) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBackendState", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(*roaring.Bitmap)
	return ret0, ret1
}

// GetBackendState indicates an expected call of GetBackendState.
func (mr *MockBackendMockRecorder) GetBackendState(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBackendState", reflect.TypeOf((*MockBackend)(nil).GetBackendState), arg0)
}

// GetChapInfo mocks base method.
func (m *MockBackend) GetChapInfo(arg0 context.Context, arg1, arg2 string) (*utils.IscsiChapInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetChapInfo", arg0, arg1, arg2)
	ret0, _ := ret[0].(*utils.IscsiChapInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetChapInfo indicates an expected call of GetChapInfo.
func (mr *MockBackendMockRecorder) GetChapInfo(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetChapInfo", reflect.TypeOf((*MockBackend)(nil).GetChapInfo), arg0, arg1, arg2)
}

// GetDebugTraceFlags mocks base method.
func (m *MockBackend) GetDebugTraceFlags(arg0 context.Context) map[string]bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDebugTraceFlags", arg0)
	ret0, _ := ret[0].(map[string]bool)
	return ret0
}

// GetDebugTraceFlags indicates an expected call of GetDebugTraceFlags.
func (mr *MockBackendMockRecorder) GetDebugTraceFlags(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDebugTraceFlags", reflect.TypeOf((*MockBackend)(nil).GetDebugTraceFlags), arg0)
}

// GetDriverName mocks base method.
func (m *MockBackend) GetDriverName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDriverName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetDriverName indicates an expected call of GetDriverName.
func (mr *MockBackendMockRecorder) GetDriverName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDriverName", reflect.TypeOf((*MockBackend)(nil).GetDriverName))
}

// GetPhysicalPoolNames mocks base method.
func (m *MockBackend) GetPhysicalPoolNames(arg0 context.Context) []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPhysicalPoolNames", arg0)
	ret0, _ := ret[0].([]string)
	return ret0
}

// GetPhysicalPoolNames indicates an expected call of GetPhysicalPoolNames.
func (mr *MockBackendMockRecorder) GetPhysicalPoolNames(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPhysicalPoolNames", reflect.TypeOf((*MockBackend)(nil).GetPhysicalPoolNames), arg0)
}

// GetProtocol mocks base method.
func (m *MockBackend) GetProtocol(arg0 context.Context) config.Protocol {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProtocol", arg0)
	ret0, _ := ret[0].(config.Protocol)
	return ret0
}

// GetProtocol indicates an expected call of GetProtocol.
func (mr *MockBackendMockRecorder) GetProtocol(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProtocol", reflect.TypeOf((*MockBackend)(nil).GetProtocol), arg0)
}

// GetSnapshot mocks base method.
func (m *MockBackend) GetSnapshot(arg0 context.Context, arg1 *storage.SnapshotConfig, arg2 *storage.VolumeConfig) (*storage.Snapshot, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSnapshot", arg0, arg1, arg2)
	ret0, _ := ret[0].(*storage.Snapshot)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSnapshot indicates an expected call of GetSnapshot.
func (mr *MockBackendMockRecorder) GetSnapshot(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSnapshot", reflect.TypeOf((*MockBackend)(nil).GetSnapshot), arg0, arg1, arg2)
}

// GetSnapshots mocks base method.
func (m *MockBackend) GetSnapshots(arg0 context.Context, arg1 *storage.VolumeConfig) ([]*storage.Snapshot, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSnapshots", arg0, arg1)
	ret0, _ := ret[0].([]*storage.Snapshot)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSnapshots indicates an expected call of GetSnapshots.
func (mr *MockBackendMockRecorder) GetSnapshots(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSnapshots", reflect.TypeOf((*MockBackend)(nil).GetSnapshots), arg0, arg1)
}

// GetUpdateType mocks base method.
func (m *MockBackend) GetUpdateType(arg0 context.Context, arg1 storage.Backend) *roaring.Bitmap {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUpdateType", arg0, arg1)
	ret0, _ := ret[0].(*roaring.Bitmap)
	return ret0
}

// GetUpdateType indicates an expected call of GetUpdateType.
func (mr *MockBackendMockRecorder) GetUpdateType(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUpdateType", reflect.TypeOf((*MockBackend)(nil).GetUpdateType), arg0, arg1)
}

// GetVolumeExternal mocks base method.
func (m *MockBackend) GetVolumeExternal(arg0 context.Context, arg1 string) (*storage.VolumeExternal, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVolumeExternal", arg0, arg1)
	ret0, _ := ret[0].(*storage.VolumeExternal)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVolumeExternal indicates an expected call of GetVolumeExternal.
func (mr *MockBackendMockRecorder) GetVolumeExternal(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVolumeExternal", reflect.TypeOf((*MockBackend)(nil).GetVolumeExternal), arg0, arg1)
}

// HasVolumes mocks base method.
func (m *MockBackend) HasVolumes() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasVolumes")
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasVolumes indicates an expected call of HasVolumes.
func (mr *MockBackendMockRecorder) HasVolumes() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasVolumes", reflect.TypeOf((*MockBackend)(nil).HasVolumes))
}

// ImportVolume mocks base method.
func (m *MockBackend) ImportVolume(arg0 context.Context, arg1 *storage.VolumeConfig) (*storage.Volume, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ImportVolume", arg0, arg1)
	ret0, _ := ret[0].(*storage.Volume)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ImportVolume indicates an expected call of ImportVolume.
func (mr *MockBackendMockRecorder) ImportVolume(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ImportVolume", reflect.TypeOf((*MockBackend)(nil).ImportVolume), arg0, arg1)
}

// InvalidateNodeAccess mocks base method.
func (m *MockBackend) InvalidateNodeAccess() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "InvalidateNodeAccess")
}

// InvalidateNodeAccess indicates an expected call of InvalidateNodeAccess.
func (mr *MockBackendMockRecorder) InvalidateNodeAccess() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InvalidateNodeAccess", reflect.TypeOf((*MockBackend)(nil).InvalidateNodeAccess))
}

// IsCredentialsFieldSet mocks base method.
func (m *MockBackend) IsCredentialsFieldSet(arg0 context.Context) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsCredentialsFieldSet", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsCredentialsFieldSet indicates an expected call of IsCredentialsFieldSet.
func (mr *MockBackendMockRecorder) IsCredentialsFieldSet(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsCredentialsFieldSet", reflect.TypeOf((*MockBackend)(nil).IsCredentialsFieldSet), arg0)
}

// Name mocks base method.
func (m *MockBackend) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockBackendMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockBackend)(nil).Name))
}

// Online mocks base method.
func (m *MockBackend) Online() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Online")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Online indicates an expected call of Online.
func (mr *MockBackendMockRecorder) Online() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Online", reflect.TypeOf((*MockBackend)(nil).Online))
}

// PublishVolume mocks base method.
func (m *MockBackend) PublishVolume(arg0 context.Context, arg1 *storage.VolumeConfig, arg2 *utils.VolumePublishInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PublishVolume", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// PublishVolume indicates an expected call of PublishVolume.
func (mr *MockBackendMockRecorder) PublishVolume(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PublishVolume", reflect.TypeOf((*MockBackend)(nil).PublishVolume), arg0, arg1, arg2)
}

// ReconcileNodeAccess mocks base method.
func (m *MockBackend) ReconcileNodeAccess(arg0 context.Context, arg1 []*utils.Node, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileNodeAccess", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReconcileNodeAccess indicates an expected call of ReconcileNodeAccess.
func (mr *MockBackendMockRecorder) ReconcileNodeAccess(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileNodeAccess", reflect.TypeOf((*MockBackend)(nil).ReconcileNodeAccess), arg0, arg1, arg2)
}

// RemoveCachedVolume mocks base method.
func (m *MockBackend) RemoveCachedVolume(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RemoveCachedVolume", arg0)
}

// RemoveCachedVolume indicates an expected call of RemoveCachedVolume.
func (mr *MockBackendMockRecorder) RemoveCachedVolume(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveCachedVolume", reflect.TypeOf((*MockBackend)(nil).RemoveCachedVolume), arg0)
}

// RemoveVolume mocks base method.
func (m *MockBackend) RemoveVolume(arg0 context.Context, arg1 *storage.VolumeConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveVolume", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveVolume indicates an expected call of RemoveVolume.
func (mr *MockBackendMockRecorder) RemoveVolume(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveVolume", reflect.TypeOf((*MockBackend)(nil).RemoveVolume), arg0, arg1)
}

// RenameVolume mocks base method.
func (m *MockBackend) RenameVolume(arg0 context.Context, arg1 *storage.VolumeConfig, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RenameVolume", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// RenameVolume indicates an expected call of RenameVolume.
func (mr *MockBackendMockRecorder) RenameVolume(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RenameVolume", reflect.TypeOf((*MockBackend)(nil).RenameVolume), arg0, arg1, arg2)
}

// ResizeVolume mocks base method.
func (m *MockBackend) ResizeVolume(arg0 context.Context, arg1 *storage.VolumeConfig, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResizeVolume", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// ResizeVolume indicates an expected call of ResizeVolume.
func (mr *MockBackendMockRecorder) ResizeVolume(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResizeVolume", reflect.TypeOf((*MockBackend)(nil).ResizeVolume), arg0, arg1, arg2)
}

// RestoreSnapshot mocks base method.
func (m *MockBackend) RestoreSnapshot(arg0 context.Context, arg1 *storage.SnapshotConfig, arg2 *storage.VolumeConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RestoreSnapshot", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// RestoreSnapshot indicates an expected call of RestoreSnapshot.
func (mr *MockBackendMockRecorder) RestoreSnapshot(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RestoreSnapshot", reflect.TypeOf((*MockBackend)(nil).RestoreSnapshot), arg0, arg1, arg2)
}

// SetBackendUUID mocks base method.
func (m *MockBackend) SetBackendUUID(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetBackendUUID", arg0)
}

// SetBackendUUID indicates an expected call of SetBackendUUID.
func (mr *MockBackendMockRecorder) SetBackendUUID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetBackendUUID", reflect.TypeOf((*MockBackend)(nil).SetBackendUUID), arg0)
}

// SetConfigRef mocks base method.
func (m *MockBackend) SetConfigRef(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetConfigRef", arg0)
}

// SetConfigRef indicates an expected call of SetConfigRef.
func (mr *MockBackendMockRecorder) SetConfigRef(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetConfigRef", reflect.TypeOf((*MockBackend)(nil).SetConfigRef), arg0)
}

// SetDriver mocks base method.
func (m *MockBackend) SetDriver(arg0 storage.Driver) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetDriver", arg0)
}

// SetDriver indicates an expected call of SetDriver.
func (mr *MockBackendMockRecorder) SetDriver(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetDriver", reflect.TypeOf((*MockBackend)(nil).SetDriver), arg0)
}

// SetName mocks base method.
func (m *MockBackend) SetName(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetName", arg0)
}

// SetName indicates an expected call of SetName.
func (mr *MockBackendMockRecorder) SetName(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetName", reflect.TypeOf((*MockBackend)(nil).SetName), arg0)
}

// SetOnline mocks base method.
func (m *MockBackend) SetOnline(arg0 bool) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetOnline", arg0)
}

// SetOnline indicates an expected call of SetOnline.
func (mr *MockBackendMockRecorder) SetOnline(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetOnline", reflect.TypeOf((*MockBackend)(nil).SetOnline), arg0)
}

// SetState mocks base method.
func (m *MockBackend) SetState(arg0 storage.BackendState) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetState", arg0)
}

// SetState indicates an expected call of SetState.
func (mr *MockBackendMockRecorder) SetState(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetState", reflect.TypeOf((*MockBackend)(nil).SetState), arg0)
}

// SetStorage mocks base method.
func (m *MockBackend) SetStorage(arg0 map[string]storage.Pool) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetStorage", arg0)
}

// SetStorage indicates an expected call of SetStorage.
func (mr *MockBackendMockRecorder) SetStorage(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetStorage", reflect.TypeOf((*MockBackend)(nil).SetStorage), arg0)
}

// SetVolumes mocks base method.
func (m *MockBackend) SetVolumes(arg0 map[string]*storage.Volume) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetVolumes", arg0)
}

// SetVolumes indicates an expected call of SetVolumes.
func (mr *MockBackendMockRecorder) SetVolumes(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetVolumes", reflect.TypeOf((*MockBackend)(nil).SetVolumes), arg0)
}

// State mocks base method.
func (m *MockBackend) State() storage.BackendState {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "State")
	ret0, _ := ret[0].(storage.BackendState)
	return ret0
}

// State indicates an expected call of State.
func (mr *MockBackendMockRecorder) State() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "State", reflect.TypeOf((*MockBackend)(nil).State))
}

// StateReason mocks base method.
func (m *MockBackend) StateReason() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StateReason")
	ret0, _ := ret[0].(string)
	return ret0
}

// StateReason indicates an expected call of StateReason.
func (mr *MockBackendMockRecorder) StateReason() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StateReason", reflect.TypeOf((*MockBackend)(nil).StateReason))
}

// Storage mocks base method.
func (m *MockBackend) Storage() map[string]storage.Pool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Storage")
	ret0, _ := ret[0].(map[string]storage.Pool)
	return ret0
}

// Storage indicates an expected call of Storage.
func (mr *MockBackendMockRecorder) Storage() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Storage", reflect.TypeOf((*MockBackend)(nil).Storage))
}

// Terminate mocks base method.
func (m *MockBackend) Terminate(arg0 context.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Terminate", arg0)
}

// Terminate indicates an expected call of Terminate.
func (mr *MockBackendMockRecorder) Terminate(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Terminate", reflect.TypeOf((*MockBackend)(nil).Terminate), arg0)
}

// UnpublishVolume mocks base method.
func (m *MockBackend) UnpublishVolume(arg0 context.Context, arg1 *storage.VolumeConfig, arg2 *utils.VolumePublishInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnpublishVolume", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UnpublishVolume indicates an expected call of UnpublishVolume.
func (mr *MockBackendMockRecorder) UnpublishVolume(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnpublishVolume", reflect.TypeOf((*MockBackend)(nil).UnpublishVolume), arg0, arg1, arg2)
}

// Volumes mocks base method.
func (m *MockBackend) Volumes() map[string]*storage.Volume {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Volumes")
	ret0, _ := ret[0].(map[string]*storage.Volume)
	return ret0
}

// Volumes indicates an expected call of Volumes.
func (mr *MockBackendMockRecorder) Volumes() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Volumes", reflect.TypeOf((*MockBackend)(nil).Volumes))
}

// MockPool is a mock of Pool interface.
type MockPool struct {
	ctrl     *gomock.Controller
	recorder *MockPoolMockRecorder
}

// MockPoolMockRecorder is the mock recorder for MockPool.
type MockPoolMockRecorder struct {
	mock *MockPool
}

// NewMockPool creates a new mock instance.
func NewMockPool(ctrl *gomock.Controller) *MockPool {
	mock := &MockPool{ctrl: ctrl}
	mock.recorder = &MockPoolMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPool) EXPECT() *MockPoolMockRecorder {
	return m.recorder
}

// AddStorageClass mocks base method.
func (m *MockPool) AddStorageClass(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddStorageClass", arg0)
}

// AddStorageClass indicates an expected call of AddStorageClass.
func (mr *MockPoolMockRecorder) AddStorageClass(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddStorageClass", reflect.TypeOf((*MockPool)(nil).AddStorageClass), arg0)
}

// Attributes mocks base method.
func (m *MockPool) Attributes() map[string]storageattribute.Offer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Attributes")
	ret0, _ := ret[0].(map[string]storageattribute.Offer)
	return ret0
}

// Attributes indicates an expected call of Attributes.
func (mr *MockPoolMockRecorder) Attributes() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Attributes", reflect.TypeOf((*MockPool)(nil).Attributes))
}

// Backend mocks base method.
func (m *MockPool) Backend() storage.Backend {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Backend")
	ret0, _ := ret[0].(storage.Backend)
	return ret0
}

// Backend indicates an expected call of Backend.
func (mr *MockPoolMockRecorder) Backend() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Backend", reflect.TypeOf((*MockPool)(nil).Backend))
}

// ConstructExternal mocks base method.
func (m *MockPool) ConstructExternal() *storage.PoolExternal {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConstructExternal")
	ret0, _ := ret[0].(*storage.PoolExternal)
	return ret0
}

// ConstructExternal indicates an expected call of ConstructExternal.
func (mr *MockPoolMockRecorder) ConstructExternal() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConstructExternal", reflect.TypeOf((*MockPool)(nil).ConstructExternal))
}

// GetLabels mocks base method.
func (m *MockPool) GetLabels(arg0 context.Context, arg1 string) map[string]string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLabels", arg0, arg1)
	ret0, _ := ret[0].(map[string]string)
	return ret0
}

// GetLabels indicates an expected call of GetLabels.
func (mr *MockPoolMockRecorder) GetLabels(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLabels", reflect.TypeOf((*MockPool)(nil).GetLabels), arg0, arg1)
}

// GetLabelsJSON mocks base method.
func (m *MockPool) GetLabelsJSON(arg0 context.Context, arg1 string, arg2 int) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLabelsJSON", arg0, arg1, arg2)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLabelsJSON indicates an expected call of GetLabelsJSON.
func (mr *MockPoolMockRecorder) GetLabelsJSON(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLabelsJSON", reflect.TypeOf((*MockPool)(nil).GetLabelsJSON), arg0, arg1, arg2)
}

// InternalAttributes mocks base method.
func (m *MockPool) InternalAttributes() map[string]string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InternalAttributes")
	ret0, _ := ret[0].(map[string]string)
	return ret0
}

// InternalAttributes indicates an expected call of InternalAttributes.
func (mr *MockPoolMockRecorder) InternalAttributes() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InternalAttributes", reflect.TypeOf((*MockPool)(nil).InternalAttributes))
}

// Name mocks base method.
func (m *MockPool) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockPoolMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockPool)(nil).Name))
}

// RemoveStorageClass mocks base method.
func (m *MockPool) RemoveStorageClass(arg0 string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveStorageClass", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// RemoveStorageClass indicates an expected call of RemoveStorageClass.
func (mr *MockPoolMockRecorder) RemoveStorageClass(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveStorageClass", reflect.TypeOf((*MockPool)(nil).RemoveStorageClass), arg0)
}

// SetAttributes mocks base method.
func (m *MockPool) SetAttributes(arg0 map[string]storageattribute.Offer) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetAttributes", arg0)
}

// SetAttributes indicates an expected call of SetAttributes.
func (mr *MockPoolMockRecorder) SetAttributes(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetAttributes", reflect.TypeOf((*MockPool)(nil).SetAttributes), arg0)
}

// SetBackend mocks base method.
func (m *MockPool) SetBackend(arg0 storage.Backend) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetBackend", arg0)
}

// SetBackend indicates an expected call of SetBackend.
func (mr *MockPoolMockRecorder) SetBackend(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetBackend", reflect.TypeOf((*MockPool)(nil).SetBackend), arg0)
}

// SetInternalAttributes mocks base method.
func (m *MockPool) SetInternalAttributes(arg0 map[string]string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetInternalAttributes", arg0)
}

// SetInternalAttributes indicates an expected call of SetInternalAttributes.
func (mr *MockPoolMockRecorder) SetInternalAttributes(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetInternalAttributes", reflect.TypeOf((*MockPool)(nil).SetInternalAttributes), arg0)
}

// SetName mocks base method.
func (m *MockPool) SetName(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetName", arg0)
}

// SetName indicates an expected call of SetName.
func (mr *MockPoolMockRecorder) SetName(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetName", reflect.TypeOf((*MockPool)(nil).SetName), arg0)
}

// SetStorageClasses mocks base method.
func (m *MockPool) SetStorageClasses(arg0 []string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetStorageClasses", arg0)
}

// SetStorageClasses indicates an expected call of SetStorageClasses.
func (mr *MockPoolMockRecorder) SetStorageClasses(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetStorageClasses", reflect.TypeOf((*MockPool)(nil).SetStorageClasses), arg0)
}

// SetSupportedTopologies mocks base method.
func (m *MockPool) SetSupportedTopologies(arg0 []map[string]string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetSupportedTopologies", arg0)
}

// SetSupportedTopologies indicates an expected call of SetSupportedTopologies.
func (mr *MockPoolMockRecorder) SetSupportedTopologies(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetSupportedTopologies", reflect.TypeOf((*MockPool)(nil).SetSupportedTopologies), arg0)
}

// StorageClasses mocks base method.
func (m *MockPool) StorageClasses() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StorageClasses")
	ret0, _ := ret[0].([]string)
	return ret0
}

// StorageClasses indicates an expected call of StorageClasses.
func (mr *MockPoolMockRecorder) StorageClasses() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StorageClasses", reflect.TypeOf((*MockPool)(nil).StorageClasses))
}

// SupportedTopologies mocks base method.
func (m *MockPool) SupportedTopologies() []map[string]string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SupportedTopologies")
	ret0, _ := ret[0].([]map[string]string)
	return ret0
}

// SupportedTopologies indicates an expected call of SupportedTopologies.
func (mr *MockPoolMockRecorder) SupportedTopologies() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SupportedTopologies", reflect.TypeOf((*MockPool)(nil).SupportedTopologies))
}
