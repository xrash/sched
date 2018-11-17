package sched

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestJob(t *testing.T) {
	fmt.Println("TestJob")

	returnValues := make(chan time.Time, 1024)

	task := func(prev, next time.Time) {
		returnValues <- prev
		returnValues <- next
	}

	duration, err := time.ParseDuration("2s")
	if err != nil {
		panic(err)
	}

	scheduler := NewLoopScheduler(duration)

	job := LaunchJob(task, scheduler)

	assert.Equal(t, scheduler, job.Scheduler(), "Schedulers must be the same.")

	prev1 := <-returnValues
	next1 := <-returnValues
	d := next1.Sub(prev1)
	assert.Equal(t, 2.0, d.Seconds(), "Seconds after first run must be 2")

	gotPrev2 := job.Prev()
	gotNext2 := job.Next()
	prev2 := <-returnValues
	next2 := <-returnValues
	d = next2.Sub(prev2)
	assert.Equal(t, gotPrev2, prev2, "Seconds after second run must be 2")
	assert.Equal(t, gotNext2, next2, "Prev and next must be the same")

	prev3 := <-returnValues
	next3 := <-returnValues
	d = next3.Sub(prev3)
	assert.Equal(t, 2.0, d.Seconds(), "Seconds after third run must be 2")
}
