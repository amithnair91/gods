package queue_test

import (
	"testing"

	q "github.com/amithnair91/gods/datastructures/queue"

	"github.com/stretchr/testify/assert"
)

func TestShouldBeAbleToCreateNewQueue(t *testing.T) {
	queue := q.NewQueue()

	assert.IsType(t, q.Queue{}, queue)
}

func TestSizeReturnsZeroIfQueueIsEmpty(t *testing.T) {
	queue := q.NewQueue()

	size := queue.Size()

	assert.Equal(t, 0, size)
}

func TestFrontReturnsErrorWhenQueueIsEmpty(t *testing.T) {
	queue := q.NewQueue()

	_, err := queue.Front()

	assert.Error(t, err)
	assert.Equal(t, "queue is empty", err.Error())
}

func TestRearReturnsErrorWhenQueueIsEmpty(t *testing.T) {
	queue := q.NewQueue()

	_, err := queue.Rear()

	assert.Error(t, err)
	assert.Equal(t, "queue is empty", err.Error())
}

func TestEnqueueShouldNotPanic(t *testing.T) {
	queue := q.NewQueue()

	assert.NotPanics(t, func() { queue.Enqueue() })
}

func TestDequeueReturnsErrorWhenQueueIsEmpty(t *testing.T) {
	queue := q.NewQueue()

	_, err := queue.Dequeue()

	assert.Error(t, err)
	assert.Equal(t, "queue is empty", err.Error())
}
