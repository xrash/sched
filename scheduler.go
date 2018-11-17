package sched

import (
	"time"
)

type Scheduler interface {
	Next(time.Time) time.Time
}
