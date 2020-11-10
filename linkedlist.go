package gtimer

import (
	"strconv"
)

type linkedlist struct {
	head *linknode
	tail *linknode
}

type linknode struct {
	idx   int
	prev  *linknode
	next  *linknode
	value *alertNode
}

var nodekey = 0

func (l *linkedlist) AddNode(value *alertNode) {

	nodekey++
	if l.head == nil {

		node := linknode{}
		node.value = value
		node.idx = nodekey
		l.head = &node
		l.tail = &node
	} else {
		tail := l.tail
		node := linknode{}
		node.value = value
		node.idx = nodekey
		node.prev = l.tail

		l.tail = &node
		tail.next = l.tail

	}
}

func (l *linkedlist) AddNodeIndex(value *alertNode, idx int) {
	if l.CheckIndex(idx) {
		panic("AddNode already exist index : " + strconv.Itoa(idx))
	}

	if l.head == nil {
		node := linknode{}
		node.idx = idx
		node.value = value
		l.head = &node
		l.tail = &node
	} else {
		tail := l.tail

		node := linknode{}
		node.idx = idx
		node.value = value
		node.prev = l.tail

		l.tail = &node
		tail.next = &node
	}
}

func (l *linkedlist) CheckIndex(index int) bool {

	node := l.head

	for {
		if node == nil {
			break
		}

		if node.idx == index {
			return true
		}

		node = node.next
	}

	return false
}

func (l *linkedlist) RemoveIndex(index int) *linknode {
	node := l.FindIndex(index)

	if node.next != nil {
		node.next.prev = node.prev
	}

	if node.prev != nil {
		node.prev.next = node.next
	}

	return node
}

func (l *linkedlist) RemoveNode(node *linknode) {

	if l.head == node {
		nextnode := node.next

		if nextnode != nil {
			nextnode.prev = nil
		}

		l.head = nextnode

	} else if l.tail == node {
		prenode := node.prev
		l.tail = prenode

		prenode.next = nil
		node.prev = nil
	} else {
		prenode := node.prev
		nexnode := node.next

		prenode.next = nexnode
		nexnode.prev = prenode

		node.prev = nil
		node.next = nil
	}

	if node.prev != nil {
		node.prev.next = node.next
	}
	if node.next != nil {
		node.next.prev = node.prev
	}

}

func (l *linkedlist) FindIndex(index int) *linknode {

	node := l.head

	for i := 0; i < index; i++ {
		node = node.next
	}

	return node
}

func (l *linkedlist) FindNode(fnode *linknode) int {

	size := l.Size()

	node := l.head

	for i := 0; i < size; i++ {
		if node == fnode {
			return i
		}

		node = node.next
	}

	return -1
}

func (l *linkedlist) ScanFunc() {

	node := l.head

	for {
		if node == nil {
			break
		}

		// println("index : ", node.idx)
		// node.value.AlertFunc()
		// println(node.value.AlertFunc)

		println("----scan----")
		println("cur : ", node.idx)
		if node.next != nil {
			println("next : ", node.next.idx)
		}

		if node.prev != nil {
			println("prev : ", node.prev.idx)
		}

		node = node.next
	}
}

func (l *linkedlist) Size() int {

	len := 0

	node := l.head

	for {
		if node == nil {
			break
		}
		len++
		node = node.next
	}

	return len
}
