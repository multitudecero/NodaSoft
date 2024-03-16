package worker

import (
	"NodaSoft/hr/golang/consts"
	"NodaSoft/hr/golang/imodels"
	"NodaSoft/hr/golang/models"
	"log"
	"time"
)

var _ imodels.Worker = &Worker{}

type Worker struct{}

func (w *Worker) Start(input chan *models.Task) chan *models.Task {
	output := make(chan *models.Task)
	go w.start(input, output)
	return output
}

func (w *Worker) start(input, output chan *models.Task) {
	defer func() {
		close(output)
		log.Println("(INFO) the worker background routine completed")
	}()

	for task := range input {
		output <- w.work(task)
	}
}

func (w *Worker) work(task *models.Task) *models.Task {
	if task.CreatedOn.After(time.Now().Add(-consts.TaskExpiration)) {
		task.Result = []byte("task completed successfully")
	} else {
		task.Result = []byte("something went wrong")
	}
	task.StartedOn = time.Now()

	time.Sleep(consts.WorkerWaiting)

	return task
}
