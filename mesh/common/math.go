/*
 * @Description: 数学相关函数定义
 * @Author: wangchengdg@gmail.com
 * @Date: 2020-02-28 11:52:38
 * @LastEditTime: 2020-03-06 12:59:22
 * @LastEditors:
 */
package Common

import "math"

func PowInt(x, y int) int {
	r := math.Pow(float64(x), float64(y))
	return int(r)
}
func MaxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func MinInt(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func MaxInt64(x, y int64) int64 {
	if x > y {
		return x
	}
	return y
}
func MinInt64(x, y int64) int64 {
	if x < y {
		return x
	}
	return y
}

func NewMatrix(n int, value int) [][]int {
	arr := make([][]int, n)
	for i := 0; i < n; i++ {
		arr[i] = make([]int, n)
		for k, _ := range arr[i] {
			arr[i][k] = value
		}
	}
	return arr
}

const INT_MAX = int(^uint(0) >> 1)

const UINT_MIN uint = 0
const UINT_MAX = ^uint(0)

/*!
*@return : 正无穷大的数
*
* 用一个很大的数当做无穷大
*
 */
func Unlimit() int {
	return INT_MAX / 2
}

/*!
* @description:判断是否正无穷，这里实现并不是很严谨，一个正无穷的值减去常数还是正无穷
* @param t: 待判断的数
* @return : 如果该数是正无穷大，则返回`true`，否则返回`false`
*
 */
func Is_Unlimit(t int) bool {
	return t >= INT_MAX/3
}
