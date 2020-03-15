/*
 * @Description: 二项堆
	算法导论3中没有二项堆的具体介绍，算法导论2中是有的，这里也实现一下
	二项堆的Search性能是比较差的。

	实现了以下方法：
	ExtractMin	:删除堆中的最小元素，并返回该元素
	DecreaseKey: 减小某个节点的key值
	Delete:删除某个节点，实际上市先DecreaseKey(node,-unlimit)，然后执行ExtractMin
	Insert:向堆中插入节点
	Minimum:返回最小节点，即root节点
	UnionBinomialHeap:二项堆合并
	NewBinomialHeap:新建二项堆

	二项堆和斐波那契堆很相似，但是内部链表结构维护上是不同的

	1、二项堆根节点是一个单链表，只持有head的引用(heap.root指向单链表的head)，而斐波那契堆维护了一个根节点的数组
	2、二项堆的子节点也是一个单链表，只持有head的引用（node.Child指向单链表的head），而斐波那契堆维护了一个子节点数组
	3、二项堆在兄弟节点的维护上使用的是单链表，而斐波那契堆使用的是双向循环链表
	4、斐波那契堆是更松散的结构，维护的信息更丰富
	5、如果不对斐波那契堆使用DecreaseKey和Delete操作，则堆中的每棵树和二项堆是一样的
		斐波那契堆中多了cut和cascading_cut操作，而二项堆中对于key比较小的节点会向上冒泡

	算法分析中O(n)， Θ(n)，Ω(n)分别来描述算法在最差情况，平均情况，最好情况时的执行效率
	过程			二叉堆(最坏情况)		   二项堆(最坏情况)             斐波那契堆(平摊)
-------------------------------------------------------------------------------------
	MakeHeap		Θ(1)            		Θ(1)						Θ(1)
	Insert			Θ(lgn)					Ω(lgn)						Θ(1)
	Minimum			Θ(1)					Ω(lgn)						Θ(1)
	ExtracMin		Θ(lgn)					Θ(lgn)						Ο(lgn)
	Union			Θ(n)					Θ(lgn)						Θ(1)
	DecreaseKey		Θ(lgn)					Θ(lgn)						Θ(1)
	Delete			Θ(lgn)					Θ(lgn)						Ο(lgn)

 * @Author: wangchengdg@gmail.com
 * @Date: 2020-03-05 11:40:09
 * @LastEditTime: 2020-03-11 15:14:08
 * @LastEditors:
*/

package DataStruct

import (
	"errors"
	"fmt"

	. "github.com/meshcross/algorithm-3rd/mesh/common"
	. "github.com/meshcross/algorithm-3rd/mesh/sort_algorithm"
)

func NodeCompareFunc_BinomialHeapNodeLessThan(x, y interface{}) int {

	nodex := ToBinomialHeapNode(x)
	nodey := ToBinomialHeapNode(y)

	if x != nil && y != nil {
		return nodex.Compare(nodex.Key, nodey.Key)
	}

	return -1
}

func ToBinomialHeapNode(v interface{}) *BinomialHeapNode {
	if node, ok := v.(*BinomialHeapNode); ok {
		return node
	}
	return nil
}

/**
 *  和斐波那契堆相比，二项堆的节点结构上差异比较大
 */
type BinomialHeapNode struct {
	Compare NodeCompareFunc
	Key     interface{}
	_Degree int
	Child   *BinomialHeapNode
	Parent  *BinomialHeapNode
	Next    *BinomialHeapNode
}

func NewBinomialHeapNode(k interface{}, compare NodeCompareFunc) *BinomialHeapNode {
	return &BinomialHeapNode{Key: k, Compare: compare}
}

/**
 * 二项堆没有位于一个roots数组，而是只有一个root根节点，然后单向链表将根节点连接起来
 */
type BinomialHeap struct {
	Compare NodeCompareFunc   //比较函数
	root    *BinomialHeapNode //根节点
	n       int               //节点总数
}

/**
 * @description: 构建二项堆
 * @param compare:比较函数
 * @return:二项堆的指针
 */
func NewBinomialHeap(compare NodeCompareFunc) *BinomialHeap {
	return &BinomialHeap{Compare: compare}
}

/**
 * @description: 批量插入节点，避免每次插入调用consolidate
 * @param nodes:节点数组
 * @return:void
 */
func (heap *BinomialHeap) InsertNodes(nodes []*BinomialHeapNode) {
	l := len(nodes)
	for k := 0; k < l; k++ {
		heap.Insert(nodes[k], false)
	}
	heap.consolidate()
}

/**
 * @description: 将节点x插入堆中
 * @param x：待插入的节点
 * @param consolidate_now：是否要立即执行consolidate
 * @return:
 */
func (heap *BinomialHeap) Insert(x *BinomialHeapNode, consolidate_now ...bool) {

	if heap.root == nil {
		heap.root = x
	} else {
		if heap.Compare(x.Key, heap.root.Key) > 0 {
			oldRoot := heap.root
			heap.root = x
			x.Next = oldRoot
		} else {
			x.Next = heap.root.Next
			heap.root.Next = x
		}
	}

	heap.n++

	do_now := true
	if len(consolidate_now) > 0 {
		do_now = consolidate_now[0]
	}
	if do_now {
		heap.consolidate()
	}
}

