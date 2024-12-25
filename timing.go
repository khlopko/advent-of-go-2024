package main

import (
	"errors"
	"time"
)

type Timing struct {
	start *time.Time
	key   *string

	records map[string]time.Duration
}

func NewTiming() *Timing {
	return &Timing{nil, nil, make(map[string]time.Duration)}
}

func (t *Timing) Start(key string) {
	now := time.Now()
	t.start = &now
	t.key = &key
}

func (t *Timing) Finish() error {
	if t.start == nil || t.key == nil {
		return errors.New("timing did not start")
	}
	duration := time.Now().Sub(*t.start)
	t.records[*t.key] = duration
	t.start = nil
	t.key = nil
	return nil
}

func (t *Timing) Accum() time.Duration {
	var sum time.Duration
	for _, value := range t.records {
		sum += value
	}
	return sum
}

