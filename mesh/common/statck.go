/*
 * @Description: 定义stack数据结构，后进先出
 * @Author: wangchengdg@gmail.com
 * @Date: 2020-02-20 11:08:03
 * @LastEditTime: 2020-03-06 13:00:07
 * @LastEditors:
 */
package Common

//后进先出
type Stack struct {
	head *Node
}

func NewStack() *Stack {
	s := &Stack{nil}
	return s
}

func (s *Stack) Push(data interface{}) {
	n := &Node{data: data, next: s.head}
	s.head = n
}

func (s *Stack) Pop() (interface{}, bool) {
	n := s.head
	if s.head == nil {
		return nil, false
	}
	s.head = s.head.next
	return n.data, true
}
