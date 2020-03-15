/*
 * @Description: 第25章25.1节 所有节点对的最短路径 之 矩阵乘法
 *
 * ### 算法原理
 *
 * matrix_shortest_path采用动态规划算法求解。考虑从结点i到j的一条最短路径p，假定p最多包含m条边，假定没有权重为负值的环路，且m为有限值。
 * 如果 i=j，则p的权重为0且不包含任何边；如果i和j不同，则将路径p分解为 i-->k(经过路径p')-->j，其中路径p'最多包含m-1条边。
 *
 * 定义 l_i_j<m>为从结点i到j的最多包含m条边的任意路径中的最小权重，则有：l_i_j<m>=
 *
 * - 0：如果i=j
 * - 正无穷:如果 i!=j
 *
 * 对于m>=1，我们有： l_i_j<m>=min(l_i_j<m-1>，min_(1<=k<=n){l_i_k<m-1>+w_k_j})=min_(1<=k<=n){l_i_k<m-1>+w_k_j}。
 * 如果图G不包含负值的环路，则对于每一对结点i,j，如果他们delt(i,j)<正无穷，则从i到j之间存在一条最短路径。由于该路径是简单路径，
 * 则包含的边最多为n-1条。因此delt(i,j)=l_i_j<n-1>=l_i_j<n>=...
 *
 * matrix_shortest_path算法根据输入矩阵W=(w_i_j)，计算出矩阵序列 L<1>，L<2>,...L<n-1>。最后的矩阵L<n-1>包含的是最短路径的权重。
 * 其中L<1>=W。算法的核心是extend_path函数，它将最近计算出的路径扩展了一条边
 *
 *
 *
 * @Author: wangchengdg@gmail.com
 * @Date: 2020-02-28 22:41:31
 * @LastEditTime: 2020-03-14 23:35:20
 * @LastEditors:
 */
package AllNodePairShortestPath

import (
	"errors"

	. "github.com/meshcross/algorithm-3rd/mesh/common"
	. "github.com/meshcross/algorithm-3rd/mesh/graph_algorithm/graph_struct"
)

type MatrixSP struct {
	_N int
}

func NewMatrixSP() *MatrixSP {
	return &MatrixSP{}
}

/**
* @description:扩展一条边
* @param L:初始L矩阵
* @param W: 图的权重矩阵
* @return: 扩展之后的L矩阵
*
* 算法步骤如下：
*
* - 外层循环 i 从 0...N-1(N次)
*   - 内层循环 j 从 0...N-1(N次)
*       - 将newL[i][j]设为正无穷，对所有的k,k 从 0...N-1(N次)，选取 L[i][k]+W[k][j]的最小值赋值给newL[i][j]
* - 最终返回 newL
*
* 性能：时间复杂度 O(n^3)
*
* 这里要充分理解扩展一条边的意思，对于一个图而言，如果边数最多为1时候计算每个点之间的最短路径值，
* 则所有edge(u,v)，只有u->v的值，跟图的权重图是一样的
*
 */
func (a *MatrixSP) extend_path(L [][]int, W [][]int) ([][]int, error) {

	newL := NewMatrix(a._N, 0)
	row_num := a._N
	col_num := a._N
	unlimit := Unlimit()
	for i := 0; i < row_num; i++ {
		for j := 0; j < col_num; j++ {
			newL[i][j] = unlimit //先置为无穷大，表明不通，一旦k循环时候能找到合适的值，则表明是通的，直接更新掉
			for k := 0; k < row_num; k++ {
				//比如当前L表示至多3条边(记为m-1条)的时候的最短路径(i->j的每条边的weight值求和)矩阵，则newL表示至多4条边(记为m条)时候的最短路径矩阵
				//下式的表意为：如果[i,j]中间能找到一个分隔点k,让i到j的距离经过k之后更小，则更新[i,j]最短路径值
				//注意:k遍历时候L[i,k]可能不通(unlimit)，W[k,j]也可能不通(unlimit)
				//L[i][k]+W[k][j]意思为，3条边时候的最短路径L[i][k]+当前这条边E(k,j)共4条边构成的路径的权重是多少
				newL[i][j] = MinInt(newL[i][j], L[i][k]+W[k][j])
			}
		}
	}
	return newL, nil
}

