// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/openfeature/hooks.go

// Package openfeature is a generated GoMock package.
package openfeature

import (
	gomock "github.com/golang/mock/gomock"
)

// MockHook is a mock of Hook interface.
type MockHook struct {
	ctrl     *gomock.Controller
	recorder *MockHookMockRecorder
}

// MockHookMockRecorder is the mock recorder for MockHook.
type MockHookMockRecorder struct {
	mock *MockHook
}

// NewMockHook creates a new mock instance.
func NewMockHook(ctrl *gomock.Controller) *MockHook {
	mock := &MockHook{ctrl: ctrl}
	mock.recorder = &MockHookMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHook) EXPECT() *MockHookMockRecorder {
	return m.recorder
}