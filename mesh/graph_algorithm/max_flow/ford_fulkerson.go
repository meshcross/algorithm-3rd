package MaxFlow

import (
	"errors"

	. "github.com/meshcross/algorithm-3rd/mesh/common"
	. "github.com/meshcross/algorithm-3rd/mesh/graph_algorithm/basic_graph"
	. "github.com/meshcross/algorithm-3rd/mesh/graph_algorithm/graph_struct"
	. "github.com/meshcross/algorithm-3rd/mesh/graph_algorithm/graph_struct/graph_vertex"
)

//!MaxFlow：最大流的ford_fulkerson算法。算法导论26章26.2节
/*!
* 最大流问题：给定流网络G、一个源结点s、一个汇点t,找出值最大的一个流
*
* ## ford_fulkerson算法
*
* ### 算法原理
*
* ford_fulkerson 算法更准确的称呼是一种方法而不是算法，因为它包含了几种运行时间各不相同的具体实现。
* 它依赖于三种重要思想：残余网络、增广路径和切割。
*
* #### 残余网络
*
*  给定流网络G(V,E)和流量f，残余网络Gf(V,Ef)由那些仍有空间对流量进行调整的边构成。
*  Gf中的顶点就是原图G中的顶点。残余网络Gf中的边可以由以下组成：
*
* - 若(u,v)属于E，且f(u,v)<c(u,v)，则存在边(u,v)属于Ef，且cf(u,v)=c(u,v)-f(u,v)，表示沿着该方向图G中还能流通cf大小的流量
* - 若(u,v)属于E，则存在边(v,u)属于Ef,且 cf(v,u)=c(u,v)，表示沿着(u,v)的反方向（即(v,u)方向)可以压入cf大小的反向流量
* > 这里f(u,v)为边(u,v)上的流；c(u,v)为图G的边(u,v)的容量；cf(u,v)为残余网络Gf上的边(u,v)的容量
*
* #### 增广路径
*
* 给定流网络G=(V,E)和流f，增广路径p是残余网络Gf中一条源结点s到汇点t的简单路径。我们称一条增广路径
* p上能够为每条边增加的流量的最大值为路径p的残余容量，定义为 cf(p)= min {cf(u,v):(u,v)属于路径p}
*
* #### 最大流和最小切割定理
*
* 流网络G=(V,E)的一个切割(S,T)是将结点集合V划分为S和T=V-S两个集合，使得s属于S，t属于T。切割(S,T)
* 的容量c(S,T)=c(u,v)的累加，其中u属于S,v属于T
*
* 设f为流网络G=(V,E)中的一个流，该流网络的源点为s，汇点为t，则下面的条件是等价的：
*
* - f是G的一个最大流
* - 残余网络Gf不包括任何增广路径
* - |f|=c(S,T)，其中(S,T)是流网络G的某个切割
*
* ### 算法步骤
*
* ford_fulkerson算法循环增加流的值。在开始的时候对于所有的结点u,v属于V，f(u,v)=0。在每一轮迭代中，
* 我们将图G的流值进行增加，方法就是在G的残余网络Gf中寻找一条增广路径。算法步骤如下：
*
* - 初始化： 创建流矩阵flow，flow[i][j]均初始化为0
* - 循环迭代：
*   - 根据flow和G创建残余网络Gf，寻找增广路径
*   - 若增广路径不存在，则跳出迭代，证明现在的流就是最大流
*   - 若增广路径存在，则更新流矩阵。更新方法为：
*       - 取出增广路径的残余容量 cf(p)= min {cf(u,v):(u,v)属于路径p}，然后更新增广路径p中的边(u,v)：
*           - 若(u,v)属于图G的边E，则 flow[u][v]=flow[u][v]+cf(p)
*           - 若(v,u)属于图G的边E，则 flow[u][v]=flow[u][v]-cf(p)
* - 返回流矩阵 flow
*
* ### 算法性能
*
* ford_fulkerson算法运行时间取决于如何寻找增广路径。如果选择不好，算法可能不会终止：流的值会随着后续的递增而增加，
* 但是它却不一定收敛于最大的流值。如果用广度优先搜索来寻找增广路径，算法的运行时间将会是多项式数量级。
*
* 假定最大流问题中的容量均为整数，由于流量值每次迭代中最少增加一个单位，因此算法运行时间为 O(E|f*|) ,|f*|为最大流的值
*
*
 */
type FordFulkerson struct {
}

type FordFulkersonActionFunc func(id int)

func NewFordFulkerson() *FordFulkerson {
	return &FordFulkerson{}
}

