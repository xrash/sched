package sched

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestOnceScheduler(t *testing.T) {
	fmt.Println("TestOnceScheduler")

	now := time.Now()

	when := now.Add(time.Second * 2)
	scheduler := NewOnceScheduler(when)

	nextNow := scheduler.Next(now)
	nextAfterAWhile := scheduler.Next(now.Add(time.Second * 1))
	nextOnWhen := scheduler.Next(when)
	nextAfterWhen := scheduler.Next(when.Add(time.Second * 30))

	assert.Equal(t, when.Unix(), nextNow.Unix(), "nextNow must equal when.")
	assert.Equal(t, when.Unix(), nextAfterAWhile.Unix(), "nextNow must equal when.")
	assert.Equal(t, true, nextOnWhen.Year() > 3000, "nextOnWhen must be very far.")
	assert.Equal(t, true, nextAfterWhen.Year() > 3000, "nextOnWhen must be very far.")
}
