package MinimumSpanningTree

import (
	"errors"

	. "github.com/meshcross/algorithm-3rd/mesh/common"
	. "github.com/meshcross/algorithm-3rd/mesh/graph_algorithm/graph_struct"
	. "github.com/meshcross/algorithm-3rd/mesh/graph_algorithm/graph_struct/graph_vertex"
	. "github.com/meshcross/algorithm-3rd/mesh/queue_algorithm"
)

//!prim：最小生成树的Prim算法，算法导论23章23.2节
/*!
* \param graph:图，必须非空
* \param source_id：最小生成树的根结点`id`，必须有效。若无效则抛出异常
* \param pre_action:一个可调用对象，在每次从最小优先级队列中弹出最小顶点时立即调用，调用参数为该顶点的`id`。默认为空操作，即不进行任何操作
* \param post_action:一个可调用对象，在每次从最小优先级队列中弹出最小顶点并处理完它的边时立即调用，调用参数为该顶点的`id`。默认为空操作，即不进行任何操作
* \return: 最小生成树的权重
*
* `source_id`在以下情况下无效：
*
* - `source_id`不在区间`[0,N)`之间时，`source_id`无效
* - `graph`中不存在某个顶点的`id`等于`source_id`时，`source_id`无效
*
* ## 最小生成树
*
* 最小生成树：对于一个连通无向图G=(V,E)，对于每一条边(u,v)属于E都赋予了一个权重w(u,v)。我们希望找出一个无环子集T，其中T为E的子集，使得所有的顶点V位于T中，
* 同时T具有最小的权重。由于T是无环的，且连通所有结点因此T必然是一棵树。我们称这样的树为生成树。称从G中求取该生成树的问题为最小生成树问题。
*
* 通用的最小生成树使用贪心策略。该策略在每个时刻找到最小生成树的一条边，并在整个策略过程中维持循环不变式：边的集合A在每次循环之前是某棵最小生成树的一个子集。
*
* 在每一步，选择一条边(u,v)，将其加入集合A中，使得A不违反循环不变式。称这样的边(u,v)为边集合A的安全边。
*
* ## Prim算法
*
* ### 算法原理
*
* 在Prim算法所具有的一个性质是集合A中的边总是构成一棵树。这棵树从一个任意的根结点r开始，一直长大到覆盖V中的所有结点为止。算法每一步在连接集合A和A之外的结点的所有边中，
* 选择一条边加入到A中（经特殊选择的边）。
*
* 为了有效地实现Prim算法，需要一种快速的方法来选择一条新的边以便加入到由集合A中的边所构成的树中。在算法执行过程中，所有不在树A中的结点都存放在一个基于key的属性的最小优先级队列Q中。
* 对于每个结点v，属性v.key保存的是连接v和树中结点的所有边中最小边的权重。若这样的边不存在则权重为正无穷。属性v.pai给出的是结点v在树中的父结点。
*
* ### 算法步骤
*
* - 初始化：将所有结点的key设为正无穷，所有结点的父结点置为空(结点构造时，父结点默认为空）
* - 设置源点：将源点的key设为0，
* - 构造最小优先级队列：将所有顶点放入最小优先级队列Q中
* - 循环，直到最小优先级队列为空。循环中执行下列操作：
*   - 弹出最小优先级队列的头部顶点u
*   - 从结点u出发的所有边，找出它的另一端的结点v。如果v也在Q中，且w(u,v)<v.key，则证明(u,v)是v到集合A的最短边，因此设置v.pai=u,v.key=w(u,v)
*   >这里隐含着一个最小优先级队列的decreate_key操作
*
* ### 算法性能
*
* Prim总时间代价为O(VlgV+ElgV)=O(ElgV)(使用最小堆实现的最小优先级队列），或者O(E+VlgV)（使用斐波那契堆实现最小优先级队列）
 */

type PrimMST struct {
}

func NewPrimMST() *PrimMST {
	return &PrimMST{}
}

type PrimMSTActionFunc func(id int)

// func NodeCompareFunc_VertexLessThan(vtx1, vtx2 IVertex) int {
// 	if vtx1 == nil || vtx2 == nil {
// 		return -1
// 	}
// 	if vtx1.GetKey() < vtx2.GetKey() {
// 		return 1
// 	}
// 	if vtx1.GetKey() == vtx2.GetKey() {
// 		return 0
// 	}
// 	return -1
// }

func (a *PrimMST) Generate(graph *Graph, source_id int, pre_action, post_action PrimMSTActionFunc) (int, error) {

	if graph == nil {
		return -1, errors.New("prim error: graph must not be nil!")
	}

	num := graph.N()
	vertex := graph.Vertexes[source_id]

	if source_id < 0 || source_id > num || vertex == nil {
		return -1, errors.New("prim error: source_id is not in limit!")
	}

	//最小优先队列
	q := NewMinQueue(NodeCompareFunc_VertexLessThan, nil)
	for i := 0; i < num; i++ {
		vertex := graph.Vertexes[i]
		if vertex != nil {

			vertex.SetParent(nil)
			if i == source_id {
				vertex.SetKey(0)
			} else {
				vertex.SetKey(Unlimit())
			}

			q.Insert(vertex)
		}
	}

	weight := 0
	for !q.IsEmpty() {

		u, _ := q.ExtractMin()
		minNode, ok := u.(*Vertex)
		if !ok || minNode == nil {
			continue
		}

		if pre_action != nil {
			pre_action(minNode.GetID())
		}
		edges := graph.EdgeTuples()
		for _, edge := range edges {
			other_id := edge.Second
			other_vtx := graph.Vertexes[other_id]
			other_weight := edge.Third

			index := q.ElementIndex(other_vtx)
			if index >= 0 && other_weight < other_vtx.GetKey() {
				other_vtx.SetParent(minNode)
				other_vtx.SetKey(other_weight)

				//q.DecreateKey(index, other_weight)
				q.DecreateKey(index, other_vtx)
			}
		}

		if post_action != nil {
			post_action(minNode.GetID())
		}
		weight += minNode.GetKey()
	}

	return weight, nil
}
