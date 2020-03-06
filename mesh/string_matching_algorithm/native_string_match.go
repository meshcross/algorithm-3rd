/*
 * @Description: 32.1 朴素字符串匹配

 * @Author: wangchengdg@gmail.com
 * @Date: 2020-03-05 15:48:11
 * @LastEditTime: 2020-03-05 23:38:37
 * @LastEditors:
 */

package StringMatchingAlgorithm

type NativeStringMatch struct {
}

func NewNativeStringMatch() *NativeStringMatch {
	return &NativeStringMatch{}
}

/**
 * @description: t中是否包含key
			朴素字符串匹配算法是通过一个循环找到所有有效偏移，该循环对n-m+1个可能的s值进行检测，看是否满足条件Key[-...m] = T[s+1...s+m]
 * @param t 字符串
 * @param key 想要在t中查找的关键词
 * @return: 返回t中首次找到的key的位置
 *
 * golang中字符串是以ptr方式作为参数传递的，所以没有性能问题
 *
 *
*/
func (a *NativeStringMatch) MatchOne(t, key string) int {
	n := len(t)
	m := len(key)
	if m > n {
		return -1
	}

	//没有使用t[s:s+i]，因为golang会生成新的string，这样会生成n-m个新的字符串，可能有性能问题
	for s := 0; s < n-m; s++ {
		match := true
	INNER:
		for i := 0; i < m; i++ {
			if t[s+i] != key[i] {
				match = false
				break INNER
			}
		}
		if match {
			return s
		}
	}
	return -1
}

/**
 * @description: 查找所有的匹配，而不只是一个
 * @param {type}
 * @return:
 */
func (a *NativeStringMatch) Match(text, pattern string) []int {
	n := len(text)
	m := len(pattern)

	result := []int{}
	if m > n {
		return result
	}

	//没有使用t[s:s+i]，因为golang会生成新的string，这样会生成n-m个新的字符串，可能有性能问题
	for s := 0; s < n-m; s++ {
		match := true
	INNER:
		for i := 0; i < m; i++ {
			if text[s+i] != pattern[i] {
				match = false
				break INNER
			}
		}
		//这里继续查找，而不是中断
		if match {
			result = append(result, s)
		}
	}
	return result
}
