/*
 * @Description: 16.5 用拟阵求解任务调度问题
	单处理器上的的单位时间任务最优调度问题，其中每个任务有一个截止时间以及错过截止时间后的惩罚值。
 * @Author: wangchengdg@gmail.com
 * @Date: 2020-03-03 14:23:55
 * @LastEditTime: 2020-03-08 17:25:47
 * @LastEditors:
*/
package GreedyAlgorithm

import (
	"sort"

	. "github.com/meshcross/algorithm-3rd/mesh/common"
)

type TaskSchedule struct {
}

func NewTaskSchedule() *TaskSchedule {
	return &TaskSchedule{}
}

/**
 * @description: 用拟阵求解任务调度问题，输入为任务
 * @param tasks包含3个信息，First:taskID,Second:taskD截止时间,Third:taskW惩罚
 * @return:返回为最小的惩罚值，以及调度方案
 *
 */
func (t *TaskSchedule) Run(tasks []*Tuple) (int, []int) {

	maxD := 0
	for _, tsk := range tasks {
		if tsk.Second > maxD {
			maxD = tsk.Second
		}
	}

	N := len(tasks)

	//w进行降序排列
	compareW_Desc := func(x, y *Tuple) int {
		if x.Third > y.Third {
			return 1
		}
		if x.Third == y.Third {
			return 0
		}
		return -1
	}
	compareD_Asc := func(x, y *Tuple) int {
		if x.Second < y.Second {
			return 1
		}
		if x.Second == y.Second {
			return 0
		}
		return -1
	}
	wapper := &TupleWapper{
		Tuples:  tasks,
		Compare: compareW_Desc,
	}
	sort.Sort(wapper)

	//独立任务
	I := make([]int, maxD+1)

	//到这里，tasks已经按照任务惩罚进行了降序排列
	line := N
	for i := 0; i < line; i++ {
		if !t.taskIsIndependence(tasks, i, I) {
			line--
			//如果不独立，则要交换i和line的位置，直到找到左边全部独立，右边都不独立的分隔位置
			tmp := tasks[i]
			tasks[i] = tasks[line]
			tasks[line] = tmp
			i--
		}
	}
	task1 := tasks[0:line]
	task2 := tasks[line:]
	wapper1 := &TupleWapper{
		Tuples:  task1,
		Compare: compareD_Asc,
	}
	wapper2 := &TupleWapper{
		Tuples:  task2,
		Compare: compareD_Asc,
	}
	//数组的[]操作会产生新的数组，查看tasks和task1的地址并不一样，数组内部
	//fmt.Println(fmt.Sprintf("%v,%v,%v", unsafe.Pointer(&task1), unsafe.Pointer(&task2), unsafe.Pointer(&tasks)))
	//line左侧的部分当然需要按照d进行増序排列，以让截止时间先到的任务排在前面，否则会造成时间到了有些任务还没有完成
	sort.Sort(wapper1)
	sort.Sort(wapper2)

	ids := []int{}
	w := 0
	for _, v := range task1 {
		ids = append(ids, v.First)
	}
	for _, v := range task2 {
		ids = append(ids, v.First)
		w += v.Third
	}

	return w, ids
}

/**
 * @description: 判断任务是否独立
	对于任务集合A，如果存在一个调度方案，使A中所有任务都不延迟，则称A是独立的
	对于t=0,1,...,n 有N_t(A)<=t
	即如果集合A按照任务截止时间递增的方式调度，则不会有任务延迟
 * @param tasks 任务列表
 * @param i
 * @param {type}
 * @return:独立则返回true，不独立返回false
*/
func (t *TaskSchedule) taskIsIndependence(tasks []*Tuple, i int, I []int) bool {
	Ilen := len(I)
	for j := tasks[i].Second; j < Ilen; j++ {
		if I[j]+1 > j {
			return false
		}
	}
	for j := tasks[i].Second; j < Ilen; j++ {
		I[j]++
	}
	return true
}
