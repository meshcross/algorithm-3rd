/*
 * @Description: golang algorithm
 * @Author: wangchengdg@gmail.com
 * @Date: 2020-02-10 22:18:53
 * @LastEditTime: 2020-03-05 12:02:00
 * @LastEditors:
 */
package SingleSourceShortestPath

import (
	"errors"

	. "github.com/meshcross/algorithm-3rd/mesh/common"
	. "github.com/meshcross/algorithm-3rd/mesh/graph_algorithm/graph_struct"

	. "github.com/meshcross/algorithm-3rd/mesh/graph_algorithm/graph_struct/graph_vertex"
)

//!bellman_ford：单源最短路径的bellman_ford算法，算法导论24章24.1节
/*!
* \param graph:图
* \param source_id：最小生成树的根结点`id`
* \return: 是否不包含可以从源结点可达的权重为负值的环路。若返回值为true，则说明不包含可以从源结点可达的权重为负值的环路
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
*
* ## Bellman-Ford算法
*
* ### 算法原理
*
* Bellman-Ford算法解决的是一般情况下的单源最短路径问题。在这里边的权重可以为负值。给定带权的有向图G=(V,E)和权重函数w:E->R，Bellman-Ford
* 算法返回一个bool值，表明是否存在一个从源结点可达的权重为负值的环路。若存在这样的一个环路，算法告诉我们不存在解决方案；若不存在这样的环路，
* 算法将给出最短路径和它们的权重。
*
* Bellman-Ford算法通过对边的松弛操作来渐近的降低从源s到每个结点v的最短路径估计值v.Key，直到该估计值与实际的最短路径权重相同为止。
*
* ### 算法步骤
*
* - 执行单源最短路径的初始化过程
* - 进行|V|-1次处理，每次处理过程为：对图的每一条边进行一次松弛操作
* - 检查图中是否存在权重为负的环路并返回与之相适应的布尔值
*
*
* ### 算法性能
*
* 时间复杂度为O(VE)
*
 */
type BellmanFordShortestPath struct {
}

func NewBellmanFordShortestPath() *BellmanFordShortestPath {
	return &BellmanFordShortestPath{}
}

/**
运算完成之后能把每个点的Key更新为source_id到改点的最短路径值，并且相应设定各节点的Parent属性
**/
func (a *BellmanFordShortestPath) ShortestPath(graph *Graph, source_id int) (bool, error) {
	if graph == nil {
		return false, errors.New("initialize_single_source error: graph must not be nil!")
	}

	num := graph.N()
	if source_id < 0 || source_id >= num || graph.Vertexes[source_id] == nil {
		return false, errors.New("initialize_single_source error: source_id muse be in [0,N) and source vertex must not be nullptr!")
	}
	a.initializeSingleSource(graph, source_id)

	//************* 第一阶段 循环处理遍历所有的边  ***************
	//relax执行了n-1次，每次都relax所有edges ;relax会调整Parent属性和Key属性
	for i := 0; i < num-1; i++ {
		edges := graph.EdgeTuples()
		for _, edge := range edges {
			from := graph.Vertexes[edge.First]
			to := graph.Vertexes[edge.Second]
			wt := edge.Third

			a.relax(from, to, wt)
		}
	}
	//**********  第二阶段 检验是否存在从源点可达的【权重为负的环路】 *************
	for _, edge := range graph.EdgeTuples() {
		v1 := graph.Vertexes[edge.First]
		v2 := graph.Vertexes[edge.Second]
		wt := edge.Third
		if v2.GetKey() > v1.GetKey()+wt {
			return false, nil
		}
	}
	return true, nil
}

//!initialize_single_source：单源最短路径的初始化操作，算法导论24章24.1节
/*!
* \param graph:图，必须非空
* \param source_id：最小生成树的根结点`id`，必须有效。若无效则抛出异常
* \return: void
*
* `source_id`在以下情况下无效：
*
* - `source_id`不在区间`[0,N)`之间时，`source_id`无效
* - `graph`中不存在某个顶点的`id`等于`source_id`时，`source_id`无效
*
* 单源最短路径的初始化操作将所有的结点的`key`设置为正无穷，将所有结点的`parent`设为空。然后将源结点的`key`设为0。
*
* 性能：时间复杂度O(V)
*
 */
func (a *BellmanFordShortestPath) initializeSingleSource(graph *Graph, source_id int) error {

	if graph == nil {
		return errors.New("initialize_single_source error: graph must not be nil!")
	}

	num := graph.N()
	if source_id < 0 || source_id >= num || graph.Vertexes[source_id] == nil {
		return errors.New("initialize_single_source error: source_id muse belongs [0,N) and source vertex must not be nil!")
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

//!relax：单源最短路径的松弛操作，算法导论24章24.1节
/*!
* \param from:松弛有向边的起始结点，
* \param to：松弛有向边的终止结点，必须非空且不等于from
* \param weight:有向边的权重
* \return: error
*
*
* 对每一个结点v来说，我们维持一个属性v.key，它记录了从源结点s到结点v的最短路径权重的上界。我们称v.key为s到v的最短路径估计。
*
* 松弛过程是测试一下是否可以对从s到v的最短路径进行改善的过程，测试方法为：
* 将结点s到u之间的最短路径估计加上(u,v)边的权重，并与当前的s到v之间的最短路径估计进行比较。如果前者较小则对v.key和v.parent进行更新。
*
* 性能：时间复杂度O(1)
*
*
 */
func (a *BellmanFordShortestPath) relax(from, to IVertex, weight int) error {
	if from == nil || to == nil {
		return errors.New("relax error: from_vertex and to_vertex must not be nil!")
	}

	if from == to {
		return errors.New("relax error: from_vertex must not be to_vertex!")
	}

	//u.key+weight为正无穷，则不可能松弛
	if Is_Unlimit(from.GetKey() + weight) {
		return errors.New("weight is max")
	}

	if to.GetKey() > from.GetKey()+weight {
		to.SetKey(from.GetKey() + weight)
		to.SetParent(from)
	}
	return nil
}
