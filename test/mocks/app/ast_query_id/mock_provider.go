// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/checkmarxDev/ast-sast-export/internal/app/interfaces (interfaces: ASTQueryIDProvider)

// Package mock_app_ast_query_id is a generated GoMock package.
package mock_app_ast_query_id

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockASTQueryIDProvider is a mock of ASTQueryIDProvider interface.
type MockASTQueryIDProvider struct {
	ctrl     *gomock.Controller
	recorder *MockASTQueryIDProviderMockRecorder
}

// MockASTQueryIDProviderMockRecorder is the mock recorder for MockASTQueryIDProvider.
type MockASTQueryIDProviderMockRecorder struct {
	mock *MockASTQueryIDProvider
}

// NewMockASTQueryIDProvider creates a new mock instance.
func NewMockASTQueryIDProvider(ctrl *gomock.Controller) *MockASTQueryIDProvider {
	mock := &MockASTQueryIDProvider{ctrl: ctrl}
	mock.recorder = &MockASTQueryIDProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockASTQueryIDProvider) EXPECT() *MockASTQueryIDProviderMockRecorder {
	return m.recorder
}

// GetQueryID mocks base method.
func (m *MockASTQueryIDProvider) GetQueryID(arg0, arg1, arg2 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetQueryID", arg0, arg1, arg2)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetQueryID indicates an expected call of GetQueryID.
func (mr *MockASTQueryIDProviderMockRecorder) GetQueryID(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetQueryID", reflect.TypeOf((*MockASTQueryIDProvider)(nil).GetQueryID), arg0, arg1, arg2)
}
