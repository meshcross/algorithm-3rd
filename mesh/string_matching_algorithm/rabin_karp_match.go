/*
 * @Description: 32.2 Rabin-Karp算法
 * @Author: wangchengdg@gmail.com
 * @Date: 2020-02-10 22:19:49
 * @LastEditTime: 2020-03-05 17:45:47
 * @LastEditors:
 */
package StringMatchingAlgorithm

import (
	. "github.com/meshcross/algorithm-3rd/mesh/common"
)

type RabinKarpMatch struct{}

func NewRabinKarpMatch() *RabinKarpMatch {
	return &RabinKarpMatch{}
}

/**
 * @description:找出text中所有能匹配到pattern的偏移位置
 * @param text 源字符串，可能有很多字符
 * @param pattern关键词，需要查找在text中是否有pattern
 * @param radix_d进制
 * @param mod_q取余
 * @return:所有匹配位置的偏移
*
* 所有的字符假设都是d进制的数字
* mod_q 为取余的底
*
* 关于根据t[s]求t[s+1]的问题，t[s+1]和t[s]首位不一样，中间是一样的
* t[s]表示T[s]+T[s+1] + ... + T[s+m-1]范围内的t值，
* 即 	t[s] = T[s+m-1] + d*T[s+m-2] + ... + Pow(d,m-1)*T[s]
*    	t[s+1] = T[s+m] + d*T[s+m-1] + ... + Pow(d,m-1)*T[s+1]
			   = T[s+m] + (d*T[s+m-1] + Pow(d,2)T[s+m-2] ... + Pow(d,m-1)*T[s+1])
				=T[s+m] + d * (T[s+m-1] + d* T[s+m-2]+ ... + Pow(d,m-2)*T[s+1])
				=T[s+m] + d * (T[s+m-1] + d* T[s+m-2]+ ... + Pow(d,m-2)*T[s+1] + Pow(d,m-1)*T[s] - Pow(d,m-1)*T[s])
				=T[s+m] + d*(ts - Pow(d,m-1)*T[s])
*
* 于是有以下迭代关系式：t[s+1] = d * (t[s] - Pow(d,m-1)*T[s]) + T[s+m]
*/
func (a *RabinKarpMatch) Match(text, pattern string, radix_d, mod_q int) []int {
	result := []int{}
	n := len(text)
	m := len(pattern)
	if m > n {
		return result
	}
	h := PowInt(radix_d, m-1) % mod_q
	var p int = 0
	var t0 int = 0

	//预处理，先求出边界位置的t值才能使用递推式
	//通过t[s]得出t[s+1]，所以必然需要计算出第一个值t[0]
	//同时要将计算出的p值和t[i]进行比较，所以p也是要计算出来的，p只需要计算一次，以后跟每个t[i]比较即可
	for i := 0; i < m; i++ {
		p = (radix_d*p + int(pattern[i])) % mod_q
		t0 = (radix_d*t0 + int(text[i])) % mod_q
	}

	//************* 遍历匹配   ****************
	var ts int = t0
	for s := 0; s <= n-m; s++ {
		//伪命中点，需要做进一步的检查。排除掉的一定不是，没有排除的不一定是
		if p == ts {
			//********* 初步排除之后需要做精确匹配  **********
			matched := true
			for j := 0; j < m; j++ {
				if text[s+j] != pattern[j] {
					matched = false
					break
				}
			}
			if matched {
				result = append(result, s)
			}
		}
		//当 s!=n-m时，向后循环推进
		if s < n-m {
			s_int := int(text[s])
			sm_int := int(text[s+m])
			ts = (radix_d*(ts+s_int*h*(mod_q-1)) + sm_int) % mod_q
		}
	}
	return result
}
