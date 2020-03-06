package MinimumSpanningTree

import (
	"fmt"
	"strconv"
	"testing"

	. "github.com/meshcross/algorithm-3rd/mesh/common"
	. "github.com/meshcross/algorithm-3rd/mesh/graph_algorithm/graph_struct"
	. "github.com/meshcross/algorithm-3rd/mesh/graph_algorithm/graph_struct/graph_vertex"
)

func TestKruskal(t *testing.T) {
	K_NUM := 10
	creator := func(key, id int) IVertex {
		ptr := NewSetVertex(key, id)
		return ptr
	}
	////////////////start setup/////////////////////////

	_1v_graph := NewGraph(-1, K_NUM, creator) //边的无效权重为-1
	_1v_graph.AddVertex(0)                    //该图只有一个顶点
	_1v_graph.AddEdge(NewTuple(0, 0, 1))      //该图只有一个顶点，必须有一条边指向自己

	_1e_graph := NewGraph(-1, K_NUM, creator) //边的无效权重为-1
	_1e_graph.AddVertex(0)
	_1e_graph.AddVertex(0)
	_1e_graph.AddEdge(NewTuple(0, 1, 1)) //该图只有一条边

	//****  含顶点图和边图：10个顶点，9条边   ****
	_list_graph := NewGraph(-1, K_NUM, creator) //边的无效权重为-1
	for i := 0; i < K_NUM; i++ {
		_list_graph.AddVertex(0)
	}
	for i := 0; i < K_NUM-1; i++ {
		_list_graph.AddEdge(NewTuple(i, i+1, i+1)) //该图的边是从左到右组成一个链条
	}

	//****  含顶点图和边图：10个顶点，90条边   ****
	_all_edges_graph := NewGraph(-1, K_NUM, creator) //边的无效权重为-1
	for i := 0; i < K_NUM; i++ {
		_all_edges_graph.AddVertex(0)
	}
	for i := 0; i < K_NUM; i++ {
		for j := 0; j < K_NUM; j++ {
			if i == j {
				continue
			} else {
				_all_edges_graph.AddEdge(NewTuple(i, j, i+j)) //该图中任意一对顶点之间都有边
			}
		}
	}

	////////////////end setup/////////////////////////

	//************* 单点图 **************
	kruskal := NewKruskalMST()
	{
		edges := []*Pair{}
		pre_action := func(id1, id2 int) {
			edges = append(edges, NewPair(id1, id2))
		}
		weight, _ := kruskal.Generate(_1v_graph, pre_action, nil)
		fmt.Println(fmt.Sprintf("a-EXPECT_EQ(%d,%d)", weight, 0))
		fmt.Println(fmt.Sprintf("a-EXPECT_EQ(%d,%d)", len(edges), 0))
	}
	//************** 单边图 ******************
	{
		edges := []*Pair{}
		pre_action := func(id1, id2 int) {
			edges = append(edges, NewPair(id1, id2))
		}
		weight, _ := kruskal.Generate(_1e_graph, pre_action, nil)
		fmt.Println(fmt.Sprintf("b-EXPECT_EQ(%d,%d)", weight, 1))
		fmt.Println(fmt.Sprintf("b-EXPECT_EQ(%v,%v)", edges, NewPair(0, 1)))

		// auto node0=_1e_graph->vertexes.at(0)->node;
		// auto node1=_1e_graph->vertexes.at(1)->node;
		// ASSERT_TRUE(node0&&node0->parent);
		// ASSERT_TRUE(node1&&node1->parent);
		// EXPECT_EQ(node0->parent,node1); //tree_root:u, tree_root:v, rank小的挂在rank大的之下；若二者rank相等，则u挂在v下
		// EXPECT_EQ(node1->parent,node1); //所以结点1是最终根结点
	}
	//************** 单链图 ******************
	{
		edges := []*Pair{}
		pre_action := func(id1, id2 int) {
			edges = append(edges, NewPair(id1, id2))
		}
		weight, _ := kruskal.Generate(_list_graph, pre_action, nil)
		fmt.Println(fmt.Sprintf("c-EXPECT_EQ(%d,%d)", weight, (K_NUM-1)*K_NUM/2))

		result_edges := []*Pair{}
		for i := 1; i < K_NUM; i++ {
			result_edges = append(result_edges, NewPair(i-1, i))
		}

		for i := 0; i < K_NUM; i++ { //tree_root:u, tree_root:v, rank小的挂在rank大的之下；若二者rank相等，则u挂在v下
			vtx := ToSetVertex(_list_graph.Vertexes[i])

			iParent := vtx.Node.Parent
			vtx_1 := ToSetVertex(_list_graph.Vertexes[1])
			node1 := vtx_1.Node
			fmt.Println(fmt.Sprintf("c-EXPECT_EQ(%v,%v)", iParent, node1)) //所以结点1是最终根结点mt.Println("c-EXPECT_EQ(_list_graph->vertexes.at(i)->node->parent,_list_graph->vertexes.at(1)->node)") //所以结点1是最终根结点
		}
		//检验添加边的数量和顺序
		fmt.Println(fmt.Sprintf("c-EXPECT_EQ(%d,%d)", len(edges), len(result_edges)))
		for i := 0; i < len(edges); i++ {
			x := edges[i].First == result_edges[i].First && edges[i].Second == result_edges[i].Second
			y := edges[i].First == result_edges[i].Second && edges[i].Second == result_edges[i].First
			fmt.Println(fmt.Sprintf("c-EXPECT_TRUE(%d,%s)", i, strconv.FormatBool(x || y)))
		}
	}
	//***************** 任意一对顶点之间都有边的图  **************
	{
		//边(u,v)的权重为(u.id+v.id)
		edges := []*Pair{}
		pre_action := func(id1, id2 int) {
			edges = append(edges, NewPair(id1, id2))
		}
		weight, _ := kruskal.Generate(_all_edges_graph, pre_action, nil)
		fmt.Println(fmt.Sprintf("d-EXPECT_EQ(%d,%d)", weight, (K_NUM-1)*K_NUM/2))
		result_edges := []*Pair{}
		for i := 1; i < K_NUM; i++ {
			result_edges = append(result_edges, NewPair(0, i))
		}

		for i := 0; i < K_NUM; i++ { //tree_root:u, tree_root:v, rank小的挂在rank大的之下；若二者rank相等，则u挂在v下
			vtx := _all_edges_graph.Vertexes[i].(*SetVertex)
			iParent := vtx.Node.Parent
			vtx_1 := _all_edges_graph.Vertexes[1].(*SetVertex)
			node1 := vtx_1.Node
			fmt.Println(fmt.Sprintf("d-EXPECT_EQ(%v,%v)", iParent, node1)) //所以结点1是最终根结点
		}

		//检验添加边的数量和顺序
		fmt.Println(fmt.Sprintf("d-EXPECT_EQ(%d,%d)", len(edges), len(result_edges)))
		for i := 0; i < len(edges); i++ {
			x := edges[i].First == result_edges[i].First && edges[i].Second == result_edges[i].Second
			y := edges[i].First == result_edges[i].Second && edges[i].Second == result_edges[i].First
			fmt.Println(fmt.Sprintf("d-EXPECT_TRUE(%d,%s,edge %v result_edge %v)", i, strconv.FormatBool(x || y), edges[i], result_edges[i]))
		}
	}
}

