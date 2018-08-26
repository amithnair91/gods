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
