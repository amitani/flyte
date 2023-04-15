// Code generated by mockery v1.0.1. DO NOT EDIT.

package mocks

import (
	storage "github.com/flyteorg/flytestdlib/storage"
	mock "github.com/stretchr/testify/mock"

	time "time"
)

// MutableTaskNodeStatus is an autogenerated mock type for the MutableTaskNodeStatus type
type MutableTaskNodeStatus struct {
	mock.Mock
}

type MutableTaskNodeStatus_GetBarrierClockTick struct {
	*mock.Call
}

func (_m MutableTaskNodeStatus_GetBarrierClockTick) Return(_a0 uint32) *MutableTaskNodeStatus_GetBarrierClockTick {
	return &MutableTaskNodeStatus_GetBarrierClockTick{Call: _m.Call.Return(_a0)}
}

func (_m *MutableTaskNodeStatus) OnGetBarrierClockTick() *MutableTaskNodeStatus_GetBarrierClockTick {
	c := _m.On("GetBarrierClockTick")
	return &MutableTaskNodeStatus_GetBarrierClockTick{Call: c}
}

func (_m *MutableTaskNodeStatus) OnGetBarrierClockTickMatch(matchers ...interface{}) *MutableTaskNodeStatus_GetBarrierClockTick {
	c := _m.On("GetBarrierClockTick", matchers...)
	return &MutableTaskNodeStatus_GetBarrierClockTick{Call: c}
}

// GetBarrierClockTick provides a mock function with given fields:
func (_m *MutableTaskNodeStatus) GetBarrierClockTick() uint32 {
	ret := _m.Called()

	var r0 uint32
	if rf, ok := ret.Get(0).(func() uint32); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint32)
	}

	return r0
}

type MutableTaskNodeStatus_GetLastPhaseUpdatedAt struct {
	*mock.Call
}

func (_m MutableTaskNodeStatus_GetLastPhaseUpdatedAt) Return(_a0 time.Time) *MutableTaskNodeStatus_GetLastPhaseUpdatedAt {
	return &MutableTaskNodeStatus_GetLastPhaseUpdatedAt{Call: _m.Call.Return(_a0)}
}

func (_m *MutableTaskNodeStatus) OnGetLastPhaseUpdatedAt() *MutableTaskNodeStatus_GetLastPhaseUpdatedAt {
	c := _m.On("GetLastPhaseUpdatedAt")
	return &MutableTaskNodeStatus_GetLastPhaseUpdatedAt{Call: c}
}

func (_m *MutableTaskNodeStatus) OnGetLastPhaseUpdatedAtMatch(matchers ...interface{}) *MutableTaskNodeStatus_GetLastPhaseUpdatedAt {
	c := _m.On("GetLastPhaseUpdatedAt", matchers...)
	return &MutableTaskNodeStatus_GetLastPhaseUpdatedAt{Call: c}
}

// GetLastPhaseUpdatedAt provides a mock function with given fields:
func (_m *MutableTaskNodeStatus) GetLastPhaseUpdatedAt() time.Time {
	ret := _m.Called()

	var r0 time.Time
	if rf, ok := ret.Get(0).(func() time.Time); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(time.Time)
	}

	return r0
}

type MutableTaskNodeStatus_GetPhase struct {
	*mock.Call
}

func (_m MutableTaskNodeStatus_GetPhase) Return(_a0 int) *MutableTaskNodeStatus_GetPhase {
	return &MutableTaskNodeStatus_GetPhase{Call: _m.Call.Return(_a0)}
}

func (_m *MutableTaskNodeStatus) OnGetPhase() *MutableTaskNodeStatus_GetPhase {
	c := _m.On("GetPhase")
	return &MutableTaskNodeStatus_GetPhase{Call: c}
}

func (_m *MutableTaskNodeStatus) OnGetPhaseMatch(matchers ...interface{}) *MutableTaskNodeStatus_GetPhase {
	c := _m.On("GetPhase", matchers...)
	return &MutableTaskNodeStatus_GetPhase{Call: c}
}

