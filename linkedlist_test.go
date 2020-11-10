package gtimer

import (
	"testing"
	"time"
)

func TestList(t *testing.T) {

	list := linkedlist{}

	list.AddNode(newAlertNode(&time.Time{}, func() {
		println("1")
	}, None))

	list.AddNode(newAlertNode(&time.Time{}, func() {
		println("2")
	}, None))

	list.AddNode(newAlertNode(&time.Time{}, func() {
		println("3")
	}, None))

	list.AddNode(newAlertNode(&time.Time{}, func() {
		println("4")
	}, None))

	list.AddNode(newAlertNode(&time.Time{}, func() {
		println("5")
	}, None))

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
