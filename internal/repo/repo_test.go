package repo_test

import (
	"context"

	"github.com/jmoiron/sqlx"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/ozoncp/ocp-resume-api/internal/achievement"
	"github.com/ozoncp/ocp-resume-api/internal/repo"
	"github.com/ozoncp/ocp-resume-api/internal/resume"
)

var _ = Describe("Repo", func() {
	r := repo.NewRepo(sqlx.DB{})
	achsA := make([]achievement.Achievement, 10)
	achsR := make([]resume.Resume, 10)
	ctx := context.Background()
	err := r.AddAchievements(ctx, achsA)
	Expect(err).ShouldNot(BeNil())
	err = r.AddResumes(ctx, achsR)
	Expect(err).ShouldNot(BeNil())
})
