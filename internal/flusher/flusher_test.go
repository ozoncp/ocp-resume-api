package flusher_test

import (
	"context"
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
		err1            error
		err2            error
		ctrl            *gomock.Controller
		mockRepo        *mocks.MockRepo
		f               flusher.Flusher
		achievements    []achievement.Achievement
		retAchievements []achievement.Achievement
		resumes         []resume.Resume
		retResumes      []resume.Resume
		achBsize        uint64
		resumeBsize     uint64
		ctx             context.Context
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		ctx = context.Background()
		mockRepo = mocks.NewMockRepo(ctrl)
	})

	JustBeforeEach(func() {
		retAchievements, err1 = f.FlushAchievements(ctx, achievements)
		retResumes, err2 = f.FlushResumes(ctx, resumes)
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	When("only achievements", func() {
		Context("got empty array", func() {
			BeforeEach(func() {
				achievements = []achievement.Achievement{}
				achBsize = 32
				f = flusher.NewFlushAchievementsOnly(mockRepo, achBsize)
				mockRepo.EXPECT().AddAchievements(ctx, achievements).Return(nil).Times(0)
				mockRepo.EXPECT().AddResumes(ctx, resumes).Return(nil).Times(0)
			})
			It("return empty array and no error for achievements flush and error for resume flush", func() {
				//Expect(err).Should(BeNil())
				Expect(retAchievements).Should(BeEmpty())
				Expect(retResumes).Should(BeEmpty())
				Expect(err1).Should(BeNil())
				Expect(err2).Should(MatchError("can't split resumes to batches"))
			})
		})
		Context("got elements in array less then bsize", func() {
			BeforeEach(func() {
				achBsize = 32
				achievements = make([]achievement.Achievement, 2)
				resumes = make([]resume.Resume, 3)
				f = flusher.NewFlushAchievementsOnly(mockRepo, achBsize)
				mockRepo.EXPECT().AddAchievements(ctx, achievements).Return(nil).Times(1)
				mockRepo.EXPECT().AddResumes(ctx, resumes).Return(errors.New("Resume not created")).Times(0)
			})
			It("return empty array and no error for achievements flush and input array and error for resume flush", func() {
				//Expect(err).Should(BeNil())
				Expect(retAchievements).Should(BeEmpty())
				Expect(retResumes).Should(Equal(resumes))
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
				resumeBsize = 32
				f = flusher.NewFlusherResumeOnly(mockRepo, resumeBsize)
				mockRepo.EXPECT().AddAchievements(ctx, achievements).Return(nil).Times(0)
				mockRepo.EXPECT().AddResumes(ctx, resumes).Return(nil).Times(0)
			})
			It("return empty array and no error for resume flush and error for achievement flush", func() {
				//Expect(err).Should(BeNil())
				Expect(retAchievements).Should(BeEmpty())
				Expect(retResumes).Should(BeEmpty())
				Expect(err1).Should(MatchError("can't split achievements to batches"))
				Expect(err2).Should(BeNil())
			})
		})
		Context("got elements in array less then bsize", func() {
			BeforeEach(func() {
				resumeBsize = 16
				achievements = make([]achievement.Achievement, 2)
				resumes = make([]resume.Resume, 3)
				f = flusher.NewFlusherResumeOnly(mockRepo, resumeBsize)
				mockRepo.EXPECT().AddAchievements(ctx, achievements).Return(errors.New("Achievements not created")).Times(0)
				mockRepo.EXPECT().AddResumes(ctx, resumes).Return(nil).Times(1)
			})
			It("return empty array and no error for resumes flush and input array and error for achievements flush", func() {
				//Expect(err).Should(BeNil())
				Expect(retAchievements).Should(Equal(achievements))
				Expect(retResumes).Should(BeEmpty())
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
				resumeBsize = 16
				achBsize = 32
				f = flusher.NewFlusher(mockRepo, resumeBsize, achBsize)
				mockRepo.EXPECT().AddAchievements(ctx, achievements).Return(nil).Times(0)
				mockRepo.EXPECT().AddResumes(ctx, resumes).Return(nil).Times(0)
			})
			It("return empty array and no error for achievements and resumes", func() {
				//Expect(err).Should(BeNil())
				Expect(retAchievements).Should(BeEmpty())
				Expect(retResumes).Should(BeEmpty())
				Expect(err1).Should(BeNil())
				Expect(err2).Should(BeNil())
			})
		})
		Context("got elements in array less then bsize", func() {
			BeforeEach(func() {
				achBsize = 32
				resumeBsize = 16
				achievements = make([]achievement.Achievement, 2)
				resumes = make([]resume.Resume, 3)
				f = flusher.NewFlusher(mockRepo, resumeBsize, achBsize)
				mockRepo.EXPECT().AddAchievements(ctx, achievements).Return(nil).Times(1)
				mockRepo.EXPECT().AddResumes(ctx, resumes).Return(nil).Times(1)
			})
			It("return empty array and no error for achievements and resumes (functions add calls ones)", func() {
				//Expect(err).Should(BeNil())
				Expect(retAchievements).Should(BeEmpty())
				Expect(retResumes).Should(BeEmpty())
				Expect(err1).Should(BeNil())
				Expect(err2).Should(BeNil())
			})
		})
		Context("got elements in array more then bsize", func() {
			BeforeEach(func() {
				achBsize = 4
				resumeBsize = 4
				achievements = make([]achievement.Achievement, 8)
				resumes = make([]resume.Resume, 7)
				f = flusher.NewFlusher(mockRepo, resumeBsize, achBsize)
				mockRepo.EXPECT().AddAchievements(ctx, gomock.Any()).Return(nil).Times(2)
				mockRepo.EXPECT().AddResumes(ctx, gomock.Any()).Return(nil).Times(2)
			})
			It("return empty array and no error for achievements and resumes (functions add calls twice)", func() {
				//Expect(err).Should(BeNil())
				Expect(retAchievements).Should(BeEmpty())
				Expect(retResumes).Should(BeEmpty())
				Expect(err1).Should(BeNil())
				Expect(err2).Should(BeNil())
			})
		})
		Context("got elements in array and have an error in repo", func() {
			BeforeEach(func() {
				achBsize = 4
				resumeBsize = 3
				achievements = make([]achievement.Achievement, 16)
				resumes = make([]resume.Resume, 9)
				f = flusher.NewFlusher(mockRepo, resumeBsize, achBsize)
				mockRepo.EXPECT().AddAchievements(ctx, gomock.Any()).Return(nil).Times(2)
				mockRepo.EXPECT().AddAchievements(ctx, gomock.Any()).Return(errors.New("error")).Times(2)
				mockRepo.EXPECT().AddResumes(ctx, gomock.Any()).Return(errors.New("error")).Times(3)
			})
			It("return not empty arrays and no error for achievements and resumes (in case with resumes returned array like input)", func() {
				//Expect(err).Should(BeNil())
				Expect(retAchievements).ShouldNot(BeEmpty())
				Expect(retResumes).Should(BeEquivalentTo(resumes))
				Expect(err1).Should(BeNil())
				Expect(err2).Should(BeNil())
			})
		})
	})

})
