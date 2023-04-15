// Code generated by mockery v1.0.1. DO NOT EDIT.

package mocks

import (
	v1alpha1 "github.com/flyteorg/flytepropeller/pkg/apis/flyteworkflow/v1alpha1"
	mock "github.com/stretchr/testify/mock"
)

// BaseWorkflow is an autogenerated mock type for the BaseWorkflow type
type BaseWorkflow struct {
	mock.Mock
}

type BaseWorkflow_FromNode struct {
	*mock.Call
}

func (_m BaseWorkflow_FromNode) Return(_a0 []string, _a1 error) *BaseWorkflow_FromNode {
	return &BaseWorkflow_FromNode{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *BaseWorkflow) OnFromNode(name string) *BaseWorkflow_FromNode {
	c := _m.On("FromNode", name)
	return &BaseWorkflow_FromNode{Call: c}
}

func (_m *BaseWorkflow) OnFromNodeMatch(matchers ...interface{}) *BaseWorkflow_FromNode {
	c := _m.On("FromNode", matchers...)
	return &BaseWorkflow_FromNode{Call: c}
}

// FromNode provides a mock function with given fields: name
func (_m *BaseWorkflow) FromNode(name string) ([]string, error) {
	ret := _m.Called(name)

	var r0 []string
	if rf, ok := ret.Get(0).(func(string) []string); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type BaseWorkflow_GetID struct {
	*mock.Call
}

func (_m BaseWorkflow_GetID) Return(_a0 string) *BaseWorkflow_GetID {
	return &BaseWorkflow_GetID{Call: _m.Call.Return(_a0)}
}

func (_m *BaseWorkflow) OnGetID() *BaseWorkflow_GetID {
	c := _m.On("GetID")
	return &BaseWorkflow_GetID{Call: c}
}

func (_m *BaseWorkflow) OnGetIDMatch(matchers ...interface{}) *BaseWorkflow_GetID {
	c := _m.On("GetID", matchers...)
	return &BaseWorkflow_GetID{Call: c}
}

// GetID provides a mock function with given fields:
func (_m *BaseWorkflow) GetID() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

type BaseWorkflow_GetNode struct {
	*mock.Call
}

func (_m BaseWorkflow_GetNode) Return(_a0 v1alpha1.ExecutableNode, _a1 bool) *BaseWorkflow_GetNode {
	return &BaseWorkflow_GetNode{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *BaseWorkflow) OnGetNode(nodeID string) *BaseWorkflow_GetNode {
	c := _m.On("GetNode", nodeID)
	return &BaseWorkflow_GetNode{Call: c}
}

func (_m *BaseWorkflow) OnGetNodeMatch(matchers ...interface{}) *BaseWorkflow_GetNode {
	c := _m.On("GetNode", matchers...)
	return &BaseWorkflow_GetNode{Call: c}
}

// GetNode provides a mock function with given fields: nodeID
func (_m *BaseWorkflow) GetNode(nodeID string) (v1alpha1.ExecutableNode, bool) {
	ret := _m.Called(nodeID)

	var r0 v1alpha1.ExecutableNode
	if rf, ok := ret.Get(0).(func(string) v1alpha1.ExecutableNode); ok {
		r0 = rf(nodeID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1alpha1.ExecutableNode)
		}
	}

	var r1 bool
	if rf, ok := ret.Get(1).(func(string) bool); ok {
		r1 = rf(nodeID)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

type BaseWorkflow_StartNode struct {
	*mock.Call
}

func (_m BaseWorkflow_StartNode) Return(_a0 v1alpha1.ExecutableNode) *BaseWorkflow_StartNode {
	return &BaseWorkflow_StartNode{Call: _m.Call.Return(_a0)}
}

func (_m *BaseWorkflow) OnStartNode() *BaseWorkflow_StartNode {
	c := _m.On("StartNode")
	return &BaseWorkflow_StartNode{Call: c}
}

func (_m *BaseWorkflow) OnStartNodeMatch(matchers ...interface{}) *BaseWorkflow_StartNode {
	c := _m.On("StartNode", matchers...)
	return &BaseWorkflow_StartNode{Call: c}
}

// StartNode provides a mock function with given fields:
func (_m *BaseWorkflow) StartNode() v1alpha1.ExecutableNode {
	ret := _m.Called()

	var r0 v1alpha1.ExecutableNode
	if rf, ok := ret.Get(0).(func() v1alpha1.ExecutableNode); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v1alpha1.ExecutableNode)
		}
	}

	return r0
}

type BaseWorkflow_ToNode struct {
	*mock.Call
}

func (_m BaseWorkflow_ToNode) Return(_a0 []string, _a1 error) *BaseWorkflow_ToNode {
	return &BaseWorkflow_ToNode{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *BaseWorkflow) OnToNode(name string) *BaseWorkflow_ToNode {
	c := _m.On("ToNode", name)
	return &BaseWorkflow_ToNode{Call: c}
}

func (_m *BaseWorkflow) OnToNodeMatch(matchers ...interface{}) *BaseWorkflow_ToNode {
	c := _m.On("ToNode", matchers...)
	return &BaseWorkflow_ToNode{Call: c}
}

// ToNode provides a mock function with given fields: name
func (_m *BaseWorkflow) ToNode(name string) ([]string, error) {
	ret := _m.Called(name)

	var r0 []string
	if rf, ok := ret.Get(0).(func(string) []string); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
