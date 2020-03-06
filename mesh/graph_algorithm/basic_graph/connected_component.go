package BasicGraph

import (
	"errors"

	. "github.com/meshcross/algorithm-3rd/mesh/graph_algorithm/graph_struct"
	. "github.com/meshcross/algorithm-3rd/mesh/graph_algorithm/graph_struct/graph_vertex"
	. "github.com/meshcross/algorithm-3rd/mesh/set_algorithm"
)

//!connected_component：无向图的连通分量，算法导论21章21.1节
/*!
* \param graph:无向图，必须非空
* \return:error
*
* connected_component函数使用不相交集合操作来计算一个无向图的连通分量。一旦connected_component函数处理了该图，same_component
* 函数就会回答两个顶点是否在同一个连通分量。
*
* connected_component算法步骤：
*
* - 将每个顶点v放入它自己的集合中
* - 对每一条边(u,v)，它将包含u和v的集合进行合并
*
* 在处理完搜有边之后，两个顶点在相同的连通分量当且仅当与之对应的对象在相同的集合中
 */

type ConnectedComponent struct {
}

func NewConnectedComponent() *ConnectedComponent {
	return &ConnectedComponent{}
}
func (a *ConnectedComponent) toSetVetex(vtx IVertex) *SetVertex {
	if v, ok := vtx.(*SetVertex); ok {
		return v
	}
	return nil
}

func (a *ConnectedComponent) SetConnectedComponent(graph *Graph) error {
	if graph == nil {
		return errors.New("connected_component error: graph must not be nil!")
	}
	//*********** 初始化 ****************
	sets := []*DisJointSetNode{}

	num := graph.N()
	for i := 0; i < num; i++ {
		vertex := a.toSetVetex(graph.Vertexes[i])
		if vertex != nil { //添加顶点到`disjoint_set`中
			set_node := NewDisJointSetNode(vertex)
			sets = append(sets, set_node)
			vertex.Node = set_node
			MakeSet(set_node)
		}
	}
	//****************** 循环  ************************
	edges := graph.EdgeTuples()
	for _, edge := range edges {
		from_vtx := a.toSetVetex(graph.Vertexes[edge.First])
		to_vtx := a.toSetVetex(graph.Vertexes[edge.Second])

		from_vertex_set_node := from_vtx.Node
		to_vertex_set_node := to_vtx.Node
		ret_from, _ := FindSet(from_vertex_set_node)
		ret_to, _ := FindSet(to_vertex_set_node)
		if ret_from != ret_to {
			UnionSet(from_vertex_set_node, to_vertex_set_node)
		}
	}
	return nil
}

//!same_component：返回无向图的两个顶点是否位于同一个连通分量中。算法导论21章21.1节
/*!
*
* \param graph:指向图的强指针，必须非空。若为空则抛出异常
* \param id1:第一个顶点，必须有效。若无效则抛出异常
* \param id2:第二个顶点，必须有效。若无效则抛出异常
*
* 当满足以下条件之一时，id无效的情况：
*
* - id小于0或者大于等于`GraphType::NUM`
* - `graph->vertexes.at(id1)`为空
*
* 在执行 same_component函数之前必须先执行 connected_component函数对无向图进行预处理。
*
 */
func (a *ConnectedComponent) InSameComponent(graph *Graph, id1, id2 int) (bool, error) {

	if graph == nil {
		return false, errors.New("same_component error: graph must not be nullptr!")
	}

	num := graph.N()
	if id1 < 0 || id1 >= num || graph.Vertexes[id1] == nil || id2 < 0 || id2 >= num || graph.Vertexes[id2] == nil {
		return false, errors.New("same_component error: id muse belongs [0,N) and graph->vertexes[id] must not be nullptr!")
	}

	vtx1 := a.toSetVetex(graph.Vertexes[id1])
	vtx2 := a.toSetVetex(graph.Vertexes[id2])

	ret1, _ := FindSet(vtx1.Node)
	ret2, _ := FindSet(vtx2.Node)

	if ret1 == ret2 {
		return true, nil
	} else {
		return false, nil
	}

}
