/*
 * @Description: 第五部分，高级数据结构

 * @Author: wangchengdg@gmail.com
 * @Date: 2020-03-05 11:37:55
 * @LastEditTime: 2020-03-11 14:48:04
 * @LastEditors:
 */

package DataStruct

import (
	"fmt"
	"testing"
)

func TestBTree(t *testing.T) {
	NUM := 10
	btree := NewBTree(2, NodeCompareFunc_IntBiggerThan)
	pairs := make([]*Pair, NUM)

	fmt.Println()
	for i := 0; i < NUM; i++ {
		r := i + 1
		p := NewPair(r, r)

		btree.Insert(p)

		pairs[i] = p
	}
	btree.Print()
	/**
		4
		2
		6 8
		1
		3
		5
		7
		9 10
	**/
	fmt.Println("----------------------")
	btree.Delete(2)
	btree.Print()
	/**
		6
		4
		8
		1 3
		5
		7
		9 10
	**/
}

func TestBPlusTreeInsert(t *testing.T) {
	order := 5
	btree := NewBPlusTree(order, NodeCompareFunc_IntLessThan)

	//a
	btree.Insert(NewPairAny(39, 39))
	btree.Print()
	//(39,)
	fmt.Println("-------------a-------------")

	//b
	btree.Insert(NewPairAny(22, 22))
	btree.Insert(NewPairAny(97, 97))
	btree.Insert(NewPairAny(41, 41))
	btree.Print()
	// (22,39,41,97,)
	fmt.Println("--------------b------------")

	//c
	btree.Insert(NewPairAny(53, 53))
	btree.Print()
	// (41,)
	// (22,39,) (41,53,97,)
	fmt.Println("--------------c------------")

	//d
	btree.Insert(NewPairAny(13, 13))
	btree.Insert(NewPairAny(21, 21))
	btree.Insert(NewPairAny(40, 40))
	btree.Print()
	// (22,41,)
	// (13,21,) (22,39,40,) (41,53,97,)
	fmt.Println("---------------d-----------")

	//e
	btree.Insert(NewPairAny(30, 30))
	btree.Insert(NewPairAny(27, 27))
	btree.Insert(NewPairAny(33, 33))
	btree.Insert(NewPairAny(36, 36))
	btree.Insert(NewPairAny(35, 35))
	btree.Insert(NewPairAny(34, 34))
	btree.Insert(NewPairAny(24, 24))
	btree.Insert(NewPairAny(29, 29))
	btree.Print()
	// (22,30,36,41,)
	// (13,21,) (22,24,27,29,) (30,33,34,35,) (36,39,40,) (41,53,97,)
	fmt.Println("-------------e-------------")

	//f
	btree.Insert(NewPairAny(26, 26))
	btree.Print()
	// (30,)
	// (22,26,) (36,41,)
	// (13,21,) (22,24,) (26,27,29,) (30,33,34,35,) (36,39,40,) (41,53,97,)
	fmt.Println("--------------f------------")

	//g
	btree.Insert(NewPairAny(17, 17))
	btree.Insert(NewPairAny(28, 28))
	btree.Insert(NewPairAny(31, 31))
	btree.Print()
	// (30,)
	// (22,26,33,) (36,41,)
	// (13,17,21,) (22,24,) (26,27,28,29,) (33,34,35,) (30,31,) (36,39,40,) (41,53,97,)
	fmt.Println("-------------g-------------")
}

