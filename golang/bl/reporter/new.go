package reporter

import (
	"NodaSoft/hr/golang/models"
)

func New() *Reporter {
	return &Reporter{nil, nil, nil, make(map[int]*models.Task), make([]error, 0)}
}
