package creator

import "NodaSoft/hr/golang/bl/bg"

func New() *Creator {
	return &Creator{bg.New()}
}