/**
 * @description: B+树中删除节点
 * @param {type}
 * @return:void

(30,)
(22,26,) (33,36,41,)
(13,17,21,) (22,24,) (26,27,28,29,) (30,31,) (33,34,35,) (36,39,40,) (41,53,97,)
-------------finish init tree-------------
(30,)
(21,26,) (33,36,41,)
(13,17,) (21,24,) (26,27,28,29,) (30,31,) (33,34,35,) (36,39,40,) (41,53,97,)
--leaf datas--
(13,17,-,-,-,) (21,24,-,-,-,) (26,27,28,29,-,) (30,31,-,-,-,) (33,34,35,-,-,) (36,39,40,-,-,) (41,53,97,-,-,)
-------------a 22-------------
(30,)
(21,27,) (33,36,41,)
(13,17,) (21,26,) (27,28,29,) (30,31,) (33,34,35,) (36,39,40,) (41,53,97,)
--leaf datas--
(13,17,-,-,-,) (21,26,-,-,-,) (27,28,29,-,-,) (30,31,-,-,-,) (33,34,35,-,-,) (36,39,40,-,-,) (41,53,97,-,-,)
-------------b 24-------------
(33,)
(27,30,) (36,41,)
(13,21,26,) (27,28,29,) (30,31,) (33,34,35,) (36,39,40,) (41,53,97,)
--leaf datas--
(13,21,26,-,-,) (27,28,29,-,-,) (30,31,-,-,-,) (33,34,35,-,-,) (36,39,40,-,-,) (41,53,97,-,-,)
-------------c 17-------------
(33,)
(27,29,) (36,41,)
(13,21,26,) (27,28,) (29,31,) (33,34,35,) (36,39,40,) (41,53,97,)
-------------d 30-------------
(27,33,36,41,)
(13,21,26,) (27,28,29,) (33,34,35,) (36,39,40,) (41,53,97,)
-------------e 31-------------
(27,34,36,41,)
(13,21,26,) (27,28,29,) (34,35,) (36,39,40,) (41,53,97,)
-------------f del33-------------
(27,29,36,41,)
(13,21,26,) (27,28,) (29,35,) (36,39,40,) (41,53,97,)
-------------g 34-------------
(27,29,39,41,)
(13,21,26,) (27,28,) (29,36,) (39,40,) (41,53,97,)
-------------h 35-------------
(27,29,39,)
(13,21,26,) (27,28,) (29,36,) (39,41,)
-------------g 40,53,97-------------
(26,39,)
(13,21,) (26,36,) (39,41,)
*/
func TestBPlusTreeDelete(t *testing.T) {

	order := 5
	tree := NewBPlusTree(order, NodeCompareFunc_IntLessThan)

	tree.Insert(NewPairAny(39, 39))
	tree.Insert(NewPairAny(22, 22))
	tree.Insert(NewPairAny(97, 97))
	tree.Insert(NewPairAny(41, 41))
	tree.Insert(NewPairAny(53, 53))
	tree.Insert(NewPairAny(13, 13))
	tree.Insert(NewPairAny(21, 21))
	tree.Insert(NewPairAny(40, 40))
	tree.Insert(NewPairAny(30, 30))
	tree.Insert(NewPairAny(27, 27))
	tree.Insert(NewPairAny(33, 33))
	tree.Insert(NewPairAny(36, 36))
	tree.Insert(NewPairAny(35, 35))
	tree.Insert(NewPairAny(34, 34))
	tree.Insert(NewPairAny(24, 24))
	tree.Insert(NewPairAny(29, 29))
	tree.Insert(NewPairAny(26, 26))
	tree.Insert(NewPairAny(17, 17))
	tree.Insert(NewPairAny(28, 28))
	tree.Insert(NewPairAny(31, 31))
	tree.Print()
	fmt.Println("-------------finish init tree-------------")

	tree.Delete(22)
	tree.Print(true)
	fmt.Println("-------------a 22-------------")
	tree.Delete(24)
	tree.Print(true)
	fmt.Println("-------------b 24-------------")

	tree.Delete(17)
	tree.Print(true)
	fmt.Println("-------------c 17-------------")

	tree.Delete(30)
	tree.Print()
	fmt.Println("-------------d 30-------------")
	tree.Delete(31)
	tree.Print()
	fmt.Println("-------------e 31-------------")

	tree.Delete(33)
	tree.Print()
	fmt.Println("-------------f del33-------------")

	tree.Delete(34)
	tree.Print()
	fmt.Println("-------------g 34-------------")

	tree.Delete(35)
	tree.Print()
	fmt.Println("-------------h 35-------------")

	tree.Delete(40)
	tree.Delete(53)
	tree.Delete(97)
	tree.Print()
	fmt.Println("-------------g 40,53,97-------------")

	tree.Delete(27)
	tree.Delete(28)
	tree.Delete(29)
	tree.Print()
	fmt.Println("-------------i 27,28,29-------------")

	tree.Delete(21)
	tree.Delete(39)
	tree.Print()
	fmt.Println("-------------j 21,39-------------")

	tree.Delete(26)
	tree.Print()
	fmt.Println("-------------k 26-------------")
	tree.Delete(36)
	tree.Print()
	fmt.Println("-------------l 36-------------")

	tree.Delete(13)
	tree.Delete(14)
	tree.Print()
	fmt.Println("-------------m 13,14-------------")

	tree.Delete(41)
	tree.Print()
	fmt.Println("-------------n 41-------------")

	tree.LoopList(" loop all leaf node ")
}

