package imodels

import (
	"NodaSoft/hr/golang/models"
	"time"
)

type Reporter interface {
	Start(chan *models.Task, chan error)
	Wait(timeout time.Duration)
}
