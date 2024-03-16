package sorter

import (
	"NodaSoft/hr/golang/imodels"
	"NodaSoft/hr/golang/models"
	"fmt"
	"log"
	"time"
)

var _ imodels.Sorter = &Sorter{}

type Sorter struct{}

func (s *Sorter) Start(input chan *models.Task) (chan *models.Task, chan error) {
	doneTasks, errs := make(chan *models.Task), make(chan error)
	go s.start(input, doneTasks, errs)
	return doneTasks, errs
}

func (s *Sorter) start(input chan *models.Task, doneTasks chan *models.Task, errs chan error) {
	defer func() {
		close(doneTasks)
		close(errs)
		log.Println("(INFO) the sorter background routine completed")
	}()

	for task := range input {
		s.sort(task, doneTasks, errs)
	}
}

func (s *Sorter) sort(task *models.Task, doneTasks chan *models.Task, errs chan error) {
	if task.Err == nil {
		doneTasks <- task
	} else {
		errs <- fmt.Errorf("taskID: %d, time: %s, error: %v", task.ID, task.CreatedOn.Format(time.RFC3339), task.Err)
	}
}
