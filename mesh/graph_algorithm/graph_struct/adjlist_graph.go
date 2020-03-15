/*
 * @Description: 第22章22.1节 图的邻接表表示法
 * @Author: wangchengdg@gmail.com
 * @Date: 2020-02-18 10:29:08
 * @LastEditTime: 2020-03-14 11:34:30
 * @LastEditors:
 */
package GraphStruct

import (
	"errors"

	. "github.com/meshcross/algorithm-3rd/mesh/common"
)

/*!
 * 图的邻接表主要包含一个数据：
 *
 * - `array`：邻接表，每一行代表一个节点：
 *
 * 为了便于计算，这里并不管理边和顶点，只是维护邻接表。边、顶点与邻接表的同步由使用者确保。
 *
 * 稀疏图使用邻接表示法
 */

type ADJListGraph struct {
	array [][]*Pair //邻接表示法
	_N    int
}

func NewADJListGraph(n int) *ADJListGraph {
	arr := make([][]*Pair, n)
	for k := 0; k < n; k++ {
		//arr[k] = make([]*Pair, n)
		arr[k] = []*Pair{}
	}
	return &ADJListGraph{_N: n, array: arr}
}

/*!
* @description:添加一条边
* @param  edge_tuple:一条边的三元素元组
*
 */
func (a *ADJListGraph) AddEdge(edge_tuple *Tuple) error {
	id1 := edge_tuple.First
	id2 := edge_tuple.Second
	wt := edge_tuple.Third

	if id1 < 0 || id1 >= a._N || id2 < 0 || id2 >= a._N {
		return errors.New("edge add param error")
	}
	has, _ := a.HasEdge(id1, id2)
	if has {
		return errors.New("edge add error,edge has already exist.")
	}

	a.array[id1] = append(a.array[id1], NewPair(id2, wt))
	return nil
}

/*!
* @description:添加一组边
* @param  edges:一组边
*
 */
func (a *ADJListGraph) AddEdges(edges []*Tuple) {
	for _, v := range edges {
		a.AddEdge(v)
	}
}

/*!
* @description:修改一条边的权重
* @param  id1:待修改边的第一个顶点
* @param  id2:待修改边的第二个顶点
* @param  wt:新的权重
*
* > 要求`id1`和`id2`均在`[0,N)`这个半闭半开区间。如果任何一个值超过该区间则认为顶点`id`无效，直接返回而不作权重修改
*
 */
func (a *ADJListGraph) AdjustEdge(id1, id2, wt int) error {
	if id1 < 0 || id1 >= a._N || id2 < 0 || id2 >= a._N {
		return errors.New("adjust edge params error")
	}

	has, _ := a.HasEdge(id1, id2)
	if !has {
		return errors.New("edge adjust error,edge does not exist.")
	}

	vec := a.array[id1] //这里必须用引用类型，因为要修改邻接表
	for _, pair := range vec {
		if pair.First == id2 {
			pair.Second = wt
			break
		}
	}
	return nil
}

/*!
* @description:返回图中所有边的三元素元组集合
* @return  :图中所有边的三元素元组集合
*
 */
func (a *ADJListGraph) EdgeTuples() []*Tuple {
	result := []*Tuple{}
	for i := 0; i < a._N; i++ {
		for _, pair := range a.array[i] {
			result = append(result, NewTuple(i, pair.First, pair.Second))
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
func (a *ADJListGraph) VertexEdgeTuples(id int) ([]*Tuple, error) {
	if id < 0 || id >= a._N {
		return nil, errors.New("vertex_edge_tuples: id must belongs [0,N),")
	}
	result := []*Tuple{}
	for _, pair := range a.array[id] {
		result = append(result, NewTuple(id, pair.First, pair.Second))
	}
	return result, nil
}

/*!
* @description:顶点之间是否存在边
* @param id_from: 第一个顶点的`id`
* @param id_to: 第二个顶点的`id`
* @return  :第一个顶点和第二个顶点之间是否存在边
*
* - 当`id_from`与`id_to`无效时，抛出异常
* >这里的无效值得是`id_from`、`id_to`不在区间`[0,N)`之间
* - 当`id_from`与`id_to`之间有边时，返回`true`
* - 当`id_from`与`id_to`之间没有边时，返回`false`
 */
func (a *ADJListGraph) HasEdge(id_from, id_to int) (bool, error) {
	if id_from < 0 || id_from >= a._N || id_to < 0 || id_to >= a._N {
		return false, errors.New("has_edge: id_from  and id _to must belongs [0,N),")
	}
	vec := a.array[id_from]
	for _, pair := range vec {
		if pair.First == id_to {
			return true, nil
		}
	}
	return false, nil
}

/*!
* @description:返回图中指定顶点之间的边的权重
* @param id_from: 第一个顶点的`id`
* @param id_to: 第二个顶点的`id`
* @return  :第一个顶点和第二个顶点之间的边的权重
*
 */
func (a *ADJListGraph) Weight(id_from, id_to int) (int, error) {
	has, _ := a.HasEdge(id_from, id_to)
	if has {
		vec := a.array[id_from]
		for _, pair := range vec {
			if pair.First == id_to {
				return pair.Second, nil
			}
		}
	} else {
		return -1, errors.New("weight error: the edge does not exist.")
	}
	return -1, nil
}
