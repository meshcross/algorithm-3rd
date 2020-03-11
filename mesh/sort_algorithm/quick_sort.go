/*
 * @Description: 算法导论第7章 快速排序
		操作1：选取数组s最后一个元素作为参考x，将<x的值放到左边，>x的值放到右边，然后返回x在数组中的索引位置pos，然后执行操作2
		操作2：pos左边和右边各得到一个数组，对新数组重复进行操作1，直到拆分出来的子数组元素个数<=1
 * @Author: wangchengdg@gmail.com
 * @Date: 2020-02-10 13:49:03
 * @LastEditTime: 2020-03-10 16:40:35
 * @LastEditors:
*/
package SortAlgorithm

import (
	. "github.com/meshcross/algorithm-3rd/mesh/common"
)

func NewQuickSort() *QuickSort {
	return &QuickSort{}
}

type QuickSort struct {
}

/**
 * @description: 快速排序
 * @param s:数据源
 * @param compares :比较函数，可选；默认使用LessThan,如果使用BiggerThan，会得到顺序相反的结果
 * @return:
 */
func (a *QuickSort) Sort(s []int, compares ...SortCompareFunc) {
	size := len(s)
	if size <= 1 {
		return
	}
	start := 0
	end := size

	var compare SortCompareFunc = nil
	if len(compares) > 0 {
		compare = compares[0]
	} else {
		compare = LessThan
	}

	// 以最后一个元素作为划分元素，返回其顺序统计量
	p := a.partition(s, start, end, end-1, compare)
	a.SortPiece(s, start, p, compare)
	a.SortPiece(s, p+1, end, compare)
}

/**
 * @description: 对s的某一个区间s[start,end]进行排序
 * @param s:数据源
 * @param start:开始下标
 * @param end:截止下标
 * @param compare:比较函数
 * @return: void
 */
func (a *QuickSort) SortPiece(s []int, start, end int, compare SortCompareFunc) {
	size := end - start
	if size <= 1 {
		return
	}

	if start >= end-1 || end > len(s) {
		return
	}
	p := a.partition(s, start, end, end-1, compare) // 以最后一个元素作为划分元素，返回其顺序统计量
	a.SortPiece(s, start, p, compare)
	a.SortPiece(s, p+1, end, compare)
}

/**
 * @description: partition不能是数组额最后一个元素
			start start>=0,end为结束符的位置，end为1到size
			该函数主要作用是选取partition位置的元素作为参考x，将<x的值放到左边，>x的值放到右边，然后返回x在数组中的索引位置
 * @param s:数据源
 * @param begin:开始下标
 * @param end:截止下标
 * @param partition:分隔s[begin,end]的参考值
 * @param compare:比较函数
 * @return: void
*/
func (a *QuickSort) partition(s []int, begin, end, partition int, compare SortCompareFunc) int {

	size := end - begin

	if size <= 0 {
		return size - 1
	}
	if size == 1 {
		return begin
	}

	partition_ok := partition >= begin && end > partition
	if !partition_ok {
		panic("partition value error!")
	}

	smaller_next := begin - 1 //指向比key小的元素区间的下一个(即大于等于key元素区间的第一个），其中key为序列最后一个元素
	current := begin          //指向当前待处理的元素

	for current != end-1 {
		if compare(s[current], s[end-1]) > 0 {
			smaller_next++
			Swap(s, smaller_next, current)
		}
		current++
	}
	Swap(s, smaller_next+1, end-1) //交换partition元素到它的位置
	return smaller_next + 1
}

/********        以上是对[]int进行排序，以下是对任意的对象进行快速排序          **********/

/**
 * @description: 任意数组的快速排序，s中可以是int，也可以是各种指针，只要NodeCompareFunc能比较即可
 * @param
 * @return:
 */
func (a *QuickSort) SortAny(s []interface{}, compares ...NodeCompareFunc) {
	size := len(s)
	if size <= 1 {
		return
	}
	start := 0
	end := size

	var compare NodeCompareFunc = nil
	if len(compares) > 0 {
		compare = compares[0]
	} else {
		compare = NodeCompareFunc_IntLessThan
	}

	// 以最后一个元素作为划分元素，返回其顺序统计量
	p := a.partitionAny(s, start, end, end-1, compare)
	a.SortPieceAny(s, start, p, compare)
	a.SortPieceAny(s, p+1, end, compare)
}

/**
 * @description: 对s的某一个区间s[start,end]进行排序
 * @param s:数据源
 * @param start:开始下标
 * @param end:截止下标
 * @param compare:比较函数
 * @return: void
 */
func (a *QuickSort) SortPieceAny(s []interface{}, start, end int, compare NodeCompareFunc) {
	size := end - start
	if size <= 1 {
		return
	}

	if start >= end-1 || end > len(s) {
		return
	}
	p := a.partitionAny(s, start, end, end-1, compare) // 以最后一个元素作为划分元素，返回其顺序统计量
	a.SortPieceAny(s, start, p, compare)
	a.SortPieceAny(s, p+1, end, compare)
}

/**
 * @description: partition不能是数组额最后一个元素
			start start>=0,end为结束符的位置，end为1到size
			该函数主要作用是选取partition位置的元素作为参考x，将<x的值放到左边，>x的值放到右边，然后返回x在数组中的索引位置
 * @param s:数据源
 * @param begin:开始下标
 * @param end:截止下标
 * @param partition:分隔s[begin,end]的参考值
 * @param compare:比较函数
 * @return: void
*/
func (a *QuickSort) partitionAny(s []interface{}, begin, end, partition int, compare NodeCompareFunc) int {

	size := end - begin

	if size <= 0 {
		return size - 1
	}
	if size == 1 {
		return begin
	}

	partition_ok := partition >= begin && end > partition
	if !partition_ok {
		panic("partition value error!")
	}

	smaller_next := begin - 1 //指向比key小的元素区间的下一个(即大于等于key元素区间的第一个），其中key为序列最后一个元素
	current := begin          //指向当前待处理的元素

	for current != end-1 {
		if compare(s[current], s[end-1]) > 0 {
			smaller_next++
			SwapAny(s, smaller_next, current)
		}
		current++
	}
	SwapAny(s, smaller_next+1, end-1) //交换partition元素到它的位置
	return smaller_next + 1
}
