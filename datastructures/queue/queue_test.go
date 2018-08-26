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

	assert.NotPanics(t, func() { queue.Enqueue("item") })
}

func TestDequeueReturnsErrorWhenQueueIsEmpty(t *testing.T) {
	queue := q.NewQueue()

	_, err := queue.Dequeue()

	assert.Error(t, err)
	assert.Equal(t, "queue is empty", err.Error())
}

func TestFrontShouldReturnTheFirstElement(t *testing.T) {
	expectedItem := "firstItem"
	queue := q.NewQueue()

	queue.Enqueue(expectedItem)
	queue.Enqueue("lastItem")
	frontItem, _ := queue.Front()

	assert.Equal(t, expectedItem, frontItem)
}

func TestRearShouldReturnTheLastElement(t *testing.T) {
	expectedItem := "rearItem"
	queue := q.NewQueue()

	queue.Enqueue("firstItem")
	queue.Enqueue("secondItem")
	queue.Enqueue(expectedItem)
	rearItem, _ := queue.Rear()

	assert.Equal(t, expectedItem, rearItem)
}

func TestDequeueShouldRemoveTheFirstElement(t *testing.T) {
	expectedItem := "firstItem"
	queue := q.NewQueue()

	queue.Enqueue(expectedItem)
	queue.Enqueue("secondItem")
	queue.Enqueue("thirdItem")

	removedItem, _ := queue.Dequeue()

	assert.Equal(t, expectedItem, removedItem)
}

func TestShouldReturnSizeOfItemsInQueue(t *testing.T) {
	queue := q.NewQueue()

	queue.Enqueue("firstItem")
	queue.Enqueue("secondItem")
	queue.Enqueue("thirdItem")

	size := queue.Size()

	assert.Equal(t, 3, size)
}

func TestShouldBeAbleToEnqueueItemsOfDifferentDataTypes(t *testing.T) {
	firstItem := "firstItem"
	secondItem := 2
	thirdItem := 3.14
	queue := q.NewQueue()

	queue.Enqueue(firstItem)
	queue.Enqueue(secondItem)
	queue.Enqueue(thirdItem)

	firstDequeue, _ := queue.Dequeue()
	secondDequeue, _ := queue.Dequeue()
	thirdDequeue, _ := queue.Dequeue()

	assert.Equal(t, firstItem, firstDequeue)
	assert.Equal(t, secondItem, secondDequeue)
	assert.Equal(t, thirdItem, thirdDequeue)
}

func TestShouldBeAbleToEnqueueAndDequeue(t *testing.T) {
	firstItem := "firstItem"
	secondItem := "secondItem"
	thirdItem := "thirdItem"
	queue := q.NewQueue()

	queue.Enqueue(firstItem)
	queue.Enqueue(secondItem)
	queue.Enqueue(thirdItem)

	assert.Equal(t, 3, queue.Size())

	firstDequeue, _ := queue.Dequeue()

	assert.Equal(t, firstItem, firstDequeue)
	assert.Equal(t, 2, queue.Size())

	secondDequeue, _ := queue.Dequeue()
	assert.Equal(t, secondItem, secondDequeue)
	assert.Equal(t, 1, queue.Size())

	thirdDequeue, _ := queue.Dequeue()
	assert.Equal(t, thirdItem, thirdDequeue)
	assert.Equal(t, 0, queue.Size())
}
