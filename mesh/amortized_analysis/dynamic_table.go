/*
 * @Description: 17.4 动态表，这里和各种语言中的数据机构底层实现原理是一样的，
				 比如golang中的map和slice等，数据量增加放不下时候，会重新开辟新的内存，然后将旧数据copy到新的内存，旧内存并不会立马回收，
				 一方面是内存空间浪费，另一方面频繁实时分配内存也影响性能
				 这就是为什么数据量很大的时候需要使用固定长度的map和slice的原因
				 优先最小级队列(MinQueue)也有类似的维护容量的操作
 * @Author: wangchengdg@gmail.com
 * @Date: 2020-03-03 16:42:58
 * @LastEditTime: 2020-03-06 12:58:03
 * @LastEditors:
*/
package AmortizedAnalysis

import (
	. "github.com/meshcross/algorithm-3rd/mesh/common"
)

type DynamicTalbe struct {
	slot_capacity int           //总共有多少槽位
	size          int           //当前有多少数据,size<=slot_num
	slots         []interface{} //存储数据的地方
}

//初始时候需要有个初始值，这里暂且取8位,
func NewDynamicTable(cnt ...int) *DynamicTalbe {
	capacity := 8
	if len(cnt) > 0 {
		capacity = cnt[0]
	}

	//设定的太小，或者为负数，都是不允许的
	capacity = MinInt(capacity, 8)
	arr := make([]interface{}, capacity)
	return &DynamicTalbe{size: 0, slot_capacity: capacity, slots: arr}
}

/**
 * @description: 向动态表中插入一个对象，如果槽位已满，则需要扩容到原来的两倍，如果有空间，则直接插入
 * @param v 需要插入的元素
 * @return:
 */
func (t *DynamicTalbe) Push(v interface{}) {
	//如果需要扩容，就将容量x2，并做数据迁移
	if t.size == t.slot_capacity {
		new_slots := make([]interface{}, t.slot_capacity*2)
		for i := 0; i < t.size; i++ {
			new_slots[i] = t.slots[i]
		}
		t.slots = new_slots
	}
	{
		t.slots[t.size] = v
		t.size++
	}
}

func (t *DynamicTalbe) Pop() interface{} {
	if t.size > 0 {
		v := t.slots[t.size-1]
		t.slots[t.size-1] = nil

		//实际的使用量不到总容量的1/4，则进行缩容操作
		//锚定1/4而不是1/2，是因为防止在1/2附近频繁增删，造成多次容量的扩展和收缩
		if t.size < t.slot_capacity/4 {
			new_slots := make([]interface{}, t.slot_capacity/4)
			for i := 0; i < t.size; i++ {
				new_slots[i] = t.slots[i]
			}
			t.slots = new_slots
		}
		t.size--

		return v
	}

	return nil
}
