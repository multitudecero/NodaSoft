package imodels

import (
	"NodaSoft/hr/golang/models"
	"time"
)

type Creator interface {
	Start(bufferSize int) chan *models.Task
	Complete()
	Wait(timeout time.Duration) Creator
}
