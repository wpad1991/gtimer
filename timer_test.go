package gtimer

import (
	"testing"
	"time"
)

func TestExec(t *testing.T) {

	timer := NewTimer()

	ntime := time.Now()

	qwe := make([]func(), 10)

	qwe[0] = func() { println(0) }
	qwe[1] = func() { println(1) }
	qwe[2] = func() { println(2) }
	qwe[3] = func() { println(3) }
	qwe[4] = func() { println(4) }
	qwe[5] = func() { println(5) }
	qwe[6] = func() { println(6) }
	qwe[7] = func() { println(7) }
	qwe[8] = func() { println(8) }
	qwe[9] = func() { println(9) }

	for i := 0; i < 10; i++ {
		d := ntime.Add(time.Duration(10000 * (i + 1)))
		qq := i
		f := func() {
			println(qq)
		}

		timer.SetAlertTime(&d, f)
	}

	time.Sleep(time.Millisecond * 600000)
}

func TestNodes(t *testing.T) {

	var alertList *linkedlist
	alertList = &linkedlist{}

	ntime := time.Now()

	qwe := make([]func(), 10)

	qwe[0] = func() { println(0) }
	qwe[1] = func() { println(1) }
	qwe[2] = func() { println(2) }
	qwe[3] = func() { println(3) }
	qwe[4] = func() { println(4) }
	qwe[5] = func() { println(5) }
	qwe[6] = func() { println(6) }
	qwe[7] = func() { println(7) }
	qwe[8] = func() { println(8) }
	qwe[9] = func() { println(9) }

	for i := 0; i < 10; i++ {
		d := ntime.Add(time.Duration(1000 * (i + 1)))

		qwe := i
		f := func() {
			if qwe == 0 {
				println(123)
			} else {
				println(qwe)
			}

		}
		println("-----------111111")
		println(f)
		n := newAlertNode(&d, f)
		println(n.AlertFunc)
		println(f)
		alertList.AddNodeIndex(n, i)

		//alertList.AddNodeIndex(newAlertNode(&d, qwe[i]), i)
	}

	println("--------------33333")
	alertList.ScanFunc()

}

func TestTime(t *testing.T) {

	t1 := time.Date(2010, time.April, 2, 0, 0, 0, 0, time.UTC)
	t2 := time.Date(2010, time.April, 1, 0, 0, 0, 0, time.UTC)

	t3 := t2.Sub(t1)

	println(t3.Milliseconds())
}

func TestAppend(t *testing.T) {

	cnt := 100000000

	tt1 := time.Now()
	t1 := make([]int, 0)
	for i := 0; i < cnt; i++ {
		t1 = append(t1, i)
	}

	tt2 := time.Now()
	t2 := make([]int, cnt)
	for i := 0; i < cnt; i++ {
		t2[i] = i
	}

	tt3 := time.Now()

	aa1 := tt2.Sub(tt1)
	aa2 := tt3.Sub(tt2)

	println(aa1.Milliseconds())
	println(aa2.Milliseconds())

}

type test struct {
	qwe int
}

func TestCopy(t *testing.T) {

	t1 := make([]test, 10)
	t2 := t1

	for i := 0; i < 10; i++ {
		t2[i].qwe = i
	}

	println(&t1)
	println(&t2)

	t2 = t2[0:2]

	println(t1)
	println(t2)

	for _, val := range t1 {
		println(&val.qwe)
	}

	println("------------------")
	for _, val := range t2 {
		println(&val.qwe)
	}

}

func TestCopySlice(t *testing.T) {

	t1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	t2 := make([]int, 4)

	t2 = t1[:2]

	for _, val := range t2 {
		println(val)
	}
}

func TestCopyTest(t *testing.T) {

	t1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	t2 := t1[:0]

	for _, val := range t2 {
		println(val)
	}

}
