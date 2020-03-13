/*
 * @Description: 第15章 15.2 矩阵链乘法
		已知n个矩阵连环相乘，通过加括号的方式控制这些矩阵相乘的先后顺序，从而影响实际进行数值运算的次数，并且不改变矩阵相乘的结果
* 		该问题需要找到正确的打括号方式，以求得最少的数值运算次数
 * @Author: wangchengdg@gmail.com
 * @Date: 2020-03-01 22:00:22
 * @LastEditTime: 2020-03-13 15:29:40
 * @LastEditors:
*/
package DynamicProgrammingAlgorithm

import (
	"fmt"

	. "github.com/meshcross/algorithm-3rd/mesh/common"
)

type MatrixChainOrder struct {
}

func NewMatrixChainOrder() *MatrixChainOrder {
	return &MatrixChainOrder{}
}

/**
 * @description: 计算N个矩阵相乘的最优顺序，即按照某个顺序进行，实际的计算次数最小
 * @param p为每个矩阵的长宽，n个矩阵需要传入n+1个参数
 * @return: m记录了每种方案的运算次数，s记录了每个k值的位置
 *
 * l == 1时候，即只有一个矩阵，自己即为相乘的结果，所以计算量m=0
 * 1 <= i <= j <= n
 * i <= k < j
 */
func (a *MatrixChainOrder) Run(p []int) ([][]int, [][]int, error) {

	//矩阵的数量
	n := len(p) - 1
	nx := n + 1
	//记录每种方式的计算次数
	m := make([][]int, nx)
	for k, _ := range m {
		m[k] = make([]int, nx)
	}
	//切割位置k
	s := make([][]int, nx)
	for k, _ := range s {
		s[k] = make([]int, nx)
	}

	unlimit := Unlimit()
	for l := 2; l <= n; l++ { //l为每个子问题规模
		//循环计算每个长度l的最小代价
		//l确定的时候，[i,j]相当于一个滑动窗口，从左向右滑动，i==n-l时候到最右边
		for i := 1; i <= n-l+1; i++ {
			j := i + l - 1
			m[i][j] = unlimit

			//k在[i,j)之间移动，在i,j确定的时候，找到最合理的k值
			for k := i; k < j; k++ {

				//关键的比较函数，按照当前k的取值，计算出前后的推导关系
				//一个 axb矩阵和一个bxc矩阵相乘，新的axc矩阵中每个元素需要计算(b+1)次，b次乘积运算，(b-1)次求和运算
				//所以实际的公式为 q := m[i][k] + m[k+1][j] + p[i-1]*(2*p[k]-1)*p[j]，但是不影响打括号的位置
				//[i,j]被k切割开，分为[i,k]和[k+1,j]，分别计算这两个矩阵的开销，再加上这两个矩阵相乘的开销即为下式的q
				q := m[i][k] + m[k+1][j] + p[i-1]*p[k]*p[j]
				if q < m[i][j] {
					m[i][j] = q //i,j确定时候，最小的计算量是多少
					s[i][j] = k //i,j确定时候，在哪里进行切割，记录k的位置
				}
			}
		}
	}

	// a.Print(s, 1, n)
	// fmt.Println()
	// fmt.Println(m)
	// fmt.Println()
	// fmt.Println(s)
	return m, s, nil
}

func (a *MatrixChainOrder) Print(s [][]int, i, j int) {
	if i == j {
		fmt.Printf("A%d", i)
	} else {
		fmt.Printf("(")
		a.Print(s, i, s[i][j])
		a.Print(s, s[i][j]+1, j)
		fmt.Printf(")")
	}
}
