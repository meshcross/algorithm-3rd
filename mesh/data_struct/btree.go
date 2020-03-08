/*
 * @Description: 第18章 B树  (Balanced Tree)
			B树是为磁盘或其他直接存取的辅助存储设备而设计的一种多路平衡搜索树。B树类似于红黑树，单它在降低磁盘I/O操作数方面要更好一些。
			许多数据库系统使用B树或者B树的变种来存储信息。
			B树与红黑树最大的不同在于，B树的节点可以有很多个孩子，数个到数千个。所以对于同样的节点数量，B树的深度要比红黑树小很多。

			B树具有以下性质：
			1、每个节点x存储n个关键字，x.key1<=x.key2<=x.key3<=...<=x.keyn；如果x为叶节点，则x.leaf为true，否则为false
			2、每个叶节点x还包含x.n+1个指向其孩子的指针x.c1,x.c2,...,x.c(x+1),叶节点没有孩子
			3、关键字x.key(i)对存储在各子树中的关键字范围加以分隔：如果ki为任意一个存储在以x.c(i)为根的子树中的关键字，那么
			  k1<=x.key1<=k2<=x.key2<=...<=x.keyn<=k_(n+1)
			4、每个叶节点具有相同的深度，即树的高度h
			5、每个节点锁包含的关键字个数有上界和下界。用一个被称为B树的最小度数(minimum degree)的固定整数t>=2来表示这些界：
				a、除了根节点以外的每个节点必须至少有t-1个关键字。因此，除了根节点以外的每个内部节点至少有t个孩子。如果树非空，根节点至少有一个关键字。
				b、每个节点至多可以包含2t-1个关键词。因此，一个内部节点至多可以有2t个孩子。当一个节点恰好有2t-1个关键词时，称该节点是满的。


 * @Author: wangchengdg@gmail.com
 * @Date: 2020-03-04 23:01:26
 * @LastEditTime: 2020-03-08 17:20:40
 * @LastEditors:
*/

package DataStruct

import (
	"errors"
	"fmt"

	. "github.com/meshcross/algorithm-3rd/mesh/common"
)

type BTreeNode struct {
	Leaf      bool            //是否为叶节点
	_Entrys   []*Pair         //B树中，每个节点是存放了时间的
	_Children []*BTreeNode    //有n个孩子节点  n>2
	_Compare  NodeCompareFunc //比较函数

	_N            int
	_EntrySize    int
	_ChildrenSize int
}

/**
 * @description: 此处采用一次性分配好数组大小，不再使用动态调整的方式，因为Entrys和Children数组中存储的是指针，所以大小可控
			但是如果BTree节点中存储的是简单数据类型，并且BTree节点长期处于严重不饱和状态，则可能会造成较大浪费
			出现这种情况可以根据业务对BTreeNode做进一步的调整

			需要特别注意的是，children数量比entry数量多1
			当前节点的entry是排序好的，并且各子child中各个数据的顺序满足和entry中的数据高度相关，需要仔细分析红黑树的性质3
			entry[i]其实是一个将各个child分开的分隔符，这一点性质很重要

			一颗典型的红黑树如下：
						    30
					/		          \
				18/22/28	         50/70
			/       |       \        /     \
		5/10/15  19/20/21  24/26   35/40   55/60

		1/10/15<18, 19/20/21属于(18,22),24/26属于(22,28),还可以增加>28的节点
 * @param
 * @return:
*/
func NewBTreeNode(n int, compare NodeCompareFunc) *BTreeNode {
	endtrys := make([]*Pair, n)
	children := make([]*BTreeNode, n+1)
	return &BTreeNode{_N: n, _Compare: compare, _Children: children, _Entrys: endtrys}
}

func (node *BTreeNode) IsFull() bool {
	return node._N == node._EntrySize
}

func (node *BTreeNode) AddEntry(e *Pair) {
	if node._EntrySize < node._N {
		node._Entrys[node._EntrySize] = e
		node._EntrySize++
	}
}
func (node *BTreeNode) Size() int {
	return node._EntrySize
}

func (node *BTreeNode) RemoveEntry(index int) *Pair {
	if index >= 0 && index < node._EntrySize {
		ret := node._Entrys[index]
		size := node._EntrySize
		end := index
		//边界情况要处理，end+1的索引可能无效，所以end<size-1
		for end < size-1 {
			node._Entrys[end] = node._Entrys[end+1]
			end++
		}
		node._Entrys[size-1] = nil
		node._EntrySize--
		return ret
	}
	return nil
}
func (node *BTreeNode) EntryAt(index int) *Pair {
	if index >= 0 && index < node._EntrySize {
		return node._Entrys[index]
	}
	return nil
}

