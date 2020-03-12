/*
 * @Description: 第19章 斐波那契堆
	一个斐波那契堆是一系列具有最小堆序的有根树的集合，也就是说，每棵树均遵循最小堆性质，每个节点的关键字大于或者等于它的父节点的关键字。
	最小堆的对于子节点数量没有强制限定，所以当往最小堆插入一个key的时候，理论上是有很多个选择的，它可以插入到任何一个node.key<key的节点中作为孩子节点。

	根节点是无序的。
	可以预见斐波那契堆的Search性能是比较差的。

	斐波那契堆数据结构有两种用途
	1、它支持一系列操作，这些操作构成了所谓的可合并堆
	2、斐波那契堆的一些操作可以在常数摊还时间内完成，这使得这种数据结构非常适用于需要频繁调用这些操作的应用

	可用于最小生成树和单源最短路径

	可合并堆(mergeable heap)是支持一下5种操作的一种数据结构，其中每个元素都有一个关键字：
	Make-Heap():创建和返回一个新的不含有任何元素与的堆
	Insert(H,x):将一个已填入关键字的元素x插入堆H中
	Minimun(H):返回一个指向堆H中具有最小关键字元素的指针
	Extract-Min(H):从堆H中删除最小关键字的元素，并返回指向该元素的指针
	Union(H1,H2):创建并返回一个包含堆H1和堆H2中所有元素的新堆，堆H1和堆H2在这一操作中被销毁

	除了以上5种操作外，斐波那契堆还支持一下操作：
	Decrease-Key(H,x,k):将堆H中的元素x的关键字赋予新值k，嘉定新值k不大于当前关键字
	Delete(H,x):从堆H中删除元素x

	性质19.1：设x是斐波那契堆中的任意节点，并假定x.degree=k。设y1,y2,...,yk表示x的孩子，
			并以他们链入x的先后顺序排列，则y1.degree>=0,且对于i=2,3,...,k，有yi.degree>=i-2
	推论19.5：一个n个节点的斐波那契堆中任意节点的最大度数D(n)为O(lgn)

		过程					二项堆(最坏情况)             斐波那契堆(平摊)
------------------------------------------------------------------------
	MakeHeap				Θ(1)						Θ(1)
	Insert					Ω(lgn)						Θ(1)
	Minimum					Ω(lgn)						Θ(1)
	ExtracMin				Θ(lgn)						Ο(lgn)
	Union					Θ(lgn)						Θ(1)
	DecreaseKey				Θ(lgn)						Θ(1)
	Delete					Θ(lgn)						Ο(lgn)

 * @Author: wangchengdg@gmail.com
 * @Date: 2020-03-05 11:41:17
 * @LastEditTime: 2020-03-10 10:42:05
 * @LastEditors:
*/

package DataStruct

import (
	"errors"
	"fmt"
	"strings"

	. "github.com/meshcross/algorithm-3rd/mesh/common"
)

/**
* @description: 堆的性质维护交给consolidate完成，
	理论上不应该在外部调用node.AddChild方法，所有孩子节点的增加都应该通过heap.Insert进行,然后调用consolidate维护堆性质
*/
type FibonacciNode struct {
	Mark     bool            //自从上一次成为别人的孩子之后是否失去过child
	Left     *FibonacciNode  //左兄弟节点
	Right    *FibonacciNode  //右兄弟节点
	Parent   *FibonacciNode  //父节点
	Children *ArrayList      //子节点
	_Degree  int             //孩子的数目存在degree中
	Key      interface{}     //关键字,数字类型,int32 int64 float32 float64
	Compare  NodeCompareFunc //比较函数
}

func (node *FibonacciNode) _getNodeCount(n *FibonacciNode, c *int) {

	size := n.Children.Size()
	*c += size
	for i := 0; i < size; i++ {
		iChild := n.ChildAt(i)
		iChild._getNodeCount(iChild, c)
	}
}

func (node *FibonacciNode) GetTreeNodeCount() int {
	n := 0
	node._getNodeCount(node, &n)
	return n
}
func (node *FibonacciNode) RemoveChild(c *FibonacciNode) {
	node.Children.Delete(c)
	node.updateLinkList()
	c.Parent = nil
	// c.Destory()
}
func (node *FibonacciNode) updateDegree() {
	node._Degree = node.Children.Size()
}

/**
 * @description: 不要在外部直接调用node.AddChild，应该通过heap.Insert加入
 * @param
 * @return:
 */
func (node *FibonacciNode) addChild(c *FibonacciNode) {
	node.Children.Append(c)
	c.Parent = node
	node.updateLinkList()
}

func (node *FibonacciNode) ChildAt(index int) *FibonacciNode {
	c := node.Children.ItemAt(index)
	node, ok := c.(*FibonacciNode)
	if ok {
		return node
	}
	return nil
}

