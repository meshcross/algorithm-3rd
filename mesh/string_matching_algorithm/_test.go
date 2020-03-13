/*
 * @Description: golang algorithm
 * @Author: wangchengdg@gmail.com
 * @Date: 2020-03-05 22:48:50
 * @LastEditTime: 2020-03-13 09:42:01
 * @LastEditors:
 */
package StringMatchingAlgorithm

import (
	"fmt"
	"testing"
)

func TestFiniteAutomatonMatch(t *testing.T) {
	m := NewFiniteAutomatonMatch()

	T := "abcabaabcabac"
	P := "abaa"
	M := "abc" //字母集合，不可重复，又要覆盖T和P

	vec, err := m.Match(T, P, M)
	expect := []int{3}

	fmt.Println(vec, err, expect)
}

func TestKmpMatch(t *testing.T) {
	m := NewKmpMatch()

	text := "abcabaabcabac"
	pattern := "abaa"

	vec, err := m.Match(text, pattern)
	expect := []int{3}

	fmt.Println(vec, err, expect)

	m := NewKmpMatch()

	T := []string{
		"aabcbababacaabcbabx",
		"txasdgaabcbababacaabcbabx",
		"abcbababacaabcbabx",
		"aaaabcbababacaabcbab",
		"bcbababacaabcbabx",
	}
	P := "abcbab"

	for k, v := range T {
		vec1, err1 := m.Match(v, P)
		vec2, err2 := m.MatchX(v, P)

		fmt.Println("ver1:", k, vec1, err1)
		fmt.Println("ver2:", k, vec2, err2)
	}

}
