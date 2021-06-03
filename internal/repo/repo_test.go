package repo_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/ozoncp/ocp-resume-api/internal/achievement"
	"github.com/ozoncp/ocp-resume-api/internal/repo"
	"github.com/ozoncp/ocp-resume-api/internal/resume"
)

var _ = Describe("Repo", func() {
	r := repo.NewRepo(11)
	achsA := make([]achievement.Achievement, 10)
	achsR := make([]resume.Resume, 10)
	err := r.AddAchievements(achsA)
	Expect(err).ShouldNot(BeNil())
	err = r.AddResumes(achsR)
	Expect(err).ShouldNot(BeNil())
})
