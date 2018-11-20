// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/carapace/core/pkg/v0/auth (interfaces: KeyMarshaller)

// Package mock is a generated GoMock package.
package mock

import (
	ecdsa "crypto/ecdsa"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockKeyMarshaller is a mock of KeyMarshaller interface
type MockKeyMarshaller struct {
	ctrl     *gomock.Controller
	recorder *MockKeyMarshallerMockRecorder
}

// MockKeyMarshallerMockRecorder is the mock recorder for MockKeyMarshaller
type MockKeyMarshallerMockRecorder struct {
	mock *MockKeyMarshaller
}

// NewMockKeyMarshaller creates a new mock instance
func NewMockKeyMarshaller(ctrl *gomock.Controller) *MockKeyMarshaller {
	mock := &MockKeyMarshaller{ctrl: ctrl}
	mock.recorder = &MockKeyMarshallerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockKeyMarshaller) EXPECT() *MockKeyMarshallerMockRecorder {
	return m.recorder
}

// MarshalPrivate mocks base method
func (m *MockKeyMarshaller) MarshalPrivate(arg0 *ecdsa.PrivateKey) ([]byte, error) {
	ret := m.ctrl.Call(m, "MarshalPrivate", arg0)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MarshalPrivate indicates an expected call of MarshalPrivate
func (mr *MockKeyMarshallerMockRecorder) MarshalPrivate(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarshalPrivate", reflect.TypeOf((*MockKeyMarshaller)(nil).MarshalPrivate), arg0)
}

// MarshalPublic mocks base method
func (m *MockKeyMarshaller) MarshalPublic(arg0 *ecdsa.PublicKey) ([]byte, error) {
	ret := m.ctrl.Call(m, "MarshalPublic", arg0)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MarshalPublic indicates an expected call of MarshalPublic
func (mr *MockKeyMarshallerMockRecorder) MarshalPublic(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarshalPublic", reflect.TypeOf((*MockKeyMarshaller)(nil).MarshalPublic), arg0)
}

// UnmarshalPrivate mocks base method
func (m *MockKeyMarshaller) UnmarshalPrivate(arg0 []byte) (*ecdsa.PrivateKey, error) {
	ret := m.ctrl.Call(m, "UnmarshalPrivate", arg0)
	ret0, _ := ret[0].(*ecdsa.PrivateKey)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UnmarshalPrivate indicates an expected call of UnmarshalPrivate
func (mr *MockKeyMarshallerMockRecorder) UnmarshalPrivate(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnmarshalPrivate", reflect.TypeOf((*MockKeyMarshaller)(nil).UnmarshalPrivate), arg0)
}

// UnmarshalPublic mocks base method
func (m *MockKeyMarshaller) UnmarshalPublic(arg0 []byte) (*ecdsa.PublicKey, error) {
	ret := m.ctrl.Call(m, "UnmarshalPublic", arg0)
	ret0, _ := ret[0].(*ecdsa.PublicKey)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UnmarshalPublic indicates an expected call of UnmarshalPublic
func (mr *MockKeyMarshallerMockRecorder) UnmarshalPublic(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnmarshalPublic", reflect.TypeOf((*MockKeyMarshaller)(nil).UnmarshalPublic), arg0)
}