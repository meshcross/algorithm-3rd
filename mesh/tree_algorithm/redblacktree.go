/*
 * @Description: 第13章 红黑树
	红黑树是一颗二叉搜索树，它在每个节点上增加了一个存储位来表示节点的颜色，可以使RED或者BLACK。通过对任何一条从根到叶子的简单路径上各个节点的颜色进行约束，
	红黑树确保没有一条路径会比其他路径长出2倍，因而近似于平衡。
	树中每个节点包含5个属性：Color/Key/Left/Right/Parent
	如果一个节点没有子节点或父节点，则该节点相应 指针属性的值为NIL,NIL被视为外部节点，有Key值的视为内部节点

	一颗红黑树是满足下面红黑性质的二叉搜索树：
	1、每个节点是红色或者黑色
	2、根节点是黑色
	3、每个叶节点是黑色
	4、如果一个节点是红色的，则它的两个子节点都是黑色的，反之则不一定
	5、对每个节点，从该节点到其所有后代节点的简单路径上，均包含相同数目的黑色节点。

	红黑树的典型用途是平衡数组，map和set的底层结构
	红黑树对插入时间、删除时间和查找时间提供了最好可能的最坏情况担保O(lgn)。

 * @Author: wangchengdg@gmail.com
 * @Date: 2020-03-03 17:27:15
 * @LastEditTime: 2020-03-05 11:52:42
 * @LastEditors:
*/

package TreeAlgorithm

import (
	"fmt"

	. "github.com/meshcross/algorithm-3rd/mesh"

	. "github.com/meshcross/algorithm-3rd/mesh/common"
)

type RedBlackTreeChecker struct {
	_NIL *RedBlackTreeNode
}

func (a *RedBlackTreeChecker) IsNil(v ITreeNode) bool {
	return v == a._NIL
}
func (a *RedBlackTreeChecker) GetNil() ITreeNode {
	return a._NIL
}

type RedBlackTree struct {
	BinarySearchTree
	//nil哨兵节点，全局共用,存储在RedBlackTreeNode中
}

func NewRedBlackTree(compare NodeCompareFunc) *RedBlackTree {
	//叶节点是黑色的，所有的叶节点都是node_nil
	node_nil := &RedBlackTreeNode{
		Color:          COLOR_BLACK,
		BinaryTreeNode: BinaryTreeNode{_Parent: nil, _LChild: nil, _RChild: nil, _Key: 0},
	}
	checker := &RedBlackTreeChecker{_NIL: node_nil}

	tree := &RedBlackTree{BinarySearchTree: BinarySearchTree{Compare: compare}}
	tree.Root = node_nil
	tree._Checker = checker
	return tree
}

func (a *RedBlackTree) Insert(node *RedBlackTreeNode) {
	node.SetLChild(a.GetNil())
	node.SetRChild(a.GetNil())
	node.Color = COLOR_RED //相对于searchtree增加了颜色设置，以及对于树的颜色调整

	if a.Compare == nil {
		panic("compare is nil")
	}
	compare := a.Compare

	if a.IsNil(a.Root) {
		node.SetParent(a.GetNil())
		a.Root = node
	} else {
		temp := a.Root
		var temp_parent ITreeNode = nil
		left := true
		for !a.IsNil(temp) { //遍历

			temp_parent = temp                             //指向父节点
			if compare(node.GetKey(), temp.GetKey()) > 0 { //向左侧遍历

				temp = temp.GetLChild()
				left = true
			} else {
				temp = temp.GetRChild() //向右侧遍历
				left = false
			}
		}
		//现在temp为空，node挂载到temp_parent下
		node.SetParent(temp_parent)
		if left {
			temp_parent.SetLChild(node)
		} else {
			temp_parent.SetRChild(node)
		}

	}

	a.insert_fixup(node)
}

