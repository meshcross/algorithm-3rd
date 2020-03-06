/*
 * @Description:
 * @Author: wangchengdg@gmail.com
 * @Date: 2020-02-18 18:55:09
 * @LastEditTime: 2020-03-06 11:45:48
 * @LastEditors:
 */
package GraphVertex

import (
	"errors"
)

/*!
* @description:获取两个顶点之间的路径
* @param v_from: 起始顶点
* @param v_to: 终止顶点
* @return : 两个顶点之间的路径包含的顶点的`id`序列
*
 */
func GetPath(v_from IVertex, v_to IVertex) ([]int, error) {
	if v_from == nil || v_to == nil {
		return nil, errors.New("get_path error: vertex must not be null!")
	}

	result := []int{}
	if v_from.GetID() == v_to.GetID() {
		result = append(result, v_from.GetID())
	} else if v_to.GetParent() == nil {
		return result, nil
	} else {
		temp, _ := GetPath(v_from, v_to.GetParent())
		for _, v := range temp {
			result = append(result, v)
		}
		result = append(result, v_to.GetID())
	}
	return result, nil
}

/**
 * @description: 用于节点的比较
 * @param v1 : 参与比较的节点1
 * @param v2: 参与比较的节点2
 * @return: v1<v2则返回1，==则返回0，>则返回-1
 */
func NodeCompareFunc_VertexLessThan(v1, v2 interface{}) int {
	if v1 == nil || v2 == nil {
		return -1
	}
	vtx1 := v1.(IVertex)
	vtx2 := v2.(IVertex)
	if vtx1.GetKey() < vtx2.GetKey() {
		return 1
	}
	if vtx1.GetKey() == vtx2.GetKey() {
		return 0
	}
	return -1
}

/**
 * @description: 所有To开头的类型转换都只是简单的拆箱操作，并不涉及到unsafe.Pointer转换
		程序尽量避免了使用unsafe.Pointer
 * @param v:需要拆箱的IVertex指针
 * @return: IFlowVertex类型的指针
*/
func ToIFlowVertex(v IVertex) IFlowVertex {
	if vtx, ok := v.(IFlowVertex); ok {
		return vtx
	}
	return nil
}

/**
 * @description: 所有To开头的类型转换都只是简单的拆箱操作，并不涉及到unsafe.Pointer转换
		程序尽量避免了使用unsafe.Pointer
 * @param v:需要拆箱的IVertex指针
 * @return: FrontFlowVertex类型的指针
*/
func ToFrontFlowVertex(v IVertex) *FrontFlowVertex {
	if vtx, ok := v.(*FrontFlowVertex); ok {
		return vtx
	}
	return nil
}
