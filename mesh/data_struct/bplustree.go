/*
 * @Description: B+树

		B+树是B树的一种变形形式，B+树上的叶子结点存储关键字以及相应记录的地址，叶子结点以上各层作为索引使用。一棵m阶的B+树定义如下:
		(1)每个结点至多有m个子女；
		(2)除根结点外，每个结点至少有[m/2]个子女，根结点至少有两个子女；
		(3)有k个子女的结点必有k个关键字。

		B+树的查找与B树不同，当索引部分某个结点的关键字与所查的关键字相等时，并不停止查找，应继续沿着这个关键字左边的指针向下，一直查到该关键字所在的叶子结点为止。

		B+树是B树的一种变形，比B树具有更广泛的应用，m阶 B+树有如下特征:
		(1)每个结点的关键字个数与孩子个数相等，所有非最下层的内层结点的关键字是对应子树上的最大关键字，最下层内部结点包含了全部关键字。
		(2)除根结点以外，每个内部结点有 到m个孩子。
		(3)所有叶结点在树结构的同一层，并且不含任何信息(可看成是外部结点或查找失败的结点)，因此，树结构总是树高平衡的。

		5阶B+树每个非叶子节点有2-4个key

		跟B树相比，B+树需要转变的一点思路是：B+树中的非叶节点，只是存储了一些范围信息，方便查找到最终的叶节点，
		进行节点拆分与合并的时候，只是把操作后的状态更新到父节点存储，并没有到父节点拆借的概念

 * @Author: wangchengdg@gmail.com
 * @Date: 2020-03-07 12:45:31
 * @LastEditTime: 2020-03-08 16:28:27
 * @LastEditors:
*/
package DataStruct

import (
	"errors"
	"fmt"

	. "github.com/meshcross/algorithm-3rd/mesh"
	. "github.com/meshcross/algorithm-3rd/mesh/common"
)

/**
 * @description: B+树的中间节点
 * keys最大个数为_N-1，children的最大个数为_N
 * 比如5阶的B+树，每个节点至多有4个key，至少2个key，至多5个子节点
 */
type BPlusTreeNode struct {
	_Order        int            //多少阶的B+树
	_Parent       *BPlusTreeNode //父节点
	_Keys         []interface{}
	_KeySize      int
	_Children     []*BPlusTreeNode
	_Compare      NodeCompareFunc
	_ChildrenSize int  //当前存放了多少child
	Leaf          bool //是否为叶节点

	_Datas []*PairAny //真正存储数据的地方

	_minCount int
	_maxCount int

	Next *BPlusTreeNode //横向的下一个节点，用于链表遍历,只有叶子节点需要维护Next关系;只有叶节点拆分与合并的时候才需要改变该属性
}

//正好最大
func (a *BPlusTreeNode) CountIsMax() bool {
	return a._KeySize == a._maxCount
}
func (a *BPlusTreeNode) Datas() []*PairAny {
	return a._Datas
}
func (a *BPlusTreeNode) Keys() []interface{} {
	return a._Keys
}

//只要slice内部没有发生resize，都能对原始数据进行修改
func (a *BPlusTreeNode) Children() []*BPlusTreeNode {
	return a._Children
}
func (a *BPlusTreeNode) Destroy() {
	a._Parent = nil
	a._Children = nil
	a._Datas = nil
	a._Keys = nil
	a._Compare = nil
}

//超过最大
func (a *BPlusTreeNode) CountIsMaxX() bool {
	return a._KeySize > a._maxCount
}
func (a *BPlusTreeNode) CountIsMin() bool {
	return a._KeySize == a._minCount
}
func (a *BPlusTreeNode) CountIsMinX() bool {
	return a._KeySize < a._minCount
}

func (a *BPlusTreeNode) RemoveChildAt(index int) {

	if index >= 0 && index < a._ChildrenSize {
		for i := index; i < a._ChildrenSize-1; i++ {
			a._Children[i] = a._Children[i+1]
		}
		a._Children[a._ChildrenSize-1] = nil
		a._ChildrenSize--
	}
}

