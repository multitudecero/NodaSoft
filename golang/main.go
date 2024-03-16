package main

import (
	"NodaSoft/hr/golang/bl"
	"NodaSoft/hr/golang/consts"
)

// ЗАДАНИЕ:
// * сделать из плохого кода хороший;
// * важно сохранить логику появления ошибочных тасков;
// * сделать правильную мультипоточность обработки заданий.
// Обновленный код отправить через merge-request.

// приложение эмулирует получение и обработку тасков, пытается и получать и обрабатывать в многопоточном режиме
// В конце должно выводить успешные таски и ошибки выполнены остальных тасков

func main() {
	cr, rp := bl.New(consts.SuperChannelBufferSize)
	cr.Wait(consts.WorkingTime).Complete()
	rp.Wait(consts.ClosedChannelTimeout)
}
