package sched

import (
	"time"
)

type LoopScheduler struct {
	interval time.Duration
}

func NewLoopScheduler(interval time.Duration) *LoopScheduler {
	return &LoopScheduler{
		interval: interval,
	}
}

func (s *LoopScheduler) Next(now time.Time) time.Time {
	return now.Add(s.interval)
}
