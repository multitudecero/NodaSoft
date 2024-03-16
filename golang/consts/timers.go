package consts

import "time"

const (
	ClosedChannelTimeout = 3 * time.Second
	WorkingTime          = 3 * time.Second
	WorkerWaiting        = 150 * time.Millisecond
	TaskExpiration       = 20 * time.Second
)
