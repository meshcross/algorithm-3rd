/*
 * @Description: 第22章 图的表示,SetVertex
 * @Author: wangchengdg@gmail.com
 * @Date: 2020-02-18 10:31:09
 * @LastEditTime: 2020-03-12 13:13:01
 * @LastEditors:
 */
package GraphVertex

import . "github.com/meshcross/algorithm-3rd/mesh/set_algorithm"

/*!
* 图的顶点，它带一个node属性
*
* 它继承自Vertex，实现了IVertex接口，有一个DisjointSetNode类型的变量Node。
 */
type SetVertex struct {
	Vertex
	Node *DisJointSetNode ///*!< 顶点对应的DisjointSetNode*/
}

func NewSetVertex(k int, ids ...int) *SetVertex {
	id := -1
	if len(ids) > 0 {
		id = ids[0]
	}
	return &SetVertex{Vertex: Vertex{_ID: id, _Key: k}, Node: nil}
}
