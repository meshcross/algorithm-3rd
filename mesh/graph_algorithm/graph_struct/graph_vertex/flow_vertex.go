/*
 * @Description: 推送重贴标签算法图的顶点 FlowVertex
 * @Author: wangchengdg@gmail.com
 * @Date: 2020-02-18 10:30:56
 * @LastEditTime: 2020-03-13 22:27:11
 * @LastEditors:
 */
package GraphVertex

type IFlowVertex interface {
	GetHeight() int
	SetHeight(v int)
	GetExceed() int
	SetExceed(v int)

	GetID() int
	GetKey() int
	GetParent() IVertex
	SetID(v int)
	SetKey(v int)
	SetParent(v IVertex)
}

/*!
* 推送-重贴标签算法图的顶点
* FlowVertex 是 Vertex的子类。它比Vertex多了一个`int h` 成员变量。其中：
*
* - `KType key`：表示顶点的超额流量
* - `int h`：表示顶点的高度
 */
type FlowVertex struct {
	Vertex
	_Height int //顶点高度
	_Exceed int //超出部分
}

func (vtx *FlowVertex) GetHeight() int {
	return vtx._Height
}
func (vtx *FlowVertex) GetExceed() int {
	return vtx._Exceed
}
func (vtx *FlowVertex) SetHeight(v int) {
	vtx._Height = v
}
func (vtx *FlowVertex) SetExceed(v int) {
	vtx._Exceed = v
}

func NewFlowVertex(h int, k int, ids ...int) *FlowVertex {
	id := -1
	if len(ids) > 0 {
		id = ids[0]
	}
	return &FlowVertex{Vertex: Vertex{_ID: id, _Key: k}, _Height: h}
}