/**
 * @description: removekey的时候会同时删除相应的data
 * @param {type}
 * @return:
 */
func (a *BPlusTreeNode) RemoveKeyAt(index int) {
	if index >= 0 && index < a._KeySize {
		for i := index; i < a._KeySize-1; i++ {
			a._Keys[i] = a._Keys[i+1]
			a._Datas[i] = a._Datas[i+1]
		}
		a._Keys[a._KeySize-1] = nil
		a._Datas[a._KeySize-1] = nil
		a._KeySize--
	}
}
func (a *BPlusTreeNode) ChildAt(v int) *BPlusTreeNode {
	if v >= 0 && v <= a._ChildrenSize {
		return a._Children[v]
	}
	return nil
}
func (a *BPlusTreeNode) ChildIndex(v *BPlusTreeNode) int {
	for i := 0; i < a._ChildrenSize; i++ {
		if a._Children[i] == v {
			return i
		}
	}
	return -1
}
func (a *BPlusTreeNode) GetParent() *BPlusTreeNode {
	return a._Parent
}
func (a *BPlusTreeNode) SetParent(v *BPlusTreeNode) {
	a._Parent = v
}
func (a *BPlusTreeNode) KeySize() int {
	return a._KeySize
}
func (a *BPlusTreeNode) ChangeKeySize(dx int) int {
	a._KeySize += dx
	return a._KeySize
}
func (a *BPlusTreeNode) ChildrenSize() int {
	return a._ChildrenSize
}
func (a *BPlusTreeNode) ChangeChildrenSize(dx int) int {
	a._ChildrenSize += dx
	return a._ChildrenSize
}
func (a *BPlusTreeNode) DataAt(index int) *PairAny {

	if index >= 0 && index < a._KeySize {
		return a._Datas[index]
	}
	return nil
}

func (a *BPlusTreeNode) KeyIndex(key interface{}) int {

	for i := 0; i < a._KeySize; i++ {
		if a._Compare(a._Keys[i], key) == 0 {
			return i
		}
	}
	return -1
}

/**
 * @description: 千万要注意，_KeySize可能>_maxCount，_maxCount只是合法的最大值，当_keySize>_maxCount的时候表明要进行拆分
 * @param {type}
 * @return:
 */
func (a *BPlusTreeNode) KeyAt(index int) interface{} {

	if index >= 0 && index < a._KeySize {
		return a._Keys[index]
	}
	return nil
}
func (a *BPlusTreeNode) Find(key interface{}) (int, *PairAny) {
	//不是叶节点，则只能返回下一个查找子节点的下标
	if !a.Leaf {
		i := 0
		c := a._KeySize
		for i < c {
			if a._Compare(key, a._Keys[i]) > 0 {
				break
			}
			i++
		}
		return i, nil
	} else {
		low := 0
		high := a._KeySize - 1
		mid := 0
		for low <= high {
			mid = (low + high) / 2
			iKey := a._Keys[mid]
			if a._Compare(iKey, key) == 0 {
				break
			} else if a._Compare(iKey, key) > 0 {
				high = mid - 1
			} else {
				low = mid + 1
			}
		}

		// 说明查找成功
		if low <= high {
			// index表示元素所在的位置
			index := mid
			if a.Leaf {
				return 0, a._Datas[index]
			} else {
				child := a.ChildAt(index)
				return child.Find(key)
			}
		}
	}
	return -1, nil
}

func (a *BPlusTreeNode) InsertKey(key interface{}) {
	i := 0
	c := a._KeySize
	for i < c {
		if a._Compare(key, a._Keys[i]) > 0 {
			break
		}
		i++
	}

	for x := c; x > i; x-- {
		a._Keys[x] = a._Keys[x-1]
		a._Datas[x] = a._Datas[x-1]
	}
	a._Keys[i] = key
	a._KeySize++
}

