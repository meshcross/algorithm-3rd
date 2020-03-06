/*
 * @Description: 15.4 最长公共子序列的子算法 ，如 12345xyz和-1235xz 的最长公共子串为135xz
 * @Author: wangchengdg@gmail.com
 * @Date: 2020-02-10 22:19:06
 * @LastEditTime: 2020-03-06 13:05:17
 * @LastEditors:
 */
package DynamicProgrammingAlgorithm

import (
	"fmt"
	"strings"
)

type LongestCommonSubsequence struct {
}

func NewLongestCommonSubsequence() *LongestCommonSubsequence {
	return &LongestCommonSubsequence{}
}

/**
 * @description: 打印结构,向地位递归的时候千万要处理好边界情况，这里特指i==0 或者 j==0的情况
 * @param :
 * @return:void
 */
func (a *LongestCommonSubsequence) lcs_print(b [][]string, str_x string, i, j int, ret *[]string) {
	// if i == 0 || j == 0 {
	// 	return
	// }
	if b[i][j] == "↖" {
		if i > 0 && j > 0 {
			a.lcs_print(b, str_x, i-1, j-1, ret)
		}
		//fmt.Printf(string(str_x[i]))
		(*ret) = append(*ret, string(str_x[i]))
	} else if b[i][j] == "↑" {
		if i > 0 {
			a.lcs_print(b, str_x, i-1, j, ret)
		}
	} else {
		if j > 0 {
			a.lcs_print(b, str_x, i, j-1, ret)
		}
	}
}

/**
 * @description: 获取str_x和str_y的最大公共子串,如 12345xyz和-1235xz 的lcs为135xz
 * @param str_x :字符串1
 * @param str_y :字符串2
 * @return:
 */
func (a *LongestCommonSubsequence) Run(str_x string, str_y string, do_prints ...bool) string {
	b, c, str := a.lcs_length(str_x, str_y)

	if len(do_prints) > 0 && do_prints[0] {
		fmt.Println("b struct:")
		for _, v := range b {
			fmt.Println(v)
		}
		fmt.Println("c struct:")
		for _, v := range c {
			fmt.Println(v)
		}
		arr := []string{}
		a.lcs_print(b, str_x, len(str_x)-1, len(str_y)-1, &arr)
		// fmt.Println("ret_str:", strings.Join(arr, ""))
	}
	return str
}

/**
 * @description:使用动态规划算法求解
 * @param 传入两个字符串
 * @return:返回公共子串
 * ←↑→↓↖↙↗↘↕
 */
func (a *LongestCommonSubsequence) lcs_length(str_x string, str_y string) ([][]string, [][]int, string) {
	m := len(str_x)
	n := len(str_y)

	b := make([][]string, m)
	c := make([][]int, m)
	for k := 0; k < m; k++ {
		b[k] = make([]string, n)
		c[k] = make([]int, n)
	}

	str_arr := []string{}
	for k := 0; k < n; k++ {
		if str_x[0] == str_y[k] {
			c[0][k] = 1
			b[0][k] = "↖"
			str_arr = append(str_arr, string(str_x[0]))
		} else {
			if k > 0 {
				c[0][k] = c[0][k-1]
			}
			b[0][k] = "→"
		}
	}
	for k := 0; k < m; k++ {
		if str_y[0] == str_x[k] {
			//避免重复处理
			if c[k][0] != 1 {
				c[k][0] = 1
				b[k][0] = "↖"
				str_arr = append(str_arr, string(str_y[0]))
			}
		} else {
			if k > 0 {
				c[k][0] = c[k-1][0]
			}
			b[k][0] = "↑"
		}
	}

	//i==0 j==0的边界情况需要处理

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {

			if str_x[i] == str_y[j] {
				c[i][j] = c[i-1][j-1] + 1
				b[i][j] = "↖"

				str_arr = append(str_arr, string(str_x[i]))
			} else if c[i-1][j] >= c[i][j-1] {
				c[i][j] = c[i-1][j]
				b[i][j] = "↑"
			} else {
				c[i][j] = c[i][j-1]
				b[i][j] = "→"
			}
		}
	}

	return b, c, strings.Join(str_arr, "")
}