/**
* \param graph:指定流网络,必须非空
* \param src_id: 流的源点
* \param dst_id: 流的汇点
* \return: 最大流矩阵,error
**/
func (a *FordFulkerson) MaxFlow(graph *Graph, src_id, dst_id int) ([][]int, error) {

	if graph == nil {
		return nil, errors.New("FordFulkerson error: graph must not be nil!")
	}

	num := graph.N()
	if src_id < 0 || src_id >= num || graph.Vertexes[src_id] == nil {
		return nil, errors.New("FordFulkerson error: src_id muse belongs [0,N) and src vertex must not be nil!")
	}

	if dst_id < 0 || dst_id >= num || graph.Vertexes[dst_id] == nil {
		return nil, errors.New("FordFulkerson error: dst_id muse belongs [0,N) and dst vertex must not be nil!")
	}

	flow := make([][]int, num)
	for k, _ := range flow {
		flow[k] = make([]int, num)
	}
	for i := 0; i < num; i++ {
		for j := 0; j < num; j++ {
			flow[i][j] = 0
		}
	}

	//残余网络
	var graphF *Graph = nil
	bfs := NewGraphBFS()
	for {
		graphF, _ = a.getResidulalNetwork(graph, flow) //创建残余网络
		//************ 寻找增广路径  *************

		bfs.Search(graphF, src_id, nil, nil)
		path, _ := GetPath(graphF.Vertexes[src_id], graphF.Vertexes[dst_id])
		path_len := len(path)
		if path_len == 0 {
			break //不存在增广路径
		}

		cf := Unlimit()
		//求增广路径的残余容量,取整个路径上最小的流量
		for i := 1; i < path_len; i++ {
			u := path[i-1]
			v := path[i]
			w, _ := graphF.Weight(u, v)
			if w < cf {
				cf = w
			}
		}
		//*****************  更新流矩阵  ***************
		for i := 1; i < path_len; i++ {
			u := path[i-1]
			v := path[i]
			uv, _ := graph.HasEdge(u, v)
			vu, _ := graph.HasEdge(v, u)
			if uv { //(u,v)属于E
				flow[u][v] += cf
			} else if vu { //(v,u)属于E
				flow[v][u] -= cf
			}
		}
	}
	return flow, nil
}

//!getResidulalNetwork：根据指定流网络生成一个残余网络。算法导论26章26.2节
/*!
*
* \param graph:指定流网络
* \param flow: 一个流
* \return: 残余网络
*
* ## 残余网络
* 给定网络G(V,E)和流量f，残余网络Gf(V,Ef)由那些仍有空间对流量进行调整的边构成。Gf中的顶点就是原图G中的顶点。残余网络Gf中的边可以由以下组成：
*
* - 若(u,v)属于E，且f(u,v)<c(u,v)，则存在边(u,v)属于Ef，且cf(u,v)=c(u,v)-f(u,v)，表示沿着该方向图G中还能流通cf大小的流量
* - 若(u,v)属于E，则存在边(v,u)属于Ef,且 cf(v,u)=c(u,v)，表示沿着(u,v)的反方向（即(v,u)方向)可以压入cf大小的反向流量
* > 这里f(u,v)为边(u,v)上的流；c(u,v)为图G的边(u,v)的容量；cf(u,v)为残余网络Gf上的边(u,v)的容量
*
* 这里假定容量c>0，一旦容量c=0表示边不存在（禁止流通）；f>=0，表示流是正向的。同时假定图中不存在双向边：
* 图G中不可能同时存在边(u,v)以及边(v,u)（即管道是单向的）
*
* 要求graph的无效权重为0，否则抛出异常
*
* 性能：时间复杂度 O(V+E)
*
 */
//计算残余网络
func (a *FordFulkerson) getResidulalNetwork(graph *Graph, flow [][]int) (*Graph, error) {

	if graph == nil {
		return nil, errors.New("getResidulalNetwork error: graph must not be nil!")
	}

	if graph.Matrix.InvalidWeight() != 0 {
		return nil, errors.New("getResidulalNetwork error: graph invalid weight must be 0!")
	}

	num := graph.N()

	creator := func(key, id int) IVertex {
		ptr := NewBFSVertex(key, id)
		return ptr
	}
	new_graph := NewGraph(graph.Matrix.InvalidWeight(), graph.N(), creator)
	//*************  创建新图的顶点  ******************
	for i := 0; i < num; i++ {
		vertex := graph.Vertexes[i]
		if vertex != nil {
			new_graph.AddVertex(vertex.GetKey(), vertex.GetID())
		}
	}

	//**********   创建新图的边  ***********************
	new_edges := []*Tuple{}
	for i := 0; i < num; i++ {
		for j := 0; j < num; j++ {
			if i == j {
				continue
			}
			if graph.Vertexes[i] == nil || graph.Vertexes[j] == nil {
				continue
			}
			uv, _ := graph.HasEdge(i, j)
			vu, _ := graph.HasEdge(j, i)
			if uv { //(u,v)属于E
				wt, _ := graph.Weight(i, j)
				new_edges = append(new_edges, NewTuple(i, j, wt-flow[i][j]))
			} else if vu { //(v,u)属于E
				new_edges = append(new_edges, NewTuple(i, j, flow[j][i]))
			}
		}
	}
	new_graph.AddEdges(new_edges)
	return new_graph, nil
}
