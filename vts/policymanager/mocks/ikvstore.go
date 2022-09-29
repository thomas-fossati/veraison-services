// Code generated by MockGen. DO NOT EDIT.
// Source: ../../kvstore/ikvstore.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	viper "github.com/spf13/viper"
)

// MockIKVStore is a mock of IKVStore interface.
type MockIKVStore struct {
	ctrl     *gomock.Controller
	recorder *MockIKVStoreMockRecorder
}

// MockIKVStoreMockRecorder is the mock recorder for MockIKVStore.
type MockIKVStoreMockRecorder struct {
	mock *MockIKVStore
}

// NewMockIKVStore creates a new mock instance.
func NewMockIKVStore(ctrl *gomock.Controller) *MockIKVStore {
	mock := &MockIKVStore{ctrl: ctrl}
	mock.recorder = &MockIKVStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIKVStore) EXPECT() *MockIKVStoreMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockIKVStore) Add(key, val string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", key, val)
	ret0, _ := ret[0].(error)
	return ret0
}

// Add indicates an expected call of Add.
func (mr *MockIKVStoreMockRecorder) Add(key, val interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockIKVStore)(nil).Add), key, val)
}

// Close mocks base method.
func (m *MockIKVStore) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockIKVStoreMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockIKVStore)(nil).Close))
}

// Del mocks base method.
func (m *MockIKVStore) Del(key string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Del", key)
	ret0, _ := ret[0].(error)
	return ret0
}

// Del indicates an expected call of Del.
func (mr *MockIKVStoreMockRecorder) Del(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Del", reflect.TypeOf((*MockIKVStore)(nil).Del), key)
}

// Get mocks base method.
func (m *MockIKVStore) Get(key string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", key)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockIKVStoreMockRecorder) Get(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockIKVStore)(nil).Get), key)
}

// Init mocks base method.
func (m *MockIKVStore) Init(v *viper.Viper) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Init", v)
	ret0, _ := ret[0].(error)
	return ret0
}

// Init indicates an expected call of Init.
func (mr *MockIKVStoreMockRecorder) Init(v interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockIKVStore)(nil).Init), v)
}

// Set mocks base method.
func (m *MockIKVStore) Set(key, val string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set", key, val)
	ret0, _ := ret[0].(error)
	return ret0
}

// Set indicates an expected call of Set.
func (mr *MockIKVStoreMockRecorder) Set(key, val interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockIKVStore)(nil).Set), key, val)
}

// Setup mocks base method.
func (m *MockIKVStore) Setup() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Setup")
	ret0, _ := ret[0].(error)
	return ret0
}

// Setup indicates an expected call of Setup.
func (mr *MockIKVStoreMockRecorder) Setup() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Setup", reflect.TypeOf((*MockIKVStore)(nil).Setup))
}
