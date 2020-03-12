/*
 * @Description: 第6章 堆排序
 					原地排序
 * @Author: wangchengdg@gmail.com
 * @Date: 2020-02-10 13:47:55
 * @LastEditTime: 2020-03-11 18:11:04
 * @LastEditors:
*/
package SortAlgorithm

import (
	. "github.com/meshcross/algorithm-3rd/mesh/common"
)

/*!
 *
 * - 堆排序思想：假设对数组A[p...r]排序：首先将数组构建成一个最大堆（或者最小堆，这里的实现采用最大堆）。然后第一个元素就是堆中最大的元素。
 * 将第一个元素与最后一个元素交换，同时堆的规模缩减1，再将堆维持最大堆性质。不断循环最后得到一个排序好的序列
 * - 时间复杂度 O(nlogn)
 * - 原地排序
 *
 * 堆排序有两个重要操作：
 *
 * - heapify(index)操作：维持以index为根节点的子堆的性质。它比较index与其左右子节点的值，选取其最大的那个提升到index节点上。同时递归向下。具体见_heapify()方法说明
 * - setupHeap()操作： 建堆操作。它从堆的最低层向上层反复调用heapify操作进行建堆。
 */

type HeapSort struct {
	heap_size int
}

func (a *HeapSort) Sort(s []int, compares ...SortCompareFunc) {
	_size := len(s)
	if _size <= 1 {
		return
	}
	var compare SortCompareFunc = nil
	if len(compares) > 0 {
		compare = compares[0]
	} else {
		compare = LessThan
	}

	a.heap_size = _size

	a._setupHeap(s, compare)

	//把堆顶的元素放到数组最后，然后将heap_size减小1，重新_heapify将剩余元素的最大值再次推到新堆的堆顶
	for a.heap_size > 0 {

		Swap(s, 0, a.heap_size-1)

		a.heap_size--
		a._heapify(0, s, compare)
	}
	//Revert(s)
}

func (a *HeapSort) _setupHeap(s []int, compare SortCompareFunc) {
	_size := a.heap_size
	if _size <= 1 {
		return
	}

	index := (_size - 1) / 2
	//此处有隐藏信息：index + 1到 _size都是叶子节点
	for index >= 0 {
		a._heapify(index, s, compare)
		index--
	}
}

/*
会把堆中的最大值放到堆顶
*/
func (a *HeapSort) _heapify(elementIndex int, s []int, compare SortCompareFunc) {
	_size := a.heap_size

	if elementIndex >= _size {
		return
	}

	maxIndex := elementIndex
	left_valid := true
	right_valid := true
	leftIndex := a._lchildIndex(elementIndex, &left_valid)
	rightIndex := a._rchildIndex(elementIndex, &right_valid)

	if left_valid {
		if compare(s[maxIndex], s[leftIndex]) > 0 {
			maxIndex = leftIndex
		}
	}
	if right_valid {
		if compare(s[maxIndex], s[rightIndex]) > 0 {
			maxIndex = rightIndex
		}
	}
	if maxIndex != elementIndex {

		Swap(s, elementIndex, maxIndex)
		a._heapify(maxIndex, s, compare)
	}
}

func (a *HeapSort) _parentIndex(elementIndex int, valid *bool) int {
	_size := a.heap_size
	if elementIndex >= _size {
		*valid = false //无效结果
		return 0
	}

	*valid = true //有效结果
	if elementIndex == 0 {
		return 0 //根节点的父节点是自己
	} else {
		return (elementIndex - 1) >> 1
	}
}

func (a *HeapSort) _lchildIndex(elementIndex int, valid *bool) int {
	_size := a.heap_size
	if _size < 2 {
		*valid = false //数组元素太少无效结果
		return 0
	}
	if elementIndex > ((_size - 2) >> 1) {
		*valid = false //超出范围，无效
		return 0
	}
	*valid = true
	return (elementIndex << 1) + 1
}

func (a *HeapSort) _rchildIndex(elementIndex int, valid *bool) int {
	_size := a.heap_size
	if _size < 3 {
		*valid = false //数组元素太少无效结果
		return 0
	}
	if elementIndex > ((_size - 3) >> 1) {
		*valid = false //超出范围，无效
		return 0
	}
	*valid = true
	return (elementIndex << 1) + 2
}

// func NewHeapSort() *HeapSort {
// 	return &HeapSort{name: "heap_sort"}
// }