func (node *FibonacciNode) Destory() {
	node._Degree = 0
	node.Parent = nil
	node.Left = nil
	node.Right = nil
	node.Mark = false
	node.Children = nil
}

func (node *FibonacciNode) Init() {
	node._Degree = 0
	node.Parent = nil
	node.Children.Clear()
	node.Mark = false
}

/**
 * @description: 子节点之间要形成环状
 * @param {type}
 * @return:
 */
func (node *FibonacciNode) updateLinkList() {
	l := node.Children.Size()

	if l == 1 {
		child := node.ChildAt(0)
		child.Left = child
		child.Right = child
	} else if l > 1 {
		for i := 0; i < l-1; i++ {
			iChild := node.ChildAt(i)
			nextChild := node.ChildAt(i + 1)
			iChild.Right = nextChild
			nextChild.Left = iChild
		}
		firstChild := node.ChildAt(0)
		lastChild := node.ChildAt(l - 1)
		firstChild.Left = lastChild
		lastChild.Right = firstChild
	}
}

func (node *FibonacciNode) _x(list *[][]*FibonacciNode, deep int) {
	if len(*list) <= deep {
		*list = append(*list, []*FibonacciNode{})
	}
	l := node.Children.Size()

	for i := 0; i < l; i++ {
		(*list)[deep] = append((*list)[deep], node.ChildAt(i))
	}
	for i := 0; i < l; i++ {
		node.ChildAt(i)._x(list, deep+1)
	}

	if deep > 100 {
		panic("error")
	}
}

func (node *FibonacciNode) Print() {
	fmt.Println(node.Key, " root node")
	arr := [][]*FibonacciNode{}
	node._x(&arr, 0)

	l := len(arr)

	for i := 0; i < l; i++ {
		iarr := arr[i]
		ilen := len(iarr)
		for j := 0; j < ilen; j++ {
			fmt.Printf("%d,", iarr[j].Key)
		}
		fmt.Println()
	}

}

func NewFibonacciNode(key interface{}, compare NodeCompareFunc) *FibonacciNode {
	return &FibonacciNode{Key: key, Children: NewArrayList(8, compare), Compare: compare}
}

type FibonacciHeap struct {
	n       int             //所有节点总数
	minNode *FibonacciNode  //最小节点，即双向链表头
	roots   *ArrayList      //根节点，多个，动态数组
	Compare NodeCompareFunc //node.key的比较函数
}

/**
 * @description: 构建斐波那契堆
 * @param compare:比较函数
 * @return: 斐波那契堆的指针
 */
func NewFibonacciHeap(compare NodeCompareFunc) *FibonacciHeap {
	roots := NewArrayList(8, compare)
	return &FibonacciHeap{roots: roots, Compare: compare}
}

func (heap *FibonacciHeap) Minimum() *FibonacciNode {
	return heap.minNode
}

//每个根节点实际上是一个最小堆，这里需要维护最小堆的性质
func (heap *FibonacciHeap) heapify() {

}
func (heap *FibonacciHeap) ChildAt(index int) *FibonacciNode {
	c := heap.roots.ItemAt(index)
	node, ok := c.(*FibonacciNode)
	if ok {
		return node
	}
	return nil
}

func (heap *FibonacciHeap) Destroy() {

}
func (heap *FibonacciHeap) Print() {
	fmt.Println("--start heap print--")
	l := heap.roots.Size()
	for i := 0; i < l; i++ {
		iRoot := heap.ChildAt(i)
		iRoot.Print()
	}
	fmt.Println("--end heap print--")
}
func (heap *FibonacciHeap) UseTestData(minNode *FibonacciNode, roots *ArrayList) {
	heap.minNode = minNode
	heap.roots = roots
	heap.updateLinkList()
	heap.updateDegree()
	heap.updateN()
}

/**
 * @description:一次插入多个节点
 * @param arr:待插入的节点数组
 * @return:void
 */
func (heap *FibonacciHeap) InsertNodes(arr []*FibonacciNode) {
	l := len(arr)
	for i := 0; i < l; i++ {
		node := arr[i]
		heap.Insert(node, false)
	}
	heap.consolidate()
}

