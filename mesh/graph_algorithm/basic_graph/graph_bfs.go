package BasicGraph

import (
	"errors"

	. "github.com/meshcross/algorithm-3rd/mesh/common"
	. "github.com/meshcross/algorithm-3rd/mesh/graph_algorithm/graph_struct"
	. "github.com/meshcross/algorithm-3rd/mesh/graph_algorithm/graph_struct/graph_vertex"
)

//!breadth_first_search：图的 广度优先搜索，算法导论22章22.2节
/*!
* \param graph:指向图的强引用，必须非空。若为空则抛出异常
* \param source_id：广度优先搜索的源点`id`，必须有效。若无效则抛出异常
* \param pre_action:一个可调用对象，在每次发现一个顶点时调用，调用参数为该顶点的`id`。默认为空操作，即不进行任何操作
* \param post_action:一个可调用对象，在每次对一个顶点搜索完成时调用，调用参数为该顶点的`id`。默认为空操作，即不进行任何操作
* \return:void
*
* `source_id`在以下情况下无效：
*
* - `source_id`不在区间`[0,N)`之间时，`source_id`无效
* - `graph`中不存在某个顶点的`id`等于`source_id`时，`source_id`无效
*
*
* 广度优先搜索：该算法维护已经发现结点和未发现结点的边界，沿着其广度方向向外扩展。每个结点有三种颜色：白色、灰色、黑色。
* 白色结点表示未发现；灰色结点表示已发现但是未处理完成；黑色结点表示已处理完成。其中灰色结点就是边界。
*
* 给定图 G=(V,E)和一个可以识别的源点s。所有的结点在一开始都被涂上白色。每个结点的颜色存放在属性Color中；
* 每个结点的前驱结点放在属性Parent中。每个结点的属性Key存放的是从源点到本结点的距离。该算法使用一个先进先出的队列Q来管理灰色结点集。
*
* - 将所有结点涂为白色，Key属性设置为正无穷，父结点置为空；
* - 将源点涂为灰色，源点前驱设为空，源点的Key设为0；
* - 将源点加入队列Q中；Q中存放的都是已发现但是尚未处理完成的结点
* - 循环直到队列Q为空，在循环中执行以下操作：
*   - 取出队列Q头部的结点`v`
*   - 对结点`v`的邻接表中的白色结点进行发现操作，并将这些结点加入队列Q中
*   - 对结点`v`染成黑色；
*
* 算法的时间复杂度为 O(E+V)
*
* 最短路径：广度优先搜索能找出给定源结点s到所有可以到达的结点之间的距离。定义从源s到结点v之间的最短路径距离 delt(s,v) 为从结点s到v之间的所有路径里面最少的边数。
* 如果从s到v没有路径，则 delt(s,v)=正无穷大 。我们定义从s到v之间的长度为 delt(s,v) 的路径为 s 到 v 的最短路径。可以证明：广度优先搜索可以正确计算出最短路径距离。
*
* 广度优先树：对于G=(V,E)和源点s，定义图G的前驱子图为 G_pai=(V_pai,E_pai)，其中 V_pai={ v属于V: v.parent!=NIL}并上{s}，E_pai={(v.parent,v):v属于(V_pai-{s})}。
* 即V_pai由从源s可达的所有结点组成（包括s本身），E_pai由V_pai中去掉s之后的结点的入边组成，其中该入边的对端为结点的父结点。
* BFS算法获取的前驱子图G_pai包含一条从源结点s到结点v的唯一简单路径，而且该路径也是图G里面从源s到v之间的一条最短路径，因此前驱子图也称为广度优先树。
*
 */
type GraphBFS struct {
}

func NewGraphBFS() *GraphBFS {
	return &GraphBFS{}
}

type BFSActionFunc func(id int)

func (a *GraphBFS) toBFSVertex(vtx IVertex) *BFSVertex {
	// ptr := unsafe.Pointer(vtx)
	// v := (*BFSVertex)(ptr)
	// return v
	if v, ok := vtx.(*BFSVertex); ok {
		return v
	}
	return nil
}

func (a *GraphBFS) Search(graph *Graph, source_id int, pre_action BFSActionFunc, post_action BFSActionFunc) error {

	if graph == nil {
		return errors.New("breadth_first_search error: graph must not be nullptr!")
	}
	num := graph.N()
	if source_id < 0 || source_id >= num || graph.Vertexes[source_id] == nil {
		return errors.New("breadth_first_search error: source_id muse belongs [0,N) and graph->vertexes[source_id] must not be nullptr!")
	}
	v_queue := NewQueue()
	unlimit := Unlimit()
	//************* 初始化顶点 ****************
	for _, ver := range graph.Vertexes {
		if ver == nil {
			continue
		}
		v := a.toBFSVertex(ver)
		v.Color = COLOR_WHITE
		v.Deep = unlimit
		v.SetParent(nil)
	}
	//************* 处理源顶点 ****************
	srcVtx := a.toBFSVertex(graph.Vertexes[source_id])
	srcVtx.SetSource()
	v_queue.Push(srcVtx)
	if pre_action != nil {
		pre_action(source_id)
	}

	//************ 处理其他顶点 ***************
	for !v_queue.Empty() {
		frontItem := v_queue.Front()
		v_queue.Pop()

		var front *BFSVertex = nil
		if v, ret := frontItem.(*BFSVertex); ret {
			front = v
		}
		tuples, _ := graph.VertexEdgeTuples(front.GetID())
		for _, edge := range tuples {
			next_id := edge.Second
			next_vertex := a.toBFSVertex(graph.Vertexes[next_id])
			if next_vertex.Color == COLOR_WHITE {
				next_vertex.SetFound(front) //Deep + 1
				v_queue.Push(next_vertex)
				if pre_action != nil {
					pre_action(next_id)
				}
			}
		}
		front.Color = COLOR_BLACK
		if post_action != nil {
			post_action(front.GetID())
		}
	}

	return nil
}
