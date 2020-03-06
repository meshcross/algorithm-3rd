/*
 * @Description: golang algorithm
 * @Author: wangchengdg@gmail.com
 * @Date: 2020-02-18 10:30:45
 * @LastEditTime: 2020-03-05 11:57:06
 * @LastEditors:
 */
package GraphVertex

import . "github.com/meshcross/algorithm-3rd/mesh/common"

type DFSVertex struct {
	Vertex
	DiscoverTime int   /*!< 顶点发现时间*/
	FinishTime   int   /*!< 顶点完成时间*/
	Color        COLOR /*!< 顶点颜色*/
}

//!set_disovered：发现本顶点
/*!
*
* \param discover_t:发现时间
*
* 发现本顶点要执行两个操作：
*
* - 将本顶点的颜色设为灰色
* - 将本顶点的`discover_time`设为`discover_t`
 */
func (a *DFSVertex) SetDisovered(discover_t int) {
	a.DiscoverTime = discover_t
	a.Color = COLOR_GRAY
}

//!set_finished：设本顶点为搜索完毕状态
/*!
*
* \param finish_t:完成时间
*
* 设本顶点为搜索完毕状态要执行两个操作：
*
* - 将本顶点的颜色设为黑色
* - 将本顶点的`finish_time`设为`finish_t`
 */
func (a *DFSVertex) SetFinished(finish_t int) {
	a.Color = COLOR_BLACK
	a.FinishTime = finish_t
}
func (a *DFSVertex) GetDFSParent() *DFSVertex {
	if a._Parent != nil {
		ptr, ok := (a._Parent).(*DFSVertex)
		if ok && ptr != nil {
			return ptr
		}
	}
	return nil
}
func NewDFSVertex(k int, ids ...int) *DFSVertex {
	id := -1
	if len(ids) > 0 {
		id = ids[0]
	}
	return &DFSVertex{
		Vertex:       Vertex{_Key: k, _ID: id, _Parent: nil},
		DiscoverTime: -1,
		FinishTime:   -1,
		Color:        COLOR_WHITE,
	}
}
