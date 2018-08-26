package datastructures

type singlyLinkedList struct {
	head *Element
	tail *Element
}

type Element struct {
	value interface{}
	next  *Element
}

func NewSinglyLinkedList() singlyLinkedList {
	return singlyLinkedList{}
}

func (l *singlyLinkedList) Add(item interface{}) bool {
	if l.head == nil {
		headElem := Element{value: item, next: nil}
		l.head = &headElem
	} else if l.tail == nil {
		nextElement := Element{item, nil}
		l.tail = &nextElement
		l.head.next = &nextElement
	} else {
		nextElement := Element{item, nil}
		l.tail.next = &nextElement
		l.tail = &nextElement
	}
	return true
}

func (l *singlyLinkedList) Remove(value interface{}) bool {

	removed := findAndRemoveElement(nil, l.head, value)
	if removed {
		return true
	}

	return false
}

//recursion
//O(n) time complexity
func findAndRemoveElement(previousElement, elem *Element, value interface{}) bool {

	if elem.value.(string) == value.(string) {
		nextElement := elem.next
		previousElement.next = nextElement
		return true
	} else if elem.next != nil {
		return findAndRemoveElement(elem, elem.next, value)
	}

	return false
}

func (l *singlyLinkedList) Head() Element {
	return *l.head
}

func (l *singlyLinkedList) Tail() Element {
	if l.tail == nil {
		return Element{}
	}
	return *l.tail
}

func (l *singlyLinkedList) String() []string {

	var items []string
	toStringArray(l.head, &items)

	return items
}

func (e *Element) Next() *Element {
	return e.next
}

func (e *Element) Value() interface{} {
	return e.value
}

func toStringArray(elem *Element, items *[]string) *[]string {

	if elem.value != nil {
		*items = append(*items, elem.value.(string))
	}

	if elem.next != nil {
		return toStringArray(elem.next, items)
	}

	return items
}
