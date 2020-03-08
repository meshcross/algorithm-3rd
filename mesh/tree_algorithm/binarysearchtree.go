/*
 * @Description: 二叉搜索树，算法导论12章
 *				二叉搜索树是一种特殊的二叉树。在二叉树中的任何一个节点，该节点的左子节点值小于它；该节点的右子节点值大于它。
 *				该算法有个缺陷，经过多次插入和删除之后，有可能改变树的形态，变为类似如下的结构：
					8
					  \
					   10
						 \
						  15
							\
							 20
				这样的结构进行搜索，其搜索性能已经变为O(n)了
				为了避免出现这样的情况，一般需要在二叉搜索树中加入平衡算法，以让树左右平衡
 * @Author: wangchengdg@gmail.com
 * @Date: 2020-02-10 22:20:21
 * @LastEditTime: 2020-03-06 16:17:52
 * @LastEditors:
 */
package TreeAlgorithm

import (
	"errors"

	. "github.com/meshcross/algorithm-3rd/mesh"
)

type BinarySearchTree struct {
	BinaryTree
	Compare NodeCompareFunc //初始化时候指定Compare的值
}

func NewBinarySearchTree(compare NodeCompareFunc) *BinarySearchTree {
	tree := &BinarySearchTree{Compare: compare}
	tree._Checker = &BinaryTreeChecker{}
	return tree
}

func (a *BinarySearchTree) Search(v interface{}, nodes ...ITreeNode) ITreeNode {

	if a.Compare == nil {
		panic("compare is nil")
	}

	node := a.Root
	if len(nodes) > 0 {
		node = nodes[0]
	}

	for !a.IsNil(node) && node.GetKey() != v {
		if a.Compare(v, node.GetKey()) > 0 {
			node = node.GetLChild()
		} else {
			node = node.GetRChild()
		}
	}
	return node
}

//! min:在二叉搜索树中最小值的节点。
/*!
* \param node:从指定节点开始搜索（默认为树的根节点）
* \return : 二叉树种的最小节点
*
* 在二叉搜索树中搜索最小值的节点。其中可以指定从哪个节点开始搜索。若不指定搜索节点，则默认为树的根节点。
*
* 算法：由于二叉树的性质，搜索最小值很简单。从指定节点沿着左子节点一路向下遍历，最左下方的节点即为最小值节点
*
* 算法时间复杂度O(h)，空间复杂度O(1)。其中h为树的高度
 */
func (a *BinarySearchTree) Min(nodes ...ITreeNode) ITreeNode {
	node := a.Root
	if len(nodes) > 0 {
		node = nodes[0]
	}
	if a.IsNil(node) {
		panic(" SearchTree.Min nil")
	}
	for !a.IsNil(node.GetLChild()) {
		node = node.GetLChild()
	}
	return node
}

//! max:在二叉搜索树中最大值的节点。
/*!
* \param node:从指定节点开始搜索（默认为树的根节点）
* \return : 二叉树种的最大节点
*
* 在二叉搜索树中搜索最小值的节点。其中可以指定从哪个节点开始搜索。若不指定搜索节点，则默认为树的根节点
*
* 算法：由于二叉树的性质，搜索最大值很简单。从指定节点沿着右子节点一路向下遍历，最右下方的节点即为最小值节点
*
* 算法时间复杂度O(h)，空间复杂度O(1)。其中h为树的高度
 */
func (a *BinarySearchTree) Max(nodes ...ITreeNode) ITreeNode {
	node := a.Root
	if len(nodes) > 0 {
		node = nodes[0]
	}
	if a.IsNil(node) {
		panic(" BinarySearchTree.Max nullptr")
	}
	for !a.IsNil(node.GetRChild()) {
		node = node.GetRChild()
	}
	return node
}

//! successor:二叉搜索树指定节点的后继节点。最近一个比node大的节点
/*!
* \param node:要搜索后继的节点
* \return : 该节点的后继节点
*
* 给定二叉搜索树的某个节点，搜索其后继节点。所谓的某节点`node`的后继节点就是在二叉搜索树中，值大于等于`node`的所有节点中最小的那一个（排除它自身）。
*
* 一个节点`node`的后继有以下情况：
*
* - 如果`node`有右子节点，则以右子节点为根的子树中的最小值节点就是`node`的后继节点
* - 如果`node`没有右子节点，则查看父节点
*      - 若`node`是父节点的左子节点；则`node`的后继节点是`node`的父节点
*      - 若`node`是父节点的右子节点；则`node`设置为`node.Parent`，递归向直到`node`是它父亲的左子节点；此时`node`的后继节点是`node`的父节点
*
* 算法时间复杂度O(h)，空间复杂度O(1)。其中h为树的高度
 */
