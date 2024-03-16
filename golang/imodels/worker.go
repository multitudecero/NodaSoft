package imodels

import "NodaSoft/hr/golang/models"

type Worker interface {
	Start(input chan *models.Task) chan *models.Task
}
