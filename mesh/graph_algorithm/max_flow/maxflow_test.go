package MaxFlow

import (
	"fmt"
	"testing"

	. "github.com/meshcross/algorithm-3rd/mesh/graph_algorithm/graph_struct"
	. "github.com/meshcross/algorithm-3rd/mesh/graph_algorithm/graph_struct/graph_vertex"

	. "github.com/meshcross/algorithm-3rd/mesh/common"
)

/**
* 稍微调整Edge的权重之后，不同的算法得到的计算结果可能并不相同，但是最终都符合最大流的要求
**/
func TestFordFulkerson(t *testing.T) {
	NUM := 6

	ford := NewFordFulkerson()

	creator := func(key, id int) IVertex {
		ptr := NewVertex(key, id)
		return ptr
	}

	_graph := NewGraph(0, NUM, creator) //边的无效权重为0
	for i := 0; i < NUM; i++ {
		_graph.AddVertex(0)
	}

	_graph.AddEdge(NewTuple(0, 1, 16))
	_graph.AddEdge(NewTuple(0, 2, 13))
	_graph.AddEdge(NewTuple(1, 3, 12))
	_graph.AddEdge(NewTuple(2, 1, 4))
	_graph.AddEdge(NewTuple(2, 4, 14))
	_graph.AddEdge(NewTuple(3, 2, 9))
	_graph.AddEdge(NewTuple(3, 5, 20))
	_graph.AddEdge(NewTuple(4, 3, 7))
	_graph.AddEdge(NewTuple(4, 5, 4))

	expect_flow := [][]int{
		[]int{0, 12, 11, 0, 0, 0},
		[]int{0, 0, 0, 12, 0, 0},
		[]int{0, 0, 0, 0, 11, 0},
		[]int{0, 0, 0, 0, 0, 19},
		[]int{0, 0, 0, 7, 0, 4},
		[]int{0, 0, 0, 0, 0, 0},
	}
	flow, _ := ford.MaxFlow(_graph, 0, 5)
	fmt.Println("--expect flow-->", expect_flow)
	fmt.Println("--   get flow-->", flow)

}

func TestGenericPushRelabel(t *testing.T) {
	NUM := 6

	relabel := NewGenericPushRelabel()

	creator := func(key, id int) IVertex {
		// ptr := NewVertex(key, id)
		ptr := NewFlowVertex(0, key, id)
		return ptr
	}

	_graph := NewGraph(0, NUM, creator) //边的无效权重为0
	for i := 0; i < NUM; i++ {
		_graph.AddVertex(0)
	}

	_graph.AddEdge(NewTuple(0, 1, 16))
	_graph.AddEdge(NewTuple(0, 2, 13))
	_graph.AddEdge(NewTuple(1, 3, 12))
	_graph.AddEdge(NewTuple(2, 1, 4))
	_graph.AddEdge(NewTuple(2, 4, 14))
	_graph.AddEdge(NewTuple(3, 2, 9))
	_graph.AddEdge(NewTuple(3, 5, 20))
	_graph.AddEdge(NewTuple(4, 3, 7))
	_graph.AddEdge(NewTuple(4, 5, 4))

	expect_flow := [][]int{
		[]int{0, 12, 11, 0, 0, 0},
		[]int{0, 0, 0, 12, 0, 0},
		[]int{0, 0, 0, 0, 11, 0},
		[]int{0, 0, 0, 0, 0, 19},
		[]int{0, 0, 0, 7, 0, 4},
		[]int{0, 0, 0, 0, 0, 0},
	}

	flow, _ := relabel.MaxFlow(_graph, 0, 5)
	fmt.Println("--expect flow-->", expect_flow)
	fmt.Println("--   get flow-->", flow)

}

func TestRelabelToFront(t *testing.T) {
	NUM := 6
	relabel := NewRelabelToFront()

	creator := func(key, id int) IVertex {
		ptr := NewFrontFlowVertex(key, id)
		return ptr
	}

	_graph := NewGraph(0, NUM, creator) //边的无效权重为0
	for i := 0; i < NUM; i++ {
		_graph.AddVertex(0)
	}

	_graph.AddEdge(NewTuple(0, 1, 16))
	_graph.AddEdge(NewTuple(0, 2, 13))
	_graph.AddEdge(NewTuple(1, 3, 12))
	_graph.AddEdge(NewTuple(2, 1, 4))
	_graph.AddEdge(NewTuple(2, 4, 14))
	_graph.AddEdge(NewTuple(3, 2, 9))
	_graph.AddEdge(NewTuple(3, 5, 20))
	_graph.AddEdge(NewTuple(4, 3, 7))
	_graph.AddEdge(NewTuple(4, 5, 4))

	expect_flow := [][]int{
		[]int{0, 12, 11, 0, 0, 0},
		[]int{0, 0, 0, 12, 0, 0},
		[]int{0, 0, 0, 0, 11, 0},
		[]int{0, 0, 0, 0, 0, 19},
		[]int{0, 0, 0, 7, 0, 4},
		[]int{0, 0, 0, 0, 0, 0},
	}

	flow, _ := relabel.MaxFlow(_graph, 0, 5)
	fmt.Println("--expect flow-->", expect_flow)
	fmt.Println("--   get flow-->", flow)

}
