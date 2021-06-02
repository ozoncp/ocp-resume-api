package flusher

import (
	"errors"

	"github.com/ozoncp/ocp-resume-api/internal/achievement"
	"github.com/ozoncp/ocp-resume-api/internal/repo"
	"github.com/ozoncp/ocp-resume-api/internal/resume"
	"github.com/ozoncp/ocp-resume-api/internal/utils"
)

type FlusherResume interface {
	FlushResumes(r []resume.Resume) ([]resume.Resume, error)
}

type FlusherAchievement interface {
	FlushAchievements(a []achievement.Achievement) ([]achievement.Achievement, error)
}

type Flusher interface {
	FlusherResume
	FlusherAchievement
}

type flusher struct {
	repo                   repo.Repo
	resumes_batchsize      uint64
	achievements_batchsize uint64
}

func (f *flusher) FlushResumes(r []resume.Resume) ([]resume.Resume, error) {
	if f == nil {
		return r, errors.New("flusher not created")
	}
	batches, ok := utils.SplitResumesToBatches(r, int(f.resumes_batchsize), false)
	if !ok {
		return r, errors.New("can't split resumes to batches")
	}
	ret_arr := make([]resume.Resume, 0, len(r))
	for _, batch := range batches {
		err := f.repo.AddResumes(batch)
		if err != nil {
			ret_arr = append(ret_arr, batch...)
		}
	}
	return ret_arr, nil
}

func (f *flusher) FlushAchievements(r []achievement.Achievement) ([]achievement.Achievement, error) {
	if f == nil {
		return r, errors.New("flusher not created")
	}
	batches, ok := utils.SplitAchievementsToBatches(r, int(f.achievements_batchsize), false)
	if !ok {
		return r, errors.New("can't split achievements to batches")
	}
	ret_arr := make([]achievement.Achievement, 0, len(r))
	for _, batch := range batches {
		err := f.repo.AddAchievements(batch)
		if err != nil {
			ret_arr = append(ret_arr, batch...)
		}
	}
	return ret_arr, nil
}
