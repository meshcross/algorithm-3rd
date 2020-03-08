/*
 * @Description: 增加一些公共函数
 * @Author: wangchengdg@gmail.com
 * @Date: 2020-03-08 17:05:52
 * @LastEditTime: 2020-03-08 17:35:01
 * @LastEditors:
 */

package Common

/**
 * @description: 对于已排序的数组，使用二分查找
 * @param {type}
 * @return:
 */
func BinarySearch(key int, arr []int) int {
	pos := -1
	left, right := 0, len(arr)

	for left < right {
		mid := (left + right) / 2
		switch true {
		case arr[mid] < key:
			left = mid
		case arr[mid] == key:
			pos = mid
			break
		case arr[mid] > key:
			right = mid
		}
	}
	return pos
}

/**
 * @description: 装箱的排好序的数组执行二分查找
	需要传入arrLen，代表arr中有效长度，服务器端很多情况下一次性分配了很大的数组，但是实际只填充了一部分数据
 * @param compare: 比较函数compare(x,y) x<y返回1，x>y返回-1，x==y返回0
 * @return:下标
*/
func BinarySearchAny(key interface{}, arr []interface{}, arrLen int, compare NodeCompareFunc) int {
	pos := -1
	left, right := 0, arrLen
	if right < 0 {
		right = len(arr)
	}

	for left < right {

		mid := (left + right) / 2
		ret := compare(arr[mid], key)
		if ret > 0 {
			left = mid
		} else if ret == 0 {
			pos = mid
			break
		} else if ret < 0 {
			right = mid
		}
	}
	return pos
}
