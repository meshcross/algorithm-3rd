/*
 * @Description: 单源最短路径测试
 * @Author: wangchengdg@gmail.com
 * @Date: 2020-02-26 23:01:47
 * @LastEditTime: 2020-03-13 22:14:10
 * @LastEditors:
 */
package SingleSourceShortestPath

import (
	"fmt"
	"testing"

	. "github.com/meshcross/algorithm-3rd/mesh/graph_algorithm/graph_struct"

	. "github.com/meshcross/algorithm-3rd/mesh/common"
	. "github.com/meshcross/algorithm-3rd/mesh/graph_algorithm/graph_struct/graph_vertex"
)

/**
 * @description:BellmanFord单源最短路径
 */
func TestBellmanFord(t *testing.T) {
	NUM := 10
	creator := func(key, id int) IVertex {
		return NewVertex(key, id)
	}
	////////////////start setup/////////////////////////

	_1v_graph := NewGraph(-1, NUM, creator) //边的无效权重为-1
	_1v_graph.AddVertex(0)                  //该图只有一个顶点

	_1e_graph := NewGraph(-1, NUM, creator) //边的无效权重为-1
	_1e_graph.AddVertex(0)
	_1e_graph.AddVertex(0)
	_1e_graph.AddEdge(NewTuple(0, 1, 1)) //该图只有一条边

	//****  含顶点图和边图：10个顶点，9条边：0-->1-->2....-->9(权重均为1)   ****
	_normal_graph := NewGraph(-1, NUM, creator) //边的无效权重为-1
	for i := 0; i < NUM; i++ {
		_normal_graph.AddVertex(0)
	}
	for i := 0; i < NUM-1; i++ {
		_normal_graph.AddEdge(NewTuple(i, i+1, 1)) //该图的边是从左到右组成一个链条
	}

	//****  含顶点图和边图：10个顶点，10条边 :0-->1-->2....-->9(权重均为1)以及 9-->8(权重-2)  ****
	_minus_graph := NewGraph(-1, NUM, creator) //边的无效权重为-1
	for i := 0; i < NUM; i++ {
		_minus_graph.AddVertex(0)
	}
	for i := 0; i < NUM-1; i++ {
		_minus_graph.AddEdge(NewTuple(i, i+1, 1)) //链条
	}
	_minus_graph.AddEdge(NewTuple(9, 8, -2)) //最后构成一个负权值回路,-2可以测试出来，-1就不行，因为-1时候环路权重为0，并不是负值

	bellman := NewBellmanFordShortestPath()

	//************  只有一个顶点的图  ************
	bellman.ShortestPath(_1v_graph, 0)

	fmt.Println(fmt.Sprintf("a-EXPECT_EQ(%d,%d)", _1v_graph.Vertexes[0].GetKey(), 0))

	//**********  只有一条边的图  ***************
	bellman.ShortestPath(_1e_graph, 0)
	fmt.Println(fmt.Sprintf("b-EXPECT_EQ(%d,%d)", _1e_graph.Vertexes[0].GetKey(), 0))
	fmt.Println(fmt.Sprintf("b-EXPECT_EQ(%d,%d)", _1e_graph.Vertexes[1].GetKey(), 1))
	fmt.Println(fmt.Sprintf("b-EXPECT_EQ(%v,%v)", _1e_graph.Vertexes[1].GetParent(), _1e_graph.Vertexes[0]))

	//**********  链边的图 ***************
	bellman.ShortestPath(_normal_graph, 0)
	fmt.Println(fmt.Sprintf("c-EXPECT_EQ(%d,%d)", _normal_graph.Vertexes[0].GetKey(), 0))
	for i := 1; i < NUM; i++ {
		fmt.Println(fmt.Sprintf("c-EXPECT_EQ(%d,%d)", _normal_graph.Vertexes[i].GetKey(), i))
		fmt.Println(fmt.Sprintf("c-EXPECT_EQ(%v,%v)", _normal_graph.Vertexes[i].GetParent(), _normal_graph.Vertexes[i-1]))
	}

	//*********** 含有负权值回路的图  ***********
	d, _ := bellman.ShortestPath(_minus_graph, 0)
	fmt.Println(fmt.Sprintf("d-EXPECT_EQ(%v,%d)", d, 0))
}

// func TestDagShortestPath(t *testing.T) {
// 	NUM := 10

