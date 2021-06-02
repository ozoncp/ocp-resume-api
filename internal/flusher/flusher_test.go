package flusher_test

import (
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/ozoncp/ocp-resume-api/internal/achievement"
	"github.com/ozoncp/ocp-resume-api/internal/flusher"
	"github.com/ozoncp/ocp-resume-api/internal/mocks"
	"github.com/ozoncp/ocp-resume-api/internal/resume"
)

var _ = Describe("Flusher", func() {
	var (
		err1             error
		err2             error
		ctrl             *gomock.Controller
		mockRepo         *mocks.MockRepo
		f                flusher.Flusher
		achievements     []achievement.Achievement
		ret_achievements []achievement.Achievement
		resumes          []resume.Resume
		ret_resumes      []resume.Resume
		ach_bsize        uint64
		resume_bsize     uint64
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		mockRepo = mocks.NewMockRepo(ctrl)
	})

	JustBeforeEach(func() {
		f = flusher.NewFlusher(mockRepo, resume_bsize, ach_bsize)
		ret_achievements, err1 = f.FlushAchievements(achievements)
		ret_resumes, err2 = f.FlushResumes(resumes)
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	Context("got empty arrays", func() {
		BeforeEach(func() {
			ach_bsize = 32
			resume_bsize = 16
			achievements = []achievement.Achievement{}
			resumes = []resume.Resume{}
			mockRepo.EXPECT().AddAchievements(achievements).Return(nil).MaxTimes(0)
			mockRepo.EXPECT().AddResumes(resumes).Return(nil).MaxTimes(0)
		})
		It("result", func() {
			//Expect(err).Should(BeNil())
			Expect(ret_achievements).Should(BeEmpty())
			Expect(ret_resumes).Should(BeEmpty())
			Expect(err1).Should(BeNil())
			Expect(err2).Should(BeNil())
		})
	})
})
