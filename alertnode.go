package gtimer

import (
	"time"
)

type Cycle int

const (
	None Cycle = 0 + iota
	Daily
	Weekly
	Monthly
	Yearly
)

type alertNode struct {
	setTime   *time.Time // Set Time
	cycleType Cycle
	alertFunc func()
}

func newAlertNode(stime *time.Time, f func(), cycle Cycle) *alertNode {
	node := alertNode{}
	node.cycleType = cycle
	node.setTime = stime
	node.alertFunc = f
	return &node
}

func (a *alertNode) AlertFunc() {
	if a.alertFunc != nil {
		a.alertFunc()
	}
}
