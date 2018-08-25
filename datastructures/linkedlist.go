package datastructures

type singlyLinkedList struct {
	head Element
	tail Element
}

type Element struct {
	value interface{}
	next  *Element
}

func NewSinglyLinkedList() singlyLinkedList {
	return singlyLinkedList{}
}

func (l *singlyLinkedList) Add(item interface{}) bool {
	if l.head.value == nil {
		l.head.value = item
		l.head.next = nil
	} else if l.tail.value == nil {
		nextElement := Element{item, nil}
		l.tail = nextElement
		l.head.next = &nextElement
	} else {
		nextElement := Element{item, nil}
		l.tail.next = &nextElement
	}
	return true
}

func (l *singlyLinkedList) Remove(value interface{}) bool {

	_, exist := findElement(&l.head, value)
	if exist {
		return true
	}

	return false
}

//recursion
func findElement(elem *Element, value interface{}) (Element, bool) {
	if elem.value.(string) == value.(string) {
		return *elem, true
	} else if elem.next != nil {
		return findElement(elem.next, value)
	}

	return Element{}, false
}

func (l *singlyLinkedList) Head() Element {
	return l.head
}

func (l *singlyLinkedList) Tail() Element {
	return l.tail
}

func (e *Element) Next() *Element {
	return e.next
}

func (e *Element) Value() interface{} {
	return e.value
}