/**
* 红黑树加入新节点z(标记为红色)之后，如果违反了红黑树的规则只有可能是两种情况：
* a、违反规则2，即根节点为红色，则此时根节点必然是z，即空数加入了节点z
* b、违反规则4，必然有z和z.p都为红色，重点调整此处冲突
*
* 插入时候旋转遵循以下规律，插入节点为z，p=z.p：
* LL ——>R  p是左节点，并且z是左节点，则右旋转
* RL—->LR  p是右节点，z是左节点，则先左旋转再右旋转
* RR——>L   p是右节点，z是右节点，则左旋转
* LR——->RL p是左节点，z是右节点，则先右旋转再左旋转
**/
func (a *RedBlackTree) insert_fixup(z *RedBlackTreeNode) {

	for ToRedBlackTreeNode(z.GetParent()).Color == COLOR_RED {
		p := ToRedBlackTreeNode(z.GetParent())
		pp := ToRedBlackTreeNode(p.GetParent())
		//父节点为左节点
		if IsLeftChild(p) {
			y := ToRedBlackTreeNode(pp.GetRChild())

			if y.Color == COLOR_RED {
				p.Color = COLOR_BLACK
				y.Color = COLOR_BLACK
				pp.Color = COLOR_RED
				z = pp
			} else if IsRightChild(z) {
				z = p
				a.LeftRotate(z, &a.Root)
			} else if IsLeftChild(z) {
				new_p := ToRedBlackTreeNode(z.GetParent())
				new_pp := ToRedBlackTreeNode(new_p.GetParent())
				new_p.Color = COLOR_BLACK
				new_pp.Color = COLOR_RED
				//需要考察这里，root变化了
				a.RightRotate(new_pp, &a.Root)
			}
		} else if IsRightChild(p) {
			y := ToRedBlackTreeNode(pp.GetLChild())

			if y.Color == COLOR_RED {
				p.Color = COLOR_BLACK
				y.Color = COLOR_BLACK
				pp.Color = COLOR_RED
				z = pp
			} else if IsLeftChild(z) {
				z = p
				a.RightRotate(z, &a.Root)
			} else if IsRightChild(z) {
				new_p := ToRedBlackTreeNode(z.GetParent())
				new_pp := ToRedBlackTreeNode(new_p.GetParent())
				new_p.Color = COLOR_BLACK
				new_pp.Color = COLOR_RED
				a.LeftRotate(new_p, &a.Root)
			}
		}
		// black_ids := []int{}
		// a.Print()
		// a.PrintColor(a.Root, &black_ids)
		// sort.Ints(black_ids)
		// fmt.Println()
		// fmt.Println("black list:", black_ids)
		// fmt.Println("-------------------------------------------------")
	}
	r := ToRedBlackTreeNode(a.Root)
	r.Color = COLOR_BLACK
}

