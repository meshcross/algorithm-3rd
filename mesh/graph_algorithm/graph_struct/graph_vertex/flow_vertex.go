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

//!FlowVertex：推送-重贴标签算法图的顶点，算法导论26章26.4节
/*!
*
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