/**
 * @description: 直接使用的书上的例子，跟进书中逻辑一步步实现的
 * @param {type}
 * @return:
 */
func TestFibonacciHeap(t *testing.T) {
	compare := NodeCompareFunc_IntLessThan
	heap := NewFibonacciHeap(compare)

	node23 := NewFibonacciNode(23, compare)
	node7 := NewFibonacciNode(7, compare)
	node21 := NewFibonacciNode(21, compare)

	node3 := NewFibonacciNode(3, compare)
	node18 := NewFibonacciNode(18, compare)
	node18.Mark = true
	node52 := NewFibonacciNode(52, compare)
	node38 := NewFibonacciNode(38, compare)
	node39 := NewFibonacciNode(39, compare)
	node39.Mark = true
	node41 := NewFibonacciNode(41, compare)

	AddFibChild(node18, node39)
	AddFibChild(node38, node41)
	AddFibChild(node3, node18)
	AddFibChild(node3, node52)
	AddFibChild(node3, node38)

	node17 := NewFibonacciNode(17, compare)
	node30 := NewFibonacciNode(30, compare)
	AddFibChild(node17, node30)

	node24 := NewFibonacciNode(24, compare)
	node26 := NewFibonacciNode(26, compare)
	node26.Mark = true
	node46 := NewFibonacciNode(46, compare)
	node35 := NewFibonacciNode(35, compare)
	AddFibChild(node26, node35)
	AddFibChild(node24, node26)
	AddFibChild(node24, node46)

	// heap.Insert(node23)
	// heap.Insert(node7)
	// heap.Insert(node21)
	// heap.Insert(node3)
	// heap.Insert(node17)
	// heap.Insert(node24)
	// heap.UpdateDegree()
	// heap.UpdateN()
	arr := NewArrayList(8, NodeCompareFunc_IntLessThan)
	arr.Append(node23)
	arr.Append(node7)
	arr.Append(node21)
	arr.Append(node3)
	arr.Append(node17)
	arr.Append(node24)
	//因为直接调用heap.Insert会出现堆的形态调整，不方便构造跟书中相同的初始树，所以提供了一个heap.UseTestData函数，方便构造跟书中一样的初始树
	heap.UseTestData(node3, arr)

	// heap.Print()

	heap.ExtractMin()

	// heap.DecreaseKey(node46, 15)
	// heap.DecreaseKey(node35, 5)
	heap.Delete(node46)
	heap.Print()
}

/**
 *  二项堆
 *
 */
func TestBinomialHeap(t *testing.T) {
	compare := NodeCompareFunc_IntLessThan
	heap := NewBinomialHeap(compare)
	num := 100
	nodes := make([]*BinomialHeapNode, num+1)
	for k := 1; k <= num; k++ {
		nodes[k] = NewBinomialHeapNode(k, compare)
	}
	heap.Insert(nodes[8])
	heap.Insert(nodes[10])
	heap.Insert(nodes[50])
	heap.Insert(nodes[100])
	heap.Insert(nodes[80])
	heap.Insert(nodes[25])
	heap.Insert(nodes[15])
	heap.Insert(nodes[5])
	heap.Insert(nodes[30])
	heap.Insert(nodes[40])

	heap.Print("--insert---")
	// heap.ExtractMin()
	//heap.DecreaseKey(nodes[40], 3)
	heap.Delete(nodes[40])
	heap.Print("--delete---")
}
