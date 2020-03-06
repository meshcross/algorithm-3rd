/*
 * @Description: 动态规划算法有两个关键要素：寻找最优子结构和子问题重叠。
	如果一个问题的最优解，包含其子问题的最优解，则此问题具有最优子结构性质；
	如果递归算法反复求解相同的子问题，则称为子问题重叠。
	找到循环前后回合之间的联系，通过[i,j]滑动窗口从[1,1]拉开到[1,n]，从而获得[1,n]范围内的最优解
 * @Author: wangchengdg@gmail.com
 * @Date: 2020-02-10 20:09:34
 * @LastEditTime: 2020-03-05 12:17:23
 * @LastEditors:
*/
package DynamicProgrammingAlgorithm

import (
	"fmt"
	"testing"

	. "github.com/meshcross/algorithm-3rd/mesh/common"
)

func TestKnapsac(t *testing.T) {

	p1 := []*Pair{
		// &Pair{First: 5, Second: 10},
		&Pair{First: 2, Second: 5},
		&Pair{First: 3, Second: 7},
		&Pair{First: 5, Second: 10},
		&Pair{First: 8, Second: 16}, //15
		&Pair{First: 10, Second: 20},
	}

	p2 := []*Tuple{
		// &Pair{First: 5, Second: 10},
		&Tuple{First: 2, Second: 5, Third: 3},
		&Tuple{First: 3, Second: 7, Third: 3},
		&Tuple{First: 5, Second: 10, Third: 3},
		&Tuple{First: 8, Second: 16, Third: 3},
		&Tuple{First: 10, Second: 20, Third: 3},
	}

	bag := NewKnapsack()
	fmt.Println(bag.PackComplete(p1, 11))
	fmt.Println(bag.PackMultiple(p2, 11))
}

func TestMatrixChainOrder(t *testing.T) {
	p := []int{
		30, 35, 15, 5, 10, 20, 25,
	}

	chain := NewMatrixChainOrder()
	chain.Run(p)

	//输出：((A1(A2A3))((A4A5)A6))
}

func TestLongestCommonSequence(t *testing.T) {
	str_a := "-123bxz"
	str_b := "123|xyz"
	// str_c := "123abcdefgbb"
	// str_d := "66123abcdefgbb66"
	// str_e := ""
	// str_f := "a9"
	// str_g := "123"
	// str_h := "cdef"
	lcs := NewLongestCommonSubsequence()
	ret := lcs.Run(str_a, str_b, true)
	EXPECT_EQ(ret, "123xz", t)
}

func TestOptimalBinarySearchTree(t *testing.T) {
	worker := NewOptimalBinarySearchTree()
	p := []int{0, 15, 10, 5, 10, 20}
	q := []int{5, 10, 5, 5, 5, 10}
	// q := []int{0, 0, 0, 0, 0, 0}
	expect, weight, root := worker.RunAsBook(p, q, 5)
	fmt.Println("expect:")
	for _, v := range expect {
		fmt.Println(v)
	}
	fmt.Println("weight:")
	for _, v := range weight {
		fmt.Println(v)
	}
	fmt.Println("root:")
	for _, v := range root {
		fmt.Println(v)
	}

	// fmt.Println(expect[1][5], weight[1][5], root[1][5])

	//两个元素相同，顺序不同的数组p传入，会有不同的结果！！！
	px1 := []int{3, 2, 1, 2, 4}
	// px2 := []int{1, 2, 4, 2, 3}
	num, cs, rs := worker.Run(px1)
	fmt.Println("num:", num)
	fmt.Println("cs:", cs)
	fmt.Println("rs:", rs)
	fmt.Println("root:", rs[1][5])
}
