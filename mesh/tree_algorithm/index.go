/*
 * @Description: golang algorithm
 * @Author: wangchengdg@gmail.com
 * @Date: 2020-03-04 16:10:31
 * @LastEditTime: 2020-03-04 16:11:40
 * @LastEditors:
 */

package TreeAlgorithm

func IsLeftChild(node ITreeNode) bool {
	p := node.GetParent()
	return p != nil && p.GetLChild() == node
}

func IsRightChild(node ITreeNode) bool {
	p := node.GetParent()
	return p != nil && p.GetRChild() == node
}
