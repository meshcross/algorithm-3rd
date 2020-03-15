/*
 * @Description: 第21章 21.1 无向图的连通分量
 * @Author: wangchengdg@gmail.com
 * @Date: 2020-02-18 10:21:23
 * @LastEditTime: 2020-03-14 11:39:49
 * @LastEditors:

 此处使用不相交集合操作来计算一个无向图的连通分量。一旦SetConnectedComponent函数与处理了该图，SameComponent
* 函数就会回答两个顶点是否在同一个连通分量。
*
* connected_component算法步骤：
*
* - 将每个顶点v放入它自己的集合中
* - 对每一条边(u,v)，它将包含u和v的集合进行合并
*
* 在处理完搜有边之后，两个顶点在相同的连通分量当且仅当与之对应的对象在相同的集合中
*/
package BasicGraph

import (
	"errors"

	. "github.com/meshcross/algorithm-3rd/mesh/graph_algorithm/graph_struct"
	. "github.com/meshcross/algorithm-3rd/mesh/graph_algorithm/graph_struct/graph_vertex"
	. "github.com/meshcross/algorithm-3rd/mesh/set_algorithm"
)

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

/**
 * @description: 图的连通分量
 *	connected_component算法步骤：
 * 		- 将每个顶点v放入它自己的集合中
 * 		- 对每一条边(u,v)，它将包含u和v的集合进行合并
 *
 * @param graph 图
 * @return: error

	如下所示的图G(V,E)：V包含7个结点，E包含6条边，G(V,E)有两个连通分量(a,b,c,d)和(e,f,g)
			a---------b          e---------f
			|        /|          |
			|      /  |          |
			|    /    |          |
			|  /      |          |
			c         d          g

*/
func (a *ConnectedComponent) SetConnectedComponent(graph *Graph) error {
	if graph == nil {
		return errors.New("SetConnectedComponent error: graph must not be nil!")
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

/*!
* @description:返回无向图的两个顶点是否位于同一个连通分量中
* @param graph:图
* @param id1:第一个顶点
* @param id2:第二个顶点
* @return bool:两个顶点是否在同一个连通分量中；error
*
* 当满足以下条件之一时，id无效的情况：
*
* - id小于0或者大于等于`num`
* - `graph.Vertexes[id1]`为空
*
* 在执行 InSameComponent函数之前必须先执行 SetConnectedComponent函数对无向图进行预处理。
*
 */
func (a *ConnectedComponent) InSameComponent(graph *Graph, id1, id2 int) (bool, error) {

	if graph == nil {
		return false, errors.New("InSameComponent error: graph must not be nil!")
	}

	num := graph.N()
	if id1 < 0 || id1 >= num || graph.Vertexes[id1] == nil || id2 < 0 || id2 >= num || graph.Vertexes[id2] == nil {
		return false, errors.New("InSameComponent error: id muse belongs [0,N) and graph.Vertexes[id] must not be nil!")
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
