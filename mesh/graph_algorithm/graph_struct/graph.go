/*
 * @Description: 算法导论22章22.1节 图
 * @Author: wangchengdg@gmail.com
 * @Date: 2020-02-18 10:28:31
 * @LastEditTime: 2020-03-06 11:35:41
 * @LastEditors:
 */
package GraphStruct

import (
	"errors"
	"sort"

	. "github.com/meshcross/algorithm-3rd/mesh/common"
	. "github.com/meshcross/algorithm-3rd/mesh/graph_algorithm/graph_struct/graph_vertex"
)

/*!
* 图的矩阵主要包含四个数据：
*
* - `matrix`：图的矩阵表示
* - `adjList`：图的邻接表表示
* - `vertexes`：顶点集合，其元素类型为指向顶点的强引用
* - `next_empty_vertex`：顶点集合中，下一个为空的位置，它用于添加顶点。

* 图支持插入、修改顶点操作，插入、修改边操作（由图的矩阵以及图的邻接表来代理），以及返回边、返回权重（由图的矩阵以及图的邻接表来代理）。
*
 */

type VertexCreatorFunc func(key int, id int) IVertex

type Graph struct {
	Vertexes          []IVertex
	next_empty_vertex int
	Matrix            *MatrixGraph
	AdjList           *ADJListGraph
	_N                int
	VertexCreator     VertexCreatorFunc
}

func NewGraph(invalidWeight int, n int, creator VertexCreatorFunc) *Graph {
	matrix := NewMatrixGraph(invalidWeight, n)
	adjList := NewADJListGraph(n)
	vers := make([]IVertex, n)
	return &Graph{next_empty_vertex: 0, Matrix: matrix, _N: n, AdjList: adjList, Vertexes: vers, VertexCreator: creator}
}

func (a *Graph) N() int {
	return a._N
}

/*!
* @description:添加一个顶点
* @param  key:顶点存放的数据
* @return: 顶点的id
*
 */
func (a *Graph) AddVertex(key int, ids ...int) (int, error) {
	if len(ids) > 0 {
		id := ids[0]
		if id < 0 || id >= a._N {
			return -1, errors.New("add_vertex error:id must >=0 and <N.")
		}

		if a.Vertexes[id] != nil {
			return -1, errors.New("add_vertex error: vertex of id has existed.")
		}

		a.Vertexes[id] = a.VertexCreator(key, id)
		return id, nil
	} else {
		for a.next_empty_vertex < a._N && a.Vertexes[a.next_empty_vertex] != nil {
			a.next_empty_vertex++
		}
		if a.next_empty_vertex >= a._N {
			return 0, errors.New("add_vertex error:Graph Vertex is full, can not add vertex.")
		}
		v_id := a.next_empty_vertex

		a.Vertexes[a.next_empty_vertex] = a.VertexCreator(key, v_id)
		a.next_empty_vertex++
		return v_id, nil
	}
}

/*!
* @description:修改一个顶点的数据
* @param  newkey:新的数据
* @param id:指定该顶点的`id`
*
 */
func (a *Graph) ModifyVertex(newkey, id int) error {
	if id < 0 || id >= a._N {
		return errors.New("modify_vertex error:id must >=0 and <N.")
	}

	if a.Vertexes[id] == nil {
		return errors.New("modify_vertex error: vertex of id does not exist.")
	}

	a.Vertexes[id].SetKey(newkey)
	return nil
}

/*!
* @description:添加一条边
* @param  edge_tuple:一条边的三元素元组
*
* 在添加边时，同时向图的矩阵、图的邻接表中添加边
*
* 如果添加的边是无效权重，则直接返回而不添加
 */
func (a *Graph) AddEdge(edge_tuple *Tuple) error {
	id1 := edge_tuple.First
	id2 := edge_tuple.Second
	wt := edge_tuple.Third

	if id1 < 0 || id1 >= a._N || id2 < 0 || id2 >= a._N {
		return errors.New("add edge error:id must >=0 and <N.")
	}

	if a.Vertexes[id1] == nil || a.Vertexes[id2] == nil {
		return errors.New("add edge error: vertex of id does not exist.")
	}

	if wt == a.Matrix.InvalidWeight() {
		return errors.New("invalid weight")
	}
	a.Matrix.AddEdge(edge_tuple)
	a.AdjList.AddEdge(edge_tuple)
	return nil
}

/*!
* @description:添加一组边
* @param  begin:边容器的起始迭代器
* @param  end:边容器的终止迭代器
* @return:void
*
* 在添加边时，同时向图的矩阵、图的邻接表中添加边
 */
func (a *Graph) AddEdges(edges []*Tuple) {
	for _, edge := range edges {
		a.AddEdge(edge)
	}
}

/*!
* @description:修改一条边的权重
* @param  id1:待修改边的第一个顶点
* @param  id2:待修改边的第二个顶点
* @param  wt:新的权重
* @return error
*
 */
