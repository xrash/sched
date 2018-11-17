package sched

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestLoopScheduler(t *testing.T) {
	fmt.Println("TestLoopScheduler")

	duration, err := time.ParseDuration("2s")
	if err != nil {
		panic(err)
	}

	scheduler := NewLoopScheduler(duration)

	before := time.Now()
	after := before.Add(duration)
	next := scheduler.Next(before)

	assert.Equal(t, after.Unix(), next.Unix(), "time.After()'s value and LoopScheduler.Next()'s value must be equal.")
}
