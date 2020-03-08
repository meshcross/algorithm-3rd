/*
 * @Description: 树结构，包含二叉树，二叉搜索树，红黑树
 * @Author: wangchengdg@gmail.com
 * @Date: 2020-02-14 20:03:01
 * @LastEditTime: 2020-03-08 17:24:45
 * @LastEditors:
 */
package TreeAlgorithm

import (
	"fmt"
	"testing"

	. "github.com/meshcross/algorithm-3rd/mesh/common"
)

/*           0               中序遍历: 7;3;8;1;4;0;5;2;6;前序遍历: 0;1;3;7;8;4;2;5;6;后序遍历: 7;8;3;4;1;5;6;2;0;
*          /   \
*         1     2
*        / \   /  \
*       3   4  5   6
*      / \
*     7   8
 */
func getBinaryTree() *BinaryTree {
	nodes := make([]*BinaryTreeNode, 9)

	for i := 0; i < 9; i++ {
		nodes[i] = &BinaryTreeNode{_Key: i}
	}
	nodes[1].SetParent(nodes[0])
	nodes[2].SetParent(nodes[0])
	nodes[0].SetLChild(nodes[1])
	nodes[0].SetRChild(nodes[2])

	nodes[3].SetParent(nodes[1])
	nodes[4].SetParent(nodes[1])
	nodes[1].SetLChild(nodes[3])
	nodes[1].SetRChild(nodes[4])

	nodes[5].SetParent(nodes[2])
	nodes[6].SetParent(nodes[2])
	nodes[2].SetLChild(nodes[5])
	nodes[2].SetRChild(nodes[6])

	nodes[7].SetParent(nodes[3])
	nodes[8].SetParent(nodes[3])
	nodes[3].SetLChild(nodes[7])
	nodes[3].SetRChild(nodes[8])

	tree := NewBinaryTree()
	tree.Root = nodes[0]

	return tree
}

func TestTreeOrder(t *testing.T) {
	str := ""
	action := func(node ITreeNode) {
		str = fmt.Sprintf("%s%d;", str, node.GetKey())
	}

	empty_str := ""
	empty_action := func(node ITreeNode) {
		empty_str = fmt.Sprintf("%s;", empty_str)
	}

	empty_tree := NewBinaryTree()
	tree := getBinaryTree()
	tree.PreOrder(tree.Root, action)
	empty_tree.PreOrder(empty_tree.Root, empty_action)
	EXPECT_EQ(str, "0;1;3;7;8;4;2;5;6;", t)
	EXPECT_EQ(empty_str, "", t)
	fmt.Println("tree.preorder:", str)
	fmt.Println("empty_tree.postorder:", empty_str)

	str = ""
	empty_str = ""
	tree.InOrder(tree.Root, action)
	empty_tree.InOrder(empty_tree.Root, empty_action)
	EXPECT_EQ(str, "7;3;8;1;4;0;5;2;6;", t)
	EXPECT_EQ(empty_str, "", t)
	fmt.Println("tree.inorder:", str)
	fmt.Println("empty_tree.postorder:", empty_str)

	str = ""
	empty_str = ""
	tree.PostOrder(tree.Root, action)
	empty_tree.PostOrder(empty_tree.Root, empty_action)
	EXPECT_EQ(str, "7;8;3;4;1;5;6;2;0;", t)
	EXPECT_EQ(empty_str, "", t)

	fmt.Println("tree.postorder:", str)
	fmt.Println("empty_tree.postorder:", empty_str)

}

func TestSearchTree_Search(t *testing.T) {
	NODE_NUM := 9
	_empty_tree := NewBinarySearchTree(NodeCompareFunc_IntLessThan)
	_normal_tree := getSearchTree()

	// ***  空树  *******
	result1 := _empty_tree.Search(0)
	fmt.Println("EXPECT_FALSE ", result1)
	//***   非空树 *******
	result2 := _normal_tree.Search(5)
	fmt.Println("ASSERT_TRUE ", result2)
	fmt.Println("EXPECT_EQ ", result2, _normal_tree.Root)

	result3 := _normal_tree.Search(NODE_NUM) //无效的值
	fmt.Println("EXPECT_FALSE ", result3)

	result4 := _normal_tree.Search(NODE_NUM - 1)
	fmt.Println("ASSERT_TRUE ", result4)
	fmt.Println("EXPECT_EQ ", result4.GetKey(), NODE_NUM-1)
}

