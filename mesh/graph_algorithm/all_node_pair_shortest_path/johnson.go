/*
 * @Description: 第25章25.3节 所有节点对的最短路径之 johnson算法 ，扩展了一个新的节点
 * @Author: wangchengdg@gmail.com
 * @Date: 2020-02-28 22:41:21
 * @LastEditTime: 2020-03-15 13:14:25
 * @LastEditors:



* ## 所有结点对的最短路径
*
* 给定一个带权重的有向图G=(V,E)，其权重函数为w:E->R，该函数将边映射到实值权重上。我们希望找到对于所有的结点对u,v属于V，找出一条从结点u
* 到结点v的最短路径，以及这条路径的权重。
*
* 与单源最短路径不同中使用邻接表来表示图不同，本章的多数算法使用邻接矩阵来表示图。该矩阵代表的是一个有n个结点的有向图G=(V,E)的边的权重W=(w_i_j)，
* 其中 w_i_j =:
*
*   - 0:若i=j
*   - 有向边(i,j)的权重，若i!=j且(i,j)属于E
*   - 正无穷，若 i!=j且(i,j)不属于E
*
* 我们允许存在负权重的边，目前仍然假定图中不存在权重为负值的环路。
*
* 本章讨论的所有结点对最短路径的算法的输出也是一个n*n的矩阵D=(d_i_j)，其中 d_i_j 代表的是结点i到j的一条最短路径的权重。
*
* 有时候为了解决所有结点对最短路径问题，我们不仅要计算出最短路径权重，还需要计算出前驱结点矩阵 II=(pai_i_j)，其中 pai_i_j在i=j或者从i到j
* 不存在路径时为NIL，在其他情况下给出的是从结点i到结点j的某条最短路径上结点j的前驱结点。由矩阵II的第i行所诱导的子图应当是一棵根节点为i
* 的最短路径树。
*
* ## Johnson算法
*
* ### 算法原理
*
* 对于稀疏图来说Johnson算法的渐近表现要优于重复平方法和Floyd-Warshall算法。Johnson算法要么返回一个包含所有结点对的最短路径权重的矩阵，
* 要么报告输入图中包含一个权重为负值的环路。Johnson算法在运行中需要使用Dijkstra算法和Bellman-Ford算法作为自己的子程序。
*
* Johnson算法使用一种称为重赋权重的技术。若图G=(V,E)中所有边权重w都为非负值，则可以通过对每一对顶点运行一次Dijkstra
* 算法来找到所有顶点对直接的最短路径。如果图G包含负权重的边，但没有权重为负值的环路，则需要计算出一组新的非负权重值，然后用同样的方法。
* 新赋的权重w'要求满足一下性质：
*
* - 对于所有结点对u,v属于V，一条路径p实在使用权重函数w时从u到v的一条最短路径，当且仅当p实在使用权重函数w'
* 时从u到v的一条最短路径
* - 对于所有的边(u,v)，新权重w'(u,v)非负
*
* 对于给定的有向图G=(V,E)，权重函数w：E->R 。我们制作一副新图G'=(V',E')，其中V'=V并上{s}，s是一个新结点，s不属于V。
* E'=E并上{(s,v);v属于V}。我们对权重函数w进行扩展，使得对于所有结点v属于V，有w(s,v)=0。
* 对所有的结点v属于V'，我们定义h(v)=delt(s,v)。定义w'(u,v)=w(u,v)+h(u)-h(v)
* w'(p)=w(p)+h(v_0)-h(v_k)
*

*/
package AllNodePairShortestPath

import (
	"errors"

	. "github.com/meshcross/algorithm-3rd/mesh/graph_algorithm/graph_struct"

	. "github.com/meshcross/algorithm-3rd/mesh/common"
	. "github.com/meshcross/algorithm-3rd/mesh/graph_algorithm/single_source_shortest_path"
)

type JohnsonSP struct {
}

func NewJohnsonSP() *JohnsonSP {
	return &JohnsonSP{}
}

/**
* @description:根据图graph生成一个新图
* @param graph:有向图
* @return: 一个新图
*
* 根据图graph生成一个新图new_graph,在原图基础上添加一个新的节点作为源节点，该节点和其他所有边相连，且边权重为0。
* 主要用于稀疏图
*
* new_graph的顶点 graph的顶点加上一个新顶点s；
* new_graph的边为graph的边加上{(s,v):v属于graph的顶点};
* new_graph的边的权重为graph的边权重，以及 w(s,v)=0
*
 */