/**
 * @description: 插入一个新的节点到堆中
	千万要注意，不允许直接向堆中插入一棵树
 * @param x:需要插入的节点
 * @param consolidate_now:是否要理解执行consolidate？多次调用的时候有影响
 * @return:
*/
func (heap *FibonacciHeap) Insert(x *FibonacciNode, consolidate_now ...bool) error {
	if x == nil {
		return errors.New("Insert:param is nil!")
	}
	//千万要注意这里，这里表明不允许向堆中插入一棵树，会对插入的x进行清理
	x.Init()

	if heap.minNode == nil {
		heap.minNode = x
		x.Left = x
		x.Right = x
		heap.roots.Append(x)
	} else {
		//x添加到minNode的左边
		heap.roots.Append(x)
		x.Left = heap.minNode.Left
		x.Right = heap.minNode
		heap.minNode.Left = x
		if heap.Compare(x.Key, heap.minNode.Key) > 0 {
			heap.minNode = x
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

	return nil
}

func (heap *FibonacciHeap) _getRootKeys(bysize ...bool) string {
	arr := []string{}

	if len(bysize) > 0 && bysize[0] {
		size := heap.roots.Size()
		for i := 0; i < size; i++ {
			arr = append(arr, fmt.Sprintf("%d", heap.ChildAt(i).Key))
		}
	} else {
		//有可能heap.minNode没有变，但是该节点被并掉了
		var end *FibonacciNode = nil
		current := heap.minNode
		for current != end {
			if end == nil {
				end = heap.minNode
			}
			arr = append(arr, fmt.Sprintf("%d", current.Key))
			current = current.Right
		}
	}

	return strings.Join(arr, ",")
}

/**
 * @description: 调整堆的健康状态，minNode务必是最小的节点
	每个根节点有不同的degree，如果有重复的degree，则需要进行link操作
 * @param
 * @return:
*/
func (heap *FibonacciHeap) consolidate() {
	degrees := make([]*FibonacciNode, heap.n+1)

	tmpList := NewArrayList(heap.roots.Capacity(), heap.Compare)
	var end *FibonacciNode = nil
	current := heap.minNode
	for current != end {
		if end == nil {
			end = heap.minNode
		}
		tmpList.Append(current)
		current = current.Right
	}

	//千万要注意这里的循环结束方式，链表的结构随时在变化，所以需要可靠的中止条件
	for !tmpList.Empty() {
		_, start := tmpList.DeleteAt(0)
		x, ok := start.(*FibonacciNode)
		if !ok {
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

	heap.minNode = nil

	for i := 0; i < heap.n; i++ {
		if degrees[i] != nil {
			if heap.minNode == nil {
				//create a root list for heap containing just A[i]
				heap.roots = NewArrayList(8, heap.Compare)
				heap.roots.Append(degrees[i])
				heap.minNode = degrees[i]
			} else {
				heap.roots.Append(degrees[i])
				if heap.Compare(degrees[i].Key, heap.minNode.Key) > 0 {
					heap.minNode = degrees[i]
				}
			}
		}
	}
}

/**
 * @description: 将节点y链接到节点x，即y变为x的子节点
 * @param
 * @return:
 */
func (heap *FibonacciHeap) link(y, x *FibonacciNode) {

	y.Left.Right = y.Right
	y.Right.Left = y.Left

	heap.roots.Delete(y)
	y.Parent = x
	x.Children.Append(y)
	x._Degree++
	y.Mark = false
	x.updateLinkList()
}

/**
ok
 * @description: 取出minNode节点，之后堆中不再有该节点
	不论执行ExtractMin之前堆中各根节点的度数是多少，在执行完之后，每个根节点有不一样的度数
	有根节点的增删，一定要维护环形链表
 * @return: 返回minNode节点
*/
func (heap *FibonacciHeap) ExtractMin() *FibonacciNode {
	fmt.Println("###### roots:", heap._getRootKeys(false), " ----- ", heap._getRootKeys(true))
	z := heap.minNode
	if z != nil {
		lastLeft := z.Left
		lastRight := z.Right
		heap.roots.Delete(z)

		size := z.Children.Size()
		for i := 0; i < size; i++ {
			child := z.ChildAt(i)
			heap.roots.Append(child)
			child.Parent = nil
			child.Left = lastLeft

			if i == 0 {
				lastLeft.Right = child
				child.Right = z.ChildAt(i + 1)
			} else if i < size-1 {
				child.Right = z.ChildAt(i + 1)
			} else {
				child.Right = lastRight
				lastRight.Left = child
			}
			lastLeft = child
			child.updateDegree()
		}

		if z == z.Right {
			heap.minNode = nil
		} else {
			heap.minNode = z.Right

			heap.consolidate()
		}
		heap.n--
	}

	return z
}

/**
 * @description: 关键字减小
 * @param {type}
 * @return:
 */
func (heap *FibonacciHeap) DecreaseKey(x *FibonacciNode, key interface{}) error {
	if heap.Compare(key, x.Key) < 0 {
		return errors.New("DecreaseKey:new key is greater than current key!")
	}
	x.Key = key
	y := x.Parent

	if y != nil && heap.Compare(x.Key, y.Key) > 0 {
		heap.cut(x, y)
		heap.cascading_cut(y)
	}

	//没有改变环形链表的结构，所以这里直接修改指针即可
	if heap.Compare(x.Key, heap.minNode.Key) > 0 {
		heap.minNode = x
	}

	return nil
}

/**
 * @description: 切断x与其父节点y之间的连接,使x成为根节点
	cut之后放到minNode的左侧
 * @param {type}
 * @return:
*/
func (heap *FibonacciHeap) cut(x, y *FibonacciNode) {
	y.Children.Delete(x)
	heap.roots.Append(x)
	x.Parent = nil
	x.Mark = false

	heap.minNode.Left.Right = x
	x.Left = heap.minNode.Left.Right
	heap.minNode.Left = x
	x.Right = heap.minNode

}

/**
 * @description: 级联切断，当前节点的父节点也被标记了，则需要对父节点执行cut，该操作沿着树一直向上递归
 * @param
 * @return:
 */
func (heap *FibonacciHeap) cascading_cut(y *FibonacciNode) {
	z := y.Parent
	if z != nil {
		if y.Mark == false {
			y.Mark = true
		} else {
			heap.cut(y, z)
			heap.cascading_cut(z)
		}
	}
}

/**
 * @description:删除某个节点
	分为两步：1、先将节点DecreaseKey到无穷小 2、然后ExtractMin取出最小节点
 * @param {type}
 * @return:
*/
func (heap *FibonacciHeap) Delete(x *FibonacciNode) {
	heap.DecreaseKey(x, -Unlimit())
	heap.ExtractMin()
}

func (heap *FibonacciHeap) updateDegree() {
	l := heap.roots.Size()
	for i := 0; i < l; i++ {
		iRoot := heap.ChildAt(i)
		iRoot.updateDegree()
	}
}

func (heap *FibonacciHeap) updateN() {
	n := 0
	l := heap.roots.Size()
	for i := 0; i < l; i++ {
		iRoot := heap.ChildAt(i)
		n += iRoot.GetTreeNodeCount()
	}
	heap.n = n
}

func (heap *FibonacciHeap) updateLinkList() {
	l := heap.roots.Size()
	if l == 1 {
		first := heap.ChildAt(0)
		first.Left = first
		first.Right = first
	} else if l > 1 {
		first := heap.ChildAt(0)
		last := heap.ChildAt(l - 1)
		for i := 0; i < l-1; i++ {
			iRoot := heap.ChildAt(i)
			nextRoot := heap.ChildAt(i + 1)
			iRoot.Right = nextRoot
			nextRoot.Left = iRoot
		}
		first.Left = last
		last.Right = first
	}
}

/**
 * @description: 斐波那契堆的势函数
 * @param {type}
 * @return:
 */
func (heap *FibonacciHeap) GetPotential() int {
	return heap.n + 2*heap.GetMarkCount()
}

/**
 * @description: 获取标记节点的数量
 * @param
 * @return:
 */
func (heap *FibonacciHeap) GetMarkCount() int {
	ms := 0
	size := heap.roots.Size()
	for i := 0; i < size; i++ {
		m := 0
		heap._nodeMarkCount(heap.ChildAt(i), &m)
		ms += m
	}
	return ms
}
func (heap *FibonacciHeap) _nodeMarkCount(node *FibonacciNode, c *int) {
	if node.Mark {
		*c += 1
	}
	size := node.Children.Size()
	for i := 0; i < size; i++ {
		heap._nodeMarkCount(node.ChildAt(i), c)
	}
}

/**
 * @description: 斐波那契堆合并，销毁h1,h2，返回新的堆
 * @param {type}
 * @return:
 */
func FibonacciHeapUnion(h1, h2 *FibonacciHeap) *FibonacciHeap {
	newHeap := NewFibonacciHeap(h1.Compare)
	newHeap.minNode = h1.minNode

	newHeap.roots.Concate(h1.roots)
	newHeap.roots.Concate(h2.roots)

	if h1.minNode == nil || h2.minNode != nil && newHeap.Compare(h2.minNode.Key, h1.minNode.Key) > 0 {
		newHeap.minNode = h2.minNode
	}
	newHeap.n = h1.n + h2.n

	//h1,h2的根节点合并之后，degree可能会有冲突，需要处理一下
	//比如h1有3个根节点，度分别为1，2，3，h2有2个根节点，度分别为2，3，
	//则进行上述合并操作之后，根节点变为5个，度分别为1,2,3,2,3，所以要进行consolidate操作
	newHeap.consolidate()

	h1.Destroy()
	h2.Destroy()
	return newHeap
}

func AddFibChild(parentNode *FibonacciNode, childNode *FibonacciNode) {
	if parentNode != nil && childNode != nil {
		parentNode.addChild(childNode)
	}
}
