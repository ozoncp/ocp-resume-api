package flusher_test

import (
	"errors"

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
		ret_achievements, err1 = f.FlushAchievements(achievements)
		ret_resumes, err2 = f.FlushResumes(resumes)
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	When("only achievements", func() {
		Context("got empty array", func() {
			BeforeEach(func() {
				achievements = []achievement.Achievement{}
				ach_bsize = 32
				f = flusher.NewFlushAchievementsOnly(mockRepo, ach_bsize)
				mockRepo.EXPECT().AddAchievements(achievements).Return(nil).Times(0)
				mockRepo.EXPECT().AddResumes(resumes).Return(nil).Times(0)
			})
			It("return empty array and no error for achievements flush and error for resume flush", func() {
				//Expect(err).Should(BeNil())
				Expect(ret_achievements).Should(BeEmpty())
				Expect(ret_resumes).Should(BeEmpty())
				Expect(err1).Should(BeNil())
				Expect(err2).Should(MatchError("can't split resumes to batches"))
			})
		})
		Context("got elements in array less then bsize", func() {
			BeforeEach(func() {
				ach_bsize = 32
				achievements = make([]achievement.Achievement, 2)
				resumes = make([]resume.Resume, 3)
				f = flusher.NewFlushAchievementsOnly(mockRepo, ach_bsize)
				mockRepo.EXPECT().AddAchievements(achievements).Return(nil).Times(1)
				mockRepo.EXPECT().AddResumes(resumes).Return(errors.New("Resume not created")).Times(0)
			})
			It("return empty array and no error for achievements flush and input array and error for resume flush", func() {
				//Expect(err).Should(BeNil())
				Expect(ret_achievements).Should(BeEmpty())
				Expect(ret_resumes).Should(Equal(resumes))
				Expect(err1).Should(BeNil())
				Expect(err2).Should(MatchError("can't split resumes to batches"))
			})
		})
	})

	When("only resumes", func() {
		Context("got empty array", func() {
			BeforeEach(func() {
				achievements = []achievement.Achievement{}
				resumes = []resume.Resume{}
				resume_bsize = 32
				f = flusher.NewFlusherResumeOnly(mockRepo, resume_bsize)
				mockRepo.EXPECT().AddAchievements(achievements).Return(nil).Times(0)
				mockRepo.EXPECT().AddResumes(resumes).Return(nil).Times(0)
			})
			It("return empty array and no error for resume flush and error for achievement flush", func() {
				//Expect(err).Should(BeNil())
				Expect(ret_achievements).Should(BeEmpty())
				Expect(ret_resumes).Should(BeEmpty())
				Expect(err1).Should(MatchError("can't split achievements to batches"))
				Expect(err2).Should(BeNil())
			})
		})
		Context("got elements in array less then bsize", func() {
			BeforeEach(func() {
				resume_bsize = 16
				achievements = make([]achievement.Achievement, 2)
				resumes = make([]resume.Resume, 3)
				f = flusher.NewFlusherResumeOnly(mockRepo, resume_bsize)
				mockRepo.EXPECT().AddAchievements(achievements).Return(errors.New("Achievements not created")).Times(0)
				mockRepo.EXPECT().AddResumes(resumes).Return(nil).Times(1)
			})
			It("return empty array and no error for resumes flush and input array and error for achievements flush", func() {
				//Expect(err).Should(BeNil())
				Expect(ret_achievements).Should(Equal(achievements))
				Expect(ret_resumes).Should(BeEmpty())
				Expect(err1).Should(MatchError("can't split achievements to batches"))
				Expect(err2).Should(BeNil())
			})
		})
	})

	When("full flusher", func() {
		Context("got empty array", func() {
			BeforeEach(func() {
				achievements = []achievement.Achievement{}
				resumes = []resume.Resume{}
				resume_bsize = 16
				ach_bsize = 32
				f = flusher.NewFlusher(mockRepo, resume_bsize, ach_bsize)
				mockRepo.EXPECT().AddAchievements(achievements).Return(nil).Times(0)
				mockRepo.EXPECT().AddResumes(resumes).Return(nil).Times(0)
			})
			It("return empty array and no error for achievements and resumes", func() {
				//Expect(err).Should(BeNil())
				Expect(ret_achievements).Should(BeEmpty())
				Expect(ret_resumes).Should(BeEmpty())
				Expect(err1).Should(BeNil())
				Expect(err2).Should(BeNil())
			})
		})
		Context("got elements in array less then bsize", func() {
			BeforeEach(func() {
				ach_bsize = 32
				resume_bsize = 16
				achievements = make([]achievement.Achievement, 2)
				resumes = make([]resume.Resume, 3)
				f = flusher.NewFlusher(mockRepo, resume_bsize, ach_bsize)
				mockRepo.EXPECT().AddAchievements(achievements).Return(nil).Times(1)
				mockRepo.EXPECT().AddResumes(resumes).Return(nil).Times(1)
			})
			It("return empty array and no error for achievements and resumes (functions add calls ones)", func() {
				//Expect(err).Should(BeNil())
				Expect(ret_achievements).Should(BeEmpty())
				Expect(ret_resumes).Should(BeEmpty())
				Expect(err1).Should(BeNil())
				Expect(err2).Should(BeNil())
			})
		})
		Context("got elements in array more then bsize", func() {
			BeforeEach(func() {
				ach_bsize = 4
				resume_bsize = 4
				achievements = make([]achievement.Achievement, 8)
				resumes = make([]resume.Resume, 7)
				f = flusher.NewFlusher(mockRepo, resume_bsize, ach_bsize)
				mockRepo.EXPECT().AddAchievements(gomock.Any()).Return(nil).Times(2)
				mockRepo.EXPECT().AddResumes(gomock.Any()).Return(nil).Times(2)
			})
			It("return empty array and no error for achievements and resumes (functions add calls twice)", func() {
				//Expect(err).Should(BeNil())
				Expect(ret_achievements).Should(BeEmpty())
				Expect(ret_resumes).Should(BeEmpty())
				Expect(err1).Should(BeNil())
				Expect(err2).Should(BeNil())
			})
		})
		Context("got elements in array and have an error in repo", func() {
			BeforeEach(func() {
				ach_bsize = 4
				resume_bsize = 3
				achievements = make([]achievement.Achievement, 16)
				resumes = make([]resume.Resume, 9)
				f = flusher.NewFlusher(mockRepo, resume_bsize, ach_bsize)
				mockRepo.EXPECT().AddAchievements(gomock.Any()).Return(nil).Times(2)
				mockRepo.EXPECT().AddAchievements(gomock.Any()).Return(errors.New("error")).Times(2)
				mockRepo.EXPECT().AddResumes(gomock.Any()).Return(errors.New("error")).Times(3)
			})
			It("return not empty arrays and no error for achievements and resumes (in case with resumes returned array like input)", func() {
				//Expect(err).Should(BeNil())
				Expect(ret_achievements).ShouldNot(BeEmpty())
				Expect(ret_resumes).Should(BeEquivalentTo(resumes))
				Expect(err1).Should(BeNil())
				Expect(err2).Should(BeNil())
			})
		})
	})

})
