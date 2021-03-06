package main

import (
	"time"

	"github.com/wpad1991/gtimer"
)

func main() {

	timer := gtimer.NewTimer()

	ntime := time.Now()

	for i := 0; i < 10000; i++ {
		d := ntime.Add(time.Duration(time.Millisecond * (1000 * time.Duration(i/2+1))))
		val := i
		timer.SetAlertTime(&d, gtimer.None, 0, func() { println("value : ", val) })
	}

	time.Sleep(time.Millisecond * 101000)
}
