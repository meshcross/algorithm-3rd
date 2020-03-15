/*
 * @Description: 第22章22.5节 有向图的强连通分量
 * @Author: wangchengdg@gmail.com
 * @Date: 2020-02-18 10:21:57
 * @LastEditTime: 2020-03-14 11:40:31
 * @LastEditors:


* 有向图G=(V,E)的强连通分量是一个最大结点集合C，C是V的子集。对于C中的任意一对结点u,v来说，路径u-->v和路径v-->u同时存在。即结点u和结点v之间相互可以到达。
*
* 在强连通分量的算法中，需要用到图G的转置G_T。定义G_T=(V,E_T),其中E_T={(u,v):(v,u)属于E}，即G_T中的边是G中的边进行反向获得。
*
* - 图G和图G_T的强连通分量相同
* - 可以证明，`scc`算法得到的就是强连通分量。
*
* 强连通分量算法步骤：
*
*   - 对原图G执行深度优先搜索，并获取每个结点的完成时间 finish_time
*   - 对转置图G_T执行深度优先搜索，但是按照 G中结点的一个排序来搜索（这个排序是按照finish_time的降序）
*   - G_T的深度优先森林就是强连通分量
*
* 性能：时间复杂度O(V+E)
*
*/
package BasicGraph

import (
	"errors"

	. "github.com/meshcross/algorithm-3rd/mesh/graph_algorithm/graph_struct"
)

type StrongConnectedComponent struct {
}

/**
 * @description: 强连通分量
* 强连通分量算法步骤：
*
*   - 对原图G执行深度优先搜索，并获取每个结点的完成时间 finish_time
*   - 对转置图G_T执行深度优先搜索，但是按照 G中结点的一个排序来搜索（这个排序是按照finish_time的降序）
*   - G_T的深度优先森林就是强连通分量
 * @param graph 图
 * @return:
*/
func (a *StrongConnectedComponent) SetStrongConnectedComponent(graph *Graph) ([][]int, error) {
	if graph == nil {
		return nil, errors.New("scc error: graph must not be nil!")
	}

	//*********  原图的深度优先搜索 ***************
	finished_order := []int{}
	empty_action := func(id, time int) {}
	finish_action := func(v_id, time int) {
		finished_order = append([]int{v_id}, finished_order...)
	} //完成时间逆序

	dfs := NewGraphDFS()
	dfs.Search(graph, empty_action, finish_action, empty_action, empty_action, nil)

	//*********** 转置图的深度优先搜索*********
	result := [][]int{}
	current_root_id := -1

	pre_root_action := func(v_id, time int) {
		result = append(result, []int{})
		current_root_id++
	} //遇到深度优先森林的树根，开启一个新的强连通分量
	pre_action := func(v_id, time int) {
		result[current_root_id] = append(result[current_root_id], v_id)
	} //将结点`id`添加到强连通分量中

	//转置图
	inverse_G := graph.Inverse()
	dfs.Search(inverse_G, pre_action, empty_action, pre_root_action, empty_action, finished_order)

	//**********  剔除单根树 *************
	real_result := [][]int{}
	for _, item := range result {
		if len(item) > 1 {
			real_result = append(real_result, item)
		}
	}
	return real_result, nil
}
