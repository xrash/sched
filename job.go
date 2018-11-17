package sched

import (
	"time"
)

type Job struct {
	task        Task
	scheduler   Scheduler
	stopChannel chan bool

	prev time.Time
	next time.Time
}

func LaunchJob(task Task, scheduler Scheduler) *Job {
	j := &Job{
		task:        task,
		scheduler:   scheduler,
		stopChannel: make(chan bool, 10),
	}

	j.prev = time.Now()
	j.next = j.scheduler.Next(j.prev)

	go j.start()

	return j
}

func (j *Job) Stop() (time.Time, time.Time) {
	j.stopChannel <- true
	return j.prev, j.next
}

func (j *Job) Prev() time.Time {
	return j.prev
}

func (j *Job) Next() time.Time {
	return j.next
}

func (j *Job) Scheduler() Scheduler {
	return j.scheduler
}

func (j *Job) Task() Task {
	return j.task
}

func (j *Job) start() {
	for {
		d := j.next.Sub(j.prev)
		select {
		case <-time.After(d):
			j.task(j.prev, j.next)
			j.prev = j.next
			j.next = j.scheduler.Next(j.prev)
		case <-j.stopChannel:
			return
		}
	}
}
