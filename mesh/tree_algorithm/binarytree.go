/*
 * @Description: 二叉树，以及二叉树的前序遍历/中序遍历/后序遍历，左旋转/右旋转，剪切
 * @Author: wangchengdg@gmail.com
 * @Date: 2020-02-10 22:20:07
 * @LastEditTime: 2020-03-13 12:41:11
 * @LastEditors:
 */
package TreeAlgorithm

import (
	"errors"
	"fmt"
	"math"
	"strings"

	. "github.com/meshcross/algorithm-3rd/mesh/common"
)

type BinaryTreeNodeFunc func(n ITreeNode)

type BinaryTree struct {
	Root     ITreeNode
	_Checker INilChecker
}
type NodePrintInfo struct {
	Value  string
	Dx     int
	Dy     int
	Level  int
	IsLine bool
}
type INilChecker interface {
	IsNil(v ITreeNode) bool
	GetNil() ITreeNode
}
type BinaryTreeChecker struct {
}

func (a *BinaryTreeChecker) IsNil(v ITreeNode) bool {
	return v == nil
}
func (a *BinaryTreeChecker) GetNil() ITreeNode {
	return nil
}
func NewBinaryTree() *BinaryTree {
	return &BinaryTree{_Checker: &BinaryTreeChecker{}}
}

/**
 * @description: 红黑树中有设定伪结点，并不是真的为Nil，所以抽象出一层判定Nil
 * @param v :ITreeNode
 * @return:
 */
func (a *BinaryTree) IsNil(v ITreeNode) bool {
	return a._Checker.IsNil(v)
}
func (a *BinaryTree) GetNil() ITreeNode {
	return a._Checker.GetNil()
}

/*
* @description:二叉树的中序遍历
* @param root：从根节点开始遍历
* @param  action:回调函数，每次有节点被访问到就会调用该回调函数
* @return void
*
* 本函数执行对二叉树的中序遍历，遍历时执行指定操作
* 算法：
* - 对左子节点前序遍历
* - 对本节点执行操作
* - 对右子节点前序遍历
*
* 时间复杂度O(n)，空间复杂度O(1)
 */
func (a *BinaryTree) InOrder(root ITreeNode, action BinaryTreeNodeFunc) {
	if !a.IsNil(root) {
		if !a.IsNil(root.GetLChild()) {
			a.InOrder(root.GetLChild(), action)
		}
		if action != nil {
			action(root)
		}
		if !a.IsNil(root.GetRChild()) {
			a.InOrder(root.GetRChild(), action)
		}
	}
}

/*!
* @description:二叉树的前序遍历
* @param root：从根节点开始遍历
* @param  action:回调函数，每次有节点被访问到就会调用该回调函数
* @return void
* 本函数执行对二叉树的前序遍历，遍历时执行指定操作
* 算法：
* - 对本节点执行操作
* - 对左子节点前序遍历
* - 对右子节点前序遍历
*
* 时间复杂度O(n)，空间复杂度O(1)
 */
func (a *BinaryTree) PreOrder(root ITreeNode, action BinaryTreeNodeFunc) {
	if !a.IsNil(root) {
		if action != nil {
			action(root)
		}
		if !a.IsNil(root.GetLChild()) {
			a.PreOrder(root.GetLChild(), action)
		}
		if !a.IsNil(root.GetRChild()) {
			a.PreOrder(root.GetRChild(), action)
		}
	}
}

/*!
* @description:二叉树的后序遍历
* @param root：从根节点开始遍历
* @param  action:回调函数，每次有节点被访问到就会调用该回调函数
* @return void
* 本函数执行对二叉树的后序遍历，遍历时执行指定操作。
* 算法：
*
* - 对左子节点后续遍历
* - 对右子节点后续遍历
* - 对本节点执行操作
*
* 时间复杂度O(n)，空间复杂度O(1)
 */
func (a *BinaryTree) PostOrder(root ITreeNode, action BinaryTreeNodeFunc) {
	if !a.IsNil(root) {
		if !a.IsNil(root.GetLChild()) {
			a.PostOrder(root.GetLChild(), action)
		}
		if !a.IsNil(root.GetRChild()) {
			a.PostOrder(root.GetRChild(), action)
		}
		if action != nil {
			action(root)
		}
	}
}

