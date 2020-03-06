/*
 * @Description: 32.3 利用有限自动机进行字符串匹配
 * @Author: wangchengdg@gmail.com
 * @Date: 2020-02-10 22:19:27
 * @LastEditTime: 2020-03-06 12:55:37
 * @LastEditors:
 */
package StringMatchingAlgorithm

import (
	"errors"
	"fmt"

	. "github.com/meshcross/algorithm-3rd/mesh/common"
)

type FiniteAutomatonMatch struct {
}

func NewFiniteAutomatonMatch() *FiniteAutomatonMatch {
	return &FiniteAutomatonMatch{}
}

/*!
* @description : 返回返回字符c的位置
* @param strM : 字符集序列M,为中的字符是不可重合的
* @param c: 字符c
* @return: 返回字符c的位置
*
* 通过逐一比较来返回a在M中的位置，M是T+P中所有字符的集合，M是一个字典，没有重复项
*
 */
func (a *FiniteAutomatonMatch) charIndexInM(strM string, c byte) (int, error) {
	lenM := len(strM)

	pos := 0
	result_iter := 0
	//找到c出现的位置
	for pos = 0; pos < lenM; pos++ {
		if strM[pos] == c {
			result_iter = pos
			pos++
			break
		}
	}

	if result_iter == lenM {
		return -1, errors.New(fmt.Sprintf("charIndexInM error : M has no charactor %s", string(c)))
	}
	return result_iter, nil
}

/*!
* @description 返回Pk是否是( Pq a)的后缀
* @param begin : 模式序列P的起始迭代器
* @param k: Pk的终止位置
* @param q : Pq的终止位置
* @param a: 字符 a
* @return: Pk是否是( Pq a)的后缀
*
* 通过逐一比较来返回Pk是否是( Pq a)的后缀。
*
 */
func (f *FiniteAutomatonMatch) endWith(str string, k, q int, a byte) bool {

	if k < 0 || q < 0 {
		panic("endWith params error")
		//return false
	}

	if k == 0 {
		return true // 空字符串是所有字符串的后缀
	}

	if a != str[k-1] {
		return false //P[k]!=a
	}

	for i := 0; i < k-1; i++ {
		if str[k-2-i] != str[q-1-i] {
			return false //P[k-i-1]!=P[q-i]
		}
	}

	return true
}

/*!
* @description 获取模式字符串的转移函数
* @param strP : 查找的关键词序列
* @param strM: M 相当于字典，P和T中可能一个字符出现多次，在M中都会出现一次，且只有一次
* @return: delta [][]int
*
* 步骤：
*
* - 遍历P，q从0到m (因为q=0时，P_0=空字符串):
*   - 对每个字符a属于有限字母集合M，寻找Pk是 (Pq a) 后缀的最大的k，则 delt(q,a)=k
*
*
 */
func (a *FiniteAutomatonMatch) getDelta(strP, strM string) [][]int {

	lenP := len(strP)
	lenM := len(strM)

	delta := make([][]int, lenP+1)
	for k, _ := range delta {
		delta[k] = make([]int, lenM)
	}

	for q := 0; q <= lenP; q++ {

		for m := 0; m < lenM; m++ {
			//*********** 寻找P[k]是 (P[q] a) 后缀的最大的k ***********
			k := MinInt(lenP+1, q+2)
			k--
			for !a.endWith(strP, k, q, strM[m]) {
				k--
			}

			//******  delt(q,a)=k *******
			delta[q][m] = k
		}
	}

	return delta

}

