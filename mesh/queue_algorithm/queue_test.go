/*
 * @Description: golang algorithm
 * @Author: wangchengdg@gmail.com
 * @Date: 2020-02-10 20:09:31
 * @LastEditTime: 2020-03-05 12:06:30
 * @LastEditors:
 */
package QueueAlgorithm

import (
	"testing"

	. "github.com/meshcross/algorithm-3rd/mesh"
	. "github.com/meshcross/algorithm-3rd/mesh/common"
)

func TestTask(t *testing.T) {

	Q_NUM := 10
	_strcut_minqueue := NewMinQueue(NodeCompareFunc_IntLessThan, nil)
	for i := Q_NUM - 1; i >= 0; i-- {

		_strcut_minqueue.Insert(i)
		m, _ := _strcut_minqueue.Min()
		//fmt.Println(fmt.Sprintf("a-min:%v,%d", m, i))
		EXPECT_EQ(m, i, t)
	}
	_strcut_minqueue.Print()
	//m, _ := _strcut_minqueue.Min()
	//fmt.Println(fmt.Sprintf("b-min:%v,%d", m, 0))

	for i := 0; i < Q_NUM; i++ {
		m, _ := _strcut_minqueue.ExtractMin()
		//fmt.Println(fmt.Sprintf("c-min:%v,%d", m, i))
		EXPECT_EQ(m, i, t)
	}
}
