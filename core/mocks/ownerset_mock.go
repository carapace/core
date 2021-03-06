// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/carapace/core/core (interfaces: OwnerSet)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	sql "database/sql"
	proto "github.com/carapace/core/api/v0/proto"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockOwnerSet is a mock of OwnerSet interface
type MockOwnerSet struct {
	ctrl     *gomock.Controller
	recorder *MockOwnerSetMockRecorder
}

// MockOwnerSetMockRecorder is the mock recorder for MockOwnerSet
type MockOwnerSetMockRecorder struct {
	mock *MockOwnerSet
}

// NewMockOwnerSet creates a new mock instance
func NewMockOwnerSet(ctrl *gomock.Controller) *MockOwnerSet {
	mock := &MockOwnerSet{ctrl: ctrl}
	mock.recorder = &MockOwnerSetMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOwnerSet) EXPECT() *MockOwnerSetMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *MockOwnerSet) Get(arg0 context.Context, arg1 *sql.Tx) (*proto.OwnerSet, error) {
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(*proto.OwnerSet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockOwnerSetMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockOwnerSet)(nil).Get), arg0, arg1)
}

// Put mocks base method
func (m *MockOwnerSet) Put(arg0 context.Context, arg1 *sql.Tx, arg2 *proto.OwnerSet) error {
	ret := m.ctrl.Call(m, "Put", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Put indicates an expected call of Put
func (mr *MockOwnerSetMockRecorder) Put(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Put", reflect.TypeOf((*MockOwnerSet)(nil).Put), arg0, arg1, arg2)
}