/*
* @description 有限自动机字符串匹配算法
* @param T : text，比较多的字符串，在其中进行查找操作
* @param P : pattern，需要被查找的字符串，通常为一个词
* @param M : T和P中的元素都来自集合M
* @return: T中所有能匹配P的偏移位置，即T中能找到几个P，每个P在哪里
*
* ## 字符串匹配
*
* 字符串匹配的形式化定义如下：假设文本是一个长度为n的数组 T[1...n]，而模式是一个长度为m的数组P[1...m]，
* 其中m<=n。进一步假设P和T的元素都是来自一个有限字母集合M的字符。如M={0,1}或者M={a,b,c,...z}。
* 字符数组P和T通常称为字符串。
*
* ## 有限自动机 字符串匹配算法
*
* ### 算法原理
*
*
* #### 字符串匹配自动机
*
* 对于一个给定的模式P，我们可以在预处理阶段构造出一个字符串匹配自动机，根据模式构造出的自动机后，再利用它来搜寻文本字符串。
* 首先定义一个辅助函数 sigma，称之为P的后缀函数(endWith)。函数 sigma是一个 M*到{0,1,....m}上的映射：
*
* - sigma(x)=max{k:P[k]是x的后缀}，即sigma(x)是x的后缀中，P的最长前缀的长度。
*
* 因为空字符串P0=e是每一个字符串的后缀，因此sigma(e)=0。对于一个长度为m的模式P，sigma(x)=m当且仅当P是x的后缀。
*
* 给定模式P[1...m]，其相应的字符串匹配自动机定义如下：
*
* - 状态集合Q为{0,1,...m}。开始状态q_0为0状态，并且只有状态m是唯一被接受的状态。
* - 对任意状态q和字符a，转移函数 delt定义为： delt(q,a)=sigma(P[q] a)  注：q是一个位置索引
*
*
* 考虑最近一次扫描T的字符。为了使得T的一个子串（以T[i]结尾的子串）能够和P的某些前缀Pj匹配，则前缀Pj必须是Ti的一个后缀。
* 假设q=phai(Ti)，则读完Ti之后，自动机处于状态q(处于位置q)。转移函数delt使用状态数q表示P的前缀和Ti后缀的最长匹配长度。也就是说，
* 在状态q是， Pq是Ti的后缀，且q=sigma(Ti)。
*
*    ----------------------------------------------------------
* T   1 , 2 , 3 ,....., i-q+1 ,..........., i ,.............., n   :Ti=T[1...i]
*    ----------------------------------------------------------
*                        |<-----长度为q---->|
*                      --------------------------------
*                P       1  , 2 ,......., q ,...., m    :Pq=P[1...q]
*                      --------------------------------
*
*
*  ### 算法步骤
*
* #### 预处理算法（构造delt函数)
*
* - 遍历P，q从0到m (因为q=0时，P_0=空字符串):
*   - 对每个字符a属于有限字母集合a，寻找Pk是 (Pq a) 后缀的最大的k，则 delt(q,a)=k
* - 返回 delt
*
* #### 匹配算法
* - 遍历T，i从1到n:
*   - 计算 q=delt(q,T[i])。如果 q==m，则偏移 i-m是有效偏移点，将 i-m 加入结果result中
*
* ### 算法性能
*
* 有限自动机字符串匹配算法的预处理时间为O(m^3 |M|)，其中|M| 为有限字母集合的大小，匹配时间为O(n)
*
* golang 中string作为参数传递的时候，虽然会new一个string，但是底层实际存储字符串的数据结构不会new，会把旧的指针传给新的string
* 所以string作为参数传递的时候，机会和string*的性能是一样的
 */
func (a *FiniteAutomatonMatch) Match(strT, strP, strM string) ([]int, error) {

	lenT := len(strT)
	lenP := len(strP)
	lenM := len(strM)

	if lenT < 0 {
		return nil, errors.New("match error:T error")
	}

	if lenP <= 0 {
		return nil, errors.New("match error:P error")
	}

	if lenM <= 0 {
		return nil, errors.New("match error:M error")
	}
	//#### M作为字典，不可以有重复项
	dict := map[byte]bool{}
	for i := 0; i < lenM; i++ {
		dict[strM[i]] = true
	}
	if len(dict) < lenM {
		return nil, errors.New("match error:M error,some items is duplicate!")
	}

	result := []int{}
	if lenT < lenP {
		return result, nil
	}

	//**********  预处理设定Delta  **************
	delta := a.getDelta(strP, strM)

	//*********** 匹配 ***************
	q := 0
	for i := 0; i < lenT; i++ {
		tiIndex, _ := a.charIndexInM(strM, strT[i])
		q = delta[q][tiIndex]
		if q == lenP {
			//[0,1,...i] ，其右侧长度为lenP的区间为[i-lenP+1,i-lenP+2,...i]
			result = append(result, i-lenP+1)
		}
	}
	return result, nil
}
