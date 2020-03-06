/*
 * @Description: golang algorithm
 * @Author: wangchengdg@gmail.com
 * @Date: 2020-02-10 13:26:43
 * @LastEditTime: 2020-03-05 12:17:14
 * @LastEditors:
 */
package mesh

import (
	. "github.com/meshcross/algorithm-3rd/mesh/common"
)

type IAlgorithm interface {
	GetName() string
}

type SortCompareFunc func(x, y int) int
type NodeCompareFunc func(x, y interface{}) int

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
