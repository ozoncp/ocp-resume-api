package flusher

import (
	"context"
	"errors"

	"github.com/ozoncp/ocp-resume-api/internal/achievement"
	"github.com/ozoncp/ocp-resume-api/internal/repo"
	"github.com/ozoncp/ocp-resume-api/internal/resume"
	"github.com/ozoncp/ocp-resume-api/internal/utils"
)

type FlusherResume interface {
	FlushResumes(ctx context.Context, r []resume.Resume) ([]resume.Resume, error)
}

type FlusherAchievement interface {
	FlushAchievements(ctx context.Context, a []achievement.Achievement) ([]achievement.Achievement, error)
}

type Flusher interface {
	FlusherResume
	FlusherAchievement
}

type flusher struct {
	repoResume            repo.RepoResume
	repoAchievement       repo.RepoAchievement
	resumesBatchSize      uint64
	achievementsBatchSize uint64
}

func NewFlusher(fullRepo repo.Repo, resumesBatchSize uint64, achievementsBatchSize uint64) Flusher {
	return &flusher{
		repoResume:            fullRepo,
		repoAchievement:       fullRepo,
		resumesBatchSize:      resumesBatchSize,
		achievementsBatchSize: achievementsBatchSize,
	}
}

func NewFlusherResumeOnly(resumeRepo repo.RepoResume, resumesBatchSize uint64) Flusher {
	return &flusher{
		repoResume:            resumeRepo,
		repoAchievement:       nil,
		resumesBatchSize:      resumesBatchSize,
		achievementsBatchSize: 0,
	}
}

func NewFlushAchievementsOnly(achievementRepo repo.Repo, achievementsBatchSize uint64) Flusher {
	return &flusher{
		repoResume:            nil,
		repoAchievement:       achievementRepo,
		resumesBatchSize:      0,
		achievementsBatchSize: achievementsBatchSize,
	}
}

func (f *flusher) FlushResumes(ctx context.Context, r []resume.Resume) ([]resume.Resume, error) {
	batches, ok := utils.SplitResumesToBatches(r, int(f.resumesBatchSize), false)
	if !ok {
		return r, errors.New("can't split resumes to batches")
	}
	retArr := make([]resume.Resume, 0, len(r))
	for _, batch := range batches {
		_, err := f.repoResume.AddResumesBatch(ctx, batch)
		if err != nil {
			retArr = append(retArr, batch...)
		}
	}
	return retArr, nil
}

func (f *flusher) FlushAchievements(ctx context.Context, r []achievement.Achievement) ([]achievement.Achievement, error) {
	batches, ok := utils.SplitAchievementsToBatches(r, int(f.achievementsBatchSize), false)
	if !ok {
		return r, errors.New("can't split achievements to batches")
	}
	retArr := make([]achievement.Achievement, 0, len(r))
	for _, batch := range batches {
		err := f.repoAchievement.AddAchievements(ctx, batch)
		if err != nil {
			retArr = append(retArr, batch...)
		}
	}
	return retArr, nil
}
