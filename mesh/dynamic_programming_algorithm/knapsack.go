/*
 * @Description: 背包问题，并不在书中，但是也属于经典的动态规划问题，所以这里做了实现
 			背包问题，属于经典的动态规划问题之一
*
* 			主要包含3个背包问题：
* 			1、0/1背包问题
*    			一组物品goods要放入背包，重量各不相同cost，价格也不一样value，背包最大承重capacity，每种物品最多选一个，求max(sum(value))
* 			2、完全背包问题
*    			前提同1，但是每个物品无限量供应，求max(sum(value))
* 			3、多重背包问题
*    			前提同1，但是每个物品数量为count(i)，求max(sum(value))
*
* 			要在纸上画出DP图方便分析
*
 * @Author: wangchengdg@gmail.com
 * @Date: 2020-02-29 23:43:10
 * @LastEditTime: 2020-03-13 16:36:40
 * @LastEditors:
*/
package DynamicProgrammingAlgorithm

import (
	"errors"

	. "github.com/meshcross/algorithm-3rd/mesh/common"
)

type Knapsack struct {
}

/**
 * @description:
 * @param {type}
 * @return:
 */
func NewKnapsack() *Knapsack {
	return &Knapsack{}
}

/**
* @description 0/1背包问题，每个物品选择1或者0个
* @param :goods 商品的消耗和价值 First为消耗，Second为价值
* @param : capacity 背包容量
* @return: 最大价值,选中的物品
* 为了提高性能，使用位运算，这里先限制len(goods)<=64，如果需要goods长度无限制，可以改进,vtxs用其他数据结构存储
**/
func (a *Knapsack) Pack01(goods []*Pair, capacity int) (int, []int, error) {
	//存储的是从1到capacity的每个容量的最优选择
	f := make([]int, capacity+1)
	last_f := make([]int, capacity+1)

	if len(goods) > 64 {
		return 0, nil, errors.New("goods count must be less than 64")
	}

	//用于存储哪些节点被选中
	vtxs := make([]uint64, capacity+1)

	n := len(goods)
	for i := 0; i < n; i++ {
		cost_i := goods[i].First
		value_i := goods[i].Second
		//ans1:j从capacity开始循环可以保证下一层由上一层推导出来
		// for j := capacity; j >= cost_i; j-- {
		// 	new_value := f[j-cost_i] + value_i
		// 	//f[j] = MaxInt(f[j], f[j-cost_i]+value_i)
		// 	if new_value > f[j] {
		// 		f[j] = new_value

		// 		m := vtxs[j-cost_i]
		// 		vtxs[j] = SetUintBit(m, i+1, 1)
		// 	}
		// }

		//ans2：缓存上一层数据的方式也可以实现，看起来更清晰一些
		for j := cost_i; j <= capacity; j++ {
			//last_f是上一层轮询goods所得(goods[i-1])，上一层不可能包含当前这一层的goods_i，
			//在整个j循环中，都一直拿last_f跟这一层的数据相比，所以对于这一层的goods_i，只有选中和没有选中的问题，不存在j-1选了再问j要不要选的问题
			new_value := last_f[j-cost_i] + value_i
			if new_value > last_f[j] {
				f[j] = new_value

				m := vtxs[j-cost_i]
				vtxs[j] = SetUintBit(m, i+1, 1)
			}
		}
		//下一层的f状态一定由上一层推导出来,这里缓存住i层的状态，到for进入i+1层的时候，可以使用
		for k, v := range f {
			last_f[k] = v
		}
	}

	selected := vtxs[capacity]
	ids := []int{}
	for i := 0; i < n; i++ {
		ok := selected & (1 << i)
		if ok > 0 {
			ids = append(ids, i)
		}
	}
	return f[capacity], ids, nil
}

