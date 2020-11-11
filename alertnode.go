package gtimer

import (
	"time"
)

type Cycle int
type WeekCycle int

const (
	None Cycle = 0 + iota
	Daily
	Weekly
	Monthly
	Yearly
)

const (
	WeekNone      WeekCycle = 0
	WeekMonday              = 1
	WeekTuesday             = 2
	WeekWednesday           = 4
	WeekThursday            = 8
	WeekFriday              = 16
	WeekSaturday            = 32
	WeekSunday              = 64
	WeekWeekday             = 128
	WeekWeekend             = 256
)

type alertNode struct {
	setTime       *time.Time // Set Time
	cycleType     Cycle
	weekCycleType WeekCycle
	alertFunc     func()
}

func newAlertNode(stime *time.Time, cycle Cycle, weekcycle WeekCycle, f func()) *alertNode {
	node := alertNode{}
	node.cycleType = cycle
	node.weekCycleType = weekcycle
	node.setTime = stime
	node.alertFunc = f
	return &node
}

func (a *alertNode) AlertFunc() {
	if a.alertFunc != nil {
		a.alertFunc()
	}
}
