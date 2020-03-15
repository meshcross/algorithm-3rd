/*
 * @Description: 图的结点，做了一个公共实现，相当于基类，实际上各种Vertex都是遵循IVertex接口，实际使用中并没有使用类的继承
 * @Author: wangchengdg@gmail.com
 * @Date: 2020-02-18 10:31:34
 * @LastEditTime: 2020-03-13 22:29:40
 * @LastEditors:
 */
package GraphVertex

/*!
*
* 顶点是一个结构体的模板，模板类型为它存储的数据的类型。它主要包含两个数据：
*
*   - `key`:顶点存放的数据
*   - `id`:顶点的编号（从0开始，负编号的节点为无效节点）。它是一个`const int`，一旦顶点初始化完毕就不可更改。
*
* 默认情况下：`id`为-1,`key`为`T()`。
*
 */

type IVertex interface {
	GetID() int
	GetKey() int
	GetParent() IVertex
	SetID(v int)
	SetKey(v int)
	SetParent(v IVertex)
}
type Vertex struct {
	_ID     int
	_Key    int
	_Parent IVertex
}

func (vtx *Vertex) GetID() int {
	return vtx._ID
}
func (vtx *Vertex) GetKey() int {
	return vtx._Key
}
func (vtx *Vertex) GetParent() IVertex {
	return vtx._Parent
}

func (vtx *Vertex) SetID(v int) {
	vtx._ID = v
}
func (vtx *Vertex) SetKey(v int) {
	vtx._Key = v
}
func (vtx *Vertex) SetParent(v IVertex) {
	vtx._Parent = v
}

func NewVertex(k int, ids ...int) *Vertex {
	id := -1
	if len(ids) > 0 {
		id = ids[0]
	}
	return &Vertex{_ID: id, _Key: k, _Parent: nil}
}
