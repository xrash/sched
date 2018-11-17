package sched

import (
	"time"
)

type Task func(prev, next time.Time)
