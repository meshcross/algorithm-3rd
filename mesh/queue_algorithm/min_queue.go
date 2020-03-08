/*
 * @Description: 算法导论第6章6.5节 最小优先级队列
 * @Author: wangchengdg@gmail.com
 * @Date: 2020-02-10 22:17:56
 * @LastEditTime: 2020-03-06 12:05:50
 * @LastEditors:
 */
package QueueAlgorithm

import (
	"errors"
	"fmt"

	. "github.com/meshcross/algorithm-3rd/mesh/common"
)

/*!
* 最小堆，是一种经过排序的完全二叉树，其中任一非终端节点的数据值均不大于其左子节点和右子节点的值。
*
* 优先级队列是一种用来维护由一组元素构成集合S的数据结构，其中每个元素都有一个相关的值，称之为关键字。一个最小优先级队列支持以下操作：
*
* - Insert(S,x):将元素x插入到集合S中。
* - Min(S):返回S中具有最小关键字的元素
* - ExtractMin(S):去掉并返回S中具有最小关键字的元素
* - DecreaseKey(S,x,k):将元素x的关键字值减小到k,这里要求k的值小于x的原始关键字
*
 */
type MinQueue struct {
	_size   int
	_data   []interface{}
	Compare NodeCompareFunc
}

/**
 * @description: 最小堆的创建函数
			当前_data由MinQueue内部维护，当size特别大的时候，Insert次数很大，会发生多次resize,造成性能损失
			所以此处支持传入data，这样在外部初始化好了再传进来，避免多次resize
 * @param compare：data传进来的数据是进过装箱操作的，不知道明确的类型，所以需要指定比较函数，以比较data中各元素的顺序
 * @param data:外部传进来的数据,让该数据符合最小堆的性质
 * @return:最小堆的指针
*/
func NewMinQueue(compare NodeCompareFunc, data []interface{}) *MinQueue {
	size := 0
	if data == nil {
		data = []interface{}{}
	} else {
		size = len(data)
	}
	return &MinQueue{Compare: compare, _data: data, _size: size}
}

/**
 * @description: 数组扩容，数组需要的是连续内存，所以不够用的时候需要重新非配整块内存，并且将旧数组的数据拷贝到新的数组
 * @param newCapacity:最小堆的新的容量，通常需要扩容时候都是将容量x2
 * @return: void
 */
func (a *MinQueue) resize(newCapacity int) {
	newArr := make([]interface{}, newCapacity)
	for i := 0; i < a._size; i++ {
		newArr[i] = a._data[i]
	}
	a._data = newArr
}

/**
 * @description: 获取最小堆头部的元素，即为最小元素
 * @return: 最小堆头部的元素
 */
func (a *MinQueue) Min() (interface{}, error) {
	if a._size <= 0 {
		return 0, errors.New("no data")
	}
	return a._data[0], nil
}

/*!
* @description: 删除队列中最小值，并返回最小值
* @return  最小值的强引用 ；error
*
* 根据最小堆的性质，队列的第一个值就是最小值。
* 取出最小值之后再进行以下操作：
*   - 如果队列为空，则返回一个空
*   - 如果队列非空，则执行以下操作：
*       - 交换队列的第一个元素和最后一个元素
*       - 队列的`_size`自减
*       - 此时队列的首个元素违反了最小堆性质，因此执行`heapify(0)`保持性质
*       - 返回旧的首个元素
*
* 一旦队列长度不足容量的1/4，则将队列容量设置为队列长度x2
*
* - 时间复杂度 O(h),h为堆的高度
* - 原地操作
 */

func (a *MinQueue) ExtractMin() (interface{}, error) {
	if a._size <= 0 {
		return 0, errors.New("no data")
	}
	result := a._data[0]
	a._data[0] = a._data[a._size-1]
	a._size--
	a.heapify(0)
	//当数据量<容量的1/4的时候，需要进行缩容
	//如果数据量<容量1/2的时候进行缩容，可能会引起频繁的resize,所以锚点是1/4的位置
	if a._size <= len(a._data)/4 {
		a.resize(a._size*2 + 2)
	}
	return result, nil
}

/*!
* @description:向队列中插入一个元素
* @param element: 待插入元素，如果元素为空则直接返回
* @return: 插入的元素在队列中的位置。若元素为空则返回-1
*
* 插入之前首先判断队列是否已满。若队列已满，则先扩容。
*
* 插入过程为：
*
* - 保留待插入元素的`key`，同时将待插入语元素的`key`设置为无穷大，并将待插入元素插入到队尾
* - 执行`decreate_key(..)`操作，插入之后需要执行该操作，以保持最大堆性质
*
* - 时间复杂度 O(h)
* - 原地操作
 */

func (a *MinQueue) Insert(element interface{}) (int, error) {

	if a._size == len(a._data) {
		a.resize(a._size*2 + 2)
	}

	index := a._size
	a._size++
	a._data[index] = element
	a.DecreateKey(index, element)
	return index, nil
}

/**
 * @description: 最小堆是否为空
 * @return: bool
 */
func (a *MinQueue) IsEmpty() bool {
	return a._size == 0
}

/**
 * @description: 返回最小堆的容量
 * @return: int
 */
func (a *MinQueue) Capacity() int {
	return len(a._data)
}

/**
 * @description: 最小堆中有多少元素
 * @return: int
 */
func (a *MinQueue) Size() int {
	return a._size
}

