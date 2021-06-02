// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ozoncp/ocp-resume-api/internal/repo (interfaces: Repo)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	achievement "github.com/ozoncp/ocp-resume-api/internal/achievement"
	repo "github.com/ozoncp/ocp-resume-api/internal/repo"
	resume "github.com/ozoncp/ocp-resume-api/internal/resume"
)

// MockRepo is a mock of Repo interface.
type MockRepo struct {
	ctrl     *gomock.Controller
	recorder *MockRepoMockRecorder
}

// MockRepoMockRecorder is the mock recorder for MockRepo.
type MockRepoMockRecorder struct {
	mock *MockRepo
}

// NewMockRepo creates a new mock instance.
func NewMockRepo(ctrl *gomock.Controller) *MockRepo {
	mock := &MockRepo{ctrl: ctrl}
	mock.recorder = &MockRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepo) EXPECT() *MockRepoMockRecorder {
	return m.recorder
}

// AddAchievements mocks base method.
func (m *MockRepo) AddAchievements(arg0 []achievement.Achievement) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddAchievements", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddAchievements indicates an expected call of AddAchievements.
func (mr *MockRepoMockRecorder) AddAchievements(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddAchievements", reflect.TypeOf((*MockRepo)(nil).AddAchievements), arg0)
}

// AddResumes mocks base method.
func (m *MockRepo) AddResumes(arg0 []resume.Resume) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddResumes", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddResumes indicates an expected call of AddResumes.
func (mr *MockRepoMockRecorder) AddResumes(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddResumes", reflect.TypeOf((*MockRepo)(nil).AddResumes), arg0)
}

// GetAchievementById mocks base method.
func (m *MockRepo) GetAchievementById(arg0 uint) (*achievement.Achievement, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAchievementById", arg0)
	ret0, _ := ret[0].(*achievement.Achievement)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAchievementById indicates an expected call of GetAchievementById.
func (mr *MockRepoMockRecorder) GetAchievementById(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAchievementById", reflect.TypeOf((*MockRepo)(nil).GetAchievementById), arg0)
}

// GetAchievementByNdx mocks base method.
func (m *MockRepo) GetAchievementByNdx(arg0 uint64) (*achievement.Achievement, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAchievementByNdx", arg0)
	ret0, _ := ret[0].(*achievement.Achievement)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAchievementByNdx indicates an expected call of GetAchievementByNdx.
func (mr *MockRepoMockRecorder) GetAchievementByNdx(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAchievementByNdx", reflect.TypeOf((*MockRepo)(nil).GetAchievementByNdx), arg0)
}

// GetResumeById mocks base method.
func (m *MockRepo) GetResumeById(arg0 uint) (*resume.Resume, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetResumeById", arg0)
	ret0, _ := ret[0].(*resume.Resume)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetResumeById indicates an expected call of GetResumeById.
func (mr *MockRepoMockRecorder) GetResumeById(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetResumeById", reflect.TypeOf((*MockRepo)(nil).GetResumeById), arg0)
}

// GetResumeByNdx mocks base method.
func (m *MockRepo) GetResumeByNdx(arg0 uint64) (*resume.Resume, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetResumeByNdx", arg0)
	ret0, _ := ret[0].(*resume.Resume)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetResumeByNdx indicates an expected call of GetResumeByNdx.
func (mr *MockRepoMockRecorder) GetResumeByNdx(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetResumeByNdx", reflect.TypeOf((*MockRepo)(nil).GetResumeByNdx), arg0)
}

// ListAchievements mocks base method.
func (m *MockRepo) ListAchievements(arg0, arg1 uint64) ([]repo.RepoAchievement, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAchievements", arg0, arg1)
	ret0, _ := ret[0].([]repo.RepoAchievement)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAchievements indicates an expected call of ListAchievements.
func (mr *MockRepoMockRecorder) ListAchievements(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAchievements", reflect.TypeOf((*MockRepo)(nil).ListAchievements), arg0, arg1)
}

// ListResumes mocks base method.
func (m *MockRepo) ListResumes(arg0, arg1 uint64) ([]repo.RepoResume, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListResumes", arg0, arg1)
	ret0, _ := ret[0].([]repo.RepoResume)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListResumes indicates an expected call of ListResumes.
func (mr *MockRepoMockRecorder) ListResumes(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListResumes", reflect.TypeOf((*MockRepo)(nil).ListResumes), arg0, arg1)
}

// RemoveAchievementById mocks base method.
func (m *MockRepo) RemoveAchievementById(arg0 uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveAchievementById", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveAchievementById indicates an expected call of RemoveAchievementById.
func (mr *MockRepoMockRecorder) RemoveAchievementById(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveAchievementById", reflect.TypeOf((*MockRepo)(nil).RemoveAchievementById), arg0)
}

// RemoveAchievementByNdx mocks base method.
func (m *MockRepo) RemoveAchievementByNdx(arg0 uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveAchievementByNdx", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveAchievementByNdx indicates an expected call of RemoveAchievementByNdx.
func (mr *MockRepoMockRecorder) RemoveAchievementByNdx(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveAchievementByNdx", reflect.TypeOf((*MockRepo)(nil).RemoveAchievementByNdx), arg0)
}

// RemoveResumeById mocks base method.
func (m *MockRepo) RemoveResumeById(arg0 uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveResumeById", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveResumeById indicates an expected call of RemoveResumeById.
func (mr *MockRepoMockRecorder) RemoveResumeById(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveResumeById", reflect.TypeOf((*MockRepo)(nil).RemoveResumeById), arg0)
}

// RemoveResumeByNdx mocks base method.
func (m *MockRepo) RemoveResumeByNdx(arg0 uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveResumeByNdx", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveResumeByNdx indicates an expected call of RemoveResumeByNdx.
func (mr *MockRepoMockRecorder) RemoveResumeByNdx(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveResumeByNdx", reflect.TypeOf((*MockRepo)(nil).RemoveResumeByNdx), arg0)
}
