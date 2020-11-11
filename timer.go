package gtimer

import (
	"sync"
	"time"
)

type timer struct {
	UTCMode   bool  // UTC or LTC
	ClockTick int64 // Clock Interval
	isStop    bool
	mutex     *sync.Mutex
	alertList *linkedlist
}

func NewTimer() *timer {
	t := timer{}
	t.ClockTick = 1000
	t.UTCMode = false
	t.mutex = new(sync.Mutex)
	t.alertList = new(linkedlist)
	t.start()
	return &t
}

func (t *timer) start() {

	go func(stop *bool, mutex *sync.Mutex, nodelist *linkedlist, delay *int64) {

		pretime := time.Time{}

		for !*stop {

			checkTime := time.Now()
			mutex.Lock()
			size := nodelist.Size()
			ntime := time.Now()
			if size > 0 {

				node := nodelist.head

				for {
					if node == nil {
						break
					}

					nextnode := node.next

					if node.value.cycleType == None {
						dtime := ntime.Sub(*node.value.setTime)
						if dtime >= 0 {
							nodelist.RemoveNode(node)
							if dtime.Milliseconds() > ntime.Sub(pretime).Milliseconds() {
								node = node.next
								continue
							}
							node.value.AlertFunc()
						}
					} else {
						dtime := GetDayTimeMilisecond(ntime) - GetDayTimeMilisecond(*node.value.setTime)
						if int64(dtime) < ntime.Sub(pretime).Milliseconds() && dtime > 0 {
							switch node.value.cycleType {
							case Daily:
								node.value.AlertFunc()
							case Weekly:
								ExecWeekFunc(node.value.weekCycleType, ntime, node.value.AlertFunc)
							case Monthly:
								if ntime.Day() == node.value.setTime.Day() {
									if ntime.Month() != node.value.setTime.Month() {
										node.value.AlertFunc()
									}
								}
							case Yearly:
								if ntime.Day() == node.value.setTime.Day() {
									if ntime.Month() == node.value.setTime.Month() {
										if ntime.Year() == node.value.setTime.Year() {
											node.value.AlertFunc()
										}
									}
								}
							}
						}
					}
					node = nextnode
				}
			}

			mutex.Unlock()
			//dd := time.Now().Sub(checkTime)
			//println("Time : ", dd.Milliseconds(), ", cnt : ", size)
			pretime = ntime

			time.Sleep(time.Millisecond * time.Duration(*delay))
		}

	}(&t.isStop, t.mutex, t.alertList, &t.ClockTick)
}

func (t *timer) Stop() {
	t.isStop = false
}

func (t *timer) SetAlertTime(stime *time.Time, cycle Cycle, weekcycle WeekCycle, f func()) {
	node := newAlertNode(stime, cycle, weekcycle, f)
	t.mutex.Lock()
	t.alertList.AddNode(node)
	t.mutex.Unlock()
}

func GetDayTimeMilisecond(t time.Time) int {

	totalmili := 0

	totalmili += t.Hour() * 1000 * 60 * 60
	totalmili += t.Minute() * 1000 * 60
	totalmili += t.Second() * 1000
	totalmili += t.Nanosecond() / 1000000

	return totalmili
}

func ExecWeekFunc(wc WeekCycle, dt time.Time, f func()) {

	if wc&WeekMonday != 0 {
		if dt.Weekday() == time.Monday {
			f()
		}
	}
	if wc&WeekTuesday != 0 {
		if dt.Weekday() == time.Tuesday {
			f()
		}
	}
	if wc&WeekWednesday != 0 {
		if dt.Weekday() == time.Wednesday {
			f()
		}
	}
	if wc&WeekThursday != 0 {
		if dt.Weekday() == time.Thursday {
			f()
		}
	}
	if wc&WeekFriday != 0 {
		if dt.Weekday() == time.Friday {
			f()
		}
	}
	if wc&WeekSaturday != 0 {
		if dt.Weekday() == time.Saturday {
			f()
		}
	}
	if wc&WeekSunday != 0 {
		if dt.Weekday() == time.Sunday {
			f()
		}
	}
	if wc&WeekWeekday != 0 {
		if dt.Weekday() != time.Saturday && dt.Weekday() != time.Sunday {
			f()
		}
	}
	if wc&WeekWeekend != 0 {
		if dt.Weekday() == time.Saturday || dt.Weekday() == time.Sunday {
			f()
		}
	}
}
