package saver

import (
	"errors"
	"sync"
	"time"

	"github.com/ozoncp/ocp-resume-api/internal/achievement"
	"github.com/ozoncp/ocp-resume-api/internal/flusher"
	"github.com/ozoncp/ocp-resume-api/internal/resume"
	"github.com/ozoncp/ocp-resume-api/internal/utils"
)

type Saver interface {
	SaveAchievements(achievements []achievement.Achievement) error
	SaveResumes(resumes []resume.Resume) error
	Init(timeout int64, smartOtherCapDel bool) error
	Close() error
}

type saver struct {
	flusher           flusher.Flusher
	achievementsCap   int
	resumesCap        int
	mutexAchievements sync.Mutex
	achievements      []achievement.Achievement
	mutexResumes      sync.Mutex
	resumes           []resume.Resume
	timer             *time.Ticker
	closeChannel      chan struct{}
	channelClosed     chan struct{}
	smartOtherCapDel  bool
}

func NewSaver(flusher flusher.Flusher, achievementsCap int, resumesCap int) Saver {
	return &saver{
		flusher:           flusher,
		mutexAchievements: sync.Mutex{},
		achievementsCap:   achievementsCap,
		mutexResumes:      sync.Mutex{},
		resumesCap:        resumesCap,
		achievements:      make([]achievement.Achievement, 0, achievementsCap),
		resumes:           make([]resume.Resume, 0, resumesCap),
		timer:             nil,
		closeChannel:      make(chan struct{}),
		channelClosed:     make(chan struct{}),
		smartOtherCapDel:  false,
	}
}

func (s *saver) Init(timeout int64, smartOtherCapDel bool) error {
	if timeout <= 0 {
		return errors.New("bad timeout. Must be > 0")
	}
	s.smartOtherCapDel = smartOtherCapDel
	s.timer = time.NewTicker(time.Duration(timeout))
	go func() {
		for {
			var isClosed bool
			select {
			case <-s.timer.C:
				isClosed = false
			case <-s.closeChannel:
				isClosed = true
			}
			s.FlushAchievements()
			s.FlushResumes()
			if isClosed {
				s.channelClosed <- struct{}{}
				return
			}
		}
	}()
	return nil
}

func (s *saver) FlushAchievements() {
	s.mutexAchievements.Lock()
	defer s.mutexAchievements.Unlock()
	arrAchievements, err := s.flusher.FlushAchievements(s.achievements)
	for tryCount := int(3); tryCount > 0; tryCount-- {
		if err == nil && len(arrAchievements) == 0 {
			break
		}
		arrAchievements, err = s.flusher.FlushAchievements(arrAchievements)
	}
}

func (s *saver) FlushResumes() {
	s.mutexResumes.Lock()
	defer s.mutexResumes.Unlock()
	arrResumes, err := s.flusher.FlushResumes(s.resumes)
	for tryCount := int(3); tryCount > 0; tryCount-- {
		if err == nil && len(arrResumes) == 0 {
			break
		}
		arrResumes, err = s.flusher.FlushResumes(arrResumes)
	}
}

func (s *saver) SaveAchievements(achievements []achievement.Achievement) error {
	s.mutexAchievements.Lock()
	s.achievements = utils.SaveAchievements(s.achievements, achievements, s.achievementsCap, s.smartOtherCapDel)
	s.mutexAchievements.Unlock()
	return nil
}
func (s *saver) SaveResumes(resumes []resume.Resume) error {
	s.mutexResumes.Lock()
	s.resumes = utils.SaveResumes(s.resumes, resumes, s.resumesCap, s.smartOtherCapDel)
	s.mutexResumes.Unlock()
	return nil
}

func (s *saver) Close() error {
	s.mutexAchievements.Lock()
	s.mutexResumes.Lock()
	s.timer.Stop()
	s.mutexAchievements.Unlock()
	s.mutexResumes.Unlock()
	s.closeChannel <- struct{}{}
	<-s.channelClosed
	return nil
}
