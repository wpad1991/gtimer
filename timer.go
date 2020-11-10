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
			ntime := time.Now()
			size := nodelist.Size()

			if size > 0 {

				node := nodelist.head

				for {
					if node == nil {
						break
					}

					dtime := ntime.Sub(*node.value.setTime)
					if dtime.Milliseconds() >= 0 {

						nodelist.RemoveNode(node)
						if dtime.Milliseconds() > ntime.Sub(pretime).Milliseconds() {
							continue
						}

						println(checkTime.Format("2006-01-02 15:04:05"), " : ", node.value.setTime.Format("2006-01-02 15:04:05"))
						node.value.AlertFunc()

						if node.value.cycleType != None {
							switch node.value.cycleType {
							case Daily:
								(*node.value.setTime) = (*node.value.setTime).AddDate(0, 0, 1)
							case Weekly:
								(*node.value.setTime) = (*node.value.setTime).AddDate(0, 0, 7)
							case Monthly:
								(*node.value.setTime) = (*node.value.setTime).AddDate(0, 1, 0)
							case Yearly:
								(*node.value.setTime) = (*node.value.setTime).AddDate(1, 0, 0)
							}
							nodelist.AddNode(node.value)
						}
					}

					node = node.next
				}
			}

			mutex.Unlock()
			dd := time.Now().Sub(checkTime)
			println("Time : ", dd.Milliseconds(), ", cnt : ", size)
			pretime = ntime
			time.Sleep(time.Millisecond * time.Duration(*delay))
		}

	}(&t.isStop, t.mutex, t.alertList, &t.ClockTick)
}

func (t *timer) Stop() {
	t.isStop = false
}

func (t *timer) SetAlertTime(stime *time.Time, f func(), cycle Cycle) {
	node := newAlertNode(stime, f, cycle)
	t.mutex.Lock()
	t.alertList.AddNode(node)
	t.mutex.Unlock()
}
