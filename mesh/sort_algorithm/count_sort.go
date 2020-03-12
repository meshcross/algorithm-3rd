/*
 * @Description: 第8章 8.2 计数排序
	计数排序假设输入数据都属于小区间内的整数

 * @Author: wangchengdg@gmail.com
 * @Date: 2020-02-10 13:48:23
 * @LastEditTime: 2020-03-11 18:13:14
 * @LastEditors:
*/
package SortAlgorithm

type CountSort struct {
}

/**
k为数组s中的最大元素
计数排序假设输入数据都属于小区间内的整数

假设n个输入元素的每一个元素都是0到k区间内的一个整数，其中k为某个整数。
**/
func (a *CountSort) Sort(s []int, k int) {
	size := len(s)
	max_val := k
	if size <= 1 {
		return
	}

	counter_arr := make([]int, max_val+1) //存放计数结果
	result_arr := make([]int, size)       //暂存排序结果

	n := 0
	for n < size {
		counter_arr[s[n]]++
		n++
	}

	counter_size := len(counter_arr)
	for i := 1; i < counter_size; i++ {
		counter_arr[i] += counter_arr[i-1]
	}

	index := size - 1

	for index >= 0 {
		data := s[index]                       //待排序的元素
		less_data_num := counter_arr[data] - 1 //比它小的元素的个数
		result_arr[less_data_num] = data       //直接定位
		counter_arr[data]--                    //此行为了防止重复元素的定位
		index--
	}

	//result_arr即为计算所得的结果，这里直接根据result_arr修改原始数组s
	result_size := len(result_arr)
	for k := 0; k < result_size; k++ {
		s[k] = result_arr[k]
	}

}
