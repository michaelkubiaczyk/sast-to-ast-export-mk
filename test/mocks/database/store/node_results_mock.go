// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/checkmarxDev/ast-sast-export/internal/database/store (interfaces: NodeResultsStore)

// Package mock_store is a generated GoMock package.
package mock_store

import (
	reflect "reflect"

	database "github.com/checkmarxDev/ast-sast-export/internal/database"
	gomock "github.com/golang/mock/gomock"
)

// MockNodeResultsStore is a mock of NodeResultsStore interface.
type MockNodeResultsStore struct {
	ctrl     *gomock.Controller
	recorder *MockNodeResultsStoreMockRecorder
}

// MockNodeResultsStoreMockRecorder is the mock recorder for MockNodeResultsStore.
type MockNodeResultsStoreMockRecorder struct {
	mock *MockNodeResultsStore
}

// NewMockNodeResultsStore creates a new mock instance.
func NewMockNodeResultsStore(ctrl *gomock.Controller) *MockNodeResultsStore {
	mock := &MockNodeResultsStore{ctrl: ctrl}
	mock.recorder = &MockNodeResultsStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNodeResultsStore) EXPECT() *MockNodeResultsStoreMockRecorder {
	return m.recorder
}

// GetByResultPathAndNode mocks base method.
func (m *MockNodeResultsStore) GetByResultPathAndNode(arg0, arg1 string, arg2 int) (*database.NodeResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByResultPathAndNode", arg0, arg1, arg2)
	ret0, _ := ret[0].(*database.NodeResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByResultPathAndNode indicates an expected call of GetByResultPathAndNode.
func (mr *MockNodeResultsStoreMockRecorder) GetByResultPathAndNode(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByResultPathAndNode", reflect.TypeOf((*MockNodeResultsStore)(nil).GetByResultPathAndNode), arg0, arg1, arg2)
}
