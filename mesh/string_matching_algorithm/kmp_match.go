/*
 * @Description: 算法导论第32章32.4节, KMP字符串匹配算法
 * @Author: wangchengdg@gmail.com
 * @Date: 2020-02-10 22:19:41
 * @LastEditTime: 2020-03-06 11:18:06
 * @LastEditors:
 */
package StringMatchingAlgorithm

import "errors"

type KmpMatch struct {
}

func NewKmpMatch() *KmpMatch {
	return &KmpMatch{}
}

/*!
* @description:KMP字符串匹配算法的预处理操作
* @param strP : pattern，需要查询的字符串
* @return: pai函数，[]int
*
* 算法步骤：
*
* - 初始化： pai[1]=0,k=0
* - 遍历 q 从 2 到 m:（因为Pk必须是Pm的真子集，因此m从2开始）
*   - 循环，条件为 k>0并且 P[k+1]!=P[q]；循环中执行 k=pai[k]（因为若P[k+1]=P[q]，则说明找到了Pk是Pm的真子集）
*   - 若 P[k+1]==P[q]，则k=k+1并且pai[q]=k
* - 返回 pai
*
* 需要注意的是pai函数只和strP相关，计算量相对而言是比较有限的
* 这里需要一个清晰的逻辑，见下图：
* “部分匹配值”就是”前缀”和”后缀”的最长的共有元素的长度。以”ABCDABD”为例，计算部分匹配值：
	-“A”的前缀和后缀都为空集，共有元素的长度为0；
	-“AB”的前缀为[A]，后缀为[B]，共有元素的长度为0；
	-“ABC”的前缀为[A, AB]，后缀为[BC, C]，共有元素的长度0；
	-“ABCD”的前缀为[A, AB, ABC]，后缀为[BCD, CD, D]，共有元素的长度为0；
	-“ABCDA”的前缀为[A, AB, ABC, ABCD]，后缀为[BCDA, CDA, DA, A]，共有元素为”A”，长度为1；
	-“ABCDAB”的前缀为[A, AB, ABC, ABCD, ABCDA]，后缀为[BCDAB, CDAB, DAB, AB, B]，共有元素为”AB”，长度为2；
	-“ABCDABD”的前缀为[A, AB, ABC, ABCD, ABCDA, ABCDAB]，后缀为[BCDABD, CDABD, DABD, ABD, BD, D]，共有元素的长度为0。
*/
func (a *KmpMatch) compute_prefix_func(strP string) []int {

	lenP := len(strP)
	pai := make([]int, lenP)

	k := 0
	for q := 1; q < lenP; q++ { //P[2...m]

		for k > 0 && strP[k] != strP[q] {
			k = pai[k]
		}

		if strP[k] == strP[q] {
			k++
		}

		pai[q] = k
	}
	return pai
}

/*!
* @description:KMP字符串匹配算法
* @param strT: 在哪个文本中查找
* @param strP: 模式pattern
* @return: 所有查找到的位置
*
* ## 字符串匹配
*
* 字符串匹配的形式化定义如下：假设文本是一个长度为n的数组 T[1...n]，而模式是一个长度为m的数组P[1...m]，
* 其中m<=n。进一步假设P和T的元素都是来自一个有限字母集合M的字符。如M={0,1}或者M={a,b,c,...z}。
* 字符数组P和T通常称为字符串。
*
* ## KMP 字符串匹配算法
*
* #### kmp 算法
*
* KMP 算法用到了辅助函数 pai，它在O(m)时间内根据模式预先计算出pai并且存放在数组pai[1...m]中。
* 数组pai能够使我们按照需要即时计算出转移函数。
*
* 计算出pai数组之后，KMP算法从左到右扫描文本序列T，并从pai中获取转移函数。当状态结果为 m时，
* 当前偏移为有效偏移点。
*
* ### 算法步骤
*
* #### 预处理算法（构造pai函数)
*
* - 初始化： pai[1]=0,k=0
* - 遍历q 从 2 到 m:（因为Pk必须是Pm的真子集，因此m从2开始）
*   - 循环，条件为 k>0并且 P[k+1]!=P[q]；循环中执行 k=pai[k]（因为若P[k+1]=P[q]，则说明找到了Pk是Pm的真子集）
*   - 若 P[k+1]==P[q]，则k=k+1并且pai[q]=k
* - 返回 pai
*
*
* #### 匹配算法
*
* - 初始化 q=0
* - 遍历i从1到n:
*   - 循环，条件为 q>0 并且 P[q+1]!=T[i]；在循环中执行 q=pai[q]
*   - 如果 P[q+1]==T[i] 则 q=q+1
*   - 如果 q==m，则找到了有效偏移点。将有效偏移加入结果result中。然后 q=pai[q](比如有这一句，否则后面P[q+1]会溢出)
* - 返回结果 result
*
*
* 计算前缀函数的运行时间为 O(m)，匹配时间为O(n)，总运行时间为 O(n)
*
* 主体的逻辑和native_string_match很像，只是引入了pai函数，这样可以少一些循环
*
 */
func (a *KmpMatch) Match(strT, strP string) ([]int, error) {

	lenT := len(strT)
	lenP := len(strP)

	if lenT < 0 {
		return nil, errors.New("match error:strT")
	}

	if lenP <= 0 {
		return nil, errors.New("match error:strP")
	}

	result := []int{}
	//模式串P较长，不用比
	if lenT < lenP {
		return result, nil
	}

	//**********  预处理  **************
	pai := a.compute_prefix_func(strP)
	//*********** 匹配 ***************
	q := 0
	for i := 0; i < lenT; i++ {
		//右移直到P[q+1]==T[i]，这里从0计数
		for q > 0 && strP[q] != strT[i] {
			q = pai[q]
		}
		//确实发生P[q+1]==T[i]，这里从0计数
		if strP[q] == strT[i] {
			q++
		}

		if q == lenP { //找到有效偏移点
			result = append(result, i-lenP+1) //i左侧lenP的位置
			q = pai[lenP-1]
		}
	}
	return result, nil
}
