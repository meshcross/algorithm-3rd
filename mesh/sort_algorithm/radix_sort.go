/*
 * @Description: 第8章 8.3 基数排序
 *
* - 基数排序思想，假设对数组A[p...r]排序，其中数组中所有元素都为正整数，并且不超过RADIXWITH位
*   - 首先对A中所有元素按照个位数大小进行排序（原地的）
*   - 再对A中所有元素按照十位数大小进行排序（原地的）
*   - 一直到最后按照A中所有元素的最高位的数字大小进行排序（原地的）
*
* - 时间复杂度 O(d(n+k))，其中d位数字的最大位宽(即这里都是d位数的整数），k为每个位数上数字取值（这里取0，1，2，3，...9）
* - 原地排序
*
*  这里尤其要重点强调，用于对指定位上的数字进行排序时，必须要满足稳定性。
*  - 快速排序就是非稳定的
*  - 用小于比较的插入排序是稳定的；用小于等于比较的插入排序是不稳定的
*
* >这里必须对整数才能采取基数排序
*
 * @Author: wangchengdg@gmail.com
 * @Date: 2020-02-10 13:48:48
 * @LastEditTime: 2020-03-11 22:12:27
 * @LastEditors:
*/
package SortAlgorithm

import (
	. "github.com/meshcross/algorithm-3rd/mesh/common"
)

type RadixSort struct {
}

/**
radix_width指定的位数（0表示个位，1表示十位，...）
**/
func (a *RadixSort) Sort(s []int, radix_width int) {
	size := len(s)
	if size <= 1 {
		return
	}

	sorter := &InsertSort{}
	for i := 0; i < radix_width; i++ { //从最低位(第0位为个位）到 （RADIXWITH-1）位的位数进行排序（一共RADIXWITH位）
		pos := i
		compareLess := func(x, y int) int {
			x1 := Digit_On_N(x, pos)
			y1 := Digit_On_N(y, pos)

			if x1 < y1 {
				return 1
			} else if x1 == y1 {
				return 0
			}
			return -1
		}
		sorter.Sort(s, compareLess)
	}
}
