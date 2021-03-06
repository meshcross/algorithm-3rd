/*
 * @Description: 第22章22.4节 拓扑排序
 * @Author: wangchengdg@gmail.com
 * @Date: 2020-02-18 10:22:15
 * @LastEditTime: 2020-03-14 17:21:21
 * @LastEditors:
 *
 *
 * 对于一个有向无环图G=（V，E)，其拓扑排序是G中所有结点的一种线性次序，该次序满足如下条件：
 * 如果图G包含边(u,v)，则结点u在拓扑排序中处于结点v的前面。
 *
 * 拓扑排序原理：对有向无环图G进行深度优先搜索。每当完成一个结点时，将该结点插入到拓扑排序结果的头部。
 * 因此如果将结点按照完成时间降序排列，则得到的就是拓扑排序的结果。
 *
 * 引理：一个有向图G=(V,E)是无环的当且仅当对其进行深度优先搜索时不产生后向边。
 *
 * 性能：时间复杂度O(V+E)
 *
 */
package BasicGraph

import (
	"errors"

	. "github.com/meshcross/algorithm-3rd/mesh/graph_algorithm/graph_struct"

	. "github.com/meshcross/algorithm-3rd/mesh/common"
)

type TopologySort struct {
}

func NewTopologySort() *TopologySort {
	return &TopologySort{}
}

/*!
 * @description:图的拓扑排序
 * @param graph:有向无环图
 * @return:拓扑排序结果，它是顶点`id`组成的[]int，表示顶点的拓扑排序后的顺序
 *
 * 前置要求：有向无环图
 * 生成的是有向无环图的拓扑排序
**/
func (a *TopologySort) Sort(graph *Graph) ([]int, error) {

	if graph == nil {
		return nil, errors.New("topology_sort error: graph must not be nil!")
	}

	//一次分配好，免得节点数过多频繁resize消耗性能
	sorted_result := make([]int, len(graph.Vertexes))
	empty_action := func(id, time int) {}
	add_count := 0
	finish_action := func(v_id, time int) {
		//sorted_result = append(sorted_result, v_id)
		sorted_result[add_count] = v_id
		add_count++
	}

	dfs := NewGraphDFS()
	dfs.Search(graph, empty_action, finish_action, empty_action, empty_action, nil)
	Revert(sorted_result)
	return sorted_result, nil
}