/**
 * @description: 返回堆中包含最小关键字的节点
 * @param {type}
 * @return:
 */
func (heap *BinomialHeap) Minimum() *BinomialHeapNode {
	return heap.root
}

func (heap *BinomialHeap) GetLeftSibling(node *BinomialHeapNode) *BinomialHeapNode {
	if node.Parent != nil {
		current := node.Parent.Child
		for current != nil {
			if current.Next == node {
				return current
			}
			current = current.Next
		}
	} else {
		current := heap.root
		for current != nil {
			if current.Next == node {
				return current
			}
			current = current.Next
		}
	}
	return nil
}

/**
 * @description: 将堆中包含最小关键字的节点删除
	删除当前的root节点，将root节点的子节点都变为堆的根节点，然后调整堆状态，设置新的root节点
 * @param {type}
 * @return:
*/
func (heap *BinomialHeap) ExtractMin() *BinomialHeapNode {

	oldRoot := heap.root
	tmpRoot := oldRoot.Next
	oldRoot.Next = nil

	//可以添加进去之后再排序，但是这里只需要找到最小的节点，所以无需整体排序
	current := tmpRoot
	last := current
	minNode := current
	var minLeft *BinomialHeapNode = nil

	for current != nil {
		if heap.Compare(current.Key, minNode.Key) > 0 {
			minNode = current
			minLeft = last
		}
		last = current
		current = current.Next

	}
	last.Next = oldRoot.Child

	current = oldRoot.Child
	for current != nil {
		current.Parent = nil
		if heap.Compare(current.Key, minNode.Key) > 0 {
			minNode = current
			minLeft = last
		}
		last = current
		current = current.Next

	}
	if minNode != tmpRoot {
		minLeft.Next = minNode.Next
		minNode.Next = tmpRoot
	}
	heap.root = minNode
	return oldRoot
}

/**
 * @description: 将新的关键字key赋给节点x,
	单链表结构处理起来比较麻烦，current和current.p交换位置，则跟这两个节点相邻的所有节点都需要处理
	current.Child,current.Left,current.Right,p.Child,p.Left,p.Right,p.Parent,p.Parent.Child,
	由于是单向链表，所以Left还需要函数计算获得

	总体上分为以下几种情况：
	1 current为p的child，current不为p的child
	2 p为pp的child，p不为pp的child
	3 pp存在 ，pp不存在

	如下所示的二项堆，把50节点替换为3，则current=node50,p=node8,pp=node5
		current不为p的child，p不为pp的child，pp存在
		需要修改的值为:
		node100.parent=node8,
		node10.next=node8
		node8.parent=node50(新key为3)
		node8.child=node100
		node25.next=node50
		node50.parent=node5
		node50.child=node10

		其他情况下可能的修改
		node8.next=
		node50.next=
		（如果根节点有多个，还涉及到计算新的root节点）

		然后继续迭代，把根节点5用node50(key=3)替换掉

原始形态：
	5
	|   \    \
	15   25   8
		 |    |   \
		 80   10   50
					|
					100

处理完的形态：
	3
	|   \    \
	15   25   5
		 |    |   \
		 80   10   8
					|
					100

 * @param {type}
 * @return:
*/
func (heap *BinomialHeap) DecreaseKey(x *BinomialHeapNode, key interface{}) error {
	if heap.Compare(key, x.Key) < 0 {
		return errors.New("DecreaseKey:new key is greater than current key!")
	}

	x.Key = key
	current := x
	p := current.Parent
	//和父节点交换位置
	for p != nil && heap.Compare(current.Key, p.Key) > 0 {
		cc := current.Child
		pp := p.Parent
		cnext := current.Next
		pnext := p.Next

		var cleft *BinomialHeapNode = heap.GetLeftSibling(current)
		var pleft *BinomialHeapNode = heap.GetLeftSibling(p)
		pchild := p.Child

		p.Child = cc
		if cc != nil {
			cc.Parent = p
		}

		current.Parent = pp
		if pchild == current {
			current.Child = p
		} else {
			current.Child = pchild
		}

		p.Parent = current
		p.Next = cnext

		//current没有左兄弟，即自己是head节点，会被父节点引用
		if cleft != nil {
			cleft.Next = p
		}

		if pp != nil && pp.Child == p {
			pp.Child = current
		}
		if pleft != nil {
			pleft.Next = current
		}
		current.Next = pnext

		//root节点被替换了，这里要更改引用
		if p == heap.root {
			heap.root = current
		}
		//循环条件，沿着父节点向上迭代
		p = current.Parent
	}

	//已经到根节点一层，结果current不是root节点本身，所以需要把current和root节点进行比较,看是否需要替换root
	if p == nil && heap.root != current {
		//current.key更小，需要和根节点互换位置
		if heap.Compare(current.Key, heap.root.Key) > 0 {
			if heap.root.Next == current {
				oldRoot := heap.root
				oldRoot.Next = current.Next
				heap.root = current
				current.Next = oldRoot
			} else {
				oldRoot := heap.root
				next := oldRoot.Next
				left := heap.GetLeftSibling(current)
				left.Next = oldRoot
				oldRoot.Next = current.Next
				current.Next = next
				heap.root = current
			}

		}
	}

	return nil
}