func TestSearchTree_Insert(t *testing.T) {

	_empty_tree := NewBinarySearchTree(NodeCompareFunc_IntLessThan)
	_normal_tree := getSearchTree()
	// ****  空树  ********
	{
		node := NewBinaryTreeNode(1234)
		_empty_tree.Insert(node)
		fmt.Println("a-EXPECT_EQ ", _empty_tree.Root, node)
	}
	// *****  非空树 ************
	{

		node := NewBinaryTreeNode(1234)
		_normal_tree.Insert(node)
		_normal_tree.Print()
		{
			str := ""
			do_work := func(val ITreeNode) { str = fmt.Sprintf("%s%d;", str, val.GetKey()) }
			_normal_tree.InOrder(_normal_tree.Root, do_work)
			fmt.Println("b-EXPECT_EQ ", str, "0;1;2;3;4;5;6;7;8;1234;")
		}
		{
			str := ""
			do_work := func(val ITreeNode) { str = fmt.Sprintf("%s%d;", str, val.GetKey()) }
			_normal_tree.PreOrder(_normal_tree.Root, do_work)
			fmt.Println("c-EXPECT_EQ ", str, "5;3;1;0;2;4;7;6;8;1234;")
		}

		/*           5
		*          /   \
		*         3     7
		*        / \   /  \
		*       1   4 6    8
		*      / \           \
		*     0   2          1234
		 */

		node2 := NewBinaryTreeNode(10)
		_normal_tree.Insert(node2)
		_normal_tree.Print()
		{
			str := ""
			do_work := func(val ITreeNode) { str = fmt.Sprintf("%s%d;", str, val.GetKey()) }
			_normal_tree.InOrder(_normal_tree.Root, do_work)

			fmt.Println("d-EXPECT_EQ ", str, "0;1;2;3;4;5;6;7;8;10;1234;")
		}
		{
			str := ""
			do_work := func(val ITreeNode) { str = fmt.Sprintf("%s%d;", str, val.GetKey()) }
			_normal_tree.PreOrder(_normal_tree.Root, do_work)
			fmt.Println("e-EXPECT_EQ ", str, "5;3;1;0;2;4;7;6;8;1234;10;")
		}
	}

	/*           5
	*          /   \
	*         3     7
	*        / \   /  \
	*       1   4 6    8
	*      / \           \
	*     0   2          1234
	*                   /
	*                  10
	 */

}

func TestSearchTree_Delete(t *testing.T) {
	_empty_tree := NewBinarySearchTree(NodeCompareFunc_IntLessThan)
	_normal_tree := getSearchTree()
	{
		node := NewBinaryTreeNode(1234)
		err := _empty_tree.Delete(node)
		fmt.Println("EXPECT_THROW ", err)
	}
	// *****  非空树 ************
	{
		node := NewBinaryTreeNode(1234) //不是树的子节点
		err := _normal_tree.Delete(node)
		fmt.Println("EXPECT_THROW ", err)

		_normal_tree.Delete(ToBinaryTreeNode(_normal_tree.Root)) //删除一个拥有两个子节点的节点
		{
			str := ""
			do_work := func(val ITreeNode) { str = fmt.Sprintf("%s%d;", str, val.GetKey()) }
			_normal_tree.InOrder(_normal_tree.Root, do_work)
			fmt.Println("EXPECT_EQ", str, "0;1;2;3;4;6;7;8;")
		}
		{
			str := ""
			do_work := func(val ITreeNode) { str = fmt.Sprintf("%s%d;", str, val.GetKey()) }
			_normal_tree.PreOrder(_normal_tree.Root, do_work)
			fmt.Println("EXPECT_EQ", str, "6;3;1;0;2;4;7;8;")
		}
		/*           6
		 *          /   \
		 *         3     7
		 *        / \      \
		 *       1   4      8
		 *      / \
		 *     0   2
		 *
		 *
		 */

		pp_max := ToBinaryTreeNode(_normal_tree.Max().GetParent())
		_normal_tree.Delete(pp_max) //删除一个只有单子节点的节点
		{
			str := ""
			do_work := func(val ITreeNode) { str = fmt.Sprintf("%s%d;", str, val.GetKey()) }
			_normal_tree.InOrder(_normal_tree.Root, do_work)
			fmt.Println("EXPECT_EQ", str, "0;1;2;3;4;6;8;")
		}
		{
			str := ""
			do_work := func(val ITreeNode) { str = fmt.Sprintf("%s%d;", str, val.GetKey()) }
			_normal_tree.PreOrder(_normal_tree.Root, do_work)
			fmt.Println("EXPECT_EQ", str, "6;3;1;0;2;4;8;")
		}
		/*           6
		 *          /   \
		 *         3     8
		 *        / \
		 *       1   4
		 *      / \
		 *     0   2
		 *
		 *
		 */
		min_node := ToBinaryTreeNode(_normal_tree.Min(_normal_tree.Root))
		_normal_tree.Delete(min_node) //删除一个叶子节点
		{
			str := ""
			do_work := func(val ITreeNode) { str = fmt.Sprintf("%s%d;", str, val.GetKey()) }
			_normal_tree.InOrder(_normal_tree.Root, do_work)
			fmt.Println("EXPECT_EQ", str, "1;2;3;4;6;8;")
		}
		{
			str := ""
			do_work := func(val ITreeNode) { str = fmt.Sprintf("%s%d;", str, val.GetKey()) }
			_normal_tree.PreOrder(_normal_tree.Root, do_work)
			fmt.Println("EXPECT_EQ", str, "6;3;1;2;4;8;")
		}
		/*           6
		 *          /   \
		 *         3     8
		 *        / \
		 *       1   4
		 *        \
		 *         2
		 *
		 *
		 */
	}
}

