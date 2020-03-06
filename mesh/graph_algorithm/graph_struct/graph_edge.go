/*
 * @Description: 算法导论22章22.1节，图的边
 * @Author: wangchengdg@gmail.com
 * @Date: 2020-02-18 10:28:40
 * @LastEditTime: 2020-03-06 11:27:47
 * @LastEditors:
 */
package GraphStruct

type Edge struct {
	vertex1 int /*!< 边的第一个顶点*/
	vertex2 int /*!< 边的第二个顶点*/
	weight  int /*!< 边的权重*/
}
