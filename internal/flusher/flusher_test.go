package flusher_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/ozoncp/ocp-resume-api/internal/achievement"
	"github.com/ozoncp/ocp-resume-api/internal/flusher"
)

var _ = Describe("Flusher", func() {
	f := flusher.NewFlusher(nil, 11, 12)
	achs := make([]achievement.Achievement, 10)
	ret_achs, err := f.FlushAchievements(achs)
	Expect(err).ShouldNot(BeNil())
	Expect(ret_achs).Should(BeEquivalentTo(achs))
})
