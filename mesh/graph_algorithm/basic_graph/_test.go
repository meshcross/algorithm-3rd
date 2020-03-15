package BasicGraph

import (
	"fmt"
	"sort"
	"testing"

	. "github.com/meshcross/algorithm-3rd/mesh/common"
	. "github.com/meshcross/algorithm-3rd/mesh/graphAalgorithm/graph_struct/graph_vertex"
	. "github.com/meshcross/algorithm-3rd/mesh/graph_algorithm/graph_struct"
)

/**
 * @description: 无向图的连通分量
 */
func TestConnectedComponent(t *testing.T) {
	C_NUM := 10
	creator := func(key, id int) IVertex {
		ptr := NewSetVertex(key, id)
		return ptr
	}
	_graph := NewGraph(-1, C_NUM, creator)
	for i := 0; i < C_NUM; i++ {
		_graph.AddVertex(i)
	}

	for i := 1; i < C_NUM/2; i++ { //边 0--1--2--...--C_NUM/2-1
		_graph.AddEdge(NewTuple(i-1, i, 9999))
	}

	for i := C_NUM/2 + 1; i < C_NUM; i++ { //边 C_NUM/2--C_NUM/2+1...--C_NUM-1
		_graph.AddEdge(NewTuple(i-1, i, 8888))
	}

	conn := NewConnectedComponent()
	conn.SetConnectedComponent(_graph)

	for i := 0; i < C_NUM/2; i++ { //tree_root:u, tree_root:v, rank小的挂在rank大的之下；若二者rank相等，则u挂在v下
		//fmt.Println(fmt.Sprintf("EXPECT_EQ(%v,%v)", _graph.Vertexes[i].Node.Parent.Value, _graph.Vertexes[1].Node.Value)) //所以结点1是最终根结点
	}

	for i := C_NUM / 2; i < C_NUM; i++ { //tree_root:u, tree_root:v, rank小的挂在rank大的之下；若二者rank相等，则u挂在v下
		//fmt.Println(fmt.Sprintf("EXPECT_EQ(%v,%v)", _graph.Vertexes[i].Node.Parent.Value, _graph.Vertexes[C_NUM/2+1].Node.Value)) //所以结点`C_NUM/2+1`是最终根结点
	}
}

/**
 * @description: 图的广度优先搜索
 */
func TestGraphBFS(t *testing.T) {
	NUM := 10
	creator := func(key, id int) IVertex {
		ptr := NewBFSVertex(key, id)
		return ptr
	}
	graph := NewGraph(-1, NUM, creator)
	graph.AddVertex(0) //该图只有一个顶点

	graph_1e := NewGraph(-1, NUM, creator)
	graph_1e.AddVertex(0)
	graph_1e.AddVertex(1)
	graph_1e.AddEdge(NewTuple(0, 1, 1)) //该图只有一条边

	//****  含顶点图和边图：10个顶点，9条边   ****
	list_graph := NewGraph(-1, NUM, creator) //边的无效权重为-1
	for i := 0; i < NUM; i++ {
		list_graph.AddVertex(i)
	}

	for i := 0; i < NUM-1; i++ {
		list_graph.AddEdge(NewTuple(i, i+1, 10+i)) //该图的边是从左到右组成一个链条
	}

	//-------------setup end------------------//

	bfs := NewGraphBFS()
	{
		str := ""
		print_bfs := func(v_id int) {
			str = fmt.Sprintf("%s%d;", str, v_id)
		}
		//****** 测试只有一个顶点的图**********
		bfs.Search(graph, 0, print_bfs, nil)
		fmt.Println(fmt.Sprintf("EXPECT_EQ(%s,%s) ", str, "0;"))
		fmt.Println(fmt.Sprintf("EXPECT_EQ(%d, %d)", graph.Vertexes[0].GetKey(), 0))
	}

	{
		str := ""
		print_bfs := func(v_id int) {
			str = fmt.Sprintf("%s%d;", str, v_id)
		}
		//***** 测试只有一条边的图*********
		bfs.Search(graph_1e, 0, print_bfs, nil)
		fmt.Println(fmt.Sprintf("b-EXPECT_EQ(%s,%s)", str, "0;1;"))
		fmt.Println(fmt.Sprintf("b-EXPECT_EQ(%d,%d)", graph_1e.Vertexes[0].GetKey(), 0))
		fmt.Println(fmt.Sprintf("b-EXPECT_EQ(%d,%d)", graph_1e.Vertexes[1].GetKey(), 1))
	}

	{
		str := ""
		print_bfs := func(v_id int) {
			str = fmt.Sprintf("%s%d;", str, v_id)
		}
		//**** 测试链边的图**********
		real_str := ""
		for i := 0; i < NUM; i++ {
			real_str = fmt.Sprintf("%s%d;", real_str, i)
		}
		// breadth_first_search(_list_graph,0,print_bfs);
		bfs.Search(list_graph, 0, print_bfs, nil)
		fmt.Println(fmt.Sprintf("c-EXPECT_EQ(%s     %s)", str, real_str))
		for i := 0; i < NUM; i++ {
			fmt.Println(fmt.Sprintf("c-EXPECT_EQ(%d,%d)", list_graph.Vertexes[i].GetKey(), i))
		}
	}
}

