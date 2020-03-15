/*
 * @Description: 前置重贴标签算法中图的节点类型 FrontFlowVertex
 * @Author: wangchengdg@gmail.com
 * @Date: 2020-02-18 10:31:22
 * @LastEditTime: 2020-03-13 22:27:34
 * @LastEditors:
 */
package GraphVertex

/*!
*
* FrontFlowVertex 继承自 FlowVertex，它比FlowVertex顶点多了一个`N_List`数据成员，表示邻接链表
*
* relabel_to_front 算法中，每一个FrontFlowVertex顶点位于两个级别的链表中：
*
* - L 链表：最顶层的链表，L包含了所有的非源、非汇顶点
* - u.N 链表：某个顶点u的邻接链表
*
 */
type FrontFlowVertex struct {
	FlowVertex
	N_List *List //存储和本节点相邻的所有节点
}

func NewFrontFlowVertex(k int, ids ...int) *FrontFlowVertex {
	id := -1
	if len(ids) > 0 {
		id = ids[0]
	}
	newList := &List{Head: nil, Current: nil}
	return &FrontFlowVertex{FlowVertex: FlowVertex{Vertex: Vertex{_ID: id, _Key: k}}, N_List: newList}
}
