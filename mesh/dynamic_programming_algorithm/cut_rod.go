/*
 * @Description: 第15章 15.1，刚调切割问题
 * @Author: wangchengdg@gmail.com
 * @Date: 2020-02-29 18:26:09
 * @LastEditTime: 2020-03-06 13:02:19
 * @LastEditors:
 */
package DynamicProgrammingAlgorithm

import (
	. "github.com/meshcross/algorithm-3rd/mesh/common"
)

//钢条切割问题
type CutDP struct {
}

func NewCutDP() *CutDP {
	return &CutDP{}
}

/**
 * @description: 15.1中的钢条切割有n=len(p)的限制，并且长度是从1-n连续的
 * @param p：价格数组 ,钢条长度不同对应的价格不一样，长度即为索引+1,比如p=[]int{1,2,4}表示长度为1的价格为1，长度为2的价格为2，长度为3的价格为4
          n：为输入的整块钢条的长度
 * @return: 最优切割获得的价格总和
**/
func (a *CutDP) BottomUpCut(p []int, n int) int {
	len_p := len(p)
	if n > len_p {
		return -1
	}

	r := make([]int, n+1)
	r[0] = 0
	unlimit := Unlimit()
	for j := 1; j <= n; j++ {
		q := -unlimit
		for i := 1; i <= j; i++ {
			q = MaxInt(q, p[i-1]+r[j-i])
		}
		r[j] = q
	}
	return r[n]
}