/**
 * @description: 图的深度优先搜索
 */
func TestGraphDFS(t *testing.T) {
	NUM := 10
	creator := func(key, id int) IVertex {
		return NewDFSVertex(key, id)
	}
	graph := NewGraph(-1, NUM, creator)
	graph.AddVertex(0) //该图只有一个顶点

	graph_1e := NewGraph(-1, NUM, creator)
	graph_1e.AddVertex(0)
	graph_1e.AddVertex(1)
	graph_1e.AddEdge(NewTuple(0, 1, 1)) //该图只有一条边

	//****  含顶点图和边图：10个顶点，9条边   ****
	list_graph := NewGraph(-1, NUM, creator) //边的无效权重为-1
	for i := 0; i < NUM; i++ {
		list_graph.AddVertex(i)
	}

	for i := 0; i < NUM-1; i++ {
		list_graph.AddEdge(NewTuple(i, i+1, 10+i)) //该图的边是从左到右组成一个链条
	}

	rlist_graph := NewGraph(-1, NUM, creator) //边的无效权重为-1
	for i := 0; i < NUM; i++ {
		rlist_graph.AddVertex(i + NUM)
	}
	for i := 0; i < NUM-1; i++ {
		rlist_graph.AddEdge(NewTuple(i, i+1, 10+i)) //该图的边是从左到右组成一个链条
	}

	//-------------setup end------------------//

	dfs := NewGraphDFS()
	{
		str_discover := ""
		print_discover := func(v_id, time int) {
			str_discover = fmt.Sprintf("%s%d;", str_discover, v_id)
		}
		str_finish := ""
		print_finish := func(v_id, time int) {
			str_finish = fmt.Sprintf("%s%d;", str_finish, v_id)
		}
		//****** 测试只有一个顶点的图**********
		dfs.Search(graph, print_discover, print_finish, nil, nil, nil)
		fmt.Println(fmt.Sprintf("a-EXPECT_EQ(%s,%s) ", str_discover, "0;"))
		fmt.Println(fmt.Sprintf("a-EXPECT_EQ(%d, %d)", graph.Vertexes[0].GetKey(), 0))
	}

	{
		str_discover := ""
		print_discover := func(v_id, time int) {
			str_discover = fmt.Sprintf("%s%d;", str_discover, v_id)
		}
		str_finish := ""
		print_finish := func(v_id, time int) {
			str_finish = fmt.Sprintf("%s%d;", str_finish, v_id)
		}
		//***** 测试只有一条边的图*********
		dfs.Search(graph_1e, print_discover, print_finish, nil, nil, nil)
		fmt.Println(fmt.Sprintf("b-EXPECT_EQ(%s,%s)", str_discover, "0;1;"))
		fmt.Println(fmt.Sprintf("b-EXPECT_EQ(%d,%d)", graph_1e.Vertexes[0].GetKey(), 0))
		fmt.Println(fmt.Sprintf("b-EXPECT_EQ(%d,%d)", graph_1e.Vertexes[1].GetKey(), 1))
	}

	{
		str_discover := ""
		print_discover := func(v_id, time int) {
			str_discover = fmt.Sprintf("%s%d;", str_discover, v_id)
		}
		str_finish := ""
		print_finish := func(v_id, time int) {
			str_finish = fmt.Sprintf("%s%d;", str_finish, v_id)
		}
		//**** 测试链边的图**********
		real_discover_str := ""
		real_finish_str := ""
		for i := 0; i < NUM; i++ {
			real_discover_str = fmt.Sprintf("%s%d;", real_discover_str, i)
			real_finish_str = fmt.Sprintf("%s%d;", real_finish_str, i)
		}

		dfs.Search(list_graph, print_discover, print_finish, nil, nil, nil)
		fmt.Println(fmt.Sprintf("c-EXPECT_EQ(%s     %s)", str_discover, real_discover_str))
		fmt.Println(fmt.Sprintf("c-EXPECT_EQ(%s     %s)", str_finish, real_finish_str))
		for i := 0; i < NUM; i++ {
			fmt.Println(fmt.Sprintf("c-EXPECT_EQ(%d,%d)", list_graph.Vertexes[i].GetKey(), i))
		}
	}

	{
		empty_action := func(v_id, time int) {}
		str_discover := ""
		print_discover := func(v_id, time int) {
			str_discover = fmt.Sprintf("%s%d;", str_discover, v_id)
		}
		str_finish := ""
		print_finish := func(v_id, time int) {
			str_finish = fmt.Sprintf("%s%d;", str_finish, v_id)
		}
		//****** 测试测试链边的图，逆序**********
		search_order := []int{}
		real_discover_str := ""
		real_finish_str := ""
		for i := NUM - 1; i >= 0; i-- {
			real_discover_str = fmt.Sprintf("%s%d;", real_discover_str, i)
			real_finish_str = fmt.Sprintf("%s%d;", real_finish_str, i)

			search_order = append(search_order, i)
		}
		dfs.Search(rlist_graph, print_discover, print_finish, empty_action, empty_action, search_order)
		fmt.Println(fmt.Sprintf("d-EXPECT_EQ(%s  %s)", str_discover, real_discover_str))
		fmt.Println(fmt.Sprintf("d-EXPECT_EQ(%s  %s)", str_finish, real_finish_str))
	}
}

