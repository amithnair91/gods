package queue

import "errors"

type Queue struct {
	items []interface{}
}

func NewQueue() Queue {
	return Queue{}
}

func (Queue) Size() int {
	return 0
}

func (q *Queue) Front() (interface{}, error) {

	if len(q.items) < 1 {
		return nil, errors.New("queue is empty")
	}

	return nil, nil
}

func (q *Queue) Rear() (interface{}, error) {

	if len(q.items) < 1 {
		return nil, errors.New("queue is empty")
	}

	return nil, nil
}

func (q *Queue) Enqueue() {

}

func (q *Queue) Dequeue() (interface{}, error) {

	if len(q.items) < 1 {
		return nil, errors.New("queue is empty")
	}

	return nil, nil
}
