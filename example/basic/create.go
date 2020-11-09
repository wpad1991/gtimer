package main

import (
	"time"

	"github.com/wpad1991/gtimer"
)

func main() {

	timer := gtimer.NewTimer()

	ntime := time.Now()

	for i := 0; i < 100; i++ {
		d := ntime.Add(time.Duration(10000 * (i + 1)))
		timer.SetAlertTime(&d, func() { println(i) })
	}

	time.Sleep(time.Millisecond * 600000)
}
