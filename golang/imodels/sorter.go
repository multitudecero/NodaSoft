package imodels

import "NodaSoft/hr/golang/models"

type Sorter interface {
	Start(input chan *models.Task) (chan *models.Task, chan error)
}
