package models

import (
	"fmt"
	"time"
)

// Task A Ttype represents a meaninglessness of our life
type Task struct {
	ID        int
	CreatedOn time.Time // время создания
	StartedOn time.Time // время выполнения
	Err       error
	Result    []byte
}

func (t *Task) String() string {
	return fmt.Sprintf("id: %d, cT: %s, nano: %d, fT: %s, taskRESULT: %v", t.ID, t.CreatedOn.Format(time.RFC3339), t.CreatedOn.Nanosecond(), t.StartedOn.Format(time.RFC3339Nano), string(t.Result))
}