func (a *BPlusTreeNode) InsertChild(v *BPlusTreeNode) {
	i := 0
	c := a._ChildrenSize
	key := v.KeyAt(0)
	for i < c {
		child := a.ChildAt(i)
		if a._Compare(key, child.KeyAt(0)) > 0 {
			break
		}
		i++
	}

	for x := c; x > i; x-- {
		a._Children[x] = a._Children[x-1]
	}
	a._Children[i] = v
	v.SetParent(a)
	a._ChildrenSize++
}
func (a *BPlusTreeNode) InsertChildAt(v *BPlusTreeNode, index int) {
	// if a.CountIsMax() {
	// 	return
	// }
	if index >= 0 && index <= a._ChildrenSize {
		for i := a._ChildrenSize + 1; i > index; i-- {
			a._Children[i] = a._Children[i-1]
		}
		a._Children[index] = v
		v.SetParent(a)
		a._ChildrenSize++
	}
}

func (a *BPlusTreeNode) GetLeft() (*BPlusTreeNode, error) {
	if a.Leaf {
		return a, nil
	} else {
		if a._ChildrenSize > 0 {
			return a._Children[0].GetLeft()
		}
	}
	return nil, errors.New("GetLeft:node has no child!")

}

func NewBPlusTreeNode(order int, compare NodeCompareFunc) *BPlusTreeNode {
	countMin := order / 2
	countMax := order - 1
	keys := make([]interface{}, countMax+1)
	datas := make([]*PairAny, countMax+1)
	children := make([]*BPlusTreeNode, order+1)
	return &BPlusTreeNode{_Order: order, _Datas: datas, _maxCount: countMax, _minCount: countMin, _Keys: keys, _Children: children, _Compare: compare}
}

/**
B+树
**/
type BPlusTree struct {
	_Compare   NodeCompareFunc //比较函数
	_Order     int             //B+树的阶
	_MaxKeyNum int             //每一个非叶子节点至少有[m/2]个key，最多有m-1个key
	_MinKeyNum int
	Root       *BPlusTreeNode //根节点
	_LeafHead  *BPlusTreeNode //叶节点实际上组成了一张链表，_Left即为表头
}

func NewBPlusTree(order int, compare NodeCompareFunc) *BPlusTree {
	maxNum := order - 1
	minNum := order / 2
	root := NewBPlusTreeNode(order, compare)
	root.Leaf = true
	return &BPlusTree{_Compare: compare, _MaxKeyNum: maxNum, _MinKeyNum: minNum, Root: root, _Order: order}
}

/**
 * @description: B+树的叶节点其实有一个链表连接起来，这里获取链表的表头
 * @param
 * @return: 第一个节点
 */
func (tree *BPlusTree) updateLeafHead() {
	tree._LeafHead, _ = tree.Root.GetLeft()
}

/**
 * @description: 所有叶节点可以通过链表的方式访问，该函数返回链表头
 * @return: 链表头
 */
func (tree *BPlusTree) GetLeafHead() *BPlusTreeNode {
	return tree._LeafHead
}

/**
 * @description: 根据key找到存入B+树中的数据
 * @param key 根据key进行检索
 * @return: 查到的数据
 */
func (tree *BPlusTree) Find(key interface{}) *PairAny {
	return tree.findFrom(tree.Root, key)
}

func (tree *BPlusTree) findFrom(node *BPlusTreeNode, key interface{}) *PairAny {

	index, ret := node.Find(key)
	if node.Leaf {
		return ret
	} else if index >= 0 {
		child := node.ChildAt(index)
		return tree.findFrom(child, key)
	}
	return nil
}

/**
 * @description: 数据插入B+树
 * @param 需要插入的数据，其中e.First是Key
 * @return:
 */
func (tree *BPlusTree) Insert(e *PairAny) error {
	if e == nil || e.First == nil {
		return errors.New("insert param is nil!")
	}
	newRoot, err := tree.insertTo(tree.Root, e)
	if err != nil {
		return err
	}
	if newRoot != nil {
		tree.Root = newRoot
	}
	tree.updateLeafHead()
	return nil
}

