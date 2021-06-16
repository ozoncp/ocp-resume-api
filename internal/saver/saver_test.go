package saver_test

import (
	"context"
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
		err             error
		ctrl            *gomock.Controller
		mockFlusher     *mocks.MockFlusher
		s               saver.Saver
		achievements    []achievement.Achievement
		resumes         []resume.Resume
		achievementsCap int
		resumesCap      int
		timeout         int64
		waitTime        time.Duration
		smartDel        bool
		ctx             context.Context
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		mockFlusher = mocks.NewMockFlusher(ctrl)
		smartDel = true
		achievementsCap = 512
		resumesCap = 128
		achievements = make([]achievement.Achievement, 256)
		resumes = make([]resume.Resume, 32)
		ctx = context.Background()
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	It("return error for timeout <=0", func() {
		//Expect(err).Should(BeNil())
		timeout = -1
		s = saver.NewSaver(mockFlusher, achievementsCap, resumesCap)
		err = s.Init(ctx, timeout, smartDel)
		Expect(err).ShouldNot(BeNil())
	})

	When("close before timeout", func() {
		BeforeEach(func() {
			timeout = 1000 * int64(time.Millisecond)
			waitTime = 2 * time.Millisecond
			mockFlusher.EXPECT().FlushAchievements(ctx, gomock.Any()).Return([]achievement.Achievement{}, nil).Times(1)
			mockFlusher.EXPECT().FlushResumes(ctx, gomock.Any()).Return([]resume.Resume{}, nil).Times(1)
		})
		It("flusher takes 1 Call", func() {
			s = saver.NewSaver(mockFlusher, achievementsCap, resumesCap)
			err = s.Init(ctx, timeout, smartDel)
			Expect(err).Should(BeNil())
			err = s.SaveAchievements(achievements)
			Expect(err).Should(BeNil())
			err = s.SaveResumes(resumes)
			Expect(err).Should(BeNil())
			time.Sleep(waitTime)
			err = s.Close()
		})
	})

	When("close after timeout", func() {
		BeforeEach(func() {
			timeout = 100 * int64(time.Millisecond)
			waitTime = 350 * time.Millisecond
			mockFlusher.EXPECT().FlushAchievements(ctx, gomock.Any()).Return([]achievement.Achievement{}, nil).Times(4)
			mockFlusher.EXPECT().FlushResumes(ctx, gomock.Any()).Return([]resume.Resume{}, nil).Times(4)
		})
		It("flusher takes 4 Call", func() {
			s = saver.NewSaver(mockFlusher, achievementsCap, resumesCap)
			err = s.Init(ctx, timeout, smartDel)
			Expect(err).Should(BeNil())
			err = s.SaveAchievements(achievements)
			Expect(err).Should(BeNil())
			err = s.SaveResumes(resumes)
			Expect(err).Should(BeNil())
			time.Sleep(waitTime)
			err = s.Close()
		})
	})

	When("flush with errors", func() {
		BeforeEach(func() {
			timeout = 100 * int64(time.Millisecond)
			waitTime = 350 * time.Millisecond
			mockFlusher.EXPECT().FlushAchievements(ctx, achievements).Return(achievements, errors.New("error")).Times(3)
			mockFlusher.EXPECT().FlushAchievements(ctx, gomock.Any()).Return([]achievement.Achievement{}, nil).Times(4)
			mockFlusher.EXPECT().FlushResumes(ctx, resumes).Return(resumes, errors.New("error")).Times(3)
			mockFlusher.EXPECT().FlushResumes(ctx, gomock.Any()).Return([]resume.Resume{}, nil).Times(4)
		})
		It("flusher takes 4 Call", func() {
			s = saver.NewSaver(mockFlusher, achievementsCap, resumesCap)
			err = s.Init(ctx, timeout, smartDel)
			Expect(err).Should(BeNil())
			err = s.SaveAchievements(achievements)
			Expect(err).Should(BeNil())
			err = s.SaveResumes(resumes)
			Expect(err).Should(BeNil())
			time.Sleep(waitTime)
			err = s.Close()
		})
	})

})
