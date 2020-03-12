/*
 * @Description: 第21章 21.3 不相交集合森林
 *
* 一个不相交集合数据结构维护了一个不相交动态集的集合{S1,S2,...Sk}。我们用一个代表来标识每个集合，它是这个集合的某个成员。
* 不相交集合支持三个操作：（x,y等都是集合中的对象)
*
* - MakeSet(x):建立一个新的集合，它的唯一成员是x
* - Union(x,y):将包含x和y的两个动态集合(表示为Sx和Sy)合并成一个新的集合。由于我们要求各个集合不相交，因此这里要消除原有的集合Sx和Sy。实际操作中，
*   我们把其中的一个集合的元素并入另一个集合中，来代替删除操作。
* - FindSet(x):返回一个指针，该指针指向包含x的唯一集合的代表。
*
* 在某些图的算法中，图和不相交集数据结构的表示需要相互引用。即一个表示顶点的对象会包含一个指向与之对应的不相交集合对象的指针；反之亦然
*
* 不相交集合森林：不相交集合森林是不相交集合的一种更快的实现。用有根数来表示集合，树中的每个结点都包含一个成员，每棵树代表一个集合。
* 在不相交集合森林中，每个成员仅仅指向它的父节点。每棵树的根就是集合的代表并且它的父节点就是自己。
*
* 这里采用了启发式策略改进运行时间，使用了两种启发式策略：
*   - 按秩合并：每个结点x维持一个整数值属性rank,它代表了x的高度（从x到某一后代叶结点的最长简单路径上的结点数目）的一个上界。在按秩合并的union操作中，
*   我们让具有较小秩的根指向具有较大秩的根
*   - 路径压缩：在`FindSet`操作中，使查找路径中的每个结点直接指向树根
*
* 如果单独采用按秩合并或者路径压缩，它们每一个都能改善不相交集合森林上操作的运行时间；而一起使用这两种启发式策略时，这种改善更大。
* 当同时使用按秩合并和路径压缩时，最坏情况下的运行时间为O(m*alpha*n))，这里alpha(n)是一个增长非常慢的函数。在任何一个可以想得到的不相交集合数据结构的应用中，
* alpha(n)<=4。其中n为结点个数，m为操作次数（运用了摊还分析）
 *
 * @Author: wangchengdg@gmail.com
 * @Date: 2020-02-10 22:17:11
 * @LastEditTime: 2020-03-01 18:34:21
 * @LastEditors:
*/
package SetAlgorithm

import "errors"

type DisJointSetNode struct {
	Value  interface{}
	Rank   int //秩
	Parent *DisJointSetNode
}

func NewDisJointSetNode(value interface{}) *DisJointSetNode {
	return &DisJointSetNode{Value: value}
}

func MakeSet(node *DisJointSetNode) error {
	if node != nil {
		node.Parent = node
		node.Rank = 0
	} else {
		return errors.New("MakeSet error: node must not be nil!")
	}
	return nil
}

/*!
* @description:返回结点所在集合的代表结点
* @param node:要查找的结点。它必须非空，否则抛出异常
* @return: 结点所在集合的代表结点
*
* 该操作简单沿着指向父节点的指针找到树的根。树的根的特征是：它的父节点就是它本身。
* 若结点不在不相交集合森林中（当结点的父节点指针为空时），则抛出异常。
*
* 计算过程是一个两阶段的方法，当它递归时，第一趟沿着查找路径向上直到找到树根；
* 当递归回溯时，第二趟沿着搜索树向下更新每个节点，使其父节点直接指向树根
*
 */
func FindSet(node *DisJointSetNode) (*DisJointSetNode, error) {
	if node == nil {
		return nil, errors.New("FindSet error: node must not be nil!")
	}

	if node.Parent != node {
		n, err := FindSet(node.Parent)
		if err != nil {
			return nil, err
		} else {
			node.Parent = n
		}
	}

	return node.Parent, nil
}

/*!
* @description:链接集合
* @param nodeX:待链接的第一个集合中的根节点。
* @param nodeY:待合并的第二个集合中的根节点
* @return error
*
* 每个结点x维持一个整数值属性rank,它代表了x的高度（从x到某一后代叶结点的最长简单路径上的结点数目）的一个上界。在链接时我们让具有较小秩的根指向具有较大秩的根.
*
* - 如果 nodeX或者nodeY为空，则直接返回
* - 如果 nodeX 和 nodeY非空，但是nodeX或者nodeY不是根结点，抛出异常
*
 */
func LinkSet(nodeX, nodeY *DisJointSetNode) error {
	if nodeX == nil || nodeY == nil {
		return errors.New("node is nil!")
	}

	if nodeX != nodeX.Parent || nodeY != nodeY.Parent {
		return errors.New("link_set error: node must be root of the set!")
	}

	if nodeX.Rank > nodeY.Rank {
		nodeY.Parent = nodeX
	} else {
		nodeX.Parent = nodeY
		if nodeX.Rank == nodeY.Rank {
			nodeY.Rank++
		}
	}
	return nil
}

//!union_set：合并集合
/*!
*
* \param nodeX:待合并的第一个集合中的某个结点
* \param nodeY:待合并的第二个集合中的某个结点
*
* 该操作首先获取每个结点所在集合的代表结点，然后将它们合并起来
*
 */
func UnionSet(nodeX, nodeY *DisJointSetNode) error {
	x, errX := FindSet(nodeX)
	if errX != nil {
		return errX
	}
	y, errY := FindSet(nodeY)
	if errY != nil {
		return errY
	}

	LinkSet(x, y)
	return nil
}
