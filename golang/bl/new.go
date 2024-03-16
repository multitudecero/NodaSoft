package bl

import (
	"NodaSoft/hr/golang/bl/creator"
	"NodaSoft/hr/golang/bl/reporter"
	"NodaSoft/hr/golang/bl/sorter"
	"NodaSoft/hr/golang/bl/worker"
	"NodaSoft/hr/golang/imodels"
)

func New(bufferSize int) (imodels.Creator, imodels.Reporter) {
	cr, rp := creator.New(), reporter.New()
	doneTasks, errs := sorter.New().Start(
		worker.New().Start(
			cr.Start(bufferSize)))

	rp.Start(doneTasks, errs)

	return cr, rp
}
