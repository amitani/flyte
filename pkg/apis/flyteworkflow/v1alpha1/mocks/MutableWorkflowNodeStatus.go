// Code generated by mockery v1.0.1. DO NOT EDIT.

package mocks

import (
	core "github.com/lyft/flyteidl/gen/pb-go/flyteidl/core"
	mock "github.com/stretchr/testify/mock"

	v1alpha1 "github.com/lyft/flytepropeller/pkg/apis/flyteworkflow/v1alpha1"
)

// MutableWorkflowNodeStatus is an autogenerated mock type for the MutableWorkflowNodeStatus type
type MutableWorkflowNodeStatus struct {
	mock.Mock
}

type MutableWorkflowNodeStatus_GetExecutionError struct {
	*mock.Call
}

func (_m MutableWorkflowNodeStatus_GetExecutionError) Return(_a0 *core.ExecutionError) *MutableWorkflowNodeStatus_GetExecutionError {
	return &MutableWorkflowNodeStatus_GetExecutionError{Call: _m.Call.Return(_a0)}
}

func (_m *MutableWorkflowNodeStatus) OnGetExecutionError() *MutableWorkflowNodeStatus_GetExecutionError {
	c := _m.On("GetExecutionError")
	return &MutableWorkflowNodeStatus_GetExecutionError{Call: c}
}

func (_m *MutableWorkflowNodeStatus) OnGetExecutionErrorMatch(matchers ...interface{}) *MutableWorkflowNodeStatus_GetExecutionError {
	c := _m.On("GetExecutionError", matchers...)
	return &MutableWorkflowNodeStatus_GetExecutionError{Call: c}
}

// GetExecutionError provides a mock function with given fields:
func (_m *MutableWorkflowNodeStatus) GetExecutionError() *core.ExecutionError {
	ret := _m.Called()

	var r0 *core.ExecutionError
	if rf, ok := ret.Get(0).(func() *core.ExecutionError); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*core.ExecutionError)
		}
	}

	return r0
}

type MutableWorkflowNodeStatus_GetWorkflowNodePhase struct {
	*mock.Call
}

func (_m MutableWorkflowNodeStatus_GetWorkflowNodePhase) Return(_a0 v1alpha1.WorkflowNodePhase) *MutableWorkflowNodeStatus_GetWorkflowNodePhase {
	return &MutableWorkflowNodeStatus_GetWorkflowNodePhase{Call: _m.Call.Return(_a0)}
}

func (_m *MutableWorkflowNodeStatus) OnGetWorkflowNodePhase() *MutableWorkflowNodeStatus_GetWorkflowNodePhase {
	c := _m.On("GetWorkflowNodePhase")
	return &MutableWorkflowNodeStatus_GetWorkflowNodePhase{Call: c}
}

func (_m *MutableWorkflowNodeStatus) OnGetWorkflowNodePhaseMatch(matchers ...interface{}) *MutableWorkflowNodeStatus_GetWorkflowNodePhase {
	c := _m.On("GetWorkflowNodePhase", matchers...)
	return &MutableWorkflowNodeStatus_GetWorkflowNodePhase{Call: c}
}

// GetWorkflowNodePhase provides a mock function with given fields:
func (_m *MutableWorkflowNodeStatus) GetWorkflowNodePhase() v1alpha1.WorkflowNodePhase {
	ret := _m.Called()

	var r0 v1alpha1.WorkflowNodePhase
	if rf, ok := ret.Get(0).(func() v1alpha1.WorkflowNodePhase); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(v1alpha1.WorkflowNodePhase)
	}

	return r0
}

type MutableWorkflowNodeStatus_IsDirty struct {
	*mock.Call
}

func (_m MutableWorkflowNodeStatus_IsDirty) Return(_a0 bool) *MutableWorkflowNodeStatus_IsDirty {
	return &MutableWorkflowNodeStatus_IsDirty{Call: _m.Call.Return(_a0)}
}

func (_m *MutableWorkflowNodeStatus) OnIsDirty() *MutableWorkflowNodeStatus_IsDirty {
	c := _m.On("IsDirty")
	return &MutableWorkflowNodeStatus_IsDirty{Call: c}
}

func (_m *MutableWorkflowNodeStatus) OnIsDirtyMatch(matchers ...interface{}) *MutableWorkflowNodeStatus_IsDirty {
	c := _m.On("IsDirty", matchers...)
	return &MutableWorkflowNodeStatus_IsDirty{Call: c}
}

// IsDirty provides a mock function with given fields:
func (_m *MutableWorkflowNodeStatus) IsDirty() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// SetExecutionError provides a mock function with given fields: executionError
func (_m *MutableWorkflowNodeStatus) SetExecutionError(executionError *core.ExecutionError) {
	_m.Called(executionError)
}

// SetWorkflowNodePhase provides a mock function with given fields: phase
func (_m *MutableWorkflowNodeStatus) SetWorkflowNodePhase(phase v1alpha1.WorkflowNodePhase) {
	_m.Called(phase)
}
