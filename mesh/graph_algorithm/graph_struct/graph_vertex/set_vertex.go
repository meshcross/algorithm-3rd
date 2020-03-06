/*
 * @Description: golang algorithm
 * @Author: wangchengdg@gmail.com
 * @Date: 2020-02-18 10:31:09
 * @LastEditTime: 2020-03-05 11:59:39
 * @LastEditors:
 */
package GraphVertex

import . "github.com/meshcross/algorithm-3rd/mesh/set_algorithm"

//!SetVertex：图的顶点，它带一个node属性，算法导论22章22.1节
/*!
*
* 它继承自Vertex，区别在于多了一个node成员变量，这个node是指向DisjointSetNode<SetVertex>对象的强指针。
* 而DisjointSetNode<SetVertex>的value成员变量是指向SetVertex的弱指针，二者可以相互访问。
* >SetVertex没有父结点指针。要想访问SetVertex的父结点，首先要取得它对应的DisjointSetNode结点。然后获取该DisjointSetNode结点的父结点。
* 然后通过该父结点获取value指向的SetVertex即可。
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
