// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/carapace/core/core (interfaces: Authorizer)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	proto "github.com/carapace/core/api/v0/proto"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockAuthorizer is a mock of Authorizer interface
type MockAuthorizer struct {
	ctrl     *gomock.Controller
	recorder *MockAuthorizerMockRecorder
}

// MockAuthorizerMockRecorder is the mock recorder for MockAuthorizer
type MockAuthorizerMockRecorder struct {
	mock *MockAuthorizer
}

// NewMockAuthorizer creates a new mock instance
func NewMockAuthorizer(ctrl *gomock.Controller) *MockAuthorizer {
	mock := &MockAuthorizer{ctrl: ctrl}
	mock.recorder = &MockAuthorizerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAuthorizer) EXPECT() *MockAuthorizerMockRecorder {
	return m.recorder
}

// GetOwners mocks base method
func (m *MockAuthorizer) GetOwners(arg0 context.Context) (*proto.OwnerSet, error) {
	ret := m.ctrl.Call(m, "GetOwners", arg0)
	ret0, _ := ret[0].(*proto.OwnerSet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOwners indicates an expected call of GetOwners
func (mr *MockAuthorizerMockRecorder) GetOwners(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOwners", reflect.TypeOf((*MockAuthorizer)(nil).GetOwners), arg0)
}

// GrantBackupRoot mocks base method
func (m *MockAuthorizer) GrantBackupRoot(arg0 context.Context, arg1 *proto.Witness) (bool, error) {
	ret := m.ctrl.Call(m, "GrantBackupRoot", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GrantBackupRoot indicates an expected call of GrantBackupRoot
func (mr *MockAuthorizerMockRecorder) GrantBackupRoot(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GrantBackupRoot", reflect.TypeOf((*MockAuthorizer)(nil).GrantBackupRoot), arg0, arg1)
}

// GrantRoot mocks base method
func (m *MockAuthorizer) GrantRoot(arg0 context.Context, arg1 *proto.Witness) (bool, error) {
	ret := m.ctrl.Call(m, "GrantRoot", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GrantRoot indicates an expected call of GrantRoot
func (mr *MockAuthorizerMockRecorder) GrantRoot(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GrantRoot", reflect.TypeOf((*MockAuthorizer)(nil).GrantRoot), arg0, arg1)
}

// HaveOwners mocks base method
func (m *MockAuthorizer) HaveOwners(arg0 context.Context) (bool, error) {
	ret := m.ctrl.Call(m, "HaveOwners", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HaveOwners indicates an expected call of HaveOwners
func (mr *MockAuthorizerMockRecorder) HaveOwners(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HaveOwners", reflect.TypeOf((*MockAuthorizer)(nil).HaveOwners), arg0)
}

// SetOwners mocks base method
func (m *MockAuthorizer) SetOwners(arg0 context.Context, arg1 *proto.OwnerSet) error {
	ret := m.ctrl.Call(m, "SetOwners", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetOwners indicates an expected call of SetOwners
func (mr *MockAuthorizerMockRecorder) SetOwners(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetOwners", reflect.TypeOf((*MockAuthorizer)(nil).SetOwners), arg0, arg1)
}