/**
 * @description:
 * @param {type}
 * @return:
*
* 算法：
*
* - 如果`node`是一个叶子节点：则直接删除它
* - 如果`node`有左子节点，但是没有右子节点：将左子剪切到`node`所在位置
* - 如果`node`有右子节点，但是没有左子节点：将右子剪切到`node`所在位置
* - 如果`node`既有左子节点，又有右子节点：首先获取`node`的后继节点`next_node`
*      - 如果`next_node`就是`node`的右子节点，则证明`next_node`没有左子（如果`next_node`有左子，则`node`的后继节点必然不是`next_node`）。
*        此时将`next_node`剪切到`node`所在位置，并且将`node`的左子挂载到`next_node`的左子
*      - 如果`next_node`不是`node`的右子节点，则`next_node`必然位于`node`右子为根的子树中。且`next_node`必然没有左子（否则`node`的后继节点必然不是`next_node`）
*          - 把`next_node`的右子节点剪切到`next_node`的位置
*          - 将`next_node`剪切到`node`的右子位置
*          - 执行`next_node`就是`node`的右子节点的操作
*
* 算法时间复杂度O(h)，空间复杂度O(1)。其中h为树的高度
*
*/
func (a *RedBlackTree) Delete(node *RedBlackTreeNode) {
	y := node
	y_color := y.Color
	var x *RedBlackTreeNode = ToRedBlackTreeNode(a._Checker.GetNil())

	if a.IsNil(node) {
		panic("node removed is nullptr")
	}
	if a.Compare == nil {
		panic("compare is nil")
	}
	compare := a.Compare
	//*************  判定 node 必须位于树中 ***********
	temp_node := a.Root

	for !a.IsNil(temp_node) && temp_node != node {
		if compare(node.GetKey(), temp_node.GetKey()) > 0 {
			temp_node = temp_node.GetLChild()
		} else {
			temp_node = temp_node.GetRChild()
		}
	}
	if temp_node != node {
		panic("node removed must be in the tree")
	}

	//************** 执行删除过程  *****************
	if a.IsNil(node.GetLChild()) && a.IsNil(node.GetRChild()) { //node是个叶子节点

		parent := node.GetParent()
		if !a.IsNil(parent) { //父节点非空
			if IsLeftChild(node) {
				parent.SetLChild(a.GetNil())
			}
			if IsRightChild(node) {
				parent.SetRChild(a.GetNil())
			}
		} else {
			a.Root = a.GetNil()
		}
	} else if !a.IsNil(node.GetLChild()) && a.IsNil(node.GetRChild()) { //node有左子，但是没有右子
		a.Transplant(node.GetLChild(), node, &a.Root) //左子剪切到node位置
	} else if !a.IsNil(node.GetRChild()) && a.IsNil(node.GetLChild()) { //node有右子，但是没有左子
		a.Transplant(node.GetRChild(), node, &a.Root) //右子剪切到node位置
	} else { //node既有左子，又有右子

		next_node := ToRedBlackTreeNode(a.Successor(node)) //node的后继节点
		y_color = next_node.Color
		x = ToRedBlackTreeNode(next_node.GetRChild())
		//边缘处理，如果后继节点的右节点为空怎么办？这里必须要处理一下，否则可能会失去fixup的机会
		if a.IsNil(x) {
			x = ToRedBlackTreeNode(next_node.GetParent())
			x.Color = COLOR_BLACK
			y_color = COLOR_BLACK
		}
		if next_node != node.GetRChild() { //node后继不是node的右子，则转换成node的后继是node的右子情形

			a.Transplant(next_node.GetRChild(), next_node, &a.Root) //将后继的右子剪切到后继的位置
			next_node.SetRChild(node.GetRChild())                   //后继的右子设为node的右子
			node.GetRChild().SetParent(next_node)                   //后继的右子设为node的右子
			node.SetRChild(next_node)

			a.Transplant(node.GetRChild(), node, &a.Root)
			next_node.SetLChild(node.GetLChild())
			node.GetLChild().SetParent(next_node)
		} else {
			//node后继是node的右子
			a.Transplant(node.GetRChild(), node, &a.Root)
			next_node.SetLChild(node.GetLChild())
			node.GetLChild().SetParent(next_node)
		}

		y = next_node
		y.Color = node.Color
	}

	a.PrintBlackColor("end delete ")
	if y_color == COLOR_BLACK {
		a.delete_fixup(x)
	}
}

/**
 * @description: 有节点被删除之后，需要对节点的红黑颜色进行校正，使得新的配色满足红黑树的5个条件
 * @param
 * @return: void
 */
