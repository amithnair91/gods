package queue

import (
	"errors"
)

type Queue struct {
	items []interface{}
}

func NewQueue() Queue {
	return Queue{}
}

func (q *Queue) Size() int {
	return len(q.items)
}

func (q *Queue) Front() (interface{}, error) {

	if len(q.items) < 1 {
		return nil, errors.New("queue is empty")
	}

	return q.items[0], nil
}

func (q *Queue) Rear() (interface{}, error) {

	if len(q.items) < 1 {
		return nil, errors.New("queue is empty")
	}

	return q.items[len(q.items)-1], nil
}

func (q *Queue) Enqueue(item interface{}) {
	q.items = append(q.items, item)
}

func (q *Queue) Dequeue() (interface{}, error) {

	if len(q.items) < 1 {
		return nil, errors.New("queue is empty")
	}

	itemToRemove := q.items[0]
	q.items = q.items[1:]

	return itemToRemove, nil
}
