/*
 * @Description: golang algorithm
 * @Author: wangchengdg@gmail.com
 * @Date: 2020-02-10 13:48:40
 * @LastEditTime: 2020-03-10 16:35:20
 * @LastEditors:
 */
package SortAlgorithm

import (
	. "github.com/meshcross/algorithm-3rd/mesh/common"
)

//! merge：算法导论第二章 2.3.1 分治算法
/*!

* - 归并思想，假设对数组A[p...q...r]归并：
*   - 拷贝：将数组A[p...q]拷贝到数组L，将数组A[q...r]拷贝到数组R，
*   - 归并： 从左到右依次取L、R中的较小的元素，存放到A中（具体算法见代码）
* - 时间复杂度 O(n)
* - 归并时需要额外的空间 O(n)
 */

type MergeSort struct {
}

/**

**/
func (a *MergeSort) Sort(s []int, compares ...SortCompareFunc) {
	size := len(s)
	if size <= 1 {
		return
	}
	var compare SortCompareFunc = nil
	if len(compares) > 0 {
		compare = compares[0]
	} else {
		compare = LessThan
	}

	end := size
	a.sortPiece(s, 0, end, compare)
}

func (a *MergeSort) sortPiece(s []int, begin, end int, compare SortCompareFunc) {
	size := end - begin
	if size <= 1 {
		return
	}

	middle := size/2 + begin
	a.sortPiece(s, begin, middle, compare)
	a.sortPiece(s, middle, end, compare)
	a.merge(s, begin, end, middle, compare)
}

func (a *MergeSort) merge(s []int, begin, end, middle int, compare SortCompareFunc) {

	if middle <= begin || end <= middle {
		return
	}
	result := make([]int, end-begin) //暂存结果
	current := 0                     //result的游标
	left_current := begin            //左侧序列当前比较位置
	right_current := middle          //右序列当前比较位置

	for left_current != middle && right_current != end {
		if compare(s[left_current], s[right_current]) > 0 {
			result[current] = s[left_current] //左侧较小
			left_current++
		} else {
			result[current] = s[right_current] //左侧较小
			right_current++
		}
		current++
	}

	if left_current == middle && right_current != end { //当左侧序列为空
		for k := right_current; k < end; k++ {
			result[current] = s[k]
			current++
		}
	}
	if right_current == end && left_current != middle { //当右侧序列为空

		for k := left_current; k < middle; k++ {
			result[current] = s[k]
			current++
		}
	}

	for k := begin; k < end; k++ {
		s[k] = result[k-begin]
	}
}
