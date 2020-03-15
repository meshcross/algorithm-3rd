/*
 * @Description: 图的广度优先搜索需要使用的节点类型BFSVertex
 * @Author: wangchengdg@gmail.com
 * @Date: 2020-02-18 10:30:35
 * @LastEditTime: 2020-03-13 22:31:23
 * @LastEditors:
 */
package GraphVertex

import (
	"errors"

	. "github.com/meshcross/algorithm-3rd/mesh/common"
)

type BFSVertex struct {
	Vertex
	Color COLOR //顶点颜色
	Deep  int
}

/*!
* @description:设本顶点为源点
* 将本顶点设为源点要执行两个操作：
*
* - 将本顶点的颜色设为灰色
* - 将本顶点的`parent`设为空
 */
func (a *BFSVertex) SetSource() {
	a.Color = COLOR_GRAY
	a._Parent = nil
	//a.Key = 0
	a.Deep = 0
}

/*!
* @description:发现本顶点
* @param v_parent:父顶点
*
* 发现本顶点要执行两个操作：
*
* - 将本顶点的颜色设为灰色
* - 将本顶点的`parent`设为`v_parent`
*
* 这里要求`v_parent`非空。若`v_parent`为空则抛出异常
 */
func (a *BFSVertex) SetFound(v_parent *BFSVertex) error {
	if v_parent == nil {
		return errors.New("set_found error: v_parent must not be nil!")
	}

	a.Color = COLOR_GRAY
	a._Parent = v_parent
	if v_parent.Deep == Unlimit() {
		a.Deep = v_parent.Deep
	} else {
		a.Deep = v_parent.Deep + 1
	}

	return nil
}

func (a *BFSVertex) GetBFSParent() *BFSVertex {
	if a._Parent != nil {
		ptr, ok := (a._Parent).(*BFSVertex)
		if ok && ptr != nil {
			return ptr
		}
	}
	return nil
}
func NewBFSVertex(k int, ids ...int) *BFSVertex {
	id := -1
	if len(ids) > 0 {
		id = ids[0]
	}
	return &BFSVertex{
		Vertex: Vertex{_Key: k, _ID: id, _Parent: nil},
		Color:  COLOR_WHITE,
		Deep:   -1,
	}
}
