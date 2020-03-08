/*
 * @Description: 第五部分，高级数据结构

 * @Author: wangchengdg@gmail.com
 * @Date: 2020-03-05 11:37:55
 * @LastEditTime: 2020-03-08 16:29:40
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
