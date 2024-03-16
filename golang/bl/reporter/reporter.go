package reporter

import (
	"NodaSoft/hr/golang/imodels"
	"NodaSoft/hr/golang/models"
	"fmt"
	"log"
	"time"
)

var _ imodels.Reporter = &Reporter{}

type Reporter struct {
	doneTasks chan *models.Task
	errs      chan error
	done      chan struct{}

	result   map[int]*models.Task
	troubles []error
}

func (r *Reporter) Start(doneTasks chan *models.Task, errs chan error) {
	r.doneTasks, r.errs, r.done = doneTasks, errs, make(chan struct{})
	go r.start()
}

func (r *Reporter) Wait(timeout time.Duration) {
	select {
	case <-r.done:
	case <-time.After(timeout):
	}
}

func (r *Reporter) start() {
	defer log.Println("(INFO) the reporter background routine completed")
	defer r.finishAndReport()

	for {
		select {
		case doneTask, more := <-r.doneTasks:
			if !more {
				return
			}
			r.result[doneTask.ID] = doneTask
		case err, more := <-r.errs:
			if !more {
				return
			}
			r.troubles = append(r.troubles, err)
		}
	}
}

func (r *Reporter) finishAndReport() {
	defer close(r.done)

	fmt.Println("Errors:")
	for r, trouble := range r.troubles {
		fmt.Println(r, trouble)
	}

	fmt.Println("Done tasks:")
	for r, task := range r.result {
		fmt.Println(r, task, task.CreatedOn.Nanosecond())
	}
}
