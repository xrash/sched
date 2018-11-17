package sched

import (
	"time"
)

type OnceScheduler struct {
	when  time.Time
	never time.Time
}

func NewOnceScheduler(when time.Time) *OnceScheduler {
	return &OnceScheduler{
		when:  when,
		never: time.Unix(1<<63-1, 0),
	}
}

func (s *OnceScheduler) Next(now time.Time) time.Time {
	if now.Before(s.when) {
		return s.when
	}

	return s.never
}
