/*
 * @Description: 贪心算法
	贪心算法做出一系列选择来求出问题的最优解。在每个决策点，它做出在当时看来最佳的选择。这种启发式策略并不能保证总能找到最优解，但是对一些问题确实有效。
	贪心算法是局部最优，不一定是全局最优。
	动态规划算法每个步骤都做出最优选择，而且后一步依赖于前一步；而贪心算法只顾当下最优，不依赖前后。
 * @Author: wangchengdg@gmail.com
 * @Date: 2020-03-03 13:12:10
 * @LastEditTime: 2020-03-05 12:02:53
 * @LastEditors:
*/

package GreedyAlgorithm

import (
	. "github.com/meshcross/algorithm-3rd/mesh/common"
)

type ActivitySelector struct{}

func NewActivitySelector() *ActivitySelector {
	return &ActivitySelector{}
}

/**
 * @description:
 * @param acts 活动列表，包含开始时间和结束时间,输入的acts是按照结束时间排序好的,First为开始时间，Second为结束时间
 * @param k
 * @param n
 * @return:选中的活动序号
 */
func (a *ActivitySelector) Run(acts []*Pair, k, n int) []int {
	m := k + 1
	//进入for循环表示未被选中，不加等号，表示一个后一个活动的开始时间可以和前一个活动的结束时间相等
	for m < n && acts[m].First <= acts[k].Second {
		m = m + 1
	}

	if m < n {
		x := []int{m}
		more := a.Run(acts, m, n)
		return append(x, more...)
	}
	return []int{}
}