/**
 * @description: 只有分裂变换了根节点的时候，才会返回新的节点
 * @param node : 向结点node中插入entry
 * @param e : 待插入的数据
 * @return: 是否发生了根节点变更，如果变更了，则返回新的根节点，没有变更则返回nil
 */
func (tree *BPlusTree) insertTo(node *BPlusTreeNode, e *PairAny) (*BPlusTreeNode, error) {
	if e == nil || e.First == nil {
		return nil, errors.New("insertTo param entry is nil")
	}

	i := 0
	c := node.KeySize()
	for i < c {
		if tree._Compare(e.First, node.KeyAt(i)) > 0 {
			break
		}
		i++
	}

	if node.Leaf {
		keys := node.Keys()
		datas := node.Datas()
		for x := c; x > i; x-- {
			keys[x] = keys[x-1]
			datas[x] = datas[x-1]
		}
		keys[i] = e.First
		datas[i] = e
		node.ChangeKeySize(1)

		if node.CountIsMaxX() {
			//增加key之后可能要拆分节点
			return tree.splidNode(node)
		}
	} else {
		return tree.insertTo(node.ChildAt(i), e)
	}
	return nil, nil
}

/**
 * @description: 节点操作完成之后，keysize==order，需要进行拆分
	特别需要注意：叶节点的拆分和非叶节点的拆分是不一样的，非叶节点的拆分会拿走一个元素放到父节点，叶节点的拆分不会，
				这是叶节点中存储了数据，不能拆借走，非叶节点中只存储索引，没有数据

	叶子节点进行拆分的时候，会改变Next属性

 * @param	node为需要拆分的节点
 * @return: 拆分之后是否发生了根节点变更，发生了则返回新的根节点，没有发生则返回nil
*/
func (tree *BPlusTree) splidNode(node *BPlusTreeNode) (*BPlusTreeNode, error) {
	if node == nil {
		return nil, errors.New("splidNode : node is nil")
	}
	if !node.CountIsMaxX() {
		return nil, errors.New("splidNode : node is not max")
	}

	var retNode *BPlusTreeNode = nil
	p := node.GetParent()
	//根节点需要拆分，先创建一个根
	if p == nil {
		p = NewBPlusTreeNode(tree._Order, tree._Compare)
		p.Leaf = false
		node.SetParent(p)
		p.InsertChild(node)
		retNode = p
	}

	maxCount := node._maxCount
	middle := maxCount / 2

	p.InsertKey(node.KeyAt(middle))

	right := NewBPlusTreeNode(node._Order, node._Compare)
	right.Leaf = node.Leaf

	if node.Leaf { //如果是叶节点，需要转移数据

		for k := node.KeySize() - 1; k >= middle; k-- {
			tree.insertTo(right, node.DataAt(k))
			node.RemoveKeyAt(k)
		}
		oldNext := node.Next
		node.Next = right
		right.Next = oldNext
	} else { //如果是非叶节点，中间的元素不应该转移到right中
		for k := node.KeySize() - 1; k > middle; k-- {
			//如果是非叶节点，不需要转移数据，但是要拆分children
			right.InsertKey(node.KeyAt(k))
			node.RemoveKeyAt(k)
		}
		node.RemoveKeyAt(middle)

		for k := node.ChildrenSize() - 1; k >= middle+1; k-- {
			//拆分的时候，底下必定有5+个子节点，拆分之后左子节点2(或者3)个，右子节点3个，中间位置的提到父节点中
			//父节点4个key,底下子节点数为5，并且都各个子节点key都满了，则如果再添加一个key到某子节点，必然造成拆分，
			//拆分之后父节点keys数+1超过最大限额，底下子节点数变为6
			right.InsertChild(node.ChildAt(k))
			node.RemoveChildAt(k)
		}
	}

	right.SetParent(node.GetParent())
	//此处还可以优化，p.InsertKey的时候能够知道相关InsertChild的位置信息
	p.InsertChild(right)

	//拆分之后，父节点keys数增加了，所以需要对父节点做一个判断，看是否还需要进一步拆分
	if p.CountIsMaxX() {
		return tree.splidNode(p)
	}
	return retNode, nil
}

