package graph_test

import (
	"testing"

	g "github.com/amithnair91/gods/datastructures/graph"

	"github.com/stretchr/testify/assert"
)

func TestShouldBeAbleToCreateNewGraph(t *testing.T) {
	graph := g.NewGraph()

	assert.IsType(t, g.Graph{}, graph)
}