//!left_rotate：
/*!
* @description：二叉树的左旋转操作
* @param
* @param
* @return
*
* 本函数执行二叉树进行左旋转。设node为被旋转的点，l_node为它的左子节点，r_node为它的右子节点。
* 则旋转的效果是：r_node取代了node的位置；而node现在挂靠在r_node的左子；r_node原来的左子现在成为node的右子
*
* <code><pre>
*         |                                                |
*        node                                            r_node
*       /    \                                           /    \
*   l_node  r_node        -- 左旋  -->                  node  r_r_node
*          /    \                                     /    \
*     l_r_node r_r_node                           l_node  l_r_node
* </pre></code>
*
* 时间复杂度O(1)，空间复杂度O(1)
 */
func (a *BinaryTree) LeftRotate(node ITreeNode, root *ITreeNode) {
	if !a.IsNil(node) { //node非空
		right := node.GetRChild() //r_node
		if !a.IsNil(right) {      //r_node非空
			l_right := right.GetLChild()      //l_r_node;
			right.SetParent(node.GetParent()) //r_node提至node的原位置
			parent := node.GetParent()        //node的父节点
			if !a.IsNil(parent) {             //node的父节点修订
				if IsLeftChild(node) {
					parent.SetLChild(right)
				}
				if IsRightChild(node) {
					parent.SetRChild(right)
				}
			} else {
				*root = right
			}
			node.SetRChild(l_right) //r_node的左子节点剪切到 node的右侧
			if !a.IsNil(l_right) {
				l_right.SetParent(node)
			}

			node.SetParent(right) //node剪切到 r_node的左侧
			right.SetLChild(node)
		}
	}
}

/**
左右旋转之后不会改变二叉树的中序遍历结果

本函数执行二叉树进行右旋转。设node为被旋转的点，l_node为它的左子节点，r_node为它的右子节点。
* 则旋转的效果是：l_node取代了node的位置；而node现在挂靠在r_node的右子；l_node原来的右子现在成为node的左子
*
* <code><pre>
*                   |                                                |
*                  node                                            l_node
*                 /    \                                           /    \
*              l_node  r_node        -- 右旋  -->              l_l_node  node
*             /    \                                                    /    \
*        l_l_node r_l_node                                         r_l_node  r_node
* </pre></code>
* 时间复杂度O(1)，空间复杂度O(1)
*/

func (a *BinaryTree) RightRotate(node ITreeNode, root *ITreeNode) {
	if !a.IsNil(node) { //node非空
		left := node.GetLChild() //l_node非空
		if !a.IsNil(left) {      //l_node非空

			r_left := left.GetRChild()       //r_l_node
			left.SetParent(node.GetParent()) //l_node提至node的原位置
			parent := node.GetParent()
			if !a.IsNil(parent) { //node的父节点修订

				if IsRightChild(node) {
					parent.SetRChild(left)
				}
				if IsLeftChild(node) {
					parent.SetLChild(left)
				}
			} else {
				*root = left
			}
			node.SetLChild(r_left) //l_node的右子节点剪切到 node的左侧
			if !a.IsNil(r_left) {
				r_left.SetParent(node)
			}

			node.SetParent(left) //node剪切到 l_node的右侧
			left.SetRChild(node)
		}
	}
}

/**
root节点可能发生变化

本函数执行二叉树的剪切操作。剪切操作要求目的节点是非空节点（若是空节点则抛出异常）。本操作将node_src剪切到node_dst中。
下图中，两个节点之间有两个链接：一个是父-->子，另一个是子-->父
* 剪切操作可能会仅仅剪断一个连接
*
*  <code><pre>
*                 src_p         dst_p                                    src_p         dst_p             dst_p
*                   ||            ||                               (父-->子)|             ||                |(子--->父)
*                 node_src     node_dst         ---剪切 -->               node_src      node_src        node_dst
*                 //   \\      //    \\                                 //     \\      //     \\        //     \\
* </pre></code>
*
* 时间复杂度O(1)，空间复杂度O(1)
*/
func (a *BinaryTree) Transplant(src ITreeNode, dst ITreeNode, root *ITreeNode) error {
	if a.IsNil(dst) {
		return errors.New("dst is nil ")
	}
	if !a.IsNil(src) { //替换掉parent
		src.SetParent(dst.GetParent())
	}
	if a.IsNil(dst.GetParent()) { //dst是根节点
		*root = src
	} else { //dst不是根节点

		if IsLeftChild(dst) { //dst是左子节点
			dst.GetParent().SetLChild(src)
		} else { //dst是右子节点
			dst.GetParent().SetRChild(src)
		}
	}
	return nil
}

