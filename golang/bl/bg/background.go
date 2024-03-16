package bg

import (
	"NodaSoft/hr/golang/consts"
	"log"
	"sync"
	"time"
)

type Background struct {
	sync.Mutex
	done chan struct{}
}

func New() Background {
	return Background{done: make(chan struct{})}
}

func (b *Background) Complete() {
	select {
	case b.done <- struct{}{}:
	case <-time.After(consts.ClosedChannelTimeout):
		log.Println("(WARN) the channel is already closed")
	}
}

func (b *Background) Done() chan struct{} {
	b.Lock()
	defer b.Unlock()

	if b.done == nil {
		b.done = make(chan struct{})
	}

	return b.done
}
