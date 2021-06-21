// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ozoncp/ocp-resume-api/internal/saver (interfaces: Saver)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	achievement "github.com/ozoncp/ocp-resume-api/internal/achievement"
	resume "github.com/ozoncp/ocp-resume-api/internal/resume"
)

// MockSaver is a mock of Saver interface.
type MockSaver struct {
	ctrl     *gomock.Controller
	recorder *MockSaverMockRecorder
}

// MockSaverMockRecorder is the mock recorder for MockSaver.
type MockSaverMockRecorder struct {
	mock *MockSaver
}

// NewMockSaver creates a new mock instance.
func NewMockSaver(ctrl *gomock.Controller) *MockSaver {
	mock := &MockSaver{ctrl: ctrl}
	mock.recorder = &MockSaverMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSaver) EXPECT() *MockSaverMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockSaver) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockSaverMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockSaver)(nil).Close))
}

// Init mocks base method.
func (m *MockSaver) Init(arg0 context.Context, arg1 int64, arg2 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Init", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Init indicates an expected call of Init.
func (mr *MockSaverMockRecorder) Init(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockSaver)(nil).Init), arg0, arg1, arg2)
}

// SaveAchievements mocks base method.
func (m *MockSaver) SaveAchievements(arg0 []achievement.Achievement) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveAchievements", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveAchievements indicates an expected call of SaveAchievements.
func (mr *MockSaverMockRecorder) SaveAchievements(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveAchievements", reflect.TypeOf((*MockSaver)(nil).SaveAchievements), arg0)
}

// SaveResumes mocks base method.
func (m *MockSaver) SaveResumes(arg0 []resume.Resume) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveResumes", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveResumes indicates an expected call of SaveResumes.
func (mr *MockSaverMockRecorder) SaveResumes(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveResumes", reflect.TypeOf((*MockSaver)(nil).SaveResumes), arg0)
}
