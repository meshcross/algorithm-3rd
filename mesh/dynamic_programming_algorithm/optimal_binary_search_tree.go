/*
 * @Description: 第15章 15.5最优二叉搜索树
				希望查找的时候，让概率更大的节点离根更近
 *              最优搜索树有如下性质：左边子树权重一定比根小，右边子树的权重比根大
 *				该算法并没有改变数组中的节点顺序，只是从现有顺序的节点中挑选合适的节点作为根节点，挑选谁作为根节点即为每个子问题的核心问题
 *				所以不同的节点顺序输入会有差异，造成不同的结果，比如{3, 2, 1, 2, 4}和{1, 2, 4, 2, 3}分别作为输入参数就会得到不一样的结果
 *
 * 如何从最终的结果矩阵root中得到最终完整的最优二叉搜索树？
 * 先找到root[1,n]，即为最终二叉树的总根节点r0，于是问题变为找到左子树[1,r-1]和右子树[r+1,n]的根节点，root矩阵图中取出即可，依次下推即可得到完整的二叉搜索树
 *
 *  [0 0 0 0 0 0 0]
	[0 1 2 3 3 3 0]
	[0 0 2 3 3 3 0]
	[0 0 0 3 3 3 0]
	[0 0 0 0 4 5 0]
	[0 0 0 0 0 5 0]
	[0 0 0 0 0 0 0]
	从以上矩阵root中得到二叉搜索树的最终形态,其中行为i，列为j,访问方式为root[i,j]：
	a、根节点为root[1,5],即node3，则左子树为[1,2],右子树为[4,5]
	b、[1,2]的根节点为root[1,2]，即node2，[4,5]的根节点为root[4,5]，即node5，所以node3的左节点为node2，右节点为node5
	c、[1,2]被根节点node2划分为两个部分,左边为[1,1]，右边没有，所以node2的左节点为node1
	d、[4,5]被根节点node5划分为两个部分,左边为[4,4]，右边没有，所以node5的左节点为node4
	e、最终形态如下图：
		   node3
		  /    \
		node2   node5
		/         /
		node1    node4

 * @Author: wangchengdg@gmail.com
 * @Date: 2020-03-02 15:21:12
 * @LastEditTime: 2020-03-03 12:10:07
 * @LastEditors:
*/

package DynamicProgrammingAlgorithm

import (
	. "github.com/meshcross/algorithm-3rd/mesh/common"
)

type OptimalBinarySearchTree struct {
}

func NewOptimalBinarySearchTree() *OptimalBinarySearchTree {
	return &OptimalBinarySearchTree{}
}

/**
 * @description:书中引入了d[i]伪结点，复杂了
 * @param p k[i]的概率列表，k为实际输入数节点，p0,p1,p2,...,pn,p0通常是无效的数字
 * @param q k[i]的概率列表，k为实际输入数节点，q0,q1,q2,...,qn,q的有效数量p多一个，q0也是有效的
 * @param n 有效节点个数，比如demo中n=5，特别注意，p的第一个元素是没有意义的，通常来说 n==len(p)-1
 * @return:
 * 最终算法的目的是要从[1,1]循环推理到[1,n],得到概率矩阵/根节点矩阵
 * 所以expect[1][n]和root[1][n]为所求
 *
 * 如果把q中的所有元素都置为0，则RunAsBook和Run会得到相同的计算结果
 *
 */
func (a *OptimalBinarySearchTree) RunAsBook(p, q []int, n int) ([][]int, [][]int, [][]int) {
	nx := n + 1
	expect := make([][]int, nx+1) //期望比较次数，或平均比较次数
	weight := make([][]int, nx+1)
	root := make([][]int, nx+1)

	for k := 0; k < nx+1; k++ {
		expect[k] = make([]int, nx)
		weight[k] = make([]int, nx)
		root[k] = make([]int, nx)
	}

	for i := 1; i < nx+1; i++ {
		expect[i][i-1] = q[i-1]
		weight[i][i-1] = q[i-1]
	}
	for i := 1; i < nx; i++ {
		expect[i][i] = p[i]
		weight[i][i] = p[i]
	}

	//需要处理0开始的边界情况，动态规划算法涉及到由上一个推导下一个的过程，第一个(或最后一个)必然没有上一个元素可供提供推导，所以必须做边界处理
	unlimit := Unlimit()
	//需要处理边界情况，所以允许l从0开始，并没有完全按照书中所写的从1开始
	for l := 0; l <= n; l++ {
		for i := 1; i < n-l+1; i++ {
			j := i + l
			expect[i][j] = unlimit
			weight[i][j] = weight[i][j-1] + p[j] + q[j]

			for r := i; r <= j; r++ {
				t := expect[i][r-1] + expect[r+1][j] + weight[i][j]
				if t < expect[i][j] {
					expect[i][j] = t
					root[i][j] = r
				}
			}
		}
	}

	return expect, weight, root
}

/**
 * @description:简化的最优二叉树查找，没有伪结点的概念
 * @param p 每个节点的概率，数组
 * @return:最少的查找次数;期望矩阵;根节点矩阵
 *
 * 目的只有一个，通过[i,j]滑动进行推理，将范围拉开到[1,n]得到的值即为所求
 *
 */
func (a *OptimalBinarySearchTree) Run(p []int) (int, [][]int, [][]int) {
	n := len(p)

	p = append([]int{0}, p...)

	nx := n + 2
	c := make([][]int, nx)
	root := make([][]int, nx)
	for k := 0; k < nx; k++ {
		c[k] = make([]int, nx)
		root[k] = make([]int, nx)
	}

	for k := 1; k <= n; k++ {
		c[k][k-1] = 0  //j>i则范围非法，置为0
		c[k][k] = p[k] //相当于只有一个元素，所以概率就是该元素的概率
		root[k][k] = k //只有一个元素，所以该元素即为根节点
	}

	unlimit := Unlimit()
	//l即为length,表示窗口宽度
	for l := 1; l < n; l++ {
		for i := 1; i <= n-l; i++ { //行的取值范围
			j := i + l //l为i到j的距离，所以i和l固定之后，很容易求出j ，求出在对角线往上半部分的i对应的j，j>=i才有意义
			optimal_cnt := unlimit
			optimal_k := i //最优二叉树的k的取值最终在[i,j]之间
			w_i_j := 0
			for k := i; k <= j; k++ {
				//在k的循环内，sum的值最终为p[i]+p[i+1] + ... + p[j] = w[i,j]
				//在已确定的[i,j]范围内，w[i,j]就变成常数了，要求出最优的k值，只需要比较变化的部分，即minnum的值
				//实际使用的公式为 c[i][j] = min(c[i][k-1] + c[k+1][j] + w[i,j]) 其中i<=k<=j
				//核心部分，只找出合适的k作为[i,j]区域的根节点，并不改变节点在数组p中的次序
				w_i_j = w_i_j + p[k]
				if c[i][k-1]+c[k+1][j] < optimal_cnt {
					optimal_cnt = c[i][k-1] + c[k+1][j]
					optimal_k = k
				}
			}
			c[i][j] = optimal_cnt + w_i_j //得到了最小值
			root[i][j] = optimal_k        //记录取得最小值时的根节点
		}
	}

	return c[1][n], c, root
}
