/*
 * @Description: 第24章24.2节 有向无环图的单源最短路径的dag shortest path算法
 * @Author: wangchengdg@gmail.com
 * @Date: 2020-02-18 10:35:09
 * @LastEditTime: 2020-03-14 20:07:26
 * @LastEditors:
 *
 *
 *
 * ## 单源最短路径
 *
 * 单源最短路径问题：给定一个带权重的有向图G=(V,E)和权重函数w:E->R，该权重函数将每条边映射到实数值的权重上。图中一条路径p=<v0,v1,...vk>的权重
 * w(p)=w(v0,v1)+w(v1,v2)+...+w(v(k-1),vk)。定义结点u到结点v的最短路径权重 delt(u,v)为：
 *
 * - min{w(p):u-->v(通过路径p)}，如果存在一条从结点u到结点v的路径
 * - 正无穷 ，如果不存在一条从结点u到结点v的路径
 *
 * 从结点u到结点v的最短路径定义为任何一条权重w(p)=delt(u,v)的从u到v的路径p。
 *
 * 给定图G=(V,E)，对每个结点v我们维持一个前驱结点v.pai。在最短路径算法中，由pai值诱导的前驱子图G_pai=(V_pai,E_pai)，其中 V_pai={v属于V:v.pai!=nil}并上源点s，
 * E_pai是V_pai中所有结点的pai值诱导的边的集合：E_pai={(v.pai,v)属于E:v属于V_pai-{s} }。算法终止时，G_pai是一棵最短路径树：该树包含了从源结点s
 * 到每个可以从s到达的结点的一条最短路径。
 *
 * 需要指出的是：最短路径不一定是唯一的，最短路径树叶不一定是唯一的。
 *
* ## dag_shortest_path算法
*
* ### 算法原理
*
* dag shortest path算法解决的是有向无环图中的单源最短路径问题。在有向无环图中不存在环路因此也就不存在权重为负值的环路，最短路径都是存在的。
* dag shortest path根据结点的拓扑排序次数来对带权重的有向无环图进行边的松弛操作，可以再O(V+E)时间内计算出单源最短路径。
*
* 时间复杂度为O(V+E)
*/
package SingleSourceShortestPath

import (
	"errors"

	. "github.com/meshcross/algorithm-3rd/mesh/common"
	. "github.com/meshcross/algorithm-3rd/mesh/graph_algorithm/basic_graph"
	. "github.com/meshcross/algorithm-3rd/mesh/graph_algorithm/graph_struct"

	. "github.com/meshcross/algorithm-3rd/mesh/graph_algorithm/graph_struct/graph_vertex"
)

type DagShortestPath struct {
}

func NewDagShortestPath() *DagShortestPath {
	return &DagShortestPath{}
}

/*
* @description:有向无环图的单源最短路径的dag shortest path算法
* @param graph:图，必须非空
* @param source_id：最小生成树的根结点`id`，必须有效。若无效则抛出异常
* @return: void
*
* 如果有向无环图中包含u到v的路径，则在拓扑排序中，u位于v的前面（如果同时存在v->u的路径，则就有环了）
*
* ### 算法步骤
*
* - 对有向无环图进行拓扑排序
* - 执行单源最短路径的初始化过程
* - 按照顶点的拓扑排序的顺序依次处理，每次处理过程为：对该顶点出发的每一条边进行一次松弛操作
*
*
* ### 算法性能
*
* 时间复杂度为O(V+E)
*
 */
func (a *DagShortestPath) ShortestPath(graph *Graph, source_id int) error {
	if graph == nil {
		return errors.New("DagShortestPath error: graph must not be nil!")
	}

	num := graph.N()
	if source_id < 0 || source_id >= num || graph.Vertexes[source_id] == nil {
		return errors.New("DagShortestPath error: source_id muse belongs [0,N) and source vertex must not be nil!")
	}

	//与bellman_ford算法不同之处，这里要进行拓扑排序
	//如果存在u->v的路径，则拓扑排序中u一定位于v的前面
	topo := NewTopologySort()
	sorted_vertexs, _ := topo.Sort(graph)

	a.initializeSingleSource(graph, source_id)

	//************* 循环处理遍历所有的边  ***************
	//相当于沿着路径u往v的方向向后探测，注意图的拓扑排序的性质
	for _, v_id := range sorted_vertexs {
		edges, _ := graph.VertexEdgeTuples(v_id)
		for _, edge := range edges {
			from := graph.Vertexes[edge.First]
			to := graph.Vertexes[edge.Second]
			wt := edge.Third

			//看from->to是不是更优的路劲，如果是，则更新到to，否则放弃
			a.relax(from, to, wt)
		}
	}
	return nil
}

/**
 * @description:初始化，source节点key设置为0,其他节点为unlimit，所有节点parent设置为nil
 * @param graph:图
 * @param source_id:源节点id
 * @return:error
 */
func (a *DagShortestPath) initializeSingleSource(graph *Graph, source_id int) error {
	if graph == nil {
		return errors.New("initializeSingleSource error: graph must not be nil!")
	}

	num := graph.N()
	if source_id < 0 || source_id >= num || graph.Vertexes[source_id] == nil {
		return errors.New("initializeSingleSource error: source_id muse belongs [0,N) and source vertex must not be nil!")
	}

	//**************** 设置所有结点 *****************
	for i := 0; i < num; i++ {

		vertex := graph.Vertexes[i]
		if vertex != nil {
			vertex.SetKey(Unlimit())
			vertex.SetParent(nil)
		}
	}
	//**************  设置源结点 ***************
	vertex := graph.Vertexes[source_id]
	vertex.SetKey(0)

	return nil
}

func (a *DagShortestPath) relax(from, to IVertex, weight int) error {
	if from == nil || to == nil {
		return errors.New("relax error: from_vertex and to_vertex must not be nil!")
	}

	if from == to {
		return errors.New("relax error: from_vertex must not be to_vertex!")
	}

	if Is_Unlimit(from.GetKey() + weight) { //u.key+weight为正无穷，则不可能松弛
		return errors.New("distance is max")
	}

	if to.GetKey() > from.GetKey()+weight {
		to.SetKey(from.GetKey() + weight)
		to.SetParent(from)
	}
	return nil
}
