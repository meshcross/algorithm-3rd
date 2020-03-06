/*
 * @Description: 定义三元组和二元组
 * @Author: wangchengdg@gmail.com
 * @Date: 2020-02-19 08:45:35
 * @LastEditTime: 2020-03-06 13:00:44
 * @LastEditors:
 */
package Common

type CompareFunc func(x, y interface{}) int
type TupleCompareFunc func(x, y *Tuple) int

func TupleCompareFunc_Less(e1, e2 *Tuple) int {
	if e1.First == e2.First && e1.Second == e2.Second && e1.Third == e2.Third {
		return 0
	}
	if e1.First < e2.First ||
		(e1.First == e2.First) && (e1.Second < e2.Second) ||
		((e1.First == e2.First) && (e1.Second == e2.Second)) && (e1.Third < e2.Third) {
		return 1
	}

	return -1
}

type Pair struct {
	First, Second int
}

func NewPair(first, second int) *Pair {
	return &Pair{First: first, Second: second}
}

type Tuple struct {
	First, Second, Third int
}

func NewTuple(first, second, third int) *Tuple {
	return &Tuple{First: first, Second: second, Third: third}
}

type PairAny struct {
	Fisrt, Second interface{}
	Compare       CompareFunc
}

type TupleAny struct {
	First, Second, Third interface{}
	Compare              CompareFunc
}

type TupleWapper struct {
	Tuples  []*Tuple
	Compare TupleCompareFunc
}

func (tw *TupleWapper) Len() int {
	return len(tw.Tuples)
}

func (tw *TupleWapper) Swap(i, j int) {
	tw.Tuples[i], tw.Tuples[j] = tw.Tuples[j], tw.Tuples[i]
}

func (tw *TupleWapper) Less(i, j int) bool {
	return tw.Compare(tw.Tuples[i], tw.Tuples[j]) > 0
}

func NewTupleWapper(tuples []*Tuple, compare TupleCompareFunc) *TupleWapper {
	return &TupleWapper{Tuples: tuples, Compare: compare}
}
