/*
 * @Description: 第23章23.2节 最小生成树的Kruskal算法  注意：G是无环的，且连通所有节点
 * @Author: wangchengdg@gmail.com
 * @Date: 2020-02-10 22:18:45
 * @LastEditTime: 2020-03-14 12:40:43
 * @LastEditors:


 * ## 最小生成树
*
* 最小生成树：对于一个连通无向图G=(V,E)，对于每一条边(u,v)属于E都赋予了一个权重w(u,v)。我们希望找出一个无环子集T，其中T为E的子集，使得所有的顶点V位于T中，
* 同时T具有最小的权重。由于T是无环的，且连通所有结点因此T必然是一棵树。我们称这样的树为生成树。称从G中求取该生成树的问题为最小生成树问题。
*
* 通用的最小生成树使用贪心策略。该策略在每个时刻找到最小生成树的一条边，并在整个策略过程中维持循环不变式：边的集合A在每次循环之前是某棵最小生成树的一个子集。
*
* 在每一步，选择一条边(u,v)，将其加入集合A中，使得A不违反循环不变式。称这样的边(u,v)为边集合A的安全边。
*
* ## Kruskal 算法
*
* ### 算法原理
*
* 在Kruskal算法中集合A是一个森林，其结点就是G的结点。Kruskal算法找到安全边的办法是：
* 在所有连接森林中两棵不同树的边里面，找到权重最小的边(u,v)。
* Kruskal算法使用一个不相交集合数据结构来维护几个互不相交的元素集合。每个集合代表当前森林中的一棵树
*

*/
package MinimumSpanningTree

import (
	"errors"
	"sort"

	. "github.com/meshcross/algorithm-3rd/mesh/common"
	. "github.com/meshcross/algorithm-3rd/mesh/graph_algorithm/graph_struct"
	. "github.com/meshcross/algorithm-3rd/mesh/set_algorithm"
)

type KruskalMST struct {
}

func NewKruskalMST() *KruskalMST {
	return &KruskalMST{}
}

type KruskalMSTActionFunc func(v, id int)

/*
* @description:Kruskal算法
* @param graph:图
* @param pre_action:在每次从最小优先级队列中弹出最小顶点时立即调用，回调函数
* @param post_action:在每次从最小优先级队列中弹出最小顶点并处理完它的边时立即调用，回调函数
* @return: 最小生成树的权重,以及最小生成树的所有边
*

* ### 算法步骤
*
* - 初始化：将集合A置为空；对G中的每一个结点v,以它为根构造一棵单根树
* - 将G中的边E按照权重单调递增的顺序排序
* - 循环挑选E中的(u,v)，按照单调增的顺序。在循环内执行：
*   - 如果u所在的树不等于v所在的树，则将(u,v)加入A中，并且合并u所在的树与v所在的树
*
* >根据算法的特征，如果图中只有一个顶点，算法得到的集合A为空集；但是实际上集合A应该包含该顶点。这是算法在极端情况下的BUG。
*
* ### 算法性能
*
* Kruskal算法运行时间依赖于不相交集合数据结构的实现方式。如果采用算法导论21.3节讨论的不相交集合森林实现（也是我在src/set_algorithms/disjoint_set中实现的），
* 则Kruskal算法的时间为 O(ElgV)
*
 */
func (a *KruskalMST) Generate(graph *Graph, pre_action, post_action KruskalMSTActionFunc) (int, []*Tuple, error) {
	if graph == nil {
		return 0, nil, errors.New("kruskal error: graph must not be nil!")
	}
	sets := []*DisJointSetNode{}
	num := graph.N()

	for i := 0; i < num; i++ {
		wp := graph.Vertexes[i]
		if wp != nil { //添加顶点到`disjoint_set`中
			vertex := ToSetVertex(wp)
			set_node := NewDisJointSetNode(vertex)
			sets = append(sets, set_node)
			vertex.Node = set_node
			MakeSet(set_node)
		}
	}
	//****************** 循环  ************************
	weight := 0
	edges := graph.EdgeTuples()
	new_edges := []*Tuple{}

	//需要将边按照权重排序

	sorter := NewTupleWapper(edges, TupleCompareFunc_Less)
	sort.Sort(sorter)

	for _, edge := range edges {
		from_id := edge.First
		to_id := edge.Second
		edge_weight := edge.Third

		vtx_from := ToSetVertex(graph.Vertexes[from_id])
		vtx_to := ToSetVertex(graph.Vertexes[to_id])

		from_vertex_set_node := vtx_from.Node
		to_vertex_set_node := vtx_to.Node
		findFrom, _ := FindSet(from_vertex_set_node) //找到树的根节点
		findTo, _ := FindSet(to_vertex_set_node)

		//如果不是同一棵树，则进行合并操作
		if findFrom != findTo {
			if pre_action != nil {
				pre_action(from_id, to_id)
			}

			new_edges = append(new_edges, edge)
			UnionSet(from_vertex_set_node, to_vertex_set_node)
			weight += edge_weight
			if post_action != nil {
				post_action(from_id, to_id)
			}
		}
	}
	return weight, new_edges, nil

}
