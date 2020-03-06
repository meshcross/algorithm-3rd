/*
 * @Description: golang algorithm
 * @Author: wangchengdg@gmail.com
 * @Date: 2020-02-10 22:20:14
 * @LastEditTime: 2020-03-05 12:19:42
 * @LastEditors:
 */
package TreeAlgorithm

import (
	"fmt"

	. "github.com/meshcross/algorithm-3rd/mesh/common"
)

type ITreeNode interface {
	GetParent() ITreeNode
	SetParent(v ITreeNode)

	GetRChild() ITreeNode
	SetRChild(v ITreeNode)

	GetLChild() ITreeNode
	SetLChild(v ITreeNode)

	GetKey() interface{}
	SetKey(v interface{})

	GetLevel() int
	SetLevel(v int)

	// IsLeftChild() bool
	// IsRightChild() bool
	UpdateLevel()
}

type BinaryTreeNode struct {
	// _Parent *BinaryTreeNode /*!< 节点的父节点*/
	// _LChild *BinaryTreeNode /*!< 节点的左子节点*/
	// _RChild *BinaryTreeNode
	// _Key    interface{}
	// _Level  int /**第几层的节点**/
	_Parent ITreeNode /*!< 节点的父节点*/
	_LChild ITreeNode /*!< 节点的左子节点*/
	_RChild ITreeNode
	_Key    interface{}
	_Level  int /**第几层的节点**/
}

func NewBinaryTreeNode(k interface{}) *BinaryTreeNode {
	return &BinaryTreeNode{_Key: k}
}
func (a *BinaryTreeNode) GetParent() ITreeNode {
	return a._Parent
}
func (a *BinaryTreeNode) SetParent(v ITreeNode) {
	a._Parent = v
}
func (a *BinaryTreeNode) GetLChild() ITreeNode {
	return a._LChild
}
func (a *BinaryTreeNode) SetLChild(v ITreeNode) {
	a._LChild = v
}
func (a *BinaryTreeNode) GetRChild() ITreeNode {
	return a._RChild
}
func (a *BinaryTreeNode) SetRChild(v ITreeNode) {
	a._RChild = v
}
func (a *BinaryTreeNode) GetKey() interface{} {
	return a._Key
}
func (a *BinaryTreeNode) SetKey(v interface{}) {
	a._Key = v
}
func (a *BinaryTreeNode) GetLevel() int {
	return a._Level
}
func (a *BinaryTreeNode) SetLevel(v int) {
	a._Level = v
}

func (a *BinaryTreeNode) ToXML() string {

	str := fmt.Sprintf("\n<node>%s\n", a._Key)
	if a._Parent != nil {
		str += fmt.Sprintf("\t <parent>%s</parent>", a._Parent.GetKey())
	} else {
		str += "\t <parent>null_ptr</parent>"
	}
	left := ""
	right := ""
	if a._LChild != nil {
		left = fmt.Sprintf("\n\t <left_child>%s</left_child>", a._LChild.GetKey())
	} else {
		left = fmt.Sprintf("\n\t <left_child>nullptr</left_child>")
	}

	if a._RChild != nil {
		right = fmt.Sprintf("\n\t<right_child>%s</right_child>", a._RChild.GetKey())
	} else {
		right = fmt.Sprintf("\n\t<right_child>nullptr</right_child>")
	}

	str = str + left + right + "\n</node>"
	return str

}

func (a *BinaryTreeNode) ToString() string {

	parent_str := ""
	if a._Parent == nil {
		parent_str = fmt.Sprintf("parent:%s", "nullptr")
	} else {
		parent_str = fmt.Sprintf("parent:%s", a._Parent.GetKey())
	}

	left_str := ""
	right_str := ""

	if a._LChild != nil {
		left_str = fmt.Sprintf("%d", a._LChild.GetKey())
	} else {
		left_str = "nullptr"
	}

	if a._RChild != nil {
		right_str = fmt.Sprintf("%d", a._RChild.GetKey())
	} else {
		right_str = "nullptr"
	}

	str := ""
	str = fmt.Sprintf("node: \n\t%s\n\t%s\n\t%s", parent_str, left_str, right_str)

	return str
}

// func (a *BinaryTreeNode) IsLeftChild() bool {
// 	return a == a._Parent.GetLChild()
// }

//继承之后会造成数据指针类型不一样，a的类型为*BinaryTreeNode，a._Parent.GetRChild()的类型为*RedBlackTreeNode
//虽然指向同一个地址，但是指针类型不一样，所以==操作会返回false
//所以干脆把IsLeftChild和IsRightChild写成静态方法,就没有指针类型的困扰了
// func (a *BinaryTreeNode) IsRightChild() bool {

// 	return a == a._Parent.GetRChild()
// }

func (a *BinaryTreeNode) UpdateLevel() {
	if a._Parent == nil {
		a._Level = 1
	}

	level := 1
	p := a._Parent
	for p.GetParent() != nil {
		p = p.GetParent()
		level++
	}
	a._Level = level
}

type RedBlackTreeNode struct {
	BinaryTreeNode
	Color COLOR //节点颜色
}

// func (a *RedBlackTreeNode) IsLeftChild() bool {
// 	return a == a._Parent.GetLChild()
// }

// //继承之后会造成数据指针类型不一样，a的类型为*BinaryTreeNode，a._Parent.GetRChild()的类型为*RedBlackTreeNode
// //虽然指向同一个地址，但是指针类型不一样，所以==操作会返回false
// func (a *RedBlackTreeNode) IsRightChild() bool {
// 	return a == a._Parent.GetRChild()
// }

func NewRedBlackTreeNode(k interface{}, colors ...COLOR) *RedBlackTreeNode {
	color := COLOR_WHITE
	if len(colors) > 0 {
		color = colors[0]
	}
	return &RedBlackTreeNode{BinaryTreeNode: BinaryTreeNode{_Key: k}, Color: color}
}

func ToBinaryTreeNode(node ITreeNode) *BinaryTreeNode {
	if v, ok := node.(*BinaryTreeNode); ok {
		return v
	}
	return nil
}
func ToRedBlackTreeNode(node ITreeNode) *RedBlackTreeNode {
	if v, ok := node.(*RedBlackTreeNode); ok {
		return v
	}
	return nil
}
