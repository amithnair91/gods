package datastructures

type linkedList struct {
	head Element
	tail Element
}

type Element struct {
	value interface{}
	next  *Element
}

func NewLinkedList() (linkedList) {
	return linkedList{}
}

func (l *linkedList) Add(item interface{}) (bool) {
	if l.head.value == nil {
		l.head.value = item
		l.head.next = nil
	} else if l.tail.value == nil {
		nextElement := Element{item, nil}
		l.tail = nextElement
		l.head.next = &nextElement
	} else {
		nextElement := Element{item, nil,}
		l.tail.next = &nextElement
	}
	return true
}

func (l *linkedList) Head() (Element) {
	return l.head
}

func (l *linkedList) Tail() (Element) {
	return l.tail
}

func (e *Element) Next() (*Element) {
	return e.next
}

func (e *Element) Value() (interface{}) {
	return e.value
}
