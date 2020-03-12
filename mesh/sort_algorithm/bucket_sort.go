/*
 * @Description: 第8章 8.4 桶排序
 * 		- 桶排序思想，假设对数组A[p...r]排序，首先将这些元素进行hash运算，根据其hash值放入桶的对应区间中；
 *		  然后对每一个区间中的元素进行排序；最后合并桶中各区间排序好的结果得到排序的数据：
*   		- hash算法必须满足：若 a<b ，则hash(a)<hash(b)
*   		- 要求 hash的结果尽量好，使得各数据平均分布在各区间内
* 		- 期望时间复杂度 O(n)
* 		- 非原地排序
 * @Author: wangchengdg@gmail.com
 * @Date: 2020-02-10 13:48:31
 * @LastEditTime: 2020-03-11 22:20:21
 * @LastEditors:
*/
package SortAlgorithm

type BucketSort struct {
}

/**
 * @description: 假设输入数据服从均匀分布，与计数排序类似，由于对输入数据做了某种假设，所以桶排序速度也很快
 * @param s:数据源
 * @param min_val:s中的所有元素都>=min_val
 * @param max_val:s中的所有元素都<=max_val
 * @return: void
 */
func (a *BucketSort) Sort(s []int, min_val, max_val int) {
	size := len(s)
	if size <= 1 {
		return
	}
	if min_val > max_val {
		panic("max_val need big than min_val")
	}

	real_bucket_num := 10 //划分10个区间
	buckets := make([][]int, real_bucket_num)

	iter := 0
	for iter != size {
		value := s[iter]
		// 这里采取线性分布，将元素划分到每个桶
		index := (value - min_val) * real_bucket_num / (max_val - min_val)
		if index < real_bucket_num {
			buckets[index] = append(buckets[index], value)
		} else {

			buckets[real_bucket_num-1] = append(buckets[real_bucket_num-1], value)
		}
		iter++
	}

	sorter := QuickSort{}
	inserted_total := 0
	//然后对每个桶中的元素使用快速排序算法进行排序，将排序后的桶进行
	for i := 0; i < real_bucket_num; i++ {
		sorter.Sort(buckets[i])

		ksize := len(buckets[i])
		for k := 0; k < ksize; k++ {
			s[inserted_total+k] = buckets[i][k]
		}
		inserted_total += ksize
	}

}
