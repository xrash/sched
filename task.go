package sched

import (
	"time"
)

type Task func(now, next time.Time)
