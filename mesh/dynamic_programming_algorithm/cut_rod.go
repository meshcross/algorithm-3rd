/*
 * @Description: 第15章 15.1，刚调切割问题
 * @Author: wangchengdg@gmail.com
 * @Date: 2020-02-29 18:26:09
 * @LastEditTime: 2020-03-13 15:15:02
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
	if len_p <= 0 || n <= 0 {
		return 0
	}

	//r记录1-n每个长度的最大收益，用于保存每个子问题的解
	r := make([]int, n+1)
	r[0] = 0 //长度为0的钢条收益为0
	unlimit := Unlimit()
	for j := 1; j <= n; j++ { //需要求解每个规模为j的子问题
		q := -unlimit
		for i := 1; i <= j; i++ { //对于已经确定的j，循环求解最优解
			//长度为i的价格为p[i-1]，长度为i，剩余长度为j-i，j-i长度的最优解为r[j-i]
			//如果p[i-1] + r[j-1]的总价更高>q，则需要划分一段长度为i的钢条，否则不需要长度为i的钢条
			q = MaxInt(q, p[i-1]+r[j-i])
		}
		r[j] = q //将j的子问题的解存入r[j]
	}
	return r[n]
}
