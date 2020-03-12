/*
 * @Description: 第26章 最大流 26.5 前置重贴标签算法 FrontFlowVertex
 * @Author: wangchengdg@gmail.com
 * @Date: 2020-02-18 10:31:22
 * @LastEditTime: 2020-03-12 17:21:25
 * @LastEditors:
 */
package GraphVertex

import "errors"

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

/*!
* @description :链表数据结构
* 链表包含两个数据成员：
*
* - head：指向链表头部的元素
* - current:指向链表当前处理元素
*
 */
type List struct {
	Head    *ListNode
	Current *ListNode
}

func (l *List) Add(element *ListNode) {
	if element == nil {
		return
	}
	if l.Head != nil { //链表非空
		element.Next = l.Head
		l.Head = element
	} else { //链表为空
		l.Head = element
	}
}

/*!
* @description:链表指定元素前面的元素
* @param element:指定的元素
* @return:指定元素前面元素
*
* 若element为空指针则抛出异常。否则遍历列表。若找到指定元素，则返回其前面元素的指针；若找不到指定元素，
* 则抛出异常
*
 */
func (l *List) FrontOf(element *ListNode) (*ListNode, error) {
	if element == nil {
		return nil, errors.New("front_of error: element must not be null!")
	}

	var pre *ListNode = nil
	cur := l.Head
	for cur != nil {
		if element == cur {
			break
		}
		pre = cur
		cur = cur.Next
	}
	if cur == nil {
		return nil, errors.New("front_of error: element not in the list!")
	} else {
		return pre, nil
	}
}

/*!
* @description :链表结点的数据结构
*
* 链表结点包含两个数据成员：
*
* - value：链表结点保存数据
* - next:指向本链表中当前结点的下一个结点
*
 */
type ListNode struct {
	Value *FrontFlowVertex
	Next  *ListNode
}

func NewFrontFlowVertex(k int, ids ...int) *FrontFlowVertex {
	id := -1
	if len(ids) > 0 {
		id = ids[0]
	}
	newList := &List{Head: nil, Current: nil}
	return &FrontFlowVertex{FlowVertex: FlowVertex{Vertex: Vertex{_ID: id, _Key: k}}, N_List: newList}
}
