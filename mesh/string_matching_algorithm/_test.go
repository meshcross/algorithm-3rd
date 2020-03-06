/*
 * @Description: golang algorithm
 * @Author: wangchengdg@gmail.com
 * @Date: 2020-03-05 22:48:50
 * @LastEditTime: 2020-03-05 22:51:29
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
}
