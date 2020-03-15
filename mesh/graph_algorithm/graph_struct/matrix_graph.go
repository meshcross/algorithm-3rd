/*
 * @Description: 图的矩阵表示法，稠密图用该方法表示
 * @Author: wangchengdg@gmail.com
 * @Date: 2020-02-18 10:28:51
 * @LastEditTime: 2020-03-14 11:42:28
 * @LastEditors:
 */
package GraphStruct

import (
	"errors"

	. "github.com/meshcross/algorithm-3rd/mesh/common"
)

/**
* 稠密图使用矩阵表示法
**/
type MatrixGraph struct {
	Matrix        [][]int //矩阵表示法
	invalidWeight int     //	不可修改
	_N            int     //	不可修改
}

func (a *MatrixGraph) InvalidWeight() int {
	return a.invalidWeight
}
func (a *MatrixGraph) AddEdge(edge_tuple *Tuple) error {
	id1 := edge_tuple.First
	id2 := edge_tuple.Second
	wt := edge_tuple.Third
	if id1 < 0 || id1 >= a._N || id2 < 0 || id2 >= a._N {
		return errors.New("param is invalid")
	}

	b, _ := a.HasEdge(id1, id2)
	if b {
		return errors.New("edge add error,edge has already exist.")
	}

	a.Matrix[id1][id2] = wt
	return nil
}

/*!
* @description:添加一组边
* @param  edges:一组边
*
 */
func (a *MatrixGraph) AddEdges(edges []*Tuple) {
	for _, v := range edges {
		a.AddEdge(v)
	}
}

/*!
* @description:修改一条边的权重
* @param  id1:待修改边的第一个顶点
* @param  id2:待修改边的第二个顶点
* @param  wt:新的权重
* @return error
*
*
 */
func (a *MatrixGraph) AdjustEdge(id1, id2, wt int) error {
	if id1 < 0 || id1 >= a._N || id2 < 0 || id2 >= a._N {
		return errors.New("param is error")
	}

	b, _ := a.HasEdge(id1, id2)
	if !b {
		return errors.New("edge adjust error,edge does not exist.")
	}

	a.Matrix[id1][id2] = wt
	return nil
}

/*!
* @description:返回图中所有边的三元素元组集合
* @return  :图中所有边的三元素元组集合
*
 */
func (a *MatrixGraph) EdgeTuples() []*Tuple {

	result := []*Tuple{}
	for i := 0; i < a._N; i++ {
		for j := 0; j < a._N; j++ {
			val := a.Matrix[i][j]
			if val != a.invalidWeight {
				result = append(result, NewTuple(i, j, val))
			}
		}
	}
	return result
}

/*!
* @description:返回图中从指定顶点出发的边的三元素元组集合
* @param id: 指定顶点`id`
* @return  :图中指定顶点出发的边的三元素元组集合
*
 */
func (a *MatrixGraph) VertexEdgeTuples(id int) ([]*Tuple, error) {

	if id < 0 || id >= a._N {
		return nil, errors.New("vertex_edge_tuples: id must belongs [0,N),")
	}
	result := []*Tuple{}
	for j := 0; j < a._N; j++ {
		val := a.Matrix[id][j]
		if val != a.invalidWeight {
			result = append(result, NewTuple(id, j, val))
		}
	}
	return result, nil
}

/*!
* @description:返回图中指定顶点之间是否存在边
* @param id_from: 第一个顶点的`id`
* @param id_to: 第二个顶点的`id`
* @return  :第一个顶点和第二个顶点之间是否存在边
*
 */
func (a *MatrixGraph) HasEdge(id_from, id_to int) (bool, error) {

	if id_from < 0 || id_from >= a._N || id_to < 0 || id_to >= a._N {
		return false, errors.New("has_edge: id_from  and id _to must belongs [0,N),")
	}
	return a.Matrix[id_from][id_to] != a.invalidWeight, nil
}

/*!
* @description:返回图中指定顶点之间的边的权重
* @param id_from: 第一个顶点的`id`
* @param id_to: 第二个顶点的`id`
* @return  :第一个顶点和第二个顶点之间的边的权重
*
 */
func (a *MatrixGraph) Weight(id_from, id_to int) (int, error) {
	b, _ := a.HasEdge(id_from, id_to)
	if b {
		return a.Matrix[id_from][id_to], nil
	} else {
		return -1, errors.New("weight error: the edge does not exist.")
	}
}

/**
 * @description: 矩阵图的创建函数
 * @param invalidWeight:将何值设定为非法权重
 * @param n:矩阵长宽规模
 * @return 返回图矩阵的指针
 */
func NewMatrixGraph(invalidWeight int, n int) *MatrixGraph {
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
		for j := 0; j < n; j++ {
			matrix[i][j] = invalidWeight
		}
	}
	return &MatrixGraph{invalidWeight: invalidWeight, _N: n, Matrix: matrix}
}
