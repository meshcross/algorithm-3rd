package AllNodePairShortestPath

import (
	"fmt"
	"testing"

	. "github.com/meshcross/algorithm-3rd/mesh/graph_algorithm/graph_struct"
	. "github.com/meshcross/algorithm-3rd/mesh/graph_algorithm/graph_struct/graph_vertex"

	. "github.com/meshcross/algorithm-3rd/mesh/common"
)

func TestMatrxSP(t *testing.T) {
	NUM := 5

	sp := NewMatrixSP()

	creator := func(key, id int) IVertex {
		ptr := NewVertex(key, id)
		return ptr
	}

	_graph := NewGraph(0, NUM, creator) //边的无效权重为0
	for i := 0; i < NUM; i++ {
		_graph.AddVertex(0)
	}

	_graph.AddEdge(NewTuple(0, 1, 3))
	_graph.AddEdge(NewTuple(0, 2, 8))
	_graph.AddEdge(NewTuple(0, 4, -4))
	_graph.AddEdge(NewTuple(1, 3, 1))
	_graph.AddEdge(NewTuple(1, 4, 7))
	_graph.AddEdge(NewTuple(2, 1, 4))
	_graph.AddEdge(NewTuple(3, 2, -5))
	_graph.AddEdge(NewTuple(3, 0, 2))
	_graph.AddEdge(NewTuple(4, 3, 6))

	expect_result := [][]int{
		[]int{0, 1, -3, 2, -4},
		[]int{3, 0, -4, 1, -1},
		[]int{7, 4, 0, 5, 3},
		[]int{2, -1, -5, 0, -2},
		[]int{8, 5, 1, 6, 0},
	}

	result, _ := sp.ShortestPathFast(_graph)
	fmt.Println("--expect result-->", expect_result)
	fmt.Println("--   get result-->", result)
}
func TestFloydWarshallSP(t *testing.T) {
	NUM := 5

	sp := NewFloydWarshallSP()

	creator := func(key, id int) IVertex {
		ptr := NewFrontFlowVertex(key, id)
		return ptr
	}

	_graph := NewGraph(0, NUM, creator) //边的无效权重为0
	for i := 0; i < NUM; i++ {
		_graph.AddVertex(0)
	}

	_graph.AddEdge(NewTuple(0, 1, 3))
	_graph.AddEdge(NewTuple(0, 2, 8))
	_graph.AddEdge(NewTuple(0, 4, -4))
	_graph.AddEdge(NewTuple(1, 3, 1))
	_graph.AddEdge(NewTuple(1, 4, 7))
	_graph.AddEdge(NewTuple(2, 1, 4))
	_graph.AddEdge(NewTuple(3, 2, -5))
	_graph.AddEdge(NewTuple(3, 0, 2))
	_graph.AddEdge(NewTuple(4, 3, 6))

	expect_P := [][]int{
		[]int{-1, 2, 3, 4, 0},
		[]int{3, -1, 3, 1, 0},
		[]int{3, 2, -1, 1, 0},
		[]int{3, 2, 3, -1, 0},
		[]int{3, 2, 3, 4, -1},
	}

	expect_D := [][]int{
		[]int{0, 1, -3, 2, -4},
		[]int{3, 0, -4, 1, -1},
		[]int{7, 4, 0, 5, 3},
		[]int{2, -1, -5, 0, -2},
		[]int{8, 5, 1, 6, 0},
	}

	d, p, _ := sp.ShortestPath(_graph)
	fmt.Println("--expect D-->", expect_D)
	fmt.Println("--   get D-->", d)
	fmt.Println("--expect P-->", expect_P)
	fmt.Println("--   get P-->", p)

}

func TestJohnsonSP(t *testing.T) {
	NUM := 5

	sp := NewJohnsonSP()

	creator := func(key, id int) IVertex {
		ptr := NewVertex(key, id)
		return ptr
	}

	unlimit := Unlimit()
	_graph := NewGraph(unlimit, NUM, creator) //边的无效权重为0
	for i := 0; i < NUM; i++ {
		_graph.AddVertex(0)
	}

	_graph.AddEdge(NewTuple(0, 1, 3))
	_graph.AddEdge(NewTuple(0, 2, 8))
	_graph.AddEdge(NewTuple(0, 4, -4))
	_graph.AddEdge(NewTuple(1, 3, 1))
	_graph.AddEdge(NewTuple(1, 4, 7))
	_graph.AddEdge(NewTuple(2, 1, 4))
	_graph.AddEdge(NewTuple(3, 2, -5))
	_graph.AddEdge(NewTuple(3, 0, 2))
	_graph.AddEdge(NewTuple(4, 3, 6))

	expect_result := [][]int{
		[]int{0, 1, -3, 2, -4},
		[]int{3, 0, -4, 1, -1},
		[]int{7, 4, 0, 5, 3},
		[]int{2, -1, -5, 0, -2},
		[]int{8, 5, 1, 6, 0},
	}

	result, _ := sp.ShortestPath(_graph)
	fmt.Println("--expect result-->", expect_result)
	fmt.Println("--   get result-->", result)
}