func (a *Graph) AdjustEdge(id1, id2, wt int) error {
	if id1 < 0 || id1 >= a._N || id2 < 0 || id2 >= a._N {
		return errors.New("adjust edge error:id must >=0 and <N.")
	}

	if a.Vertexes[id1] == nil || a.Vertexes[id2] == nil {
		return errors.New("adjust edge error: vertex of id does not exist.")
	}

	a.Matrix.AdjustEdge(id1, id2, wt)
	a.AdjList.AdjustEdge(id1, id2, wt)
	return nil
}

/*!
* @description:返回图中所有边的三元素元组集合
* @return  :图中所有边的三元素元组集合
*
* 要求图的矩阵和图的邻接表都返回同样的结果
 */
func (a *Graph) EdgeTuples() []*Tuple {
	edges1 := a.Matrix.EdgeTuples()
	edges2 := a.AdjList.EdgeTuples()

	wapper1 := NewTupleWapper(edges1, TupleCompareFunc_Less)
	wapper2 := NewTupleWapper(edges2, TupleCompareFunc_Less)
	sort.Sort(wapper1)
	sort.Sort(wapper2)

	// if edges1 != edges2 {
	// 	return nil, errors.New("error")
	// }
	return edges1
}

/*!
* @description:返回图中从指定顶点出发的边的三元素元组集合
* @param id: 指定顶点`id`
* @return  :图中指定顶点出发的边的三元素元组集合
*
 */
func (a *Graph) VertexEdgeTuples(id int) ([]*Tuple, error) {

	if id < 0 || id >= a._N {
		return nil, errors.New("vertex_edge_tuples error:id must >=0 and <N.")
	}

	if a.Vertexes[id] == nil {
		return nil, errors.New("vertex_edge_tuples error: vertex of id does not exist.")
	}

	edges1, _ := a.Matrix.VertexEdgeTuples(id)
	edges2, _ := a.AdjList.VertexEdgeTuples(id)
	compare := TupleCompareFunc_Less

	wapper1 := NewTupleWapper(edges1, compare)
	wapper2 := NewTupleWapper(edges2, compare)
	sort.Sort(wapper1)
	sort.Sort(wapper2)

	// if edges1 != edges2 {
	// 	return nil, errors.New("error")
	// }
	return edges1, nil
}

/*!
* @description:返回图中指定顶点之间是否存在边
* @param id_from: 第一个顶点的`id`
* @param id_to: 第二个顶点的`id`
* @return  :第一个顶点和第二个顶点之间是否存在边
*
 */
func (a *Graph) HasEdge(id_from, id_to int) (bool, error) {
	if id_from < 0 || id_from >= a._N || id_to < 0 || id_to >= a._N {
		return false, errors.New("has edge error:id must >=0 and <N.")
	}

	if a.Vertexes[id_from] == nil || a.Vertexes[id_to] == nil {
		return false, errors.New("has edge error: vertex of id does not exist.")
	}

	m_has, _ := a.Matrix.HasEdge(id_from, id_to)
	a_has, _ := a.AdjList.HasEdge(id_from, id_to)

	if m_has != a_has {
		return false, errors.New("error")
	}
	return a.Matrix.HasEdge(id_from, id_to)
}

/*!
* @description:返回图中指定顶点之间的边的权重
* @param id_from: 第一个顶点的`id`
* @param id_to: 第二个顶点的`id`
* @return  :第一个顶点和第二个顶点之间的边的权重
*
 */
func (a *Graph) Weight(id_from, id_to int) (int, error) {
	if id_from < 0 || id_from >= a._N || id_to < 0 || id_to >= a._N {
		return -1, errors.New("edge weight error:id must >=0 and <N.")
	}

	if a.Vertexes[id_from] == nil || a.Vertexes[id_to] == nil {
		return -1, errors.New("edge weight error: vertex of id does not exist.")
	}

	m_wt, _ := a.Matrix.Weight(id_from, id_to)
	a_wt, _ := a.AdjList.Weight(id_from, id_to)
	if m_wt != a_wt {
		return -1, errors.New("error")
	}
	return m_wt, nil
}

/*!
* @description:返回图的一个翻转镜像
* @return  :图的一个镜像的一个强引用
*
* 图的一个镜像也是一个图，它与原图有以下关系：
*
* - 图的镜像的顶点与原图的顶点相同
* - 图的镜像的边是原图的边的反向
*
* 首先新建一个图，再根据原图的顶点来执行顶点的深拷贝。然后再获取原图的边的反向边，将该反向边作为镜像图的边
 */
func (a *Graph) Inverse() *Graph {
	graph := NewGraph(a.Matrix.invalidWeight, a._N, a.VertexCreator)

	vLen := len(a.Vertexes)
	for i := 0; i < vLen; i++ {
		v := a.Vertexes[i]
		if v != nil {
			graph.Vertexes[i] = a.VertexCreator(v.GetKey(), v.GetID())
		}
	}
	edges := a.EdgeTuples()
	for _, edge := range edges {
		tmp := edge.First
		edge.First = edge.Second
		edge.Second = tmp
	}
	graph.AddEdges(edges)
	return graph
}
