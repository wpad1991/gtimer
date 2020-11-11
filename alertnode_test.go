package gtimer

import "testing"

func TestEnum(t *testing.T) {

	weektest := WeekMonday + WeekTuesday

	switch weektest {
	case WeekMonday:
		println("It's ?")
	case WeekTuesday:
		println("It's ?")
	}

	println(weektest & WeekMonday)
	println(weektest & WeekTuesday)
	println(weektest & WeekWednesday)

}
