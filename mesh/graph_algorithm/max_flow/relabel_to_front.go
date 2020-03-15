/*
 * @Description: 第26章26.5节 最大流的 前置重贴标签算法
 *	在一般情况下，该算法性能不亚于generic_push_relabel(O(v^2.E))算法，对于稠密图，该算法性能更优O(v^3)
 * @Author: wangchengdg@gmail.com
 * @Date: 2020-02-18 10:34:12
 * @LastEditTime: 2020-03-15 18:21:06
 * @LastEditors:
 */
package MaxFlow

import (
	"errors"

	. "github.com/meshcross/algorithm-3rd/mesh/graph_algorithm/graph_struct"
	. "github.com/meshcross/algorithm-3rd/mesh/graph_algorithm/graph_struct/graph_vertex"
)

type RelabelToFront struct {
	GenericPushRelabel
}

func NewRelabelToFront() *RelabelToFront {
	return &RelabelToFront{}
}

/**
* @description:释放操作
* @param graph:流网络
* @param u_id: 图的顶点id
* @param flow: 预流
* @return: void
*
* 对于溢出结点u,如果将其所有多余的流通过许可边推送到相邻的结点上，则称该结点得到释放。
* 在释放过程中，需要对结点u进行重贴标签操作，这使得从结点u发出的边成为许可边。discharge(u) 操作步骤如下：
*
* - 循环，条件为u.e>0，循环内操作为：
*   - 获取u.current，假设为v
*   - 如果v为空，即遍历到u.N链表的末尾，则对u执行relabel操作，然后将u.current指向u.N链表的头部
*   - 如果 v非空，且满足 push 操作的条件(c_f(u,v)>0且 u.h=v.h+1)，则执行push操作
*   - 如果 v 非空，但不满足 push 操作，则 u.current指向u.N链表的下一个结点
*
* 把自己的超额流量推送到相邻结点上
 */
func (a *RelabelToFront) discharge(graph *Graph, u_id int, flow [][]int) error {

	if graph == nil {
		return errors.New("discharge error: graph must not be nil!")
	}
	num := graph.N()
	if u_id < 0 || u_id >= num {
		return errors.New("discharge error:id must >=0 and <N.")
	}

	vertex_u := ToFrontFlowVertex(graph.Vertexes[u_id])
	if vertex_u == nil {
		return errors.New("discharge error: u_id does not exist.")
	}

	//**************  开始循环  *******************
	//key代表残余流量，如果参与流量>0
	for vertex_u.GetKey() > 0 {

		node_v := vertex_u.N_List.Current //
		if node_v == nil {
			a.relabel(graph, u_id, flow)
			vertex_u.N_List.Current = vertex_u.N_List.Head
		} else {
			c_f := 0
			vertex_v := node_v.Value
			//***********  获取 c_f(u,v)  **************
			v_id := vertex_v.GetID()
			uv, _ := graph.HasEdge(u_id, v_id)
			vu, _ := graph.HasEdge(v_id, u_id)
			if uv {
				wt, _ := graph.Weight(u_id, v_id)
				c_f = wt - flow[u_id][v_id]
			} else if vu {
				c_f = flow[v_id][u_id]
			} else {
				return errors.New("discharge error: (u,v) belongs E or (v,u) belongs E must be true!")
			}

			//************  根据 c_f(u,v)以及 h函数决定是否 push  **************
			uvtx := ToFrontFlowVertex(vertex_u)
			vvtx := ToFrontFlowVertex(vertex_v)
			//残余流量>0，且u的高度==v的高度+1，则推送流量，否则向后移动Current指针
			if c_f > 0 && (uvtx.GetHeight() == vvtx.GetHeight()+1) {
				//push之后改变参与流量
				a.push(graph, u_id, v_id, flow)
			} else {
				uvtx.N_List.Current = uvtx.N_List.Current.Next
			}
		}
	}
	return nil
}

/**
* @description:创建L链表操作
* @param graph:流网络
* @param src_id: 流的源点
* @param dst_id: 流的汇点
* @return: 初始化的L链表
*
* 该操作将所有的除s、t之外的顶点加入到L链表中
 */
func (a *RelabelToFront) create_L(graph *Graph, src_id, dst_id int) (*List, error) {

	if graph == nil {
		return nil, errors.New("create_L error: graph must not be nil!")
	}
	num := graph.N()
	if src_id < 0 || src_id >= num || dst_id < 0 || dst_id >= num {
		return nil, errors.New("create_L error:id must >=0 and <N.")
	}
	if graph.Vertexes[src_id] == nil || graph.Vertexes[dst_id] == nil {
		return nil, errors.New("create_L error: vertex id does not exist.")
	}

	L := &List{}
	for i := 0; i < num; i++ {
		if i == src_id || i == dst_id {
			continue
		}
		ivtx := ToFrontFlowVertex(graph.Vertexes[i])
		node := &ListNode{Value: ivtx}
		L.Add(node)
	}
	return L, nil
}

