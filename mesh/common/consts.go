/*
 * @Description: 定义常量
 * @Author: wangchengdg@gmail.com
 * @Date: 2020-02-20 11:57:00
 * @LastEditTime: 2020-03-14 23:40:58
 * @LastEditors:
 */
package Common

type COLOR uint8

const (
	COLOR_GRAY  COLOR = 1
	COLOR_BLACK COLOR = 2
	COLOR_WHITE COLOR = 3

	COLOR_RED COLOR = 5
)

//图的表示方法
const (
	GRAPH_REPRESENTION_MATRIX = "matrix"    //矩阵表示法，用于稠密图
	GRAPH_REPRESENTION_ADJ    = "adjacency" //邻接矩阵表示法，用于稀疏图
)
