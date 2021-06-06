package saver

import (
	"errors"
	"fmt"
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
	Init(timeout int64, smart_other_cap_del bool) error
	Close() error
}

type saver struct {
	flusher             flusher.Flusher
	achievements_cap    int
	resumes_cap         int
	mutex_a             sync.Mutex
	achievements        []achievement.Achievement
	mutex_r             sync.Mutex
	resumes             []resume.Resume
	timer               *time.Ticker
	close_channel       chan struct{}
	channel_closed      chan struct{}
	smart_other_cap_del bool
}

func NewSaver(flusher flusher.Flusher, achievements_cap int, resumes_cap int) Saver {
	return &saver{
		flusher:             flusher,
		mutex_a:             sync.Mutex{},
		achievements_cap:    achievements_cap,
		mutex_r:             sync.Mutex{},
		resumes_cap:         resumes_cap,
		achievements:        make([]achievement.Achievement, 0, achievements_cap),
		resumes:             make([]resume.Resume, 0, resumes_cap),
		timer:               nil,
		close_channel:       make(chan struct{}),
		channel_closed:      make(chan struct{}),
		smart_other_cap_del: false,
	}
}

func (s *saver) Init(timeout int64, smart_other_cap_del bool) error {
	if timeout <= 0 {
		return errors.New("bad timeout. Must be > 0")
	}
	s.smart_other_cap_del = smart_other_cap_del
	s.timer = time.NewTicker(time.Duration(timeout))
	go func() {
		for {
			select {
			case <-s.timer.C:
				s.mutex_a.Lock()
				s.flusher.FlushAchievements(s.achievements)
				s.mutex_a.Unlock()
				s.mutex_r.Lock()
				s.flusher.FlushResumes(s.resumes)
				s.mutex_r.Unlock()
				fmt.Print("\n\n\ntimed\n\n\n")
			case <-s.close_channel:
				s.mutex_a.Lock()
				s.flusher.FlushAchievements(s.achievements)
				s.mutex_a.Unlock()
				s.mutex_r.Lock()
				s.flusher.FlushResumes(s.resumes)
				s.mutex_r.Unlock()
				s.channel_closed <- struct{}{}
				return
			}
		}
	}()
	return nil
}

func (s *saver) SaveAchievements(achievements []achievement.Achievement) error {
	s.mutex_a.Lock()
	s.achievements = utils.SaveAchievements(s.achievements, achievements, s.achievements_cap, s.smart_other_cap_del)
	s.mutex_a.Unlock()
	return nil
}
func (s *saver) SaveResumes(resumes []resume.Resume) error {
	s.mutex_r.Lock()
	s.resumes = utils.SaveResumes(s.resumes, resumes, s.resumes_cap, s.smart_other_cap_del)
	s.mutex_r.Unlock()
	return nil
}

func (s *saver) Close() error {
	s.mutex_a.Lock()
	s.mutex_r.Lock()
	s.timer.Stop()
	s.mutex_a.Unlock()
	s.mutex_r.Unlock()
	s.close_channel <- struct{}{}
	<-s.channel_closed
	return nil
}