/**
* @description:初始化邻接链表
* @param graph:流网络
* @param src_id: 流的源点
* @param dst_id: 流的汇点
* @return: void
*
* 该操作将初始化除了s、t之外所有顶点的邻接链表
*
 */
func (a *RelabelToFront) initial_vertex_NList(graph *Graph, src_id, dst_id int) error {

	if graph == nil {
		return errors.New("initial_vertex_NList error: graph must not be nil!")
	}
	num := graph.N()
	if src_id < 0 || src_id >= num || dst_id < 0 || dst_id >= num {
		return errors.New("initial_vertex_NList error:id must >=0 and <N.")
	}
	if graph.Vertexes[src_id] == nil || graph.Vertexes[dst_id] == nil {
		return errors.New("initial_vertex_NList error: vertex id does not exist.")
	}

	for i := 0; i < num; i++ {
		if i == src_id || i == dst_id {
			continue
		}
		vertex_u := ToFrontFlowVertex(graph.Vertexes[i])
		//************ 扫描邻接矩阵  **************
		matrix := graph.Matrix.Matrix
		invalid_weight := graph.Matrix.InvalidWeight()

		for j := 0; j < num; j++ {
			if matrix[i][j] != invalid_weight { //从u出发的边
				vvtx := ToFrontFlowVertex(graph.Vertexes[j])
				node := &ListNode{Value: vvtx}
				vertex_u.N_List.Add(node) //每个节点的邻接矩阵
			}
			if matrix[j][i] != invalid_weight { //进入u的边
				vvtx := ToFrontFlowVertex(graph.Vertexes[j])
				node := &ListNode{Value: vvtx}
				vertex_u.N_List.Add(node)
			}
		}
		vertex_u.N_List.Current = vertex_u.N_List.Head //将 u.N.current设为u.N.head
	}
	return nil
}

/**
* @description:前置重贴标签算法
* @param graph:流网络,必须非空
* @param src_id: 流的源点，必须有效否则抛出异常
* @param dst_id: 流的汇点，必须有效否则抛出异常
* @return: 最大流矩阵

* 最大流问题：给定流网络G、一个源结点s、一个汇点t,找出值最大的一个流
*
* ### 算法步骤
*
* - 初始化预流操作（与 generic_push_relabel 算法相同）
* - 对所有的非s、t的结点，将它们加入到链表L中（任意顺序）
* - 对所有的非s、t的结点u,初始化u.current为u.N.head
* - 设置u为L.head
* - 循环，条件为u!=NIL，循环中操作：
*   - 保留u.h为oldh
*   - 对u执行discharge操作
*   - 如果u.h>oldh，证明对u执行了重贴标签操作，此时将u移动到L的头部
*   - u=u.next（提取u在L中的下一个）
*
* ### 算法性能
*
* 算法性能：时间复杂度 O(V^3)
*
 */
func (a *RelabelToFront) MaxFlow(graph *Graph, src_id, dst_id int) ([][]int, error) {

	if graph == nil {
		return nil, errors.New("relabel_to_front error: graph must not be nil!")
	}
	num := graph.N()
	if src_id < 0 || src_id >= num || dst_id < 0 || dst_id >= num {
		return nil, errors.New("relabel_to_front error:id must >=0 and <N.")
	}
	if graph.Vertexes[src_id] == nil || graph.Vertexes[dst_id] == nil {
		return nil, errors.New("relabel_to_front error: vertex id does not exist.")
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
	a.initialize_preflow(graph, src_id, &flow)

	//此处即为对generic_push_relabel的优化，不是每次都进行所有vertex的扫描，
	//而是将需要检查的节点（有残存边）放入链表中，从链表取出，从而优化了性能
	L, _ := a.create_L(graph, src_id, dst_id)     // create List L，将所有节点都加入链表L，除了src和dst
	a.initial_vertex_NList(graph, src_id, dst_id) //create u.N for each u，计算每个节点的邻接链表，除了src和dst

	//************   循环 **************
	node_u := L.Head
	for node_u != nil {

		vertex_u := ToFrontFlowVertex(node_u.Value)
		old_height := vertex_u.GetHeight()         //保存旧h值
		a.discharge(graph, vertex_u.GetID(), flow) //释放u，把自己的超额流量推送到相邻结点上

		//L链表增加节点
		if vertex_u.GetHeight() > old_height { //若重贴标签则h值增加，则u前置到L头部
			if node_u != L.Head { //当u已经是L头时无需操作
				frontNode, _ := L.FrontOf(node_u)
				frontNode.Next = node_u.Next
				node_u.Next = L.Head
				L.Head = node_u
			}
		}
		node_u = node_u.Next
	}
	return flow, nil
}