/**
 * @description: 将两个节点合并，必然是一个node1.keysize<min一个节点node2.keysize==min,需要做判断
		叶节点的(合并与拆借操作)与根节点的合并逻辑上是有较大差异的：
		叶节点合并：left0,left1,right0合并，先回变成[left0,left1,right0],并更新parent.key
		非叶节点合并：left0,left1,right0合并，先变成[left0,left1,parentx,right0],
		叶节点拆借：left0,left1(del),right0,right1,right2，会变成[left0,right0]和[right1,right2]，并更新parent.key=right1
		非叶节点拆借：left0,left1(del),right0,right1,right2 会变成[left0,parentx]和[right1,right2]，并更新parent.key=right0

		叶子节点合并的时候，会改变Next属性，需要维护

 * @param node1 需要合并的两个节点
 * @param node2
 * @return:如果是根节点发生了变更，则返回新的根节点，否则返回Nil
*/
func (tree *BPlusTree) mergeNode(node1, node2 *BPlusTreeNode) (*BPlusTreeNode, error) {
	if node1 == nil && node2 == nil {
		return nil, errors.New("mergeNode:node1 or node2 is nil")
	}
	if node1.GetParent() != node2.GetParent() {
		return nil, errors.New("mergeNode:node1 is not the brother of node2!")
	}
	keysCountMatch := (node1.CountIsMinX() && node2.CountIsMin()) || (node1.CountIsMin() == node2.CountIsMinX())
	if !keysCountMatch {
		return nil, errors.New("mergeNode:node1 or node2 has too many keys!")
	}

	node1_is_small := false
	key1 := node1.KeyAt(0)
	key2 := node2.KeyAt(0)
	p := node1.GetParent()

	index1 := p.ChildIndex(node1)

	if tree._Compare(key1, key2) > 0 {
		node1_is_small = true
	}

	if node1_is_small {
		//把node2中的数据都移到node1中
		if node1.Leaf {
			size := node2.KeySize()
			for i := 0; i < size; i++ {
				tree.insertTo(node1, node2.DataAt(i))
			}
			node1.Next = node2.Next
		} else {
			size := node2.KeySize()
			for i := 0; i < size; i++ {
				node1.InsertKey(node2.KeyAt(i))
			}
			transferKey := p.KeyAt(index1)
			node1.InsertKey(transferKey)
			childrenSize := node2.ChildrenSize()
			for i := 0; i < childrenSize; i++ {
				node1.InsertChild(node2.ChildAt(i))
			}
		}

		nodeIndex1 := p.ChildIndex(node1)
		nodeIndex2 := p.ChildIndex(node2)

		p.RemoveChildAt(nodeIndex2)
		//合并之后，在p.keys中，卡在node1和node2之间的分隔符就没有意义了，要删除
		p.RemoveKeyAt(nodeIndex1)
		node2.Destroy()
		//如果是根节点，则直接返回即可
		if p.GetParent() == nil && p.KeySize() == 0 {
			node1.SetParent(nil)
			return node1, nil
		}
		//合并之后可能造成父节点的Key数量减少
		if p.CountIsMinX() {
			return tree.deleteFixup(p)
		}
	} else {
		return tree.mergeNode(node2, node1)
	}

	return nil, nil
}

/**
 * @description: 节点的删除
 * @param key 需要删除的节点的key
 * @return:void
 */
func (tree *BPlusTree) Delete(key interface{}) error {
	if key == nil {
		return errors.New("Delete:key is nil!")
	}
	newRoot, err := tree.deleteFrom(tree.Root, key)
	if err != nil {
		return err
	}
	if newRoot != nil {
		tree.Root = newRoot
	}
	tree.updateLeafHead()
	return nil
}

