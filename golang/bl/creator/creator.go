package creator

import (
	"NodaSoft/hr/golang/bl/bg"
	"NodaSoft/hr/golang/imodels"
	"NodaSoft/hr/golang/models"
	"fmt"
	"log"
	"time"
)

var _ imodels.Creator = &Creator{}

type Creator struct {
	bg.Background
}

func (c *Creator) Start(bufferSize int) chan *models.Task {
	output := make(chan *models.Task, bufferSize)
	go c.start(output)
	return output
}

func (c *Creator) start(output chan *models.Task) {
	for {
		select {
		case <-c.Done():
			close(output)
			log.Println("(INFO) the creator background routine completed")
			return
		case output <- c.newTask():
		}
	}
}

func (c *Creator) newTask() *models.Task {
	createdOn, err := time.Now(), error(nil)
	//if time.Now().Nanosecond()%2 > 0 { // вот такое условие появления ошибочных тасков
	if createdOn.Nanosecond()%200 > 0 { // в моем случае time.Now().Nanosecond() всегда возвращает числа кратные 100, поэтому %2 всегда равен 0 (не стал разбираться, есть описание проблемы)
		err = fmt.Errorf("some error occurred (CreatedOn nanosec: %d)", createdOn.Nanosecond())
	}
	return &models.Task{CreatedOn: createdOn, ID: int(createdOn.Unix()), Err: err} // передаем таск на выполнение
}

func (c *Creator) Wait(timeout time.Duration) imodels.Creator {
	fmt.Println("sleeping...")
	time.Sleep(timeout)
	fmt.Println("sleeping complete")
	return c
}
