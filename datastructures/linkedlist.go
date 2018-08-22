package datastructures


type linkedList struct{
	head interface{}
}


func NewLinkedList()(linkedList){

	return linkedList{}
}

func (l *linkedList)Add(item interface{})(bool){
	if l.head == nil{
		l.head = item
	}
	return true
}

func (l *linkedList)GetHead()(interface{}){
	return l.head
}