func (a *BinarySearchTree) Successor(nodes ...ITreeNode) ITreeNode {
	node := a.Root
	if len(nodes) > 0 {
		node = nodes[0]
	}
	if a.IsNil(node) {
		panic(" BinarySearchTree.Successor nullptr")
	}

	if !a.IsNil(node.GetRChild()) { //以右子节点为根的子树中的最小值节点就是`node`的后继节点
		return a.Min(node.GetRChild())
	}
	parent := node.GetParent()                   //取出父节点
	for !a.IsNil(parent) && IsRightChild(node) { //向上迭代，直到父节点为空或者node是父节点的左子节点，则迭代终止
		node = parent
		parent = parent.GetParent()
	}
	return parent
}

//! Predecesor:二叉搜索树指定节点的前驱。最近一个比node小的节点
/*!
* \param node:要搜索前驱的节点
* \return : 该节点的前驱节点
*
* 给定二叉搜索树的某个节点，搜索其前驱节点。所谓的某节点`node`的前驱节点就是在二叉搜索树中，值小于`node`的所有节点中最大的那一个。
*
* 一个节点`node`的前驱有以下情况：
*
* - 如果`node`有左子节点，则以左子节点为根的子树中的最大值节点就是`node`的前驱节点
* - 如果`node`没有左子节点，则查看父节点
*      - 若`node`是父节点的右子节点；则`node`的前驱节点是`node`的父节点
*      - 若`node`是父节点的左子节点；则`node`设置为`node.Parent`，递归向直到`node`是它父亲的右子节点；此时`node`的前驱节点是`node`的父节点
*
* 算法时间复杂度O(h)，空间复杂度O(1)。其中h为树的高度
 */
func (a *BinarySearchTree) Predecesor(nodes ...ITreeNode) ITreeNode {
	node := a.Root
	if len(nodes) > 0 {
		node = nodes[0]
	}
	if a.IsNil(node) {
		panic(" SearchTree.Predecesor nullptr")
	}
	if !a.IsNil(node.GetLChild()) { //以左子节点为根的子树中的最大值节点就是`node`的前驱节点
		return a.Max(node.GetLChild())
	}
	parent := node.GetParent()                  //取出父节点
	for !a.IsNil(parent) && IsLeftChild(node) { //向上迭代，直到父节点为空或者node是父节点的右子节点，则迭代终止
		node = parent
		parent = parent.GetParent()
	}
	return parent
}

//! insert:向二叉搜索树中插入节点。
/*!
* \param node:要插入的节点
* \return : void
*
* 给定新节点`node`，将该节点插入到二叉搜索树中。
*
* 算法：遍历二叉搜索树，若当前节点的值大于`node`的值，则向左侧遍历；若当前节点值小于`node`的值，则向右侧遍历。直到碰到`nullptr`则挂载该节点
*
* 算法时间复杂度O(h)，空间复杂度O(1)。其中h为树的高度
* 由于SearTree是有特性的，所以每个节点有固定的位置，这里需要找个这个位置，然后将新节点放置在该位置
 */
func (a *BinarySearchTree) Insert(node ITreeNode) {
	node.SetLChild(a.GetNil())
	node.SetRChild(a.GetNil())
	if a.Compare == nil {
		panic("compare is nil")
	}
	compare := a.Compare

	if a.IsNil(a.Root) {
		a.Root = node
		return
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
}

//! delete:从二叉搜索树中删除节点。
/*!
* \param node:要删除的节点
* \return : void
*
* 给定节点`node`，从二叉搜索树中删除它。如果`node`不在二叉搜索树中则抛出异常。
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
func (a *BinarySearchTree) Delete(node *BinaryTreeNode) error {
	if a.IsNil(node) {
		return errors.New("node removed is nullptr")
	}
	if a.Compare == nil {
		return errors.New("compare is nil")
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
		return errors.New("node removed must be in the tree")
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
	} else //node既有左子，又有右子
	{
		next_node := a.Successor(node)     //node的后继节点
		if next_node != node.GetRChild() { //node后继不是node的右子，则转换成node的后继是node的右子情形
			a.Transplant(next_node.GetRChild(), next_node, &a.Root) //将后继的右子剪切到后继的位置
			next_node.SetRChild(node.GetRChild())                   //后继的右子设为node的右子
			node.GetRChild().SetParent(next_node)                   //后继的右子设为node的右子
			node.SetRChild(next_node)                               //后继的右子设为node的右子
		}
		//node后继是node的右子
		a.Transplant(node.GetRChild(), node, &a.Root)
		next_node.SetLChild(node.GetLChild())
		node.GetLChild().SetParent(next_node)
	}

	return nil
}