// GetPhase provides a mock function with given fields:
func (_m *MutableTaskNodeStatus) GetPhase() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

type MutableTaskNodeStatus_GetPhaseVersion struct {
	*mock.Call
}

func (_m MutableTaskNodeStatus_GetPhaseVersion) Return(_a0 uint32) *MutableTaskNodeStatus_GetPhaseVersion {
	return &MutableTaskNodeStatus_GetPhaseVersion{Call: _m.Call.Return(_a0)}
}

func (_m *MutableTaskNodeStatus) OnGetPhaseVersion() *MutableTaskNodeStatus_GetPhaseVersion {
	c := _m.On("GetPhaseVersion")
	return &MutableTaskNodeStatus_GetPhaseVersion{Call: c}
}

func (_m *MutableTaskNodeStatus) OnGetPhaseVersionMatch(matchers ...interface{}) *MutableTaskNodeStatus_GetPhaseVersion {
	c := _m.On("GetPhaseVersion", matchers...)
	return &MutableTaskNodeStatus_GetPhaseVersion{Call: c}
}

// GetPhaseVersion provides a mock function with given fields:
func (_m *MutableTaskNodeStatus) GetPhaseVersion() uint32 {
	ret := _m.Called()

	var r0 uint32
	if rf, ok := ret.Get(0).(func() uint32); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint32)
	}

	return r0
}

type MutableTaskNodeStatus_GetPluginState struct {
	*mock.Call
}

func (_m MutableTaskNodeStatus_GetPluginState) Return(_a0 []byte) *MutableTaskNodeStatus_GetPluginState {
	return &MutableTaskNodeStatus_GetPluginState{Call: _m.Call.Return(_a0)}
}

func (_m *MutableTaskNodeStatus) OnGetPluginState() *MutableTaskNodeStatus_GetPluginState {
	c := _m.On("GetPluginState")
	return &MutableTaskNodeStatus_GetPluginState{Call: c}
}

func (_m *MutableTaskNodeStatus) OnGetPluginStateMatch(matchers ...interface{}) *MutableTaskNodeStatus_GetPluginState {
	c := _m.On("GetPluginState", matchers...)
	return &MutableTaskNodeStatus_GetPluginState{Call: c}
}

// GetPluginState provides a mock function with given fields:
func (_m *MutableTaskNodeStatus) GetPluginState() []byte {
	ret := _m.Called()

	var r0 []byte
	if rf, ok := ret.Get(0).(func() []byte); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	return r0
}

type MutableTaskNodeStatus_GetPluginStateVersion struct {
	*mock.Call
}

func (_m MutableTaskNodeStatus_GetPluginStateVersion) Return(_a0 uint32) *MutableTaskNodeStatus_GetPluginStateVersion {
	return &MutableTaskNodeStatus_GetPluginStateVersion{Call: _m.Call.Return(_a0)}
}

func (_m *MutableTaskNodeStatus) OnGetPluginStateVersion() *MutableTaskNodeStatus_GetPluginStateVersion {
	c := _m.On("GetPluginStateVersion")
	return &MutableTaskNodeStatus_GetPluginStateVersion{Call: c}
}

func (_m *MutableTaskNodeStatus) OnGetPluginStateVersionMatch(matchers ...interface{}) *MutableTaskNodeStatus_GetPluginStateVersion {
	c := _m.On("GetPluginStateVersion", matchers...)
	return &MutableTaskNodeStatus_GetPluginStateVersion{Call: c}
}

// GetPluginStateVersion provides a mock function with given fields:
func (_m *MutableTaskNodeStatus) GetPluginStateVersion() uint32 {
	ret := _m.Called()

	var r0 uint32
	if rf, ok := ret.Get(0).(func() uint32); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint32)
	}

	return r0
}

type MutableTaskNodeStatus_GetPreviousNodeExecutionCheckpointPath struct {
	*mock.Call
}