func (a *JohnsonSP) graph_add_one_vertex(graph *Graph) (*Graph, error) {

	if graph == nil {
		return nil, errors.New("graph_add_one_vertex error: graph must not be nil!")
	}
	num := graph.N()
	invalid_weight := graph.Matrix.InvalidWeight()
	new_graph := NewGraph(invalid_weight, num+1, graph.VertexCreator)

	//*************  创建新图的顶点  ******************
	for i := 0; i < num; i++ {
		vertex := graph.Vertexes[i]
		if vertex != nil {
			new_graph.AddVertex(vertex.GetKey(), vertex.GetID())
		}
	}
	new_graph.AddVertex(0, num) //新的顶点s，id为num，新的图多一个顶点

	//**********   创建新图的边  ***********************
	source_edges := graph.EdgeTuples()
	// 生成边 (s,v)，s只有出的边没有入的边，且w(s,v)权重为0
	for i := 0; i < num; i++ {
		source_edges = append(source_edges, NewTuple(num, i, 0))
	}
	new_graph.AddEdges(source_edges)
	return new_graph, nil
}

/**
* @description:返回所有节点对的最短路径的johnson算法
* @param graph:有向图
* @return: 一个n*n的矩阵(d_i_j)，其中 d_i_j 代表的是结点i到j的一条最短路径的权重
*
* ### 算法步骤
*
* - 重赋权重：
*   - 创建新图 new_graph
*   - 对新图执行 bellman_ford 的调用，源点为新创建的结点s
*   - 如果有负权值的环路，则抛出异常。
*   - 如果没有负权重环路，则创建h函数，并对new_graph中的所有边执行重新赋权
* - 在 new_graph上，除了新的顶点s之外的所有顶点v,以v为源顶点执行dijkstra过程。D[i][j]等于 new_graph 中以i为源点到j的最短路径的权重的修正值，
* 修正的方法就是重新赋权的逆过程。
* - 返回矩阵 D
*
* ###算法性能
*
*  时间复杂度 O（V^2 lgV + VE)
 */
func (a *JohnsonSP) ShortestPath(graph *Graph) ([][]int, error) {

	if graph == nil {
		return nil, errors.New("johnson error: graph must not be nil!")
	}

	//*******************  第一阶段 重赋权值  **************
	new_graph, _ := a.graph_add_one_vertex(graph) //新的图
	num := graph.N()
	numNew := new_graph.N()

	//*******************  第二阶段 bellmanFord对图进行调整  **************
	//bellmanFord会对各边权重进行调整
	bellmanFord := NewBellmanFordShortestPath()
	//新顶点s的id为 num，AddVertex时候设定的
	//bellmanFord算法之后，会算出每个节点到s点的最短路径，并且设定好Parent关系
	b, _ := bellmanFord.ShortestPath(new_graph, num)
	if !b {
		//不能有权重为负的环路
		return nil, errors.New("johnson error: graph has a nagative-weight circle!")
	}

	//*******************  第三阶段 bellmanFord已经获得了新的权重，H函数将new_graph的权重调整到非负  **************
	//创建h函数， h(v)=delt(s,v)
	H := make([]int, numNew)
	for i := 0; i < num; i++ {
		H[i] = new_graph.Vertexes[i].GetKey()
	}

	//通过重新赋值生成非负权重，可以对比该循环前后graph和new_graph的Matrix.Matrix属性
	for i := 0; i < numNew; i++ {
		for j := 0; j < numNew; j++ {
			has_edge, _ := new_graph.HasEdge(i, j)
			//只有边存在的情况下才调整
			if has_edge {
				wt, _ := new_graph.Weight(i, j)
				new_graph.AdjustEdge(i, j, wt+H[i]-H[j])
			}
		}
	}

	//******************  第四阶段：在新图上以每个顶点为源点，计算单源最短路径  *********
	dijkstra := NewDijkstra()
	D := NewMatrix(num, 0)
	//剔除新顶点s作为源点
	for i := 0; i < num; i++ {
		//该算法要求所有边的权重都非负，所以需要H函数调整权重，调整之后还需要恢复权重
		//dijkstra算法比bellmanFord算法性能更高一些，但是对graph有要求
		dijkstra.ShortestPath(new_graph, i)
		for j := 0; j < num; j++ {
			D[i][j] = new_graph.Vertexes[j].GetKey() + H[j] - H[i] // 恢复权值
		}
	}
	return D, nil
}
