package BasicGraph

import (
	"errors"

	. "github.com/meshcross/algorithm-3rd/mesh/common"
	. "github.com/meshcross/algorithm-3rd/mesh/graph_algorithm/graph_struct"
	. "github.com/meshcross/algorithm-3rd/mesh/graph_algorithm/graph_struct/graph_vertex"
)

//!visit：深度优先搜索的辅助函数，用于访问顶点，算法导论22章22.3节
/*!
* \param graph:指向图的强引用，必须非空。若为空则抛出异常
* \param v_id:待访问顶点的`id`，必须有效。如果无效则抛出异常
* \param time:访问时刻，是一个引用参数，确保每次`visit`都访问同一个时钟。
* \param pre_action:一个可调用对象，在每次发现一个顶点时调用，调用参数为该顶点的`id`以及发现时间`time`。默认为空操作，即不进行任何操作
* \param post_action:一个可调用对象，在每次对一个顶点搜索完成时调用，调用参数为该顶点的`id`以及完成时间`time`。默认为空操作，即不进行任何操作
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

func (a *GraphDFS) Search(graph *Graph, pre_action, post_action, pre_root_action, post_root_action DFSActionFunc, search_order []int) error {
	if graph == nil {
		return errors.New("depth_first_search error: graph must not be nullptr!")
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

//!visit：深度优先搜索的辅助函数，用于访问顶点，算法导论22章22.3节
/*!
* \param graph:指向图
* \param v_id:待访问顶点的`id`，必须有效。如果无效则抛出异常
* \param time:访问时刻，是一个引用参数，确保每次Visit都访问同一个时钟。
* \param pre_action:一个可调用对象，在每次发现一个顶点时调用，调用参数为该顶点的`id`以及发现时间`time`。默认为空操作，即不进行任何操作
* \param post_action:一个可调用对象，在每次对一个顶点搜索完成时调用，调用参数为该顶点的`id`以及完成时间`time`。默认为空操作，即不进行任何操作
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
		return errors.New("visit error: graph must not be nullptr!")
	}

	num := graph.N()
	if v_id < 0 || v_id >= num || graph.Vertexes[v_id] == nil {
		return errors.New("visit error: v_id muse belongs [0,N) and graph->vertexes[v_id] must not be nullptr!")
	}

	time++
	//*******  发现本顶点 *****************
	if pre_action != nil {
		pre_action(v_id, time)
	}
	vtx := a.toBFSVertext(graph.Vertexes[v_id])
	vtx.SetDisovered(time)

	//********  搜索本顶点相邻的顶点*************
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
	//*********** 完成本顶点的搜索
	time++
	vtx.SetFinished(time)
	post_action(v_id, time)

	return nil
}