func (_m MutableTaskNodeStatus_GetPreviousNodeExecutionCheckpointPath) Return(_a0 storage.DataReference) *MutableTaskNodeStatus_GetPreviousNodeExecutionCheckpointPath {
	return &MutableTaskNodeStatus_GetPreviousNodeExecutionCheckpointPath{Call: _m.Call.Return(_a0)}
}

func (_m *MutableTaskNodeStatus) OnGetPreviousNodeExecutionCheckpointPath() *MutableTaskNodeStatus_GetPreviousNodeExecutionCheckpointPath {
	c := _m.On("GetPreviousNodeExecutionCheckpointPath")
	return &MutableTaskNodeStatus_GetPreviousNodeExecutionCheckpointPath{Call: c}
}

func (_m *MutableTaskNodeStatus) OnGetPreviousNodeExecutionCheckpointPathMatch(matchers ...interface{}) *MutableTaskNodeStatus_GetPreviousNodeExecutionCheckpointPath {
	c := _m.On("GetPreviousNodeExecutionCheckpointPath", matchers...)
	return &MutableTaskNodeStatus_GetPreviousNodeExecutionCheckpointPath{Call: c}
}

// GetPreviousNodeExecutionCheckpointPath provides a mock function with given fields:
func (_m *MutableTaskNodeStatus) GetPreviousNodeExecutionCheckpointPath() storage.DataReference {
	ret := _m.Called()

	var r0 storage.DataReference
	if rf, ok := ret.Get(0).(func() storage.DataReference); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(storage.DataReference)
	}

	return r0
}

type MutableTaskNodeStatus_IsDirty struct {
	*mock.Call
}

func (_m MutableTaskNodeStatus_IsDirty) Return(_a0 bool) *MutableTaskNodeStatus_IsDirty {
	return &MutableTaskNodeStatus_IsDirty{Call: _m.Call.Return(_a0)}
}

func (_m *MutableTaskNodeStatus) OnIsDirty() *MutableTaskNodeStatus_IsDirty {
	c := _m.On("IsDirty")
	return &MutableTaskNodeStatus_IsDirty{Call: c}
}

func (_m *MutableTaskNodeStatus) OnIsDirtyMatch(matchers ...interface{}) *MutableTaskNodeStatus_IsDirty {
	c := _m.On("IsDirty", matchers...)
	return &MutableTaskNodeStatus_IsDirty{Call: c}
}

// IsDirty provides a mock function with given fields:
func (_m *MutableTaskNodeStatus) IsDirty() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// SetBarrierClockTick provides a mock function with given fields: tick
func (_m *MutableTaskNodeStatus) SetBarrierClockTick(tick uint32) {
	_m.Called(tick)
}

// SetLastPhaseUpdatedAt provides a mock function with given fields: updatedAt
func (_m *MutableTaskNodeStatus) SetLastPhaseUpdatedAt(updatedAt time.Time) {
	_m.Called(updatedAt)
}

// SetPhase provides a mock function with given fields: phase
func (_m *MutableTaskNodeStatus) SetPhase(phase int) {
	_m.Called(phase)
}

// SetPhaseVersion provides a mock function with given fields: version
func (_m *MutableTaskNodeStatus) SetPhaseVersion(version uint32) {
	_m.Called(version)
}

// SetPluginState provides a mock function with given fields: _a0
func (_m *MutableTaskNodeStatus) SetPluginState(_a0 []byte) {
	_m.Called(_a0)
}

// SetPluginStateVersion provides a mock function with given fields: _a0
func (_m *MutableTaskNodeStatus) SetPluginStateVersion(_a0 uint32) {
	_m.Called(_a0)
}

// SetPreviousNodeExecutionCheckpointPath provides a mock function with given fields: _a0
func (_m *MutableTaskNodeStatus) SetPreviousNodeExecutionCheckpointPath(_a0 storage.DataReference) {
	_m.Called(_a0)
}
