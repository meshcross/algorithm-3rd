/*
 * @Description: 定义公共函数
 * @Author: wangcheng
 * @Date: 2020-02-11 19:06:11
 * @LastEditTime: 2020-03-06 12:59:10
 * @LastEditors:
 */

package Common

import (
	"encoding/binary"
	"math"
	"reflect"
	"testing"
)

func EXPECT_EQ(v1 interface{}, v2 interface{}, t *testing.T) {
	if !reflect.DeepEqual(v1, v2) {
		t.Errorf("error")
	}
}

func Revert(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func SwapAny(s []interface{}, index1, index2 int) {
	l := len(s)
	if index1 < 0 || index1 >= l || index2 < 0 || index2 >= l {
		return
	}
	x1 := s[index1]
	x2 := s[index2]
	tmp := x1
	s[index1] = x2
	s[index2] = tmp
}

func Swap(s []int, index1, index2 int) {
	l := len(s)
	if index1 < 0 || index1 >= l || index2 < 0 || index2 >= l {
		return
	}
	x1 := s[index1]
	x2 := s[index2]
	tmp := x1
	s[index1] = x2
	s[index2] = tmp
}

//整数num的第n位
func Digit_On_N(num, n int) int {
	p1 := math.Pow10(n)
	p2 := math.Pow10(n + 1)
	return num/int(p1) - num/int(p2)*10
}

//n为1-64，v为0或者1
func SetUintBit(num uint64, n int, v int) uint64 {

	if v > 0 {
		ret := num | 1<<(n-1)
		return ret
	} else {
		//高位不会溢出
		hight := (num >> n) << n
		dn := 64 - n + 1
		x := (num << dn)
		low := x >> dn
		ret := hight | low
		return ret
	}
}

func Int64ToBytes(i int64) []byte {
	var buf = make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(i))
	return buf
}

func BytesToInt64(buf []byte) int64 {
	return int64(binary.BigEndian.Uint64(buf))
}
