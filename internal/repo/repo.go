package repo

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/opentracing/opentracing-go"
	"github.com/ozoncp/ocp-resume-api/internal/achievement"
	"github.com/ozoncp/ocp-resume-api/internal/resume"
	"github.com/ozoncp/ocp-resume-api/internal/utils"
	"github.com/rs/zerolog/log"
)

type Repo interface {
	RepoResume
	RepoAchievement
}

type RepoResume interface {
	AddResumes(ctx context.Context, resumeArr []resume.Resume) ([]uint64, error)
	AddResumesBatch(ctx context.Context, resumeArr []resume.Resume) ([]uint64, error)
	RemoveResumeById(ctx context.Context, resumeId uint) error
	GetResumeById(ctx context.Context, resumeId uint) (*resume.Resume, error)
	ListResumes(ctx context.Context, offset, limit uint64) ([]resume.Resume, error)
	UpdateResumeById(ctx context.Context, resumeId uint, newData resume.Resume) error
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

func (r *repo) UpdateResumeById(ctx context.Context, resumeId uint, newData resume.Resume) error {
	query := sq.Update("resumes").Where(sq.Eq{"id": resumeId}).SetMap(
		map[string]interface{}{
			"document_id": newData.DocumentId,
			"id":          newData.Id,
		},
	).RunWith(r.base).PlaceholderFormat(sq.Dollar)
	result, err := query.ExecContext(ctx)
	if err != nil {
		return err
	}
	if affected, err := result.RowsAffected(); err == nil && affected == 0 {
		return errors.New("resume id not found")
	}
	return nil
}

func (r *repo) AddResumes(ctx context.Context, resumeArr []resume.Resume) ([]uint64, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, fmt.Sprintf("Add %v resumes", len(resumeArr)))
	defer span.Finish()
	batches, err := utils.SplitResumesToBatches(resumeArr, 64, false)
	if err {
		return []uint64{}, errors.New("can't split for batches")
	}
	returnedIds := make([]uint64, 0)
	returnErr := false
	for _, batch := range batches {
		batchIds, err := r.AddResumesBatch(ctx, batch)
		if err != nil {
			log.Err(err).Msgf("%v resumes was not added", len(batch))
			returnErr = true
		}
		returnedIds = append(returnedIds, batchIds...)
	}
	if returnErr {
		return returnedIds, errors.New("some resumes was not added! wrong situation, which the programmers did not describe")
	}
	return returnedIds, nil
}

func (r *repo) AddResumesBatch(ctx context.Context, resumeArr []resume.Resume) ([]uint64, error) {
	query := sq.Insert("resumes").Columns("document_id").Suffix(
		"returning id").RunWith(r.base).PlaceholderFormat(sq.Dollar)
	for _, resume := range resumeArr {
		query = query.Values(resume.DocumentId)
	}
	rows, err := query.QueryContext(ctx)
	if err != nil {
		log.Err(err).Msgf("Error while trying to add resumes")
		return nil, err
	}
	defer rows.Close()
	inserted := make([]uint64, 0)
	for rows.Next() {
		var resumeId uint64
		if err := rows.Scan(&resumeId); err != nil {
			log.Err(err).Msgf("Error while trying to list inserted resumes")
			return inserted, err
		}
		inserted = append(inserted, resumeId)
	}
	return inserted, nil
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