// 	dag := NewDagShortestPath()
// 	creator := func(key, id int) IVertex {
// 		return NewVertex(key, id)
// 	}
// 	//**********  只有一条边的图  ***************
// 	_1e_graph := NewGraph(-1, NUM, creator) //边的无效权重为-1
// 	_1e_graph.AddVertex(0)
// 	_1e_graph.AddVertex(0)
// 	_1e_graph.AddEdge(NewTuple(0, 1, 1)) //该图只有一条边
// 	dag.ShortestPath(_1e_graph, 0)
// 	fmt.Println(fmt.Sprintf("b-EXPECT_EQ(%d,%d)", _1e_graph.Vertexes[0].GetKey(), 0))
// 	fmt.Println(fmt.Sprintf("b-EXPECT_EQ(%d,%d)", _1e_graph.Vertexes[1].GetKey(), 1))
// 	fmt.Println(fmt.Sprintf("b-EXPECT_EQ(%v,%v)", _1e_graph.Vertexes[1].GetParent(), _1e_graph.Vertexes[0]))

// 	//**********  链边的图 ***************
// 	_normal_graph := NewGraph(-1, NUM, creator) //边的无效权重为-1
// 	for i := 0; i < NUM; i++ {
// 		_normal_graph.AddVertex(0)
// 	}
// 	for i := 0; i < NUM-1; i++ {
// 		_normal_graph.AddEdge(NewTuple(i, i+1, 1))
// 	}
// 	dag.ShortestPath(_normal_graph, 0)
// 	fmt.Println(fmt.Sprintf("c-EXPECT_EQ(%d,%d)", _normal_graph.Vertexes[0].GetKey(), 0))
// 	for i := 1; i < NUM; i++ {
// 		fmt.Println(fmt.Sprintf("c-EXPECT_EQ(%d,%d)", _normal_graph.Vertexes[i].GetKey(), i))
// 		fmt.Println(fmt.Sprintf("c-EXPECT_EQ(%v,%v)", _normal_graph.Vertexes[i].GetParent(), _normal_graph.Vertexes[i-1]))
// 	}

// 	_1v_graph := NewGraph(-1, NUM, creator) //边的无效权重为-1
// 	_1v_graph.AddVertex(0)                  //该图只有一个顶点
// 	err := dag.ShortestPath(_1v_graph, 0)
// 	fmt.Println(fmt.Sprintf("a-EXPECT_EQ(%d,%d)", _1v_graph.Vertexes[0].GetKey(), 0), err)

// }

// func testDijkstra() {
// 	NUM := 10

// 	shortest := NewDijkstra()
// 	creator := func(key, id int) IVertex {
// 		return NewVertex(key, id)
// 	}
// 	//**********  只有一条边的图  ***************
// 	_1e_graph := NewGraph(-1, NUM, creator) //边的无效权重为-1
// 	_1e_graph.AddVertex(0)
// 	_1e_graph.AddVertex(0)
// 	_1e_graph.AddEdge(NewTuple(0, 1, 1)) //该图只有一条边
// 	shortest.ShortestPath(_1e_graph, 0)
// 	fmt.Println(fmt.Sprintf("b-EXPECT_EQ(%d,%d)", _1e_graph.Vertexes[0].GetKey(), 0))
// 	fmt.Println(fmt.Sprintf("b-EXPECT_EQ(%d,%d)", _1e_graph.Vertexes[1].GetKey(), 1))
// 	fmt.Println(fmt.Sprintf("b-EXPECT_EQ(%v,%v)", _1e_graph.Vertexes[1].GetParent(), _1e_graph.Vertexes[0]))

// 	//**********  链边的图 ***************
// 	_normal_graph := NewGraph(-1, NUM, creator) //边的无效权重为-1
// 	for i := 0; i < NUM; i++ {
// 		_normal_graph.AddVertex(0)
// 	}
// 	for i := 0; i < NUM-1; i++ {
// 		_normal_graph.AddEdge(NewTuple(i, i+1, 1))
// 	}
// 	shortest.ShortestPath(_normal_graph, 0)
// 	fmt.Println(fmt.Sprintf("c-EXPECT_EQ(%d,%d)", _normal_graph.Vertexes[0].GetKey(), 0))
// 	for i := 1; i < NUM; i++ {
// 		fmt.Println(fmt.Sprintf("c-EXPECT_EQ(%d,%d)", _normal_graph.Vertexes[i].GetKey(), i))
// 		fmt.Println(fmt.Sprintf("c-EXPECT_EQ(%v,%v)", _normal_graph.Vertexes[i].GetParent(), _normal_graph.Vertexes[i-1]))
// 	}

// 	_1v_graph := NewGraph(-1, NUM, creator) //边的无效权重为-1
// 	_1v_graph.AddVertex(0)                  //该图只有一个顶点
// 	err := shortest.ShortestPath(_1v_graph, 0)
// 	fmt.Println(fmt.Sprintf("a-EXPECT_EQ(%d,%d)", _1v_graph.Vertexes[0].GetKey(), 0), err)

// }
