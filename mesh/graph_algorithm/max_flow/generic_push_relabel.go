package MaxFlow

import (
	"errors"

	. "github.com/meshcross/algorithm-3rd/mesh/graph_algorithm/graph_struct"
	. "github.com/meshcross/algorithm-3rd/mesh/graph_algorithm/graph_struct/graph_vertex"

	. "github.com/meshcross/algorithm-3rd/mesh/common"
)

/**
* 该算法有一个缺陷，在给定的测试案例中(src_id=0,dst_id=5)，dst_id结点流量已满，部分流量需要从结点4回退的时候，
* 会出现结点2和结点4之间反复(push和relabel)的情况，直到结点2的Height增加到比结点0还大
* 在节点数比较多的时候，这里可能会有一些浪费
* 而且结点4在向结点5发送流量之后，自己剩余的流量(9个)还可以发送到结点3再转到结点5，但是该算法先返回2再到4再到3，主要是为了调整Height
* 当前的执行次序为：1-3，1-0，2-4，3-5，4-5，4-2，2-4，4-3，3-5，4-2，2-4，4-2，2-0
**/
//!generic_push_relabel：最大流的推送-重贴标签算法。算法导论26章26.4节
/*!
*
* \param graph:指定流网络,必须非空
* \param src_id: 流的源点
* \param dst_id: 流的汇点
* \return: 最大流矩阵,error
*
* 最大流问题：给定流网络G、一个源结点s、一个汇点t,找出值最大的一个流
*
* ## generic_push_relabel 算法
*
* ### 算法原理
*
* 目前最大流的最快实现是基于推送-重贴标签算法。推送-重贴标签算法比 Ford-fulkserson的局域性更强。它不是对整个残余网络进行检查，
* 而是一个结点一个结点的查看，每一步只检查当前结点的邻结点。与 Ford-fulkerson 方法不同，
* 推送-重贴标签算法并不在整个执行过程中保持流量守恒性质，而是维持一个预流，该流是一个 V*V --> R的函数f,它满足容量限制单不满足流量守恒性质。
* 进入一个点的流量可以超过流出该结点的流量，我们称结点u的 （进入流量-流出流量）=e(u)为结点u的超额流。
* 对于结点u属于V-{s,t}，如果 e(u)>0，则称结点u溢出。
*
* 考虑一个流网络G=(V,E)，我们将有向边看作管道，连接管道的结点有两个性质：
*
* - 为了容纳额外的流e，每个结点都隐藏有一个外流的管道，该管道通向一个容量无穷大的水库
* - 每个结点都有一个高度h，高度的值随着算法的推进而增加
*
* 高度h满足下面性质：h(s)=|V|,h(t)=0,对于所有的边(u,v)属于E_f（残余网络中的边），则有 h(u)<=h(v)+1
*
* 结点的高度h决定了流的推送方向：我们只从高处向低处push流。我们将源的高度固定在|V|，将汇点的高度固定在0。其他结点的高度初始时为0，
* 但是随着时间的推移而不断增加。
*
* generic_push_relabel算法首先从源结点往下发送尽可能多的流到汇点，发送的流量为源结点所发出的所有管道的流量之和。
* 当流进入一个中间结点时，他们被收集在该结点的水库中。
*
* 算法过程中，可能发现所有离开结点u的未充满的管道高度都比u大或者和u相等。此时为了消除结点u的超额流量，必须增加结点u的高度。
* 这就是“重贴标签”操作，我们将结点u的高度增加到比其最低的邻居结点的高度多1个单位的高度，这里要求结点u到该邻居结点的管道必须未充满。
* 因此在执行重贴标签后，一个结点u至少有一个流出管道。可以通过它推送更多的流。
*
* 最终一旦所有水库为空，则预流不但是“合法”的流，也是一个最大流。
*
* ### 基本操作
*
* #### push 操作
*
* 如果一个结点 u是一个溢出结点，则至少有一条入边，假设为(v,u)，且f(v,u)>0。此时残留网络G_f中的边c_f(u,v)>0。如果此时还有h(u)=h(v)+1，
* 则可以对边(u,v)执行push操作。
*
* 因为结点u有一个正的超额流u.e，且边(u,v)的残余容量为正，因此可以增加结点u到v的流，增加的幅度为min(u.e,c_f(u,v))，这种幅度的增加不会导致
* u.e为负值或者容量 c(u,v) 被突破
*
* push 推送流量操作步骤：
*
* - 计算残余容量 c_f(u,v)
* - 计算流的增加幅度 delt_f=min(u.e,c_f(u,v))
* - 更新流：
*   - 如果 (u,v) 属于 E，则 f(u,v) += delt_f
*   - 如果 (v,u) 属于 E，则 f(v,u) -= delt_f
*
* - 更新溢出流量：
*   - u.e -= delt(u,v)
*   - v.e += delt(u,v)
*
* #### relabel 操作
*
* 如果结点u溢出，并且对于所有边(u,v)属于E_f(残留网络G_f中的边)，有u.h<=v.h，则可以对结点u执行relabel操作。
* 当调用relabel操作时，E_f必须包含至少一条从结点u出发的边。
*
* relabel 重贴标签步骤：
*
* - 计算 min{v.h:(u,v)属于E_f}
* - u.h=1+ min{v.h:(u,v)属于E_f}
*
* ### 算法步骤
*
* - 初始化操作：
*   - 初始化预流 flow: flow(u,v)=c(u,v)如果u=s;否则 flow(u,v)=0
*   - 初始化高度函数 h: h(s)=|V|;h(u)=0, u属于 V-{s}
*   - 初始化超额流量 e： e(u)=c(s,u)，u为与源s相邻的结点; e(u)=0，u为与源s不相邻的结点; e(s)初始化为所有s出发的管道之后的相反数
*
* - 执行循环，直到所有结点u属于V-{s,t}不存在超额流量e为止。循环内执行：
*   - 如果可以执行 push 操作，则执行push操作
*   - 如果不能执行 push 操作，又由于存在超额流量 e>0 的结点，因此必然可以执行 relabel 操作。则执行 relabel 操作
*
*
* ### 算法性能
*
* 算法性能：时间复杂度 O(V^2 E)
*
*
 */