/**
* @description 0/1背包问题，每个物品选择1或者0个
* @param :goods 商品的消耗和价值
* @param : capacity 背包容量
* @return: 最大价值;每种物品的数量
**/
func (a *Knapsack) Pack01_Plus(goods []*Pair, capacity int) (int, []int, error) {

	n := len(goods)

	//************** stage1 init *****************//
	f := make([][]int, n)
	for k := 0; k < n; k++ {
		f[k] = make([]int, capacity+1)
	}

	for i := 0; i < n; i++ {
		f[i][0] = 0
	}

	//************** stage2 get max value *****************//

	firstItem := goods[0]
	for j := 1; j <= capacity; j++ {
		if j >= firstItem.First {
			f[0][j] = firstItem.Second
		} else {
			f[0][j] = 0
		}
	}

	//这里的遍历不包含goods[0]，所以在前面就处理掉了
	for i := 1; i < n; i++ {
		item := goods[i]
		cost_i := item.First
		value_i := item.Second
		for j := 0; j <= capacity; j++ {
			if j < cost_i {
				f[i][j] = f[i-1][j]
			} else {
				f[i][j] = MaxInt(f[i-1][j], f[i-1][j-cost_i]+value_i)
			}
		}
	}

	max_value := f[n-1][capacity]

	//************** stage3 get goods selected *****************//
	selected := []int{}
	i := n - 1
	j := capacity
	for i >= 1 {
		cost_i := goods[i].First
		value_i := goods[i].Second
		// if j >= cost_i {
		if f[i][j] == f[i-1][j] {
			//fmt.Println(fmt.Sprintf("未选第 %d 件物品", i))
		} else if f[i][j] == f[i-1][j-cost_i]+value_i {
			//fmt.Println(fmt.Sprintf("选中第 %d 件物品", i))
			selected = append(selected, i)
			j -= cost_i
		}
		// }
		i--
	}

	if j < firstItem.First {
		//fmt.Println(fmt.Sprintf("未选第 %d 件物品", 0))
	} else {
		//fmt.Println(fmt.Sprintf("选中第 %d 件物品", 0))
		selected = append(selected, 0)
	}
	return max_value, selected, nil
}

/**
* @description:完全背包问题，每种物品不限数量
* @param :goods 商品的消耗和价值  First为消耗，Second为价值
* @param : capacity 背包容量
* @return: 最大价值，每种物品的数量
**/
func (a *Knapsack) PackComplete(goods []*Pair, capacity int) (int, []int, error) {

	//存储的是从1到capacity的每个容量的最优选择
	f := make([]int, capacity+1)
	vtxs := make([]map[int]int, capacity+1)
	for k, _ := range vtxs {
		vtxs[k] = map[int]int{}
	}

	n := len(goods)
	for i := 0; i < n; i++ {
		cost_i := goods[i].First
		value_i := goods[i].Second
		for j := cost_i; j <= capacity; j++ {
			//随着j向右推进，容量增加，则goods_i可以不断添加，从而选择多个goods_i
			//每一次的j循环，f都在被更新，这是和0/1背包问题最大的不同，一旦new_value被选中，表明又选了一个goods_i
			new_value := f[j-cost_i] + value_i
			if new_value > f[j] {
				f[j] = new_value

				old := vtxs[j-cost_i]
				for k, _ := range vtxs[j] {
					// vtxs[j][k] = 0
					delete(vtxs[j], k)
				}
				for k, v := range old {
					vtxs[j][k] = v
				}
				vtxs[j][i] += 1
			}
		}
	}

	selected := vtxs[capacity]
	ids := make([]int, n)
	for i := 0; i < n; i++ {
		count := selected[i]
		ids[i] = count
	}
	// fmt.Println(vtxs[capacity])
	return f[capacity], ids, nil
}

/**
* @description:多重背包问题，每种物品有数量限制
* @param :goods 商品的消耗和价值,First为消耗，Second为价值，Third为限制数量
* @param : capacity 背包容量
* @return: 每种物品的数量
**/
func (a *Knapsack) PackMultiple(goods []*Tuple, capacity int) (int, []int, error) {
	f := make([]int, capacity+1)
	last_f := make([]int, capacity+1)

	//每个f都得分都有对应的物品组合，对应的组合存储在这里
	vtxs := make([]map[int]int, capacity+1)
	for k, _ := range vtxs {
		vtxs[k] = map[int]int{}
	}

	n := len(goods)
	for i := 0; i < n; i++ {
		cost_i := goods[i].First   //消耗
		value_i := goods[i].Second //价值
		count_i := goods[i].Third  //数量限制

		// for j := capacity; j >= cost_i; j-- {
		for j := cost_i; j <= capacity; j++ {
			for k := 1; k <= count_i; k++ {
				if j >= k*cost_i {
					//使用上一层的f参数:last_f
					new_value := last_f[j-k*cost_i] + k*value_i

					//f[j] = MaxInt(f[j], f[j-cost_i]+value_i)
					if new_value > f[j] {
						f[j] = new_value

						// start update selected goods
						//为了避免总是重新生成map，这里直接执行delete和重新赋值操作
						old := vtxs[j-cost_i]
						for x, _ := range vtxs[j] {
							// vtxs[j][k] = 0
							delete(vtxs[j], x)
						}
						for x, v := range old {
							vtxs[j][x] = v
						}
						vtxs[j][i] = k
						// end update selected goods
					}
				}
			}
		}

		//下一层的f状态一定由上一层推导出来,这里缓存住i层的状态，到for进入i+1层的时候，可以使用
		for k, v := range f {
			last_f[k] = v
		}
	}

	selected := vtxs[capacity]
	ids := make([]int, n)
	for i := 0; i < n; i++ {
		count := selected[i]
		ids[i] = count
	}

	return f[capacity], ids, nil
}