func testPrim() {
	P_NUM := 10
	creator := func(key, id int) IVertex {
		ptr := NewVertex(key, id)
		return ptr
	}
	////////////////start setup/////////////////////////

	_1v_graph := NewGraph(-1, P_NUM, creator) //边的无效权重为-1
	_1v_graph.AddVertex(0)                    //该图只有一个顶点
	_1v_graph.AddEdge(NewTuple(0, 0, 1))      //该图只有一个顶点，必须有一条边指向自己

	_1e_graph := NewGraph(-1, P_NUM, creator) //边的无效权重为-1
	_1e_graph.AddVertex(0)
	_1e_graph.AddVertex(0)
	_1e_graph.AddEdge(NewTuple(0, 1, 1)) //该图只有一条边

	//****  含顶点图和边图：10个顶点，9条边   ****
	_list_graph := NewGraph(-1, P_NUM, creator) //边的无效权重为-1
	for i := 0; i < P_NUM; i++ {
		_list_graph.AddVertex(0)
	}
	for i := 0; i < P_NUM-1; i++ {
		_list_graph.AddEdge(NewTuple(i, i+1, i+1)) //该图的边是从左到右组成一个链条
	}

	//****  含顶点图和边图：10个顶点，90条边   ****
	_all_edges_graph := NewGraph(-1, P_NUM, creator) //边的无效权重为-1
	for i := 0; i < P_NUM; i++ {
		_all_edges_graph.AddVertex(0)
	}
	for i := 0; i < P_NUM; i++ {
		for j := 0; j < P_NUM; j++ {
			if i == j {
				continue
			} else {
				_all_edges_graph.AddEdge(NewTuple(i, j, i+j)) //该图中任意一对顶点之间都有边
			}
		}
	}

	////////////////end setup/////////////////////////

	//************* 单点图 **************
	prim := NewPrimMST()
	{
		ids := []int{}
		pre_action := func(id int) {
			ids = append(ids, id)
		}
		weight, _ := prim.Generate(_1v_graph, 0, pre_action, nil)
		fmt.Println(fmt.Sprintf("a-EXPECT_EQ(%d,%d)", weight, 0))
		fmt.Println(fmt.Sprintf("a-EXPECT_EQ(%d,%d)", len(ids), 0))
	}
	//************** 单边图 ******************
	{
		ids := []int{}
		pre_action := func(id int) {
			ids = append(ids, id)
		}
		weight, _ := prim.Generate(_1e_graph, 0, pre_action, nil)
		fmt.Println(fmt.Sprintf("b-EXPECT_EQ(%d,%d)", weight, 1))
		fmt.Println(fmt.Sprintf("b-EXPECT_EQ(%v,%v)", ids, []int{0, 1}))
	}
	//************** 单链图 ******************
	{
		ids := []int{}
		pre_action := func(id int) {
			ids = append(ids, id)
		}
		weight, _ := prim.Generate(_list_graph, 0, pre_action, nil)
		fmt.Println(fmt.Sprintf("c-EXPECT_EQ(%d,%d)", weight, (P_NUM-1)*P_NUM/2))

		result_ids := []int{}
		for i := 0; i < P_NUM; i++ {
			result_ids = append(result_ids, i)
			//EXPECT_EQ(_list_graph.Vertexes[i].Parent,_list_graph.Vertexes[i-1]);

		}

		fmt.Println(fmt.Sprintf("d-EXPECT_EQ(%v,%v)", ids, result_ids))
	}
	//***************** 任意一对顶点之间都有边的图  **************
	{
		ids := []int{}
		pre_action := func(id int) {
			ids = append(ids, id)
		}
		weight, _ := prim.Generate(_all_edges_graph, 0, pre_action, nil)
		fmt.Println(fmt.Sprintf("e-EXPECT_EQ(%d,%d)", weight, (P_NUM-1)*P_NUM/2))
		result_ids := []int{}
		for i := 0; i < P_NUM; i++ {
			result_ids = append(result_ids, i)
			//EXPECT_EQ(_all_edges_graph.Vertexes[i].Parent,_all_edges_graph.Vertexes[0])<<"i:"<<i;
		}
		fmt.Println(fmt.Sprintf("e-EXPECT_EQ(%v,%v)", ids, result_ids))

	}
}
