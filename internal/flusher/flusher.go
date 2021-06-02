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
	repoResume             repo.RepoResume
	repoAchievement        repo.RepoAchievement
	resumes_batchsize      uint64
	achievements_batchsize uint64
}

func NewFlusher(full_repo repo.Repo, resumes_batchsize uint64, achievements_batchsize uint64) Flusher {
	return &flusher{
		repoResume:             full_repo,
		repoAchievement:        full_repo,
		resumes_batchsize:      resumes_batchsize,
		achievements_batchsize: achievements_batchsize,
	}
}

func NewFlusherResumeOnly(resume_repo repo.RepoResume, resumes_batchsize uint64) Flusher {
	return &flusher{
		repoResume:             resume_repo,
		repoAchievement:        nil,
		resumes_batchsize:      resumes_batchsize,
		achievements_batchsize: 0,
	}
}

func NewFlushAchievementsOnly(achievement_repo repo.Repo, achievements_batchsize uint64) Flusher {
	return &flusher{
		repoResume:             nil,
		repoAchievement:        achievement_repo,
		resumes_batchsize:      0,
		achievements_batchsize: achievements_batchsize,
	}
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
		err := f.repoResume.AddResumes(batch)
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
		err := f.repoAchievement.AddAchievements(batch)
		if err != nil {
			ret_arr = append(ret_arr, batch...)
		}
	}
	return ret_arr, nil
}
