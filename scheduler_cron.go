package sched

import (
	"github.com/robfig/cron"
	"time"
)

type CronScheduler struct {
	cronSchedule cron.Schedule
}

func NewCronScheduler(spec string) (*CronScheduler, error) {
	parserOpts := cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow
	cronParser := cron.NewParser(parserOpts)
	s, err := cronParser.Parse(spec)
	if err != nil {
		return nil, err
	}

	return &CronScheduler{
		cronSchedule: s,
	}, nil
}

func (s *CronScheduler) Next(now time.Time) time.Time {
	return s.cronSchedule.Next(now)
}