/**
* 在节点中查找给定的键。
* 如果节点中存在给定的键，则返回给定的键在节点中的索引和给定的键关联的值；
* 如果不存在，则返回给定的键应该插入的位置，该键的关联值为nil。
*
* @param key - 给定的键值
* @return - 查找结果
 */
func (node *BTreeNode) SearchKey(key interface{}) (int, *Pair) {
	low := 0
	high := node._EntrySize - 1
	mid := 0
	for low <= high {
		mid = (low + high) / 2
		entry := node._Entrys[mid]
		if node._Compare(entry.First, key) == 0 {
			break
		} else if node._Compare(entry.First, key) > 0 {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	//如果找到，则为相应元素所在位置的索引，如果没有找到，则为插入位置的索引
	index := -1
	var value *Pair = nil
	// 说明查找成功
	if low <= high {
		// index表示元素所在的位置
		index = mid
		value = node._Entrys[index]
	} else {
		// index表示元素应该插入的位置
		index = low
	}
	return index, value
}

func (node *BTreeNode) PutEntry(e *Pair) *Pair {
	_, entry := node.SearchKey(e.First)
	if node != nil {
		// oldValue := entry.Second
		entry.Second = e.Second

		return entry
	} else {
		node.InsertEntry(e)
		return nil
	}
}
func (node *BTreeNode) InsertEntry(e *Pair) bool {
	index, entry := node.SearchKey(e.First)
	if entry != nil {
		return false
	} else {
		node.InsertEntryAt(e, index)
		return true
	}
}
func (node *BTreeNode) InsertEntryAt(e *Pair, index int) {

	if node._EntrySize < node._N {
		end := node._EntrySize
		for end > index {
			node._Entrys[end] = node._Entrys[end-1]
			end--
		}
		node._Entrys[index] = e
		node._EntrySize++
	}
}
func (node *BTreeNode) ChildAt(index int) *BTreeNode {
	size := node._ChildrenSize
	if index >= 0 && index <= size {
		return node._Children[index]
	}
	return nil
}

func (node *BTreeNode) AddChild(c *BTreeNode) {
	if node._ChildrenSize <= node._N {
		node._Children[node._ChildrenSize] = c
		node._ChildrenSize++
	}
}
func (node *BTreeNode) ChildrenSize() int {
	return node._ChildrenSize
}

func (node *BTreeNode) RemoveChild(index int) {
	if index >= 0 && index < node._ChildrenSize {
		size := node._ChildrenSize
		end := index
		for end < size-1 {
			node._Children[end] = node._Children[end+1]
			end++
		}
		node._Children[size-1] = nil
		node._ChildrenSize--
	}
}

func (node *BTreeNode) InsertChildAt(newNode *BTreeNode, index int) {

	if node._ChildrenSize <= node._N {
		end := node._ChildrenSize - 1
		//通通往后移，把index位置空出来，然后设定问newNode
		for end > index {
			node._Children[end] = node._Children[end-1]
			end--
		}
		node._Children[index] = newNode
		node._ChildrenSize++
	}
}

type BTree struct {
	Root *BTreeNode

	_minKeySize int //非根节点中最小的键值数
	_maxKeySize int //非根节点中最大的键值数
	_t          int //B树的最小度数 minmum degree

	_Compare NodeCompareFunc
}

func NewBTree(t int, compare NodeCompareFunc) *BTree {
	minSize := t - 1
	maxSize := 2*t - 1
	tree := &BTree{_t: t, _maxKeySize: maxSize, _minKeySize: minSize, _Compare: compare}
	root := NewBTreeNode(maxSize, compare)
	root.Leaf = true
	tree.Root = root
	return tree
}

/**
 * @description: 节点是否已满
 * @param {type}
 * @return:
 */
func (t *BTree) NodeIsFull(node *BTreeNode) bool {
	return node.Size() >= t._maxKeySize
}

/**
 * @description: 节点数量是否快要少于最低值限制
 * @param {type}
 * @return:
 */
func (t *BTree) NodeIsMin(node *BTreeNode) bool {
	return node.Size() <= t._minKeySize
}
func (t *BTree) Search(node *BTreeNode, key interface{}) *Pair {
	index, retNode := node.SearchKey(key)
	if retNode != nil {
		return retNode
	} else {
		if node.Leaf {
			return nil
		} else {
			cNode := node.ChildAt(index)
			return t.Search(cNode, key)
		}
	}
	// return nil
}

/**
 * @description: 分裂一个满的子节点
 * @param
 * @return:
 */
func (bt *BTree) SplitNode(parentNode, childNode *BTreeNode, index int) error {

	if parentNode == nil || childNode == nil {
		return errors.New("parent or child must not be nil!")
	}
	if childNode.Size() < bt._maxKeySize {
		return errors.New(" child must be full!")
	}
	maxKeySize := bt._maxKeySize
	minKeySize := bt._minKeySize
	t := bt._t

	siblingNode := NewBTreeNode(maxKeySize, bt._Compare)

	siblingNode.Leaf = childNode.Leaf
	// 将满子节点中索引为[t, 2t - 2]的(t - 1)个项插入新的节点中
	for i := 0; i < minKeySize; i++ {
		siblingNode.AddEntry(childNode.EntryAt(t + i))
	}
	// 提取满子节点中的中间项，其索引为(t - 1)
	entry := childNode.EntryAt(t - 1)
	// 删除满子节点中索引为[t - 1, 2t - 2]的t个项
	for i := maxKeySize - 1; i >= t-1; i-- {
		childNode.RemoveEntry(i)
	}
	if !childNode.Leaf { // 如果满子节点不是叶节点，则还需要处理其子节点

		// 将满子节点中索引为[t, 2t - 1]的t个子节点插入新的节点中
		for i := 0; i < minKeySize+1; i++ { //++i
			siblingNode.AddChild(childNode.ChildAt(t + i))
		}
		// 删除满子节点中索引为[t, 2t - 1]的t个子节点
		for i := maxKeySize; i >= t; i-- { //--i
			childNode.RemoveChild(i)
		}
	}
	// 将entry插入父节点
	parentNode.InsertEntryAt(entry, index)
	// 将新节点插入父节点
	parentNode.InsertChildAt(siblingNode, index+1)

	return nil
}

/**
* 在一个非满节点中插入给定的项。
*
* @param node  - 非满节点
* @param entry - 给定的项
* @return true，如果B树中不存在给定的项，否则false
 */
func (bt *BTree) InsertNotFull(node *BTreeNode, entry *Pair) bool {
	//必须是不满的节点才能插入
	if node.IsFull() {
		return false
	}

	if node.Leaf { // 如果是叶子节点，直接插入
		return node.InsertEntry(entry)
	} else {
		/* 找到entry在给定节点应该插入的位置
		 */
		index, retEntry := node.SearchKey(entry.First)
		// 如果存在，则直接返回失败
		if retEntry != nil {
			return false
		}
		childNode := node.ChildAt(index)
		if childNode.IsFull() { // 如果子节点是满节点

			// 则先分裂
			bt.SplitNode(node, childNode, index)
			/* 如果给定entry的键大于分裂之后新生成项的键，则需要插入该新项的右边，
			* 否则左边。
			 */
			if bt._Compare(entry.First, node.EntryAt(index).First) > 0 {
				childNode = node.ChildAt(index + 1)
			}
		}
		return bt.InsertNotFull(childNode, entry)
	}
}

func (bt *BTree) Insert(e *Pair) bool {
	root := bt.Root
	if root.IsFull() { // 如果根节点满了，则B树长高

		newRoot := NewBTreeNode(bt._maxKeySize, bt._Compare)
		newRoot.Leaf = false
		newRoot.AddChild(root)
		bt.SplitNode(newRoot, root, 0)
		bt.Root = newRoot
	}
	return bt.InsertNotFull(bt.Root, e)
}

/**
* 如果存在给定的键，则更新键关联的值，
* 否则插入给定的项。
*
* @param node  - 非满节点
* @param entry - 给定的项
* @return true，如果B树中不存在给定的项，否则false
 */
func (bt *BTree) PutNotFull(node *BTreeNode, entry *Pair) *Pair {
	if node.IsFull() {
		return nil
	}

	if node.Leaf { // 如果是叶子节点，直接插入
		return node.PutEntry(entry)
	} else {
		/* 找到entry在给定节点应该插入的位置，那么entry应该插入
		* 该位置对应的子树中
		 */
		index, retEntry := node.SearchKey(entry.First)
		// 如果存在，则更新
		if retEntry != nil {
			return node.PutEntry(entry)
		}
		childNode := node.ChildAt(index)
		if childNode.IsFull() { // 如果子节点是满节点

			// 则先分裂
			bt.SplitNode(node, childNode, index)
			/* 如果给定entry的键大于分裂之后新生成项的键，则需要插入该新项的右边，
			* 否则左边。
			 */
			if bt._Compare(entry.First, node.EntryAt(index).First) > 0 {
				childNode = node.ChildAt(index + 1)
			}
		}
		return bt.PutNotFull(childNode, entry)
	}
}

/**
* 如果B树中存在给定的键，则更新值。
* 否则插入。
*
* @param key   - 键
* @param value - 值
* @return 如果B树中存在给定的键，则返回之前的值，否则null
 */
func (bt *BTree) Put(e *Pair) *Pair {
	root := bt.Root
	if root.IsFull() { // 如果根节点满了，则B树长高

		newRoot := NewBTreeNode(bt._maxKeySize, bt._Compare)
		newRoot.Leaf = false
		newRoot.AddChild(root)
		bt.SplitNode(newRoot, root, 0)
		bt.Root = newRoot
	}
	return bt.PutNotFull(root, e)
}

/**
* 从B树中删除一个与给定键关联的项。
*
* @param key - 给定的键
* @return 如果B树中存在给定键关联的项，则返回删除的项，否则null
 */
func (bt *BTree) Delete(key interface{}) (*Pair, error) {
	return bt.DeleteFromNode(bt.Root, key)
}

/**
* 从以给定<code>node</code>为根的子树中删除与给定键关联的项。
* <p/>
* 删除的实现思想请参考《算法导论》第二版的第18章。
*
* @param node - 给定的节点
* @param key  - 给定的键
* @return 如果B树中存在给定键关联的项，则返回删除的项，否则null
 */
func (bt *BTree) DeleteFromNode(node *BTreeNode, key interface{}) (*Pair, error) {
	// 该过程需要保证，对非根节点执行删除操作时，其关键字个数至少为t。
	if node != bt.Root {
		if node.Size() < bt._t {
			return nil, errors.New("node entrys is min!")
		}
	}

	index, retEntry := node.SearchKey(key)
	/*
	* 因为这是查找成功的情况，0 <= index <= (node.Size() - 1)，
	* 因此(index + 1)不会溢出。
	 */
	if retEntry != nil {
		// 1.如果关键字在节点node中，并且是叶节点，则直接删除。
		if node.Leaf {
			return node.RemoveEntry(index), nil
		} else {
			// 2.a 如果节点node中前于key的子节点包含至少t个项
			leftChildNode := node.ChildAt(index)
			//如果左节点entry数未到底线
			if !bt.NodeIsMin(leftChildNode) {
				// 使用leftChildNode中的最后一个项代替node中需要删除的项
				node.RemoveEntry(index)
				node.InsertEntryAt(leftChildNode.EntryAt(leftChildNode.Size()-1), index)
				// 递归删除左子节点中的最后一个项
				return bt.DeleteFromNode(leftChildNode, leftChildNode.EntryAt(leftChildNode.Size()-1))
			} else {

				rightChildNode := node.ChildAt(index + 1)
				// 2.b 如果左节点entry数到底线了，右节点还没有到
				if !bt.NodeIsMin(rightChildNode) {
					// 使用rightChildNode中的第一个项代替node中需要删除的项
					node.RemoveEntry(index)
					node.InsertEntryAt(rightChildNode.EntryAt(0), index)
					// 递归删除右子节点中的第一个项
					return bt.DeleteFromNode(rightChildNode, rightChildNode.EntryAt(0))
				} else { // 2.c leftChild和rightChild节点都到底线了，这里就涉及到节点合并的问题
					//合并流程为：node(作为parent)先借一个节点(index下标)给leftChild,然后leftChild合并掉rightChild

					// node节点的index下标的entry添加到leftChlid
					leftChildNode.AddEntry(node.EntryAt(index))
					//node自身删除该entry
					node.RemoveEntry(index)
					//删除rightChild，并且把rightChild的数据移动到leftChild中
					node.RemoveChild(index + 1)
					for i := 0; i < rightChildNode.Size(); i++ {
						leftChildNode.AddEntry(rightChildNode.EntryAt(i))
					}
					// 将rightChildNode中的子节点合并进leftChildNode，如果有的话
					if !rightChildNode.Leaf {
						for i := 0; i <= rightChildNode.Size(); i++ {
							leftChildNode.AddChild(rightChildNode.ChildAt(i))
						}
					}
					return bt.DeleteFromNode(leftChildNode, key)
				}
			}
		}
	} else {
		/*
		* 因为这是查找失败的情况，0 <= index <= node.Size()，
		* 因此(index + 1)会溢出。
		 */
		if node.Leaf { // 如果关键字不在节点node中，并且是叶节点，则什么都不做，因为该关键字不在该B树中

			// logger.infoerr:=errors.New("The key  isn not in the BTree.")
			return nil, nil
		}
		//此处根据index取child，主要依据是红黑树的性质3
		childNode := node.ChildAt(index)
		if !bt.NodeIsMin(childNode) { // // 如果子节点有不少于t个项，则递归删除
			return bt.DeleteFromNode(childNode, key)
		} else { // 3
			//如果子节点的size==t-1，删除后不能满足红黑树的size>=t-1的性质，则需要如下处理:
			//找左右兄弟节点，如果左右兄弟节点的size>t-1，则表示兄弟节点有富余，于是可以借一个节点

			// 先查找右边的兄弟节点
			var siblingNode *BTreeNode = nil
			siblingIndex := -1
			if index < node.Size() { // 存在右兄弟节点

				if !bt.NodeIsMin(node.ChildAt(index + 1)) {
					siblingNode = node.ChildAt(index + 1)
					siblingIndex = index + 1
				}
			}
			// 如果右边的兄弟节点不符合条件，则试试左边的兄弟节点
			if siblingNode == nil {
				if index > 0 { // 存在左兄弟节点

					if !bt.NodeIsMin(node.ChildAt(index - 1)) {
						siblingNode = node.ChildAt(index - 1)
						siblingIndex = index - 1
					}
				}
			}
			// 3.a 有一个相邻兄弟节点至少包含t个项
			if siblingNode != nil {
				if siblingIndex < index { // 左兄弟节点满足条件
					//疑问：如果node==root，并且root只有一个entry怎么办？按照红黑树的性质，应该不会出现这种情况？

					//先找父节点借一个
					childNode.InsertEntryAt(node.EntryAt(siblingIndex), 0)
					//借走了，需要从父节点删除
					node.RemoveEntry(siblingIndex)
					//从兄弟节点划拉一个给父节点，补上数量，所以父节点上数量是持平的
					node.InsertEntryAt(siblingNode.EntryAt(siblingNode.Size()-1), siblingIndex)
					//给了父节点一个，需要从兄弟节点删除这个entry
					siblingNode.RemoveEntry(siblingNode.Size() - 1)

				} else { // 右兄弟节点满足条件

					//childNode.InsertEntryAt(node.EntryAt(index), childNode.Size()-1)
					childNode.InsertEntryAt(node.EntryAt(index), childNode.Size()) //应该防在在最后
					node.RemoveEntry(index)
					node.InsertEntryAt(siblingNode.EntryAt(0), index)
					siblingNode.RemoveEntry(0)

				}
				return bt.DeleteFromNode(childNode, key)
			} else { // 3.b 如果其相邻左右节点都包含t-1个项，即兄弟节点都不富裕，没有可供划拉的

				if index < node.Size() { // 存在右兄弟，直接在后面追加，进行节点合并

					rightSiblingNode := node.ChildAt(index + 1)
					childNode.AddEntry(node.EntryAt(index))
					node.RemoveEntry(index)
					node.RemoveChild(index + 1)
					for i := 0; i < rightSiblingNode.Size(); i++ {
						childNode.AddEntry(rightSiblingNode.EntryAt(i))
					}
					if !rightSiblingNode.Leaf {
						for i := 0; i <= rightSiblingNode.Size(); i++ {
							childNode.AddChild(rightSiblingNode.ChildAt(i))
						}
					}
				} else { // 存在左节点，在前面插入，进行节点合并

					leftSiblingNode := node.ChildAt(index - 1)
					childNode.InsertEntryAt(node.EntryAt(index-1), 0)
					node.RemoveEntry(index - 1)
					node.RemoveChild(index - 1)
					for i := leftSiblingNode.Size() - 1; i >= 0; i-- {
						childNode.InsertEntryAt(leftSiblingNode.EntryAt(i), 0)
					}
					if !leftSiblingNode.Leaf {
						for i := leftSiblingNode.Size(); i >= 0; i-- {
							childNode.InsertChildAt(leftSiblingNode.ChildAt(i), 0)
						}
					}
				}
				// 如果node是root并且node不包含任何项了
				if node == bt.Root && node.Size() == 0 {
					bt.Root = childNode
				}
				return bt.DeleteFromNode(childNode, key)
			}
		}
	}
}

func (bt *BTree) Print() {

	queue := NewQueue()
	queue.Push(bt.Root)
	for !queue.Empty() {
		item, _ := queue.Pop()

		node := item.(*BTreeNode)
		for i := 0; i < node.Size(); i++ {
			ety := node.EntryAt(i)
			if ety != nil {
				fmt.Printf("%d ", ety.Second)
			} else {
				fmt.Printf("empty node")
			}
		}
		fmt.Println()
		if !node.Leaf {
			for i := 0; i < node.ChildrenSize(); i++ {
				queue.Push(node.ChildAt(i))
			}
		}
	}
}
