/*
 * @Description: 第24章 24.3 单源最短路径的dijkstra算法
 * @Author: wangchengdg@gmail.com
 * @Date: 2020-02-18 10:35:17
 * @LastEditTime: 2020-03-14 20:23:50
 * @LastEditors:


 *					有权重的有向图上单源最短路径问题,Dijkstra算法要求所有边的权重都为非负值，算法会维护一个最小优先队列。
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
 * 需要指出的是：最短路径不一定是唯一的。
 *
 *
 * ## Dijkstra算法
 * ### 算法原理
 *
 * Dijkstra算法解决的是有向图中的单源最短路径问题。Dijkstra算法要求所有边的权重都为非负值。
 *
 * Dijkstra算法核心信息是一组结点集合 S 。从源结点s 到集合 S 中的每一个结点之间的最短路径已经被找到。算法重复从结点集合 V-S 中选择最短路径估计最小的结点u,
 * 将u加入到集合S中，然后对所有从u出发的边进行松弛。本算法利用最小优先队列Q来保存结点集合，每个结点的关键字为它的key值。
 *
 */
package SingleSourceShortestPath

import (
	"errors"

	. "github.com/meshcross/algorithm-3rd/mesh/common"
	. "github.com/meshcross/algorithm-3rd/mesh/graph_algorithm/graph_struct"
	. "github.com/meshcross/algorithm-3rd/mesh/queue_algorithm"

	. "github.com/meshcross/algorithm-3rd/mesh/graph_algorithm/graph_struct/graph_vertex"
)

type Dijkstra struct {
}

func NewDijkstra() *Dijkstra {
	return &Dijkstra{}
}

/*!
 * @description:单源最短路径的dijkstra算法
 * @param graph:图
 * @param source_id：最小生成树的根结点`id`
 * @return: error
 *
 * ### 算法步骤
 *
 * - 调用 `initializeSingleSource`函数对图的顶点进行初始化
 * - 将集合S置为空，将所有顶点放入最小优先队列Q
 * - 循环操作直到Q(MinQueue)为空，循环内执行以下操作：
 *   - 弹出最小优先级队列队首元素u
 *   - 将u对应的结点放入集合S中
 *   - 对从u出发的边且另一端在Q中的边进行松弛。松弛过程隐含着Q的一个`DecreateKey()`方法的调用
 *
 * 算法实现中，集合S并没有什么用处,所以注释掉了
 *
 * ### 算法性能
 *
 * 时间复杂度为O(V^2+E)
 *
 */
func (a *Dijkstra) ShortestPath(graph *Graph, source_id int) error {
	if graph == nil {
		return errors.New("ShortestPath error: graph must not be nil!")
	}

	num := graph.N()
	if source_id < 0 || source_id >= num || graph.Vertexes[source_id] == nil {
		return errors.New("ShortestPath error: source_id muse belongs [0,N) and source vertex must not be nil!")
	}

	//sets := []*SetVertex{}

	//************* 第一阶段 初始化  ***************
	a.initializeSingleSource(graph, source_id)

	//************* 第二阶段 构建最小优先队列  ***************
	//注意，此次可以使用斐波那契堆，能获得更好的性能
	q := NewMinQueue(NodeCompareFunc_VertexLessThan, nil)
	for i := 0; i < num; i++ {
		vertex := graph.Vertexes[i]
		if vertex != nil {
			q.Insert(vertex)
		}
	}

	//************* 第三阶段 从最小优先队列中提取元素u，不断的relax结点u的相关的边  ***************
	for !q.IsEmpty() {

		//把Key最小的从队列中提出来,所以ExtractMin和DecreaseKey对该算法的性能影响很大
		//使用斐波那契堆能改善性能
		u, _ := q.ExtractMin()
		minNode, ok := u.(IVertex)
		if !ok || minNode == nil {
			continue
		}

		edges, _ := graph.VertexEdgeTuples(minNode.GetID())
		for _, edge := range edges {
			other_id := edge.Second
			other_vtx := graph.Vertexes[other_id]
			other_weight := edge.Third

			a.relax(minNode, other_vtx, other_weight)

			index := q.ElementIndex(other_vtx)
			if index >= 0 {
				q.DecreateKey(index, other_vtx)
			}
		}
	}
	return nil
}

func (a *Dijkstra) initializeSingleSource(graph *Graph, source_id int) error {
	if graph == nil {
		return errors.New("initializeSingleSource error: graph must not be nil!")
	}

	num := graph.N()
	unlimit := Unlimit()
	if source_id < 0 || source_id >= num || graph.Vertexes[source_id] == nil {
		return errors.New("initializeSingleSource error: source_id muse belongs [0,N) and source vertex must not be nil!")
	}

	//**************** 设置所有结点 *****************
	for i := 0; i < num; i++ {

		vertex := graph.Vertexes[i]
		if vertex != nil {
			vertex.SetKey(unlimit)
			vertex.SetParent(nil)
		}
	}
	//**************  设置源结点 ***************
	vertex := graph.Vertexes[source_id]
	vertex.SetKey(0)

	return nil
}

func (a *Dijkstra) relax(from, to IVertex, weight int) error {
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
