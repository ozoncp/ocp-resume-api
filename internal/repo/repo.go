package repo

import (
	"context"
	"database/sql"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/ozoncp/ocp-resume-api/internal/achievement"
	"github.com/ozoncp/ocp-resume-api/internal/resume"
	"github.com/rs/zerolog/log"
)

type Repo interface {
	RepoResume
	RepoAchievement
}

type RepoResume interface {
	AddResumes(ctx context.Context, r []resume.Resume) error
	RemoveResumeById(ctx context.Context, resumeId uint) error
	GetResumeById(ctx context.Context, resumeId uint) (*resume.Resume, error)
	ListResumes(ctx context.Context, offset, limit uint64) ([]resume.Resume, error)
}

type RepoAchievement interface {
	AddAchievements(ctx context.Context, r []achievement.Achievement) error
	//RemoveAchievementById(AchievementId uint) error
	//RemoveAchievementByNdx(AchievementNdx uint64) error
	//GetAchievementById(AchievementId uint) (*achievement.Achievement, error)
	//GetAchievementByNdx(AchievementNdx uint64) (*achievement.Achievement, error)
	//ListAchievements(offsetNdx, limitNdx uint64) ([]RepoAchievement, error)
}

type repo struct {
	base *sqlx.DB
}

func NewRepo(db *sqlx.DB) Repo {
	return &repo{
		base: db,
	}
}

func (r *repo) AddResumes(ctx context.Context, resumeArr []resume.Resume) error {
	query := sq.Insert("resumes").Columns("document_id").RunWith(r.base).PlaceholderFormat(sq.Dollar)
	for _, resume := range resumeArr {
		query = query.Values(resume.DocumentId)
	}
	_, err := query.ExecContext(ctx)
	if err == nil {
		log.Err(err).Msgf("Error while trying to add resume %v", resumeArr)
		return err
	}
	log.Info().Msgf("%v resumes added", len(resumeArr))
	return nil
}

func (r *repo) RemoveResumeById(ctx context.Context, resumeId uint) error {
	query := sq.Delete("resumes").Where(sq.Eq{"id": resumeId}).RunWith(r.base).PlaceholderFormat(sq.Dollar)
	_, err := query.ExecContext(ctx)
	if err == nil {
		log.Err(err).Msgf("Error while trying to remove resume with id %v", resumeId)
		return err
	}
	log.Info().Msgf("Resume with id %v removed", resumeId)
	return nil
}

func (r *repo) GetResumeById(ctx context.Context, resumeId uint) (*resume.Resume, error) {
	query := sq.Select("id", "document_id").From("resumes").Where(sq.Eq{"id": resumeId}).RunWith(r.base).PlaceholderFormat(sq.Dollar)
	var selected resume.Resume
	err := query.QueryRowContext(ctx).Scan(&selected.Id, &selected.DocumentId)
	switch {
	case err == sql.ErrNoRows:
		log.Err(err).Msgf("No resume with id %v", resumeId)
		return nil, err
	case err != nil:
		log.Err(err).Msgf("Query error while trying to find resume with id %v", resumeId)
		return nil, err
	}
	return &selected, nil
}

func (r *repo) ListResumes(ctx context.Context, offset, limit uint64) ([]resume.Resume, error) {
	query := sq.Select("id", "document_id").From("resumes").Limit(limit).Offset(offset).RunWith(r.base).PlaceholderFormat(sq.Dollar)
	rows, err := query.QueryContext(ctx)
	if err != nil {
		log.Err(err).Msgf("Error while trying to list resumes")
		return nil, err
	}
	defer rows.Close()
	selected := make([]resume.Resume, 0)
	for rows.Next() {
		var resumeRow resume.Resume
		if err := rows.Scan(&resumeRow.Id, &resumeRow.DocumentId); err != nil {
			log.Err(err).Msgf("Error while trying to list resumes")
			return selected, err
		}
		selected = append(selected, resumeRow)
	}
	return selected, nil
}

func (r *repo) AddAchievements(ctx context.Context, achievementArr []achievement.Achievement) error {
	query := sq.Insert("achievements").Columns("name", "description").RunWith(r.base).PlaceholderFormat(sq.Dollar)
	for _, resume := range achievementArr {
		query = query.Values(resume.Name, resume.Description)
	}
	_, err := query.ExecContext(ctx)
	if err == nil {
		fmt.Printf("Error execute")
		return err
	}
	return nil
}
