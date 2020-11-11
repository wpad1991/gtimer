package gtimer

import (
	"testing"
	"time"
)

func TestList(t *testing.T) {

	list := linkedlist{}

	list.AddNode(newAlertNode(&time.Time{}, None, 0, func() {
		println("1")
	}))

	list.AddNode(newAlertNode(&time.Time{}, None, 0, func() {
		println("2")
	}))

	list.AddNode(newAlertNode(&time.Time{}, None, 0, func() {
		println("3")
	}))

	list.AddNode(newAlertNode(&time.Time{}, None, 0, func() {
		println("4")
	}))

	list.AddNode(newAlertNode(&time.Time{}, None, 0, func() {
		println("5")
	}))

	println("-> ", list.Size())
	list.FindIndex(0).value.AlertFunc()
	list.FindIndex(1).value.AlertFunc()
	list.FindIndex(2).value.AlertFunc()
	list.FindIndex(3).value.AlertFunc()
	list.FindIndex(4).value.AlertFunc()

	list.RemoveIndex(2)
	println("--------------------")
	println("-> ", list.Size())
	list.FindIndex(0).value.AlertFunc()
	list.FindIndex(1).value.AlertFunc()
	list.FindIndex(2).value.AlertFunc()
	list.FindIndex(3).value.AlertFunc()
}
