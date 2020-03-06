/*
 * @Description: golang algorithm
 * @Author: wangchengdg@gmail.com
 * @Date: 2020-02-20 12:48:02
 * @LastEditTime: 2020-03-05 12:01:27
 * @LastEditors:
 */
package GraphStruct

import (
	"unsafe"

	. "github.com/meshcross/algorithm-3rd/mesh/graph_algorithm/graph_struct/graph_vertex"
)

func BFSVertex2Vertex(vtx *BFSVertex) *Vertex {

	ptr := unsafe.Pointer(vtx)
	v := (*Vertex)(ptr)
	return v
}

func DFSVertex2Vertex(vtx *DFSVertex) *Vertex {

	ptr := unsafe.Pointer(vtx)
	v := (*Vertex)(ptr)
	return v
}
func SetVertex2Vertex(vtx *SetVertex) *Vertex {

	ptr := unsafe.Pointer(vtx)
	v := (*Vertex)(ptr)
	return v
}
func Vertex2SetVertex(vtx *Vertex) *SetVertex {

	ptr := unsafe.Pointer(vtx)
	v := (*SetVertex)(ptr)
	return v
}

func ToSetVertex(vtx IVertex) *SetVertex {
	if ptr, ok := vtx.(*SetVertex); ok {
		return ptr
	}
	return nil
}
func ToDFSVertex(vtx IVertex) *DFSVertex {
	if ptr, ok := vtx.(*DFSVertex); ok {
		return ptr
	}
	return nil
}
func ToBFSVertex(vtx IVertex) *BFSVertex {
	if ptr, ok := vtx.(*BFSVertex); ok {
		return ptr
	}
	return nil
}
