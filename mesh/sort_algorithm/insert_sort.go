/*
 * @Description: golang algorithm
 * @Author: wangchengdg@gmail.com
 * @Date: 2020-02-10 13:48:15
 * @LastEditTime: 2020-03-01 18:34:44
 * @LastEditors:
 */
package SortAlgorithm

import . "github.com/meshcross/algorithm-3rd/mesh"

//!insert_sort：算法导论第二章 2.1 插入排序
/*!

 * - 插入排序思想，假设对数组A[p...r]排序：
 *   - 维持不变式：设当前排序的元素是 A[q]，则保持A[p...q-1]为排好的，A[q]在A[p...q-1]中找到它的位置坐下
 * - 时间复杂度 O(n^2)
 * - 原地排序
 */
type InsertSort struct {
}

func (a *InsertSort) Sort(s []int, compares ...SortCompareFunc) {
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

	current := 0

	for current != size {
		small_next := current                                             //指向比*current小的元素中最大的那个元素
		for small_next != 0 && compare(s[current], s[small_next-1]) > 0 { //current较小，则向前寻找插入的位置插入
			small_next--
		}

		key := s[current]
		iter := current
		for iter != small_next { //插入
			s[iter] = s[iter-1] //后移
			iter--
		}
		s[iter] = key
		current++
	}

}