type GenericPushRelabel struct {
}

func NewGenericPushRelabel() *GenericPushRelabel {
	return &GenericPushRelabel{}
}

func (a *GenericPushRelabel) MaxFlow(graph *Graph, src_id, dst_id int) ([][]int, error) {

	if graph == nil {
		return nil, errors.New("generic_push_relabel MaxFlow error: graph must not be nil!")
	}
	num := graph.N()

	if src_id < 0 || src_id >= num || graph.Vertexes[src_id] == nil {
		return nil, errors.New("generic_push_relabel MaxFlow error:source id error")
	}

	if dst_id < 0 || dst_id >= num || graph.Vertexes[dst_id] == nil {
		return nil, errors.New("generic_push_relabel MaxFlow error: destination id error")
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

	a.initialize_preflow(graph, src_id, flow)

	for {
		u_id := 0
		//*********** 寻找 溢出结点 ***********
		has_overflow := false

		for _, vtx := range graph.Vertexes {
			vtx_id := vtx.GetID()
			if vtx_id == src_id || dst_id == vtx_id {
				continue
			}
			//有溢出结点
			if vtx.GetKey() > 0 {
				has_overflow = true
				u_id = vtx_id
				break
			}
		}

		//如果有溢出结点
		if has_overflow {
			v_id, _ := a.min_v_at_Ef(graph, u_id, flow)
			uvtx := ToIFlowVertex(graph.Vertexes[u_id])
			vvtx := ToIFlowVertex(graph.Vertexes[v_id])

			//当前节点的Height比临近节点高1，则执行push操作
			if uvtx.GetHeight() > vvtx.GetHeight() {
				a.push(graph, u_id, v_id, flow)
			} else {
				//调整节点u的Height，则马上满足push条件，于是可以执行push
				//push会挑选min_v_at_Ef选出的结点，relabel只修改了u_id结点的Height，所以可以立即调用a.push(graph, u_id, v_id, flow)，因为就算再次循环一次v_id也不会变
				a.relabel(graph, u_id, flow)
				a.push(graph, u_id, v_id, flow)
			}
		} else {
			break
		}
	}
	return flow, nil
}

//!initialize_preflow：generic_push_relabel算法的初始化操作。算法导论26章26.4节
/*!
*
* \param graph:指定流网络。它必须非空，否则抛出异常
* \param src: 流的源点，必须有效否则抛出异常
* \param flow: 预流的引用（执行过程中会更新预流）
* \return: void
*
* 初始化操作执行下列操作：
*
*   - 初始化预流 flow: flow(u,v)=c(u,v)如果u=s;否则 flow(u,v)=0
*   - 初始化高度函数 h: h(s)=|V|;h(u)=0, u属于 V-{s}
*   - 初始化超额流量 e： e(u)=c(s,u)，u为与源s相邻的结点; e(u)=0，u为与源s不相邻的结点; e(s)初始化为所有s出发的管道之后的相反数
* > 由顶点的`key`属性存储超额流量e
 */
func (a *GenericPushRelabel) initialize_preflow(graph *Graph, src_id int, flow [][]int) error {
	if graph == nil {
		return errors.New("initialize_preflow error: graph must not be nil!")
	}
	num := graph.N()
	if src_id < 0 || src_id >= num {
		return errors.New("initialize_preflow error:id must >=0 and <N.")
	}

	if graph.Vertexes[src_id] == nil {
		return errors.New("initialize_preflow error: vertex id does not exist.")
	}

	//*********** 所有结点的 e为0, h为0 ***********
	for _, vtx := range graph.Vertexes {
		v := ToIFlowVertex(vtx)
		v.SetHeight(0)
		//Key当v.e使用
		v.SetKey(0)
	}
	//************* 所有预流为0  **************
	for i := 0; i < num; i++ {
		for j := 0; j < num; j++ {
			flow[i][j] = 0
		}
	}
	svtx := ToIFlowVertex(graph.Vertexes[src_id])
	svtx.SetHeight(num) // h(s)= |V|，该值要设定好，否则可能for迭代次数很多
	//**************  对s出发的边调整  *************
	edges, _ := graph.VertexEdgeTuples(src_id)
	for _, edge := range edges {
		v_id := edge.Second                //{v:(s,v)属于E}
		c_s_v := edge.Third                // c(s,v)
		flow[src_id][v_id] = c_s_v         //f(s,v)
		graph.Vertexes[v_id].SetKey(c_s_v) // v.e=c(s,v)
		svtx.SetKey(svtx.GetKey() - c_s_v)
	}

	return nil
}

//!push：generic_push_relabel算法的push操作。算法导论26章26.4节
/*!
*
* \param graph:指定流网络。它必须非空，否则抛出异常
* \param u_id: 结点u的id，必须有效否则抛出异常
* \param v_id: 结点v的id，必须有效否则抛出异常
* \param flow: 预流的引用（执行过程中会更新预流）
* \return: void
*
* push操作步骤：
*
* - 计算残余容量 c_f(u,v)
* - 计算流的增加幅度 delt_f=min(u.e,c_f(u,v))
* - 更新流：
*   - 如果 (u,v) 属于 E，则 f(u,v) += delt_f
*   - 如果 (v,u) 属于 E，则 f(v,u) -= delt_f
*
* - 更新溢出流量：
*   - u.e -= delt(u,v)
*   - v.e += delt(u,v)
*
* > - 由顶点的`key`属性存储超额流量e        *
* > - 执行push(u,v)时，要求存在残留边 (u,v)属于Ef，且c_f(u,v)>0；否则抛出异常
* > - 执行push(u,v)时，要求 u.e>0；否则抛出异常
*
 */
func (a *GenericPushRelabel) push(graph *Graph, u_id, v_id int, flow [][]int) error {

	if graph == nil {
		return errors.New("push error: graph must not be nil!")
	}

	num := graph.N()
	if u_id < 0 || u_id >= num || v_id < 0 || v_id >= num {
		return errors.New("push error:id must >=0 and <N.")
	}

	if graph.Vertexes[u_id] == nil || graph.Vertexes[v_id] == nil {
		return errors.New("push error: vertex id does not exist.")
	}

	c_f := 0
	delt_f := 0

	//u.e必须有超额流量
	if graph.Vertexes[u_id].GetKey() <= 0 {
		return errors.New("push error:u.e must >0 !")
	}

	//***********  获取 残余流量c_f(u,v) ***********
	uv, _ := graph.HasEdge(u_id, v_id)
	vu, _ := graph.HasEdge(v_id, u_id)
	if uv { //(u,v)属于E
		wt, _ := graph.Weight(u_id, v_id)
		c_f = wt - flow[u_id][v_id] // 残余流量 c_f(u,v)=c(u,v)-f(u,v)
	} else if vu { //(v,u)属于E
		c_f = flow[v_id][u_id] //c_f(u,v)=f(v,u)
	} else {
		return errors.New("push error: must be (u,v) in E or (v,u) in E!")
	}

	//残余网络也必须有流量
	if c_f <= 0 {
		return errors.New("push error:must be c_f > 0!")
	}

	//************ 获取 delt_f(u,v) *********

	//在u上有超额流u.e>0，而且c_f(u,v)残余流量>0，则表示可以把u上的流量发delt_f到v，而不会导致局部系统破坏
	delt_f = MinInt(graph.Vertexes[u_id].GetKey(), c_f)

	//************ 更新 flow *************
	if uv { //(u,v)属于E
		flow[u_id][v_id] += delt_f
	} else if vu { //(v,u)属于E
		flow[v_id][u_id] -= delt_f
	} else {
		return errors.New("push error: need (u,v) in E or (v,u) in E!")
	}

	//************ 更新 更新u,v结点的e ***************
	//u上的u.e切割一部分发往v之后，u.e会减少，v.e会增加
	old_ukey := graph.Vertexes[u_id].GetKey()
	old_vkey := graph.Vertexes[v_id].GetKey()
	graph.Vertexes[u_id].SetKey(old_ukey - delt_f)
	graph.Vertexes[v_id].SetKey(old_vkey + delt_f)

	return nil
}

//!min_v_at_Ef：relabel操作中的min_v_at_Ef操作。算法导论26章26.4节
/*!
*
* \param graph:指定流网络
* \param u_id: 结点u的id
* \param flow: 预流
* \return: 所有边(u,v)属于E_f(残留网络G_f中的边)中，高度最小的结点v
*
* 该方法扫描Ef中所有从u出发的边(u,v)，找出高度最小的结点v
*
 */
func (a *GenericPushRelabel) min_v_at_Ef(graph *Graph, u_id int, flow [][]int) (int, error) {

	if graph == nil {
		return -1, errors.New("min_v_at_Ef error: graph must not be nil!")
	}
	num := graph.N()
	if u_id < 0 || u_id >= num {
		return -1, errors.New("min_v_at_Ef error:id must >=0 and <N.")
	}
	if graph.Vertexes[u_id] == nil {
		return -1, errors.New("min_v_at_Ef error: vertex id does not exist.")
	}

	Ef_v := []int{}                                // {v:(u,v)属于Ef}
	matrix_graph := graph.Matrix.Matrix            //图的矩阵描述
	invalid_weight := graph.Matrix.InvalidWeight() //无效权重
	//*************  获取所有边(u,v)属于E_f(残留网络G_f中的边)的结点 v *************
	for i := 0; i < num; i++ {
		c_u_v := matrix_graph[u_id][i]
		if c_u_v != invalid_weight { //(u,v)属于E
			if flow[u_id][i] < c_u_v { // c(u,v)-f(u,v)>0  有残余流量
				Ef_v = append(Ef_v, i)
			}
		}

		c_v_u := matrix_graph[i][u_id]
		if c_v_u != invalid_weight { //(v,u)属于E
			if flow[i][u_id] > 0 { // f(v,u)>0  使用过的指向该节点的边，正反两个方向判断条件不一样，这里用于获取退还流量的结点
				Ef_v = append(Ef_v, i)
			}
		}
	}
	min_h := Unlimit()
	v_id := -1

	//找到Height最小的临近结点
	for _, id := range Ef_v {
		ivtx := ToIFlowVertex(graph.Vertexes[id])
		if ivtx.GetHeight() < min_h {
			min_h = ivtx.GetHeight()
			v_id = id
		}
	}
	return v_id, nil
}

//!relabel：generic_push_relabel算法的relabel操作。算法导论26章26.4节
/*!
*
* \param graph:指定流网络
* \param u_id: 结点u的id
* \param flow: 预流
* \return: error
*
* relabel 步骤：
*
* - 计算 min{v.h:(u,v)属于E_f}
* - u.h=1+ min{v.h:(u,v)属于E_f}
*
* 当对于存在边(u,v)属于E_f(残留网络G_f中的边)，有u.h>v.h时，抛出异常
*
* 重贴标签操作主要是调整当前结点u的Height值，之后流量会从当前结点流向低Height的临近结点(u.h>v.h)
 */
func (a *GenericPushRelabel) relabel(graph *Graph, u_id int, flow [][]int) error {

	if graph == nil {
		return errors.New("relabel error: graph must not be nil!")
	}
	num := graph.N()
	if u_id < 0 || u_id >= num {
		return errors.New("relabel error:id must >=0 and <N.")
	}
	if graph.Vertexes[u_id] == nil {
		return errors.New("relabel error: vertex id does not exist.")
	}

	if graph.Vertexes[u_id].GetKey() <= 0 {
		return errors.New("relabel error:u.e must >0 !")
	}

	min_v_id, _ := a.min_v_at_Ef(graph, u_id, flow)
	if min_v_id < 0 {
		return errors.New("relabel error: there must be edges in E_f start from u !")
	}

	uvtx := ToIFlowVertex(graph.Vertexes[u_id])
	vvtx := ToIFlowVertex(graph.Vertexes[min_v_id])

	if uvtx.GetHeight() > vvtx.GetHeight() {
		return errors.New("relabel error: u.h  > min_v.h !")
	}

	uvtx.SetHeight(vvtx.GetHeight() + 1)
	return nil
}
