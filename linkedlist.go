package gtimer

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

func (l *linkedlist) AddNode(value *alertNode) {
	if l.head == nil {
		node := linknode{}
		node.value = value
		l.head = &node
		l.tail = &node
	} else {
		tail := l.tail

		node := linknode{}
		node.value = value
		node.idx = l.tail.idx + 1
		node.prev = l.tail

		l.tail = &node
		tail.next = l.tail

	}
}

func (l *linkedlist) AddNodeIndex(value *alertNode, idx int) {
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

		if node.prev == nil {
			println("is NILL?")
		}

	}
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