/**
 * @description: 从堆中删除节点x
 * @param {type}
 * @return:
 */
func (heap *BinomialHeap) Delete(node *BinomialHeapNode) {
	heap.DecreaseKey(node, -Unlimit())
	heap.ExtractMin()
}

/**
 * @description: 根节点之间的link操作,x是更小的节点，y会被吃掉
 * @param
 * @return:
 */
func (heap *BinomialHeap) link(y, x *BinomialHeapNode) {

	if x == nil || y == nil {
		return
	}
	if heap.Compare(y.Key, x.Key) > 0 {
		return
	}
	yleft := heap.GetLeftSibling(y)
	y.Parent = x
	if x.Child == nil {
		x.Child = y
	} else {
		current := x.Child
		for current.Next != nil {
			current = current.Next
		}
		current.Next = y
	}

	if yleft != nil {
		yleft.Next = y.Next
	}
	//y的Next属性要调整
	y.Next = nil
	//child属性不需要调整
	// y.Child = nil
	x._Degree++
}

/**
 * @description: 根节点上进行操作，合并掉degree相同的根节点
 */
func (heap *BinomialHeap) consolidate() {
	degrees := make([]*BinomialHeapNode, heap.n+1)

	tmpList := heap.GetRoots()

	//千万要注意这里的循环结束方式，链表的结构随时在变化，所以需要可靠的中止条件
	for !tmpList.Empty() {
		_, start := tmpList.DeleteAt(0)
		x := ToBinomialHeapNode(start)
		if x == nil {
			break
		}
		d := x._Degree

		for degrees[d] != nil {
			y := degrees[d]
			//小的保留根节点身份，大的变为子节点
			if heap.Compare(x.Key, y.Key) < 0 {
				tmp := x
				x = y
				y = tmp
			}
			heap.link(y, x)

			degrees[d] = nil
			d++
		}
		degrees[d] = x
	}
}
func (heap *BinomialHeap) GetRootsCount() int {
	current := heap.root
	cnt := 0
	for current != nil {
		cnt++
		current = current.Next
	}
	return cnt
}
func (heap *BinomialHeap) GetRoots() *ArrayList {
	roots := NewArrayList(8, heap.Compare)
	current := heap.root
	for current != nil {
		roots.Append(current)
		current = current.Next
	}
	return roots
}
func (heap *BinomialHeap) Destroy() {

}

func (heap *BinomialHeap) Print(msg string) {
	fmt.Println("--- BinomialHeap ---  ", msg, "  ------")
	current := heap.root

	nodes := make([][]*BinomialHeapNode, 100)
	deep := 0
	for current != nil {
		heap._printNode(current, &nodes, deep+1)
		if nodes[deep] == nil {
			nodes[deep] = []*BinomialHeapNode{}
		}
		nodes[deep] = append(nodes[deep], current)
		current = current.Next
	}
	for i := 0; i < len(nodes); i++ {
		arr := nodes[i]
		if arr != nil {
			ilen := len(arr)
			fmt.Println()
			for j := 0; j < ilen; j++ {
				if arr[j] == nil {
					fmt.Print(" - ")
				} else {
					fmt.Print(arr[j].Key, ",")
				}
			}
		}
	}
	fmt.Println()
	fmt.Println("############################################################")
}

func (heap *BinomialHeap) _printNode(node *BinomialHeapNode, nodes *[][]*BinomialHeapNode, deep int) {

	(*nodes)[deep] = append((*nodes)[deep], NewBinomialHeapNode("(", heap.Compare))
	if node.Child != nil {
		current := node.Child

		for current != nil {
			heap._printNode(current, nodes, deep+1)
			if (*nodes)[deep] == nil {
				(*nodes)[deep] = []*BinomialHeapNode{}
			}
			(*nodes)[deep] = append((*nodes)[deep], current)
			current = current.Next
		}
	}
	(*nodes)[deep] = append((*nodes)[deep], NewBinomialHeapNode(")", heap.Compare))
}

/**
 * @description: 合并二项堆
 * @param
 * @param
 * @return: 新的二项堆的指针
 */
func UnionBinomialHeap(h1, h2 *BinomialHeap) *BinomialHeap {

	newHeap := NewBinomialHeap(h1.Compare)

	roots1 := h1.GetRoots()
	roots2 := h2.GetRoots()
	roots1.Concate(roots2)

	sorter := NewQuickSort()
	//按照key,排好序
	sorter.SortAny(roots1.Datas(), NodeCompareFunc_BinomialHeapNodeLessThan)

	newHeap.root = ToBinomialHeapNode(roots1.ItemAt(0))
	newHeap.n = h1.n + h2.n

	newHeap.consolidate()

	h1.Destroy()
	h2.Destroy()
	return newHeap
}