/**
* @description:返回所有节点对的最短路径的矩阵乘法算法
* @param graph:图
* @return: 一个n*n的矩阵(d_i_j)，其中 d_i_j 代表的是结点i到j的一条最短路径的权重
*
* ### 算法步骤
*
* - 初始化：从图中获取权重矩阵 W
* - 执行循环扩展L，其中 L<0>=W, L<k>=extend_path(L<k-1>,W)
* - 最终返回 L<N-1>
*
* ### 算法性能
*
* 时间复杂度O(V^4)
 */
func (a *MatrixSP) ShortestPath(graph *Graph) ([][]int, error) {

	if graph == nil {
		return nil, errors.New("matrix_shortest_path error: graph must not be nil!")
	}

	num := graph.N()
	a._N = num
	unlimit := Unlimit()
	W := NewMatrix(num, 0)
	//**************  从图中创建权重矩阵  ***************
	//****  这里不能直接从图的矩阵描述中提取，因为这里要求 w(i,i)=0，而图的矩阵描述中，结点可能有指向自己的边
	for i := 0; i < num; i++ {
		for j := 0; j < num; j++ {
			if i == j {
				W[i][j] = 0
			} else {
				has_edge, _ := graph.HasEdge(i, j)
				if !has_edge {
					W[i][j] = unlimit
				} else {
					wt, _ := graph.Weight(i, j)
					W[i][j] = wt
				}
			}
		}
	}
	//*********  计算 L <n-1> ***********
	L := W                       //初始状态和W相同，即为1条边的时候，每个节点都只和邻接点有数据
	for i := 0; i < num-2; i++ { //扩展 N-2次
		//每次扩展一条边，比如当前L表示只有1条边的时候，然后进行一次扩展，表示有2条边的时候
		//见书中的例子
		//这里是一条一条边的逐步增加，W始终代表的是一条边时候的最短路径
		L, _ = a.extend_path(L, W)
	}
	return L, nil
}

/**
* @description:矩阵乘法复平方算法
* @param graph:指定的有向图
* @return: 一个n*n的矩阵(d_i_j)，其中 d_i_j 代表的是结点i到j的一条最短路径的权重
*
* ### 算法步骤
*
* - 初始化：从图中获取权重矩阵 W
* - 执行循环扩展L，其中 L<0>=W, L<2*k>=extend_path(L<k>,L<k>)
* - 最终返回 L<log(N-1)的上界整数>
*
* ### 算法性能
* 时间复杂度O(V^3lgV)
 */
func (a *MatrixSP) ShortestPathFast(graph *Graph) ([][]int, error) {

	if graph == nil {
		return nil, errors.New("ShortestPathFast error: graph must not be nil!")
	}

	num := graph.N()
	a._N = num
	unlimit := Unlimit()
	W := NewMatrix(num, 0)
	//**************  从图中创建权重矩阵  ***************
	//****  这里不能直接从图的矩阵描述中提取，因为这里要求 w(i,i)=0，而图的矩阵描述中，结点可能有指向自己的边
	for i := 0; i < num; i++ {
		for j := 0; j < num; j++ {
			if i == j {
				W[i][j] = 0
			} else {
				has_edge, _ := graph.HasEdge(i, j)
				if !has_edge {
					W[i][j] = unlimit
				} else {
					wt, _ := graph.Weight(i, j)
					W[i][j] = wt
				}
			}
		}
	}
	//*********  计算 L <n-1> ***********
	L := W
	m := 1
	//如果图中有num个节点，则i->j的路径如果存在，则最多有n-1条边,哪怕图中有超过n条边，当m>n-1之后的计算结果是没有意义的
	//L的边数扩张方式为1,2,4,8,16...
	for m < num-1 {
		//此处每次是x2的扩张方式，指数级的上升
		//如L表示4条边时候的最短路径，则对于两个相邻的点E(u,v)，不管L是最多几条边的最短路径，u->v的最短路径都是固定不变的，为weight(u,v)
		//此时extend_path(L, L)表示4条边再扩展4条边，最多8条边时候的最短路径
		L, _ = a.extend_path(L, L)
		m = m << 1 //mx2
	}
	return L, nil
}