func (tree *BPlusTree) deleteFrom(node *BPlusTreeNode, key interface{}) (*BPlusTreeNode, error) {
	i := 0
	c := node.KeySize()

	if node.Leaf {
		for i < c {
			if tree._Compare(key, node.KeyAt(i)) == 0 {
				break
			}
			i++
		}
		if i < c {
			keys := node.Keys()
			datas := node.Datas()
			ksize := node.KeySize()
			for x := i; x < c-1; x++ {
				keys[x] = keys[x+1]
				datas[x] = datas[x+1]
			}
			keys[ksize-1] = nil
			datas[ksize-1] = nil
			node.ChangeKeySize(-1)
			if node.CountIsMinX() { //增加key之后可能要拆分节点
				return tree.deleteFixup(node)
			} else if i == 0 { //如果是首key，则可能会对parent的keys造成影响，这里矫正一下
				p := node.GetParent()
				if p != nil {
					index := p.ChildIndex(node)
					//最左侧一个child的keys变化了，不会影响父节点的keys
					if index > 0 {
						p.RemoveKeyAt(index - 1)
						p.InsertKey(node.KeyAt(0))
					}
				}
			}
		}
	} else {
		for i < c {
			if tree._Compare(key, node.KeyAt(i)) > 0 {
				break
			}
			i++
		}

		return tree.deleteFrom(node.ChildAt(i), key)
	}
	return nil, nil
}

/**
 * @description:删除之后需要调整节点
 * @param node：需要调整的节点
 * @return:void
 */
func (tree *BPlusTree) deleteFixup(node *BPlusTreeNode) (*BPlusTreeNode, error) {
	if node == nil || node.GetParent() == nil {
		return nil, errors.New("deleteFixup:node is nil or node.parent is nil")
	}
	if !node.CountIsMinX() {
		return nil, errors.New("deleteFixup:node has too many keys!")
	}

	p := node.GetParent()
	cnt := p.ChildrenSize()
	var leftSibling *BPlusTreeNode = nil
	var rightSibling *BPlusTreeNode = nil
	leftIndex := -1
	rightIndex := -1
	index := p.ChildIndex(node)

	//右边有兄弟节点
	if index >= 0 && index < cnt-1 {
		rightIndex = index + 1
		rightSibling = p.ChildAt(rightIndex)
	}
	//左边有兄弟节点
	if index > 0 && index <= cnt-1 {
		leftIndex = index - 1
		leftSibling = p.ChildAt(leftIndex)
	}

	//左兄弟节点有富余，可以借一个
	if leftSibling != nil && !leftSibling.CountIsMin() {
		//左兄弟节点中借一个
		iKey := leftSibling.KeyAt(leftSibling.KeySize() - 1)
		iData := leftSibling.DataAt(leftSibling.KeySize() - 1)
		iChild := leftSibling.ChildAt(leftSibling.ChildrenSize() - 1)
		parentKey := p.KeyAt(index)

		leftSibling.RemoveKeyAt(leftSibling.KeySize() - 1)
		leftSibling.RemoveChildAt(leftSibling.KeySize() - 1)

		//非叶节点处理key和child
		if !leftSibling.Leaf {
			p.RemoveKeyAt(index)
			//处理key
			node.InsertKey(parentKey)
			//处理child
			node.InsertChild(iChild)
			p.InsertKey(iKey)
		} else { //叶节点处理key和data
			tree.insertTo(node, iData)
			//该分隔点位置更新到父节点中
			p.RemoveKeyAt(leftIndex)
			p.InsertKey(iKey)
		}
	} else if rightSibling != nil && !rightSibling.CountIsMin() { //右边兄弟节点有富余，可以借一个
		iKey := rightSibling.KeyAt(0)
		iData := rightSibling.DataAt(0)
		iChild := rightSibling.ChildAt(0)
		parentKey := p.KeyAt(index)

		rightSibling.RemoveKeyAt(0)
		rightSibling.RemoveChildAt(0)

		//非叶节点需要处理key和child,而且key的处理方式是不一样的
		if !rightSibling.Leaf {
			p.RemoveKeyAt(index)
			//处理key
			node.InsertKey(parentKey)
			//处理child
			node.InsertChild(iChild)
			p.InsertKey(iKey)
		} else { //叶节点需要处理key和data
			tree.insertTo(node, iData)

			//该分隔点位置更新到父节点中
			p.RemoveKeyAt(index)
			newKey := rightSibling.KeyAt(0)
			p.InsertKey(newKey)
		}
	} else { //如果左右兄弟节点都没有富余,则需要合并节点
		if leftSibling != nil { //优先合并左兄弟节点
			return tree.mergeNode(leftSibling, node)
		} else if rightSibling != nil { //没有左兄弟节点才选择合并右兄弟节点
			return tree.mergeNode(node, rightSibling)
		} else {
			fmt.Println("leftsibling is nil and rightsibling is nil")
		}
	}

	//如果只是发生拆借行为，不会造成parent的节点数减少，因而不会变更root节点，但是如果发生了合并节点的情况，则可能会造成root变更
	return nil, nil
}