/**
 * @description: 图的强连通分量
 */
func TestStrongConnectedComponent(t *testing.T) {
	SCC_N := 10
	creator := func(key, id int) IVertex {
		ptr := NewDFSVertex(key, id)
		return ptr
	}
	//****  含顶点图和边图：10个顶点，9条边  ****
	_list_graph := NewGraph(-1, SCC_N, creator) //边的无效权重为-1
	for i := 0; i < SCC_N; i++ {
		_list_graph.AddVertex(i + SCC_N)
	}

	for i := 0; i < SCC_N-1; i++ {
		_list_graph.AddEdge(NewTuple(i, i+1, 10+i)) //该图的边是从左到右组成一个链条
	}

	//****  含顶点图和边图：10个顶点，10条边  ****
	_scc_graph := NewGraph(-1, SCC_N, creator) //边的无效权重为-1
	for i := 0; i < SCC_N; i++ {
		_scc_graph.AddVertex(i + SCC_N)
	}

	for i := 0; i < SCC_N-1; i++ {
		_scc_graph.AddEdge(NewTuple(i, i+1, 10+i))
	}

	_scc_graph.AddEdge(NewTuple(SCC_N-1, 0, 10+SCC_N-1))

	setor := &StrongConnectedComponent{}
	res, _ := setor.SetStrongConnectedComponent(_list_graph)
	//****** 测试一条链边的图**********
	fmt.Println(fmt.Sprintf("a-EXPECT_EQ(%d,%d)", len(res), 0))

	//****** 测试环状图**********
	result, _ := setor.SetStrongConnectedComponent(_scc_graph)
	real_vertexes := []int{}
	for i := 0; i < SCC_N; i++ {
		real_vertexes = append(real_vertexes, i)
	}

	sort.Ints(real_vertexes)

	fmt.Println(fmt.Sprintf("b-EXPECT_EQ(%d,%d)", len(result), 1))

	scc_vertexes := result[0]
	sort.Ints(scc_vertexes)
	fmt.Println(fmt.Sprintf("c-EXPECT_EQ(%v,%v)", scc_vertexes, real_vertexes))

}
