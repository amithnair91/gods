package datastructures


type LinkedList struct{
	items *[]Item
	currentItem *Item
}

type Item struct{
	value struct{}
	next Item
}

func (l *LinkedList)GetHead()(item Item){
	return (*l.items)[0]
}


func (l *LinkedList)GetTail()(item Item){
	lastIndex := len(*l.items)
	return (*l.items)[lastIndex]
}


func (l *LinkedList)Next()(item Item){
	return l.currentItem.next
}
