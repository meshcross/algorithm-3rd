/*
 * @Description: 第25章25.2节 所有节点对的最短路径之 floyd warshall算法
 * @Author: wangchengdg@gmail.com
 * @Date: 2020-02-10 22:18:13
 * @LastEditTime: 2020-03-15 12:40:16
 * @LastEditors:
 *
 *
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
 *	//重点理解该规则，如果k点在p上，则d_i_j<k-1>不通(理解为无穷大),d_i_k<k-1>+d_k_j<k-1>为p
 *	//如果k不在p上，则d_i_j<k-1>为p，d_i_k<k-1>+d_k_j<k-1>绕道经过了非最短路径p上的点，所以和值必然比p大
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
 *
 *
 */
package AllNodePairShortestPath

import (
	"errors"

	. "github.com/meshcross/algorithm-3rd/mesh/common"
	. "github.com/meshcross/algorithm-3rd/mesh/graph_algorithm/graph_struct"
)

type FloydWarshallSP struct {
}

func NewFloydWarshallSP() *FloydWarshallSP {
	return &FloydWarshallSP{}
}

/**
* @description:floyd_warshall算法
* @param graph:有向图
* @return: 一个n*n的权重矩阵(d_i_j)与n*n的前驱矩阵(p_i_j)，其中 d_i_j 代表的是结点i到j的一条最短路径的权重,
*   p_i_j 为从结点i到j的一条最短路径上j的前驱结点
*
*  假设一条最短路径p(i,j)经过的节点为v0,v1,v2,v3,v4,则v0为出发点，v4为目的地结点，v3即为j的前驱节点
*  所以对于如下的前驱矩阵有如下解读：
	这是一个5x5的前驱矩阵，现在要求v2到v0的最短路径经过了哪些节点
	a、首先找到i=2,j=0的前驱节点，即v3
	b、找到i=2,j=3的前驱节点，即v1
	c、找到i=2,j=1的前驱节点,即v2，为i节点本身
	d、最短路径为v2,v1,v3,v0

		[-1 2 3 4 0]
		[3 -1 3 1 0]
		[3 2 -1 1 0]
		[3 2 3 -1 0]
		[3 2 3 4 -1]
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
	//**************  初始化 D 和 P(前驱矩阵) ************
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
		newD := NewMatrix(num, 0)
		newP := NewMatrix(num, 0)
		for i := 0; i < num; i++ {
			for j := 0; j < num; j++ {
				// D中存放的是D<k-1>,P中存放的是P<k-1>
				sum := 0

				//如果k节点跟i或者j不通，则一定不在p(i,j)的最短路径上
				if Is_Unlimit(D[i][k]) || Is_Unlimit(D[k][j]) {
					sum = unlimit
				} else {
					sum = D[i][k] + D[k][j]
				}

				//如果原来的最短路径值更小，则k不在p(i,j)的最短路径上
				if D[i][j] <= sum {

					newD[i][j] = D[i][j]
					newP[i][j] = P[i][j]
				} else { //否则k在p(i,j)的最短路径上，需要记录到前驱矩阵

					newD[i][j] = sum
					//k在p(i,j)的最短路劲上，所以和p(k,j)的前驱是一样的
					newP[i][j] = P[k][j]
				}
			}
		}

		D = newD
		P = newP
	}

	return D, P, nil
}
