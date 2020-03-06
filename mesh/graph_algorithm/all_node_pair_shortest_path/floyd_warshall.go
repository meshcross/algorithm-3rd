/*
 * @Description: golang algorithm
 * @Author: wangchengdg@gmail.com
 * @Date: 2020-02-10 22:18:13
 * @LastEditTime: 2020-03-05 11:55:16
 * @LastEditors:
 */
package AllNodePairShortestPath

import (
	"errors"

	. "github.com/meshcross/algorithm-3rd/mesh/common"
	. "github.com/meshcross/algorithm-3rd/mesh/graph_algorithm/graph_struct"
	// . "github.com/meshcross/algorithm-3rd/mesh/GraphAlgorithm/graph_struct/graph_vertex"
)

type FloydWarshallSP struct {
}

func NewFloydWarshallSP() *FloydWarshallSP {
	return &FloydWarshallSP{}
}

//!floyd_warshall：返回所有节点对的最短路径的floyd_warshall算法。算法导论25章25.2节
/*!
*
* \param graph:指定的有向图。它必须非空，否则抛出异常
* \return: 一个n*n的矩阵(d_i_j)与n*n的矩阵(p_i_j)，其中 d_i_j 代表的是结点i到j的一条最短路径的权重,
*   p_i_j 为从结点i到j的一条最短路径上j的前驱结点
*
* ## 所有结点对的最短路径
*
* ## floyd_warshall 算法
*
* ### 算法原理
*
* floyd_warshall也是采用一种动态规划算法来解决问题。假定图G中的所有结点为V={1,2,3...n},考虑其中的一个子集{1,2,...k}，这里k是某个小于n的整数。
* 图中可以存在负权重的边，但是不能存在权重为负的环路。对于任意结点对i,j属于V，考虑从结点i到结点j的所有中间结点均取自集合{1,2,...k}的路径，并且假设p为其中权重最小的路径。
*
* - 若结点k不是路径p的中间结点，则路径p上的所有中间结点都属于集合{1,2,3,...k-1}。则结点i到j的中间结点取自集合{1,2,...k-1}
* 的一条最短路径也是从结点i到j的中间结点取自{1,2...k)的一条最短路径
* - 若结点k是路径p上的中间结点，则路径p分解为 i-->k(经过p1)-->j(经过p2)，则路径p1是结点i到k的中间结点取自集合{1,2...k-1}
* 的一条最短路径。路径p2是结点k到j的中间结点取自集合{1,2,...k-1}的一条最短路径。
*
* 根据以上观测，设d_i_j<k>为从结点i到结点j的所有中间结点全部取自集合{1,2...k}的一条最短路径的权重。当k=0时，从结点i到j
* 的一条不包含编号大于0的中间结点的路径将没有任何中间结点。这样的路径最多只有一条边，因此d_i_j<0>=w_i_j。因此d_i_j<k>为：
*
* - w_i_j：当k=0		//表示i，j没有中间节点，所以直接取边的权重即可
* - min(d_i_j<k-1>,d_i_k<k-1>+d_k_j<k-1>：当k>0
	//重点理解该规则，如果k点在p上，则d_i_j<k-1>不通(理解为无穷大),d_i_k<k-1>+d_k_j<k-1>为p
	//如果k不在p上，则d_i_j<k-1>为p，d_i_k<k-1>+d_k_j<k-1>绕道经过了非最短路径p上的点，所以和值必然比p大
*
* 对任何路径来说，所有中间结点都属于集合{1,2,...n}，则矩阵D<n>=(d_i_j<n>)
*
* 我们可以在计算矩阵D<k>的同时计算出前驱矩阵II，即计算一个矩阵序列 II<0>,II<1>...II<k>。这里定义II<k>=(pai_i_j<k>)，
* pai_i_j<k>为从结点i到j的一条所有中间结点都取自集合{1,2,...k}的最短路径上j的前驱结点。
*
* 当k=0时，从i到j的一条最短路径没有中间结点，因此 pai_i_j<0>=:
*
* - nil:若i=j或者w_i_j=正无穷
* - i ：若i!=j且w_i_j!=正无穷
*
* 对于 k>=1，若结点k不是路径p的中间结点，则路径p上的所有中间结点都属于集合{1,2,3,...k-1}，则 pai_i_j<k>=pai_i_j<k-1>;
* 若结点k是路径p的中间结点,考虑路径 i-->k-->j，这里 k!=j，则pai_i_j<k>=pai_k_j<k-1>。因此当k>=1时，pai_i_j<k>=:
*
* - pai_i_j<k-1>：当d_i_j<k-1> <= d_i_k<k-1>+dk_j<k-1>
* - pai_k_j<k-1>: 当 d_i_j<k-1> > d_i_k<k-1>+dk_j<k-1>
*
* ### 算法步骤
*
* - 初始化：从图中获取结果矩阵D，以及父矩阵P
* - 外层循环 k 从 0..N-1(N次)
*   - 新建 D<k>,P<k>，同时执行内层循环 i 从 0..N-1(N次)
*   - 执行最内层循环 j 从 0..N-1(N次)
*       - 根据递推公式对d_i_j<k>和p_i_j<k>赋值(此时D等于D<k-1>，P等于P<k-1>
*   - 将D<k>赋值给D,P<k>赋值给P
* - 返回 D,P
*
* ### 算法性能
*
* 时间复杂度 O(V^3)
*/

func (a *FloydWarshallSP) ShortestPath(graph *Graph) ([][]int, [][]int, error) {

	if graph == nil {
		return nil, nil, errors.New("floyd_warshall error: graph must not be nil!")
	}
	//**************  初始化 D 和 P ************
	//****  这里不能直接从图的矩阵描述中提取，因为这里要求 w(i,i)=0，而图的矩阵描述中，结点可能有指向自己的边
	num := graph.N()
	D := NewMatrix(num, 0)
	P := NewMatrix(num, 0)
	unlimit := Unlimit()
	for i := 0; i < num; i++ {
		for j := 0; j < num; j++ {
			if i == j {
				D[i][j] = 0
				P[i][j] = -1
			} else {
				has_edge, _ := graph.HasEdge(i, j)
				if !has_edge {
					D[i][j] = unlimit
					P[i][j] = -1
				} else {
					wt, _ := graph.Weight(i, j)
					D[i][j] = wt
					P[i][j] = i
				}
			}
		}
	}
	//**************  计算矩阵D和前驱矩阵P ******************
	for k := 0; k < num; k++ {
		newD := NewMatrix(num, 0) //  D<k>
		newP := NewMatrix(num, 0) //  P<k>
		for i := 0; i < num; i++ {
			for j := 0; j < num; j++ {
				// D中存放的是D<k-1>,P中存放的是P<k-1>
				sum := 0
				if D[i][k] == unlimit || D[k][j] == unlimit || D[i][k]+D[k][j] >= unlimit {
					sum = unlimit
				} else {
					sum = D[i][k] + D[k][j]
				}

				if D[i][j] <= sum { // d_i_j<k-1> <= d_i_k<k-1>+d_k_j<k-1>

					newD[i][j] = D[i][j] //则 d_i_j<k> = d_i_j<k-1>
					newP[i][j] = P[i][j] //则 p_i_j<k> = p_i_j<k-1>
				} else { // d_i_j<k-1> > d_i_k<k-1>+d_k_j<k-1>

					newD[i][j] = sum     //则 d_i_j<k> = d_i_k<k-1>+d_k_j<k-1>
					newP[i][j] = P[k][j] //则 p_i_j<k> = p_k_j<k-1>
				}
			}
		}

		D = newD
		P = newP
	}

	return D, P, nil
}