func getSearchTree() *BinarySearchTree {
	NODE_NUM := 9
	nodes := make([]*BinaryTreeNode, NODE_NUM)

	for i := 0; i < NODE_NUM; i++ {
		nodes[i] = NewBinaryTreeNode(i)
	}
	nodes[3].SetParent(nodes[5])
	nodes[7].SetParent(nodes[5])
	nodes[5].SetLChild(nodes[3])
	nodes[5].SetRChild(nodes[7])

	nodes[1].SetParent(nodes[3])
	nodes[4].SetParent(nodes[3])
	nodes[3].SetLChild(nodes[1])
	nodes[3].SetRChild(nodes[4])

	nodes[6].SetParent(nodes[7])
	nodes[8].SetParent(nodes[7])
	nodes[7].SetLChild(nodes[6])
	nodes[7].SetRChild(nodes[8])

	nodes[0].SetParent(nodes[1])
	nodes[2].SetParent(nodes[1])
	nodes[1].SetLChild(nodes[0])
	nodes[1].SetRChild(nodes[2])

	tree := NewBinarySearchTree(NodeCompareFunc_IntLessThan)
	tree.Root = (nodes[5])
	return tree
}

func TestRedBlack_Insert(t *testing.T) {

	// NODE_NUM := 8

	node_1 := NewRedBlackTreeNode(1)
	node_2 := NewRedBlackTreeNode(2)
	node_5 := NewRedBlackTreeNode(5)
	node_7 := NewRedBlackTreeNode(7)
	node_8 := NewRedBlackTreeNode(8)
	node_11 := NewRedBlackTreeNode(11)
	node_14 := NewRedBlackTreeNode(14)
	node_15 := NewRedBlackTreeNode(15)
	nodes := []*RedBlackTreeNode{node_1, node_2, node_5, node_7, node_8, node_11, node_14, node_15}

	//11
	node_11.SetParent(nil)
	node_11.SetLChild(node_2)
	node_11.SetRChild(node_14)
	node_11.Color = COLOR_BLACK

	//2
	node_2.SetParent(node_11)
	node_2.SetLChild(node_1)
	node_2.SetRChild(node_7)
	node_2.Color = COLOR_RED

	//1
	node_1.SetParent(node_2)
	node_1.SetLChild(nil)
	node_1.SetRChild(nil)
	node_1.Color = COLOR_BLACK

	//7
	node_7.SetParent(node_2)
	node_7.SetLChild(node_5)
	node_7.SetRChild(node_8)
	node_7.Color = COLOR_BLACK

	//5
	node_5.SetParent(node_7)
	node_5.SetLChild(nil)
	node_5.SetRChild(nil)
	node_5.Color = COLOR_RED

	//8
	node_8.SetParent(node_7)
	node_8.SetLChild(nil)
	node_8.SetRChild(nil)
	node_8.Color = COLOR_RED

	//14
	node_14.SetParent(nodes[5])
	node_14.SetLChild(nil)
	node_14.SetRChild(node_15)
	node_14.Color = COLOR_BLACK

	//15
	node_15.SetParent(node_14)
	node_15.SetLChild(nil)
	node_15.SetRChild(nil)
	node_5.Color = COLOR_RED

	tree := NewRedBlackTree(NodeCompareFunc_IntLessThan)
	tree.Root = node_11

	for _, node := range nodes {
		if node.GetLChild() == nil {
			node.SetLChild(tree.GetNil())
		}
		if node.GetRChild() == nil {
			node.SetRChild(tree.GetNil())
		}
		if node.GetParent() == nil {
			node.SetParent(tree.GetNil())
		}
	}

	//tree.Print()
	newNode := NewRedBlackTreeNode(4)
	tree.Insert(newNode)
	tree.Print()

}
