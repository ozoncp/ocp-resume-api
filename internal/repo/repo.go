package repo

import (
	"errors"

	"github.com/ozoncp/ocp-resume-api/internal/achievement"
	"github.com/ozoncp/ocp-resume-api/internal/resume"
)

type Repo interface {
	RepoResume
	RepoAchievement
}

type RepoResume interface {
	AddResumes(r []resume.Resume) error
	//RemoveResumeById(resumeId uint) error
	//RemoveResumeByNdx(resumeNdx uint64) error
	//GetResumeById(resumeId uint) (*resume.Resume, error)
	//GetResumeByNdx(resumeNdx uint64) (*resume.Resume, error)
	//ListResumes(offsetNdx, limitNdx uint64) ([]RepoResume, error)
}

type RepoAchievement interface {
	AddAchievements(r []achievement.Achievement) error
	//RemoveAchievementById(AchievementId uint) error
	//RemoveAchievementByNdx(AchievementNdx uint64) error
	//GetAchievementById(AchievementId uint) (*achievement.Achievement, error)
	//GetAchievementByNdx(AchievementNdx uint64) (*achievement.Achievement, error)
	//ListAchievements(offsetNdx, limitNdx uint64) ([]RepoAchievement, error)
}

type repo struct {
	count uint64
}

func NewRepo(count uint64) Repo {
	return &repo{
		count: count,
	}
}

func (r *repo) AddResumes(resume_arr []resume.Resume) error {
	if r == nil {
		return errors.New("resume repo is not created")
	}
	return nil
}

func (r *repo) AddAchievements(achievement_arr []achievement.Achievement) error {
	if r == nil {
		return errors.New("achievement repo is not created")
	}
	return nil
}
