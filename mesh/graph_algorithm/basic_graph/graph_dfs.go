/*
 * @Description: 第22章22.3节 深度优先搜索
 * @Author: wangchengdg@gmail.com
 * @Date: 2020-02-18 10:22:32
 * @LastEditTime: 2020-03-14 11:40:19
 * @LastEditors:

 *
 * 深度优先搜索：深度优先搜索总是对最近才发现的结点v的出发边进行搜索，直到该结点的所有出发边都被发现为止。一旦结点v的所有出发边都被发现，则“回溯”到v的前驱结点（v是经过该结点才被发现的）。
 * 该过程一直持续到源结点可以达到的所有结点都被发现为止。如果还存在尚未发现的结点，则深度优先搜索将从这些未被发现的结点中任选一个作为新的源结点，并重复同样的搜索过程。该算法重复整个过程，
 * 直到图中的所有结点被发现为止。
 *
 * 深度优先搜索维护一个全局的时间。每个结点v有两个时间戳，DiscoverTime记录了v第一次被发现的时间（v涂上灰色的时刻）；FinishTime记录了搜索完成v的相邻结点的时间（v涂上黑色的时刻）。
 * 结点v在v.DiscoverTime之前为白色，在v.DiscoverTime之后与v.FinishTime之前为灰色，在v.FinishTime之后为黑色

 */
package BasicGraph

import (
	"errors"

	. "github.com/meshcross/algorithm-3rd/mesh/common"
	. "github.com/meshcross/algorithm-3rd/mesh/graph_algorithm/graph_struct"
	. "github.com/meshcross/algorithm-3rd/mesh/graph_algorithm/graph_struct/graph_vertex"
)

type GraphDFS struct {
}

type DFSActionFunc func(id, time int)

func NewGraphDFS() *GraphDFS {
	return &GraphDFS{}
}

func (a *GraphDFS) toBFSVertext(vtx IVertex) *DFSVertex {
	// ptr := unsafe.Pointer(vtx)
	// v := (*DFSVertex)(ptr)
	// return v
	if v, ok := vtx.(*DFSVertex); ok {
		return v
	}
	return nil
}

/**
 * @description: 深度优先搜索
* @param graph:图
* @param pre_action:在每次发现一个顶点时调用，回调函数
* @param post_action:在每次对一个顶点搜索完成时调用，回调函数
* @param pre_root_action:在每次发现一个顶点且该顶点是深度优先森林的根节点时调用，回调函数
* @param post_root_action:在每次对一个顶点搜索完成时且该顶点是深度优先森林的根节点时调用调用，回调函数
* @param search_order:指定搜索顶点的顺序（不同顺序可能形成的深度优先森林不同)，如果为空则按照顶点的`id`顺序。默认为空
* @return:error
*/
func (a *GraphDFS) Search(graph *Graph, pre_action, post_action, pre_root_action, post_root_action DFSActionFunc, search_order []int) error {
	if graph == nil {
		return errors.New("depth_first_search error: graph must not be nil!")
	}

	num := graph.N()
	//************  创建真实的 search_order ****************
	real_search_order := []int{}
	if search_order == nil || len(search_order) <= 0 {
		for i := 0; i < num; i++ {
			real_search_order = append(real_search_order, i)
		}
	} else {
		real_search_order = search_order
	}

	//************* 初始化顶点 ****************
	for _, ver := range graph.Vertexes {
		if ver == nil {
			continue
		}
		v := a.toBFSVertext(ver)
		v.Color = COLOR_WHITE
		v.SetParent(nil)
	}

	//*************** 深度优先搜索 *************
	time := 0
	for _, v_id := range real_search_order {
		if v_id < 0 || v_id >= num || graph.Vertexes[v_id] == nil { //顶点为空
			continue
		}
		ver := graph.Vertexes[v_id]
		v := a.toBFSVertext(ver)
		if v.Color == COLOR_WHITE {
			if pre_root_action != nil {
				pre_root_action(v_id, time)
			}

			a.Visit(graph, v.GetID(), time, pre_action, post_action)

			if post_root_action != nil {
				post_root_action(v_id, time)
			}
		}
	}
	return nil
}

/*!
* @description:深度优先搜索的辅助函数，用于访问每个顶点
* @param graph:图
* @param v_id:待访问顶点的`id`
* @param time:访问时刻，是一个引用参数，确保每次`visit`都访问同一个时钟。
* @param pre_action:在每次发现一个顶点时调用，回调函数
* @param post_action:在每次对一个顶点搜索完成时调用，回调函数
* @return :error
*
* `v_id`在以下情况下无效：
*
* - `v_id`不在区间`[0,N)`之间时，`v_id`无效
* - `graph`中不存在某个顶点的`id`等于`v_id`时，`v_id`无效
*
* 在每次对一个结点调用visit的过程中，结点v_id的初始颜色都是白色。然后执行下列步骤：
*
* - 将全局时间 time 递增
* - 发现结点 v_id
* - 对结点 v_id 的每一个相邻结点进行检查，在相邻结点是白色的情况下递归访问该相邻结点
* - 当结点 v_id 的相邻结点访问完毕，则全局时间 time 递增，然后将结点 v_id 设置为完成状态
*
 */
func (a *GraphDFS) Visit(graph *Graph, v_id, time int, pre_action DFSActionFunc, post_action DFSActionFunc) error {

	if graph == nil {
		return errors.New("visit error: graph must not be nil!")
	}

	num := graph.N()
	if v_id < 0 || v_id >= num || graph.Vertexes[v_id] == nil {
		return errors.New("visit error: v_id muse belongs [0,N) and graph.Vertexes[v_id] must not be nil!")
	}

	time++
	//-------stage1  发现本顶点 --------
	if pre_action != nil {
		pre_action(v_id, time)
	}
	vtx := a.toBFSVertext(graph.Vertexes[v_id])
	vtx.SetDisovered(time)

	//--------stage2 搜索本顶点相邻的顶点 --------
	edges, _ := graph.VertexEdgeTuples(v_id)
	for _, edge := range edges {
		another_id := edge.Second
		another_vertex_wp := graph.Vertexes[another_id]
		another_vertex := a.toBFSVertext(another_vertex_wp)
		if another_vertex.Color == COLOR_WHITE {
			another_vertex.SetParent(vtx)
			a.Visit(graph, another_id, time, pre_action, post_action)
		}
	}
	//--------stage3 完成本顶点的搜索--------
	time++
	vtx.SetFinished(time)
	post_action(v_id, time)

	return nil
}
