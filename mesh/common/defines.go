/*
 * @Description: 一些回调函数定义
 * @Author: wangchengdg@gmail.com
 * @Date: 2020-03-08 17:08:49
 * @LastEditTime: 2020-03-08 17:11:05
 * @LastEditors:
 */

package Common

type SortCompareFunc func(x, y int) int
type NodeCompareFunc func(x, y interface{}) int

func NodeCompareFunc_IntBiggerThan(x, y interface{}) int {

	if ToInt(x) > ToInt(y) {
		return 1
	}
	if x == y {
		return 0
	}
	return -1
}
func NodeCompareFunc_IntLessThan(x, y interface{}) int {

	if ToInt(x) < ToInt(y) {
		return 1
	}
	if x == y {
		return 0
	}
	return -1
}
func NodeCompareFunc_FloatLessThan(x, y interface{}) int {
	if ToFloat64(x) < ToFloat64(y) {
		return 1
	}
	if x == y {
		return 0
	}
	return -1
}
func LessThan(x, y int) int {
	if x < y {
		return 1
	}
	if x == y {
		return 0
	}
	return -1
}

func BigThan(x, y int) int {
	if x < y {
		return 1
	}
	if x == y {
		return 0
	}
	return 1
}
