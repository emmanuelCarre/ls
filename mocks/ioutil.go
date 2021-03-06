// Code generated by MockGen. DO NOT EDIT.
// Source: ./io/ioutil/ioutil.go

// Package mock_ioutil is a generated GoMock package.
package mock_ioutil

import (
	gomock "github.com/golang/mock/gomock"
	os "os"
	reflect "reflect"
)

// MockIoUtil is a mock of IoUtil interface
type MockIoUtil struct {
	ctrl     *gomock.Controller
	recorder *MockIoUtilMockRecorder
}

// MockIoUtilMockRecorder is the mock recorder for MockIoUtil
type MockIoUtilMockRecorder struct {
	mock *MockIoUtil
}

// NewMockIoUtil creates a new mock instance
func NewMockIoUtil(ctrl *gomock.Controller) *MockIoUtil {
	mock := &MockIoUtil{ctrl: ctrl}
	mock.recorder = &MockIoUtilMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIoUtil) EXPECT() *MockIoUtilMockRecorder {
	return m.recorder
}

// ReadDir mocks base method
func (m *MockIoUtil) ReadDir(dirname string) ([]os.FileInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadDir", dirname)
	ret0, _ := ret[0].([]os.FileInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadDir indicates an expected call of ReadDir
func (mr *MockIoUtilMockRecorder) ReadDir(dirname interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadDir", reflect.TypeOf((*MockIoUtil)(nil).ReadDir), dirname)
}