func (a *BinaryTree) GetHeight(nodes ...ITreeNode) int {
	node := a.Root
	if len(nodes) > 0 {
		node = nodes[0]
	}
	if a.IsNil(node) {
		return 0
	}
	leftHeight := a.GetHeight(node.GetLChild())
	rightHeight := a.GetHeight(node.GetRChild())

	ret := math.Max(float64(leftHeight), float64(rightHeight)) + 1.0
	return int(ret)
}

func (a *BinaryTree) updateNodeLevel() {
	nodes := a.GetNodes(a.Root)
	l := len(nodes)
	for k := 0; k < l; k++ {
		node := nodes[k]
		node.UpdateLevel()
	}
}
func (a *BinaryTree) GetNodes(node ITreeNode) []ITreeNode {

	if a.IsNil(node) {
		panic("node is nil")
	}
	arr := []ITreeNode{}
	action := func(n ITreeNode) {
		arr = append(arr, n)
	}
	a.InOrder(node, action)
	return arr
}

func (a *BinaryTree) GetNodesOfLevel(level int) []ITreeNode {
	h := a.GetHeight()
	if level <= 0 || level > h {
		return []ITreeNode{}
	}
	a.updateNodeLevel()
	//node := a.Root
	arr := []ITreeNode{}
	nodes := a.GetNodes(a.Root)
	l := len(arr)
	for k := 0; k < l; k++ {
		node := nodes[k]
		if node.GetLevel() == level {
			arr = append(arr, node)
		}
	}
	return arr
}

/**
控制台打印二叉树,必须要按层打印
dx,dy表示格子数偏移
**/
func (a *BinaryTree) GetPrintInfos(node ITreeNode, dx, dy int, level int) []*NodePrintInfo {

	h := a.GetHeight()
	unit := h - level
	unit = int(float64(unit) * 0.7)
	if unit < 1 {
		unit = 1 //简单处理可以直接设定unit=1
	}

	// unit:=1
	levelInfos := []*NodePrintInfo{}
	if node != nil {
		info := &NodePrintInfo{Value: ToString(node.GetKey()), Dx: dx, Dy: dy, Level: level, IsLine: false}
		levelInfos = append(levelInfos, info)
		left := node.GetLChild()
		right := node.GetRChild()
		if left != nil {
			lInfo := &NodePrintInfo{Value: "/", Dx: dx - unit, Dy: dy + unit, Level: level, IsLine: true}
			levelInfos = append(levelInfos, lInfo)
			arr := a.GetPrintInfos(left, dx-2*unit, dy+2*unit, level+1)
			levelInfos = append(levelInfos, arr...)
		}

		if right != nil {
			rInfo := &NodePrintInfo{Value: "\\", Dx: dx + unit, Dy: dy + unit, Level: level, IsLine: true}
			levelInfos = append(levelInfos, rInfo)
			arr := a.GetPrintInfos(right, dx+2*unit, dy+2*unit, level+1)
			levelInfos = append(levelInfos, arr...)
		}
	}

	return levelInfos
}

func (a *BinaryTree) Print() {
	fmt.Println("######start print binary tree ######")
	infos := a.GetPrintInfos(a.Root, 0, 0, 1)
	h := a.GetHeight()
	for k := 1; k <= h; k++ {
		// info := infos[k]
		line1 := a.GetLevelString(infos, k, h, false)
		line2 := a.GetLevelString(infos, k, h, true)
		fmt.Println(line1)
		fmt.Println(line2)
	}
	fmt.Println("------end print binary tree ------")
}

func (a *BinaryTree) GetLevelString(arr []*NodePrintInfo, lv, h int, isLine bool) string {
	delt := 40
	l := len(arr)
	strs := []string{}
	lastx := 0
	for k := 0; k < l; k++ {
		info := arr[k]
		if info.Level == lv && info.IsLine == isLine {
			x := delt + info.Dx
			// word_x := h - lv
			blank := GetXString(" ", x-lastx)
			strs = append(strs, blank, info.Value)
			lastx = x
		}
	}
	return strings.Join(strs, "")
}
