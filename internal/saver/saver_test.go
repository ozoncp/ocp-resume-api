package saver_test

import (
	"errors"
	"time"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/ozoncp/ocp-resume-api/internal/achievement"
	"github.com/ozoncp/ocp-resume-api/internal/mocks"
	"github.com/ozoncp/ocp-resume-api/internal/resume"
	"github.com/ozoncp/ocp-resume-api/internal/saver"
)

var _ = Describe("Saver", func() {
	var (
		err              error
		ctrl             *gomock.Controller
		mockFlusher      *mocks.MockFlusher
		s                saver.Saver
		achievements     []achievement.Achievement
		resumes          []resume.Resume
		achievements_cap int
		resumes_cap      int
		timeout          int64
		wait_time        time.Duration
		smart_del        bool
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		mockFlusher = mocks.NewMockFlusher(ctrl)
		smart_del = true
		achievements_cap = 512
		resumes_cap = 128
		achievements = make([]achievement.Achievement, 256)
		resumes = make([]resume.Resume, 32)
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	It("return error for timeout <=0", func() {
		//Expect(err).Should(BeNil())
		timeout = -1
		s = saver.NewSaver(mockFlusher, achievements_cap, resumes_cap)
		err = s.Init(timeout, smart_del)
		Expect(err).ShouldNot(BeNil())
	})

	When("close before timeout", func() {
		BeforeEach(func() {
			timeout = 1000 * int64(time.Millisecond)
			wait_time = 2 * time.Millisecond
			mockFlusher.EXPECT().FlushAchievements(gomock.Any()).Return([]achievement.Achievement{}, nil).Times(1)
			mockFlusher.EXPECT().FlushResumes(gomock.Any()).Return([]resume.Resume{}, nil).Times(1)
		})
		It("flusher takes 1 Call", func() {
			s = saver.NewSaver(mockFlusher, achievements_cap, resumes_cap)
			err = s.Init(timeout, smart_del)
			s.SaveAchievements(achievements)
			s.SaveResumes(resumes)
			time.Sleep(wait_time)
			err = s.Close()
		})
	})

	When("close after timeout", func() {
		BeforeEach(func() {
			timeout = 100 * int64(time.Millisecond)
			wait_time = 350 * time.Millisecond
			mockFlusher.EXPECT().FlushAchievements(gomock.Any()).Return([]achievement.Achievement{}, nil).Times(4)
			mockFlusher.EXPECT().FlushResumes(gomock.Any()).Return([]resume.Resume{}, nil).Times(4)
		})
		It("flusher takes 4 Call", func() {
			s = saver.NewSaver(mockFlusher, achievements_cap, resumes_cap)
			err = s.Init(timeout, smart_del)
			s.SaveAchievements(achievements)
			s.SaveResumes(resumes)
			time.Sleep(wait_time)
			err = s.Close()
		})
	})

	When("flush with errors", func() {
		BeforeEach(func() {
			timeout = 100 * int64(time.Millisecond)
			wait_time = 350 * time.Millisecond
			mockFlusher.EXPECT().FlushAchievements(achievements).Return(achievements, errors.New("error")).Times(3)
			mockFlusher.EXPECT().FlushAchievements(gomock.Any()).Return([]achievement.Achievement{}, nil).Times(4)
			mockFlusher.EXPECT().FlushResumes(resumes).Return(resumes, errors.New("error")).Times(3)
			mockFlusher.EXPECT().FlushResumes(gomock.Any()).Return([]resume.Resume{}, nil).Times(4)
		})
		It("flusher takes 4 Call", func() {
			s = saver.NewSaver(mockFlusher, achievements_cap, resumes_cap)
			err = s.Init(timeout, smart_del)
			s.SaveAchievements(achievements)
			s.SaveResumes(resumes)
			time.Sleep(wait_time)
			err = s.Close()
		})
	})

})