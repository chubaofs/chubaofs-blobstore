// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cubefs/blobstore/common/raftserver (interfaces: KVStorage)

// Package raftserver is a generated GoMock package.
package raftserver

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockKVStorage is a mock of KVStorage interface.
type MockKVStorage struct {
	ctrl     *gomock.Controller
	recorder *MockKVStorageMockRecorder
}

// MockKVStorageMockRecorder is the mock recorder for MockKVStorage.
type MockKVStorageMockRecorder struct {
	mock *MockKVStorage
}

// NewMockKVStorage creates a new mock instance.
func NewMockKVStorage(ctrl *gomock.Controller) *MockKVStorage {
	mock := &MockKVStorage{ctrl: ctrl}
	mock.recorder = &MockKVStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockKVStorage) EXPECT() *MockKVStorageMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockKVStorage) Get(arg0 []byte) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockKVStorageMockRecorder) Get(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockKVStorage)(nil).Get), arg0)
}

// Put mocks base method.
func (m *MockKVStorage) Put(arg0, arg1 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Put", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Put indicates an expected call of Put.
func (mr *MockKVStorageMockRecorder) Put(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Put", reflect.TypeOf((*MockKVStorage)(nil).Put), arg0, arg1)
}
