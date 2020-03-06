/*
 * @Description: golang algorithm
 * @Author: wangchengdg@gmail.com
 * @Date: 2020-02-10 20:09:29
 * @LastEditTime: 2020-03-01 18:34:16
 * @LastEditors:
 */
package SelectAlgorithm

import (
	"fmt"
	"testing"
)

func TestTask(t *testing.T) {

	selector := &RandomizedSelect{}

	for i := 0; i < 10; i++ {
		data1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		val, _ := selector.Select(data1, i)
		fmt.Println(fmt.Sprintf("EXPECT_EQ(%d,%d)<<<<%d", val, i+1, i))
	}
	for i := 0; i < 10; i++ {
		data2 := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
		val, _ := selector.Select(data2, i)
		fmt.Println(fmt.Sprintf("EXPECT_EQ(%d,%d)<<<<%d", val, i+1, i))
	}
	{
		data3 := []int{5, 5, 5, 5, 5, 4, 4, 4, 4, 4}
		val, _ := selector.Select(data3, 0)
		fmt.Println(fmt.Sprintf("EXPECT_EQ(%d,%d)", val, 4))
	}
	{
		data3 := []int{5, 5, 5, 5, 5, 4, 4, 4, 4, 4}
		val, _ := selector.Select(data3, 9)
		fmt.Println(fmt.Sprintf("EXPECT_EQ(%d,%d)", val, 5))
	}
	{
		data4 := []int{5}
		val, _ := selector.Select(data4, 0)
		fmt.Println(fmt.Sprintf("EXPECT_EQ(%d,%d)", val, 5))
	}
	fmt.Println("x")
}