/*!
* @description:返回指定元素是否在队列中
* @param element:待判定的元素
* @return 指定元素在队列中的下标
*
* - 时间复杂度 O(h)
 */
func (a *MinQueue) ElementIndex(element interface{}) int {
	for index := 0; index < a._size; index++ {
		if element == a._data[index] {
			return index
		}
	}
	return -1
}

/*!
* @description:缩减队列中某个元素的`key`
* @param element_index: 待缩减元素的下标
* @param new_key：待缩减元素的新`key`，即为存储的数据
* @return error
*
* 缩减过程为：
*
* - 将待缩减元素的`key`赋值为新值
* - 不断的将该元素向父节点比较：
*   - 若父节点较小，则终止比较过程
*   - 若父节点较大，则交换当前节点与父节点，并将当前节点指向父节点进行下一轮比较
*   - 若当前节点已经是队列首个元素，则终止比较过程
*
* - 时间复杂度 O(h)
* - 原地操作
 */
func (a *MinQueue) DecreateKey(elementIndex int, newKey interface{}) error {
	if elementIndex >= a._size {
		return errors.New("index error")
	}

	if a.Compare(a._data[elementIndex], newKey) > 0 {
		return errors.New("new key is too big")
	}

	a._data[elementIndex] = newKey

	//通过遍历比较，将新增加的元素放置在合适的位置上
	//实际为最小堆的维持健康操作，只是这里关注的只有新添加的元素，而heapify为整个堆
	for elementIndex != 0 {
		valid, pIndex := a._parentIndex(elementIndex)
		if !valid {
			break
		}
		if a.Compare(a._data[pIndex], a._data[elementIndex]) > 0 {
			break
		}
		SwapAny(a._data, pIndex, elementIndex)
		elementIndex = pIndex
	}
	return nil
}

/*!
* @description:返回一个节点的父节点位置
* @param elementIndex : 子节点位置
* @return 返回值1：一个bool&值，用于返回，指示父节点是否有效 ； 返回值2： 父节点位置
*
* 根据最小堆的性质，一个子节点elementIndex的父节点是它的位置(elementIndex-1)/2。
*
 */
func (a *MinQueue) _parentIndex(elementIndex int) (bool, int) {
	if elementIndex >= a._size {
		return false, 0 //无效结果
	}
	//有效结果
	index := 0
	if elementIndex == 0 { //根节点的父节点是自己
		index = 0
	} else {
		index = (elementIndex - 1) >> 1
	}

	return true, index
}

/*!
* @description:返回一个节点的左子节点位置
* @param elementIndex : 节点位置
* @return 返回值1：一个bool&值，用于返回，指示子节点是否有效；返回值2：左子节点位置
*
* 根据最小堆的性质，一个节点elementIndex的左子节点是它的位置(elementIndex/2)+1
*
 */
func (a *MinQueue) _lchildIndex(elementIndex int) (bool, int) {
	if a._size < 2 {
		//数组元素太少无效结果
		return false, 0
	}
	if elementIndex > ((a._size - 2) >> 1) {
		return false, 0 //超出范围，无效
	}
	return true, (elementIndex << 1) + 1
}

/*!
* @description:返回一个节点的右子节点位置
* @param elementIndex : 节点位置
* @return 返回值1：一个bool&值，用于返回，指示子节点是否有效，返回值2：右子节点位置
*
* 根据最小堆的性质，一个节点elementIndex的右子节点是它的位置(elementIndex/2)+2
*
 */
func (a *MinQueue) _rchildIndex(elementIndex int) (bool, int) {
	if a._size < 3 {
		return false, 0 //数组元素太少无效结果
	}
	if elementIndex > ((a._size - 3) >> 1) {
		return false, 0 //超出范围，无效
	}
	return true, (elementIndex << 1) + 2
}

/*!
* @description:建堆
* @return void
*
* 从后一半的元素开始依次向前调用heapify操作（根据最小堆性质，除了最底层它是完全充满的）
*
* - 时间复杂度 O(nlogn)
* - 原地操作
 */
func (a *MinQueue) setupHeap() {
	if a._size <= 1 {
		return
	}

	index := (a._size - 1) / 2
	for index >= 0 {
		a.heapify(index)
		index--
	}
}

/*!
* @description:维持堆性质
* @param elementIndex : 要维持以该节点为根节点的子堆的堆性质
* @return void
*
* 首先调用比较该节点与左右子节点的最小值。如果最小值为它本身，则维持了性质，返回；如果最小值不是它本身，那么必然为左、右子节点之一。
* 将该最小节点（假设为左子节点）交换到根节点，然后以左子节点递归调用heapify操作
*
* - 时间复杂度 O(n)
* - 原地操作
 */
func (a *MinQueue) heapify(elementIndex int) {
	if elementIndex >= a._size {
		return
	}

	minIndex := elementIndex

	left_valid, leftIndex := a._lchildIndex(elementIndex)
	right_valid, rightIndex := a._rchildIndex(elementIndex)

	if left_valid {
		if a.Compare(a._data[leftIndex], a._data[minIndex]) > 0 {
			minIndex = leftIndex
		}
	}
	if right_valid {
		if a.Compare(a._data[rightIndex], a._data[minIndex]) > 0 {
			minIndex = rightIndex
		}
	}

	if minIndex != elementIndex {

		SwapAny(a._data, elementIndex, minIndex)
		a.heapify(minIndex)
	}
}

func (a *MinQueue) Print() {
	fmt.Println(a._data)
}
