/*
 * @Description: 算法导论第9章9.2 顺序统计量的随机选择算法
 		- 选择算法思想，假设对数组A[p...r]选择，选择第k小的元素：
*      		- 随机选择主元：随机选取数组的一个下标q,A[q]作为划分元素
*      		- 划分：利用A[q]划分数组，获得A[q]在序列中是第 m 小
*      		- 判定：
*          		- 若m==k，则找到了这个元素，返回  A[q]
*          		- 若m<k ，则说明指定的元素在 A[q+1...r]中，且位于这个新数组的第(k-m-1)小，此时递归调用randomized_select(q+1,end,k-m-1)
*          		- 若m>k， 则说明指定的元素在 A[p...q-1]中，且位于这个新数组的第 k 小，此时递归调用randomized_select(begin,q,k)
* 		- 时间复杂度：最坏情况下为O(n^2)，期望时间为O(n)
* 		- 原地操作
 * @Author: wangchengdg@gmail.com
 * @Date: 2020-02-10 22:17:43
 * @LastEditTime: 2020-03-06 12:18:56
 * @LastEditors:
*/
package SelectAlgorithm

import (
	"errors"
	"math/rand"

	. "github.com/meshcross/algorithm-3rd/mesh"
	. "github.com/meshcross/algorithm-3rd/mesh/common"
)

type RandomizedSelect struct {
}

/*!
* @description:顺序统计量的随机选择算法
* @param s :数据源
* @param begin:开始下标>=0
* @param end:截止下标<=size-1
* @param rank: 指定选取的顺序数，0为选取最小的元素，1为选取次小的元素....n表示选取排序为n的元素（从小排序）
* @param compare: 比较函数，x:=compare(v1,v2)，v1<v2则x=1，v1=v2则x=0，v1>v2则x=-1
* @return 第rank小的元素值
*
* - 选择算法思想，假设对数组A[p...r]选择，选择第k小的元素：
*      - 随机选择主元：随机选取数组的一个下标q,A[q]作为划分元素
*      - 划分：利用A[q]划分数组，获得A[q]在序列中是第 m 小
*      - 判定：
*          - 若m==k，则找到了这个元素，返回  A[q]
*          - 若m<k ，则说明指定的元素在 A[q+1...r]中，且位于这个新数组的第(k-m-1)小，此时递归调用randomized_select(q+1,end,k-m-1)
*          - 若m>k， 则说明指定的元素在 A[p...q-1]中，且位于这个新数组的第 k 小，此时递归调用randomized_select(begin,q,k)
* - 时间复杂度：最坏情况下为O(n^2)，期望时间为O(n)
* - 原地操作
 */
func (a *RandomizedSelect) randomized_select(s []int, begin, end, rank int, compare SortCompareFunc) (int, error) {

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

	radom := rand.Intn(size - 1)
	partitioned_iter := a.partition(s, begin, end, begin+radom, compare) //随机划分
	distance := partitioned_iter - begin

	if distance == rank { //找到了该排位的数
		return s[partitioned_iter], nil
	} else if distance < rank { //已知某排位的数位次较低，则指定排位数的元素在它右侧
		return a.randomized_select(s, partitioned_iter+1, end, rank-distance-1, compare) //从partitioned_iter+1，则找右侧的第rank-distance-1位
	} else { //已知某排位的数位次较高，则指定排位数的元素在它左侧
		return a.randomized_select(s, begin, partitioned_iter, rank, compare)
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
func (a *RandomizedSelect) partition(s []int, begin, end, partition int, compare SortCompareFunc) int {
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

/**
 * @description: 顺序统计量的随机选择算法获取第rank小的元素
 * @param s:数据源
 * @param rank:第rank小
 * @param compares:比较函数，可选
 * @return: 第rank小的元素下标;error
 */
func (a *RandomizedSelect) Select(s []int, rank int, compares ...SortCompareFunc) (int, error) {
	var compare SortCompareFunc = nil
	if len(compares) > 0 {
		compare = compares[0]
	} else {
		compare = LessThan
	}
	begin := 0
	end := len(s)
	return a.randomized_select(s, begin, end, rank, compare)
}
