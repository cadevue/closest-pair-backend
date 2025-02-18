package internal

import "time"

type ExecTimer struct {
	startTime   time.Time
	elapsedTime time.Duration
}

func (t *ExecTimer) Start() {
	t.startTime = time.Now()
}

func (t *ExecTimer) Stop() {
	t.elapsedTime = time.Since(t.startTime)
}

func (t *ExecTimer) GetElapsedTime() time.Duration {
	return t.elapsedTime
}
