/*
 * @Description: 实现queue数据结构，先进先出
 * @Author: wangchengdg@gmail.com
 * @Date: 2020-02-20 11:07:44
 * @LastEditTime: 2020-03-06 13:00:14
 * @LastEditors:
 */
package Common

type Node struct {
	data interface{}
	next *Node
}

//先进先出
type Queue struct {
	head *Node
	end  *Node
}

func NewQueue() *Queue {
	q := &Queue{nil, nil}
	return q
}

func (q *Queue) Push(data interface{}) {
	n := &Node{data: data, next: nil}

	if q.end == nil {
		q.head = n
		q.end = n
	} else {
		q.end.next = n
		q.end = n
	}

	return
}

func (q *Queue) Pop() (interface{}, bool) {
	if q.head == nil {
		return nil, false
	}

	data := q.head.data
	q.head = q.head.next
	if q.head == nil {
		q.end = nil
	}

	return data, true
}

func (q *Queue) Front() interface{} {
	if q.head != nil {
		return q.head.data
	}
	return nil
}
func (q *Queue) Empty() bool {
	return q.head == nil
}
