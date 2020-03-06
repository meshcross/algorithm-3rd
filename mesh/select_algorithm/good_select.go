/*
 * @Description: 算法导论第9章9.3 最坏时间为O(n)的顺序统计量选择算法
			选择算法思想，假设对数组A[p...r]选择，选择第k小的元素：
*      				- 选择主元：
*          				- 首先将序列从前往后，按照5个元素一组分组。其中最后一组可能为1～5个元素（也就是可能不是满的）
*          				- 然后将这些分组进行排序（我采用的是快速排序）
*          				- 然后将这些分组中的中位数（即最中间的数）取出
*          				- 针对这些分组的中位数构成的序列，递归调用 good_select，找出中位数的中位数
*          				- 将这个中位数的中位数作为划分主元
*      				- 划分：根据找到的主元对原序列进行划分，假设划分完毕后，主元是第m小
*      				- 判定：
*          				- 若m==k，则找到了这个元素，返回这个主元
*          				- 若m<k ，则说明指定的元素在 A[m+1...r]中，且位于这个新数组的第(k-m-1)小，此时递归调用good_select(....)
*          				- 若m>k， 则说明指定的元素在 A[p...m-1]中，且位于这个新数组的第 k 小，此时递归调用good_select(...)
* 			- 时间复杂度：最坏情况下时间为O(n)
* 			- 非原地操作：因为这里要把所有分组的中位数构造成一个序列，然后把找到该序列的中位数作为good_select(...)的主元
 * @Author: wangchengdg@gmail.com
 * @Date: 2020-02-10 22:17:35
 * @LastEditTime: 2020-03-06 12:36:28
 * @LastEditors:
*/
package SelectAlgorithm

import (
	"errors"

	. "github.com/meshcross/algorithm-3rd/mesh"
	. "github.com/meshcross/algorithm-3rd/mesh/common"
	SortAlgorithm "github.com/meshcross/algorithm-3rd/mesh/sort_algorithm"
)

type GoodSelect struct {
}

/**
 * @description: 从数组s中找到第rank小的元素,rank从0开始
			数组s可以改为s[]interface{}，只需要使用对应的SortCompareFunc即可
 * @param s:数据源
 * @param rank:第rank小
 * @param compares:比较函数
 * @return: 第rank小的元素下标;error
*/
func (a *GoodSelect) Select(s []int, rank int, compares ...SortCompareFunc) (int, error) {

	var compare SortCompareFunc = nil
	if len(compares) > 0 {
		compare = compares[0]
	} else {
		compare = LessThan
	}
	begin := 0
	end := len(s)
	return a._select(s, begin, end, rank, compare)
}

func (a *GoodSelect) _select(s []int, begin, end, rank int, compare SortCompareFunc) (int, error) {

	size := end - begin
	if size <= 0 {
		return 0, errors.New("size is zero")
	}
	if rank > size {
		return 0, errors.New("rank is small")
	}

	if size == 1 {
		return s[begin], nil
	}

	SPAN := 5
	quick_sort := SortAlgorithm.QuickSort{}
	// *********  将序列划分为5个元素一组的区间，最后一组可能不满5个元素；对每组进行排序，取出其中值放入slice中  *********
	iter := begin
	middle_nums := []int{}
	for iter < end {
		from := iter
		to := iter + SPAN
		if to > end {
			to = end
		}

		sub := s[from:to]
		quick_sort.Sort(sub, compare)
		sub_size := len(sub)
		middle_nums = append(middle_nums, sub[(sub_size-1)/2])
		iter += SPAN
	}
	// ********* 取出这些中值的中值,如果有偶数个，默认选择较小的 ************
	mid_num_size := len(middle_nums)
	mid_of_middles, _ := a._select(middle_nums, 0, mid_num_size, (mid_num_size-1)/2, compare) //所有中值的中值
	iter = begin
	for end-iter > 0 && s[iter] != mid_of_middles { //得到中值的中值在原序列中的位置
		iter++
	}
	//********* 划分 **************
	mid_of_middles_iter := a.partition(s, begin, end, iter, compare) //以中值的中值作为一个划分值
	// ********** 判别 ***************
	mid_of_middles_rank := mid_of_middles_iter - begin //中值的中值在划分之后的排序

	if mid_of_middles_rank == rank { //找到了该排位的数
		return s[mid_of_middles_iter], nil
	} else if mid_of_middles_rank < rank { //目标排位在右侧
		return a._select(s, mid_of_middles_iter+1, end, rank-mid_of_middles_rank-1, compare) //mid_of_middles_iter+1，则找右侧的第rank-mid_of_middles_rank-1位
	} else { //目标排位在左侧
		return a._select(s, begin, mid_of_middles_iter, rank, compare)
	}
}

/**
 * @description: begin从0开始到size-1
				end为1到size，c++中为结束符
				以partition为参考值，<partition的放在左边，>partition的放在右边
 * @param s :数据源
 * @param begin:开始下标>=0
 * @param end:截止下标<=size-1
 * @param partition:将s[begin,end]分隔为两部分的参考值
 * @param compare:比较函数
 * @return:分隔点(下标)
*/
func (a *GoodSelect) partition(s []int, begin, end, partition int, compare SortCompareFunc) int {
	size := end - begin

	if size == 0 {
		return end - 1 //空序列，返回end
	}
	if size == 1 {
		return begin
	}

	partition_ok := partition >= begin && end > partition
	if !partition_ok {
		panic("partition value error!")
	}

	smaller_next := begin //指向比key小的元素区间的下一个(即大于等于key元素区间的第一个），其中key为序列最后一个元素
	current := begin      //指向当前待处理的元素

	for current != end-1 {
		if compare(s[current], s[end-1]) > 0 {
			Swap(s, smaller_next, current)
			smaller_next++
		}
		current++
	}
	Swap(s, smaller_next, end-1) //交换partition元素到它的位置
	return smaller_next
}
