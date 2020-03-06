/*
 * @Description: golang algorithm
 * @Author: wangchengdg@gmail.com
 * @Date: 2020-02-10 19:54:13
 * @LastEditTime: 2020-03-01 18:35:01
 * @LastEditors:
 */
package SortAlgorithm

import (
	"fmt"
	"testing"

	. "github.com/meshcross/algorithm-3rd/mesh/common"
)

func TestHeapSort(t *testing.T) {
	data1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	data2 := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	data3 := []int{5, 5, 5, 5, 5, 4, 4, 4, 4, 4}
	data4 := []int{5}

	expect1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expect2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expect3 := []int{4, 4, 4, 4, 4, 5, 5, 5, 5, 5}
	expect4 := []int{5}

	sorter := &HeapSort{}
	sorter.Sort(data1)
	EXPECT_EQ(data1, expect1, t)
	//fmt.Println("TestHeapSort1:", data1, expect1)
	sorter.Sort(data2)
	EXPECT_EQ(data2, expect2, t)
	//fmt.Println("TestHeapSort2:", data2, expect2)

	sorter.Sort(data3)
	EXPECT_EQ(data3, expect3, t)
	//fmt.Println("TestHeapSort3:", data3, expect3)
	sorter.Sort(data4)
	EXPECT_EQ(data4, expect4, t)
	//fmt.Println("TestHeapSort4:", data4, expect4)
}

func TestQuickSort(t *testing.T) {
	data1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	data2 := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	data3 := []int{5, 5, 5, 5, 5, 4, 4, 4, 4, 4}
	data4 := []int{5}

	expect1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expect2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expect3 := []int{4, 4, 4, 4, 4, 5, 5, 5, 5, 5}
	expect4 := []int{5}

	sorter := &QuickSort{}
	sorter.Sort(data1)
	EXPECT_EQ(data1, expect1, t)
	sorter.Sort(data2)
	EXPECT_EQ(data2, expect2, t)
	sorter.Sort(data3)
	EXPECT_EQ(data3, expect3, t)
	sorter.Sort(data4)
	EXPECT_EQ(data4, expect4, t)
}

func TestRadixSort(t *testing.T) {
	data1 := []int{5, 0, 50, 0, 3, 301, 5, 2, 1}
	data2 := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}

	expect1 := []int{0, 0, 1, 2, 3, 5, 5, 50, 301}
	expect2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	sorter := &RadixSort{}
	sorter.Sort(data1, 3)
	EXPECT_EQ(data1, expect1, t)
	sorter.Sort(data2, 2)
	EXPECT_EQ(data2, expect2, t)

	fmt.Println("data:", data1, data2)
}

func TestBucketSort(t *testing.T) {
	data1 := []int{123, 234, 345, 456, 567, 678, 789, 890, 901, 912}
	data2 := []int{912, 901, 890, 789, 678, 567, 456, 345, 234, 123}
	data3 := []int{555, 555, 544, 554, 545, 444, 455, 445, 454, 444}
	data4 := []int{555}
	comparedata1 := []int{123, 234, 345, 456, 567, 678, 789, 890, 901, 912}
	comparedata2 := []int{123, 234, 345, 456, 567, 678, 789, 890, 901, 912}
	comparedata3 := []int{444, 444, 445, 454, 455, 544, 545, 554, 555, 555}
	comparedata4 := []int{555}

	sorter := &BucketSort{}
	sorter.Sort(data1, 100, 1000)
	sorter.Sort(data2, 100, 1000)
	sorter.Sort(data3, 100, 1000)
	sorter.Sort(data4, 100, 1000)

	EXPECT_EQ(data1, comparedata1, t)
	EXPECT_EQ(data2, comparedata2, t)
	EXPECT_EQ(data3, comparedata3, t)
	EXPECT_EQ(data4, comparedata4, t)

	// fmt.Println("TestBucketSort1:", data1, comparedata1)
	// fmt.Println("TestBucketSort2:", data2, comparedata2)
	// fmt.Println("TestBucketSort3:", data3, comparedata3)
	// fmt.Println("TestBucketSort4:", data4, comparedata4)
}

func TestMergeSort(t *testing.T) {
	data1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	data2 := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	data3 := []int{5, 5, 5, 5, 5, 4, 4, 4, 4, 4}
	data4 := []int{5}
	comparedata1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	comparedata2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	comparedata3 := []int{4, 4, 4, 4, 4, 5, 5, 5, 5, 5}
	comparedata4 := []int{5}
	sorter := &MergeSort{}
	sorter.Sort(data1)
	sorter.Sort(data2)
	sorter.Sort(data3)
	sorter.Sort(data4)

	EXPECT_EQ(data1, comparedata1, t)
	EXPECT_EQ(data2, comparedata2, t)
	EXPECT_EQ(data3, comparedata3, t)
	EXPECT_EQ(data4, comparedata4, t)

	// fmt.Println("TestMergeSort1:", data1, comparedata1)
	// fmt.Println("TestMergeSort2:", data2, comparedata2)
	// fmt.Println("TestMergeSort3:", data3, comparedata3)
	// fmt.Println("TestMergeSort4:", data4, comparedata4)
}