func (a *RedBlackTree) delete_fixup(x *RedBlackTreeNode) {
	a.PrintBlackColor("delete fixup start ")
	if a.IsNil(x) {
		return
	}

	for x != a.Root && x.Color == COLOR_BLACK {
		p := ToRedBlackTreeNode(x.GetParent())

		if IsLeftChild(x) {
			w := ToRedBlackTreeNode(x.GetParent().GetRChild())
			w_left := ToRedBlackTreeNode(w.GetLChild())
			w_right := ToRedBlackTreeNode(w.GetRChild())

			if w.Color == COLOR_RED { //case 1：x的兄弟w是红色的
				w.Color = COLOR_BLACK
				p.Color = COLOR_RED
				a.LeftRotate(p, &a.Root)

				w = ToRedBlackTreeNode(x.GetParent().GetRChild())
			} else if w_left.Color == COLOR_BLACK && w_right.Color == COLOR_BLACK { //case2：w是黑色，且两个孩子都是黑色
				w.Color = COLOR_RED
				x = ToRedBlackTreeNode(x.GetParent())
			} else if w_right.Color == COLOR_BLACK { //case3：w是黑色，左孩子红色，右孩子黑色
				w_left.Color = COLOR_BLACK
				w.Color = COLOR_RED
				a.RightRotate(w, &a.Root)
				w = ToRedBlackTreeNode(x.GetParent().GetRChild())
			} else if w_right.Color == COLOR_RED { //case4：w是黑色，右孩子是红色
				w.Color = p.Color
				p.Color = COLOR_BLACK
				w_right.Color = COLOR_BLACK
				a.LeftRotate(p, &a.Root)
				x = ToRedBlackTreeNode(a.Root)
			}

		} else if IsRightChild(x) {
			w := ToRedBlackTreeNode(x.GetParent().GetLChild())
			w_left := ToRedBlackTreeNode(w.GetLChild())
			w_right := ToRedBlackTreeNode(w.GetRChild())

			if w.Color == COLOR_RED { //case 1：x的兄弟w是红色的
				w.Color = COLOR_BLACK
				p.Color = COLOR_RED
				a.RightRotate(p, &a.Root)

				w = ToRedBlackTreeNode(x.GetParent().GetLChild())
			} else if w_left.Color == COLOR_BLACK && w_right.Color == COLOR_BLACK { //case2：w是黑色，且两个孩子都是黑色
				w.Color = COLOR_RED
				x = ToRedBlackTreeNode(x.GetParent())
			} else if w_left.Color == COLOR_BLACK { //case3：w是黑色，左孩子黑色，又孩子红色
				w_right.Color = COLOR_BLACK
				w.Color = COLOR_RED
				a.LeftRotate(w, &a.Root)
				w = ToRedBlackTreeNode(x.GetParent().GetLChild())
			} else if w_left.Color == COLOR_RED { //case4：w是黑色，左孩子是红色
				w.Color = p.Color
				p.Color = COLOR_BLACK
				w_left.Color = COLOR_BLACK
				a.RightRotate(p, &a.Root)
				x = ToRedBlackTreeNode(a.Root)
			}
		}
		x.Color = COLOR_BLACK
	}

	a.PrintBlackColor("end delete_fixup")
}

/**
* @description:不同之处只是a.IsNil不同
* @param
* @return:
 */
func (a *RedBlackTree) Transplant(src ITreeNode, dst ITreeNode, root *ITreeNode) {

	if a.IsNil(dst.GetParent()) { //dst是根节点
		*root = src
	} else if IsLeftChild(dst) {
		dst.GetParent().SetLChild(src)
	} else { //dst不是根节点
		dst.GetParent().SetRChild(src)
	}
	src.SetParent(dst.GetParent())
}

func (a *RedBlackTree) PrintBlackColor(msg string) {
	ids2 := []int{}
	a.printColor(a.Root, &ids2)
	fmt.Println()
	fmt.Println(msg, "--black-ids:", ids2)
}

func (a *RedBlackTree) printColor(v ITreeNode, black_ids *[]int) {
	node := ToRedBlackTreeNode(v)
	if a.IsNil(node) {
		return
	}
	//color := "Red"
	if node.Color == COLOR_BLACK {
		//color = "Black"
		id := node.GetKey().(int)
		*black_ids = append(*black_ids, id)
	}
	//fmt.Printf("%d %s;", node.GetKey(), color)

	l := node.GetLChild()
	r := node.GetRChild()
	if !a.IsNil(l) {
		a.printColor(l, black_ids)
	}
	if !a.IsNil(r) {
		a.printColor(r, black_ids)
	}
}