func (tree *BPlusTree) GetDeepList(node *BPlusTreeNode, list [][]*BPlusTreeNode, deep int) {
	for k := 0; k < node.ChildrenSize(); k++ {
		list[deep] = append(list[deep], node.ChildAt(k))
	}

	for i := 0; i < node.ChildrenSize(); i++ {
		child := node.ChildAt(i)
		tree.GetDeepList(child, list, deep+1)
	}
}

/**
 * @description: 打印树结构，方便调试
 * @param show_leafs 是否显示叶子上的数据信息？默认不显示
 * @return:void
 */
func (tree *BPlusTree) Print(show_leafs ...bool) {
	node := tree.Root

	show_leaf := false
	if len(show_leafs) > 0 {
		show_leaf = show_leafs[0]
	}
	num := 10
	list := make([][]*BPlusTreeNode, num)
	for i := 0; i < num; i++ {
		list[i] = []*BPlusTreeNode{}
	}
	list[0] = []*BPlusTreeNode{node}
	tree.GetDeepList(node, list, 1)

	for i := 0; i < num; i++ {
		arr := list[i]
		for j := 0; j < len(arr); j++ {
			jchild := arr[j]
			fmt.Printf("(")
			for w := 0; w < jchild.KeySize(); w++ {
				fmt.Printf(fmt.Sprintf("%d,", jchild.KeyAt(w)))
			}
			leaf := ""
			if jchild.Leaf {
				leaf = "leaf"
			}
			fmt.Printf(" %s) ", leaf)
		}
		if len(arr) > 0 {
			leaf := arr[0].Leaf
			if show_leaf && leaf {
				fmt.Println()
				fmt.Println("--leaf datas--")
				for j := 0; j < len(arr); j++ {
					jchild := arr[j]
					fmt.Printf("(")
					for _, y := range jchild.Datas() {
						if y == nil {
							fmt.Printf("-,")
						} else {
							fmt.Printf(fmt.Sprintf("%d,", y.First))
						}
					}
					fmt.Printf(") ")
				}
			}
			fmt.Println()
		}

	}
}

func (tree *BPlusTree) LoopList(msg string) {
	node := tree.GetLeafHead()
	fmt.Println(fmt.Sprintf("-------------start loop list %s-------------", msg))
	for node != nil {
		size := node.KeySize()
		fmt.Printf("(")
		for i := 0; i < size; i++ {
			fmt.Printf("%d,", node.KeyAt(i))
		}
		fmt.Printf(")")
		fmt.Println()
		node = node.Next
	}
	fmt.Println("-------------end loop list-------------")
	fmt.Println()
}
