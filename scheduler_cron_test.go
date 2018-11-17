package sched

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCronScheduler(t *testing.T) {
	fmt.Println("TestCronScheduler")
	scheduler, err := NewCronScheduler("* * * * *")
	assert.Equal(t, nil, err, "Error must be nil.")

	now := time.Now()

	next1 := scheduler.Next(now)
	minutes1 := next1.Sub(now).Minutes()
	next2 := scheduler.Next(next1)
	minutes2 := next2.Sub(next1).Minutes()
	next3 := scheduler.Next(next2)
	minutes3 := next3.Sub(next2).Minutes()

	assert.Equal(t, true, minutes1 <= 1.0, "Must be between one minute.")
	assert.Equal(t, true, minutes1 > 0, "Must be between one minute.")

	assert.Equal(t, true, minutes2 <= 1.0, "Must be between one minute.")
	assert.Equal(t, true, minutes2 > 0, "Must be between one minute.")

	assert.Equal(t, true, minutes3 <= 1.0, "Must be between one minute.")
	assert.Equal(t, true, minutes3 > 0, "Must be between one minute.")
}
