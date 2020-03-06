/*
 * @Description: 第18章 B树
			B树是为磁盘或其他直接存取的辅助存储设备而设计的一种平衡搜索树。B树类似于红黑树，单它在降低磁盘I/O操作数方面要更好一些。
			许多数据库系统使用B树或者B树的变种来存储信息。
			B树与红黑树最大的不同在于，B树的节点可以有很多个孩子，数个到数千个。所以对于同样的节点数量，B树的深度要比红黑树小很多。

			B树具有以下性质：
			1、每个节点x存储n个关键字，x.key1<=x.key2<=x.key3<=...<=x.keyn；如果x为叶节点，则x.leaf为true，否则为false
			2、每个叶节点x还包含x.n+1个指向其孩子的指针x.c1,x.c2,...,x.c(x+1),叶节点没有孩子
			3、关键字x.key(i)对存储在各子树中的关键字范围加以分隔：如果ki为任意一个存储在以x.c(i)为根的子树中的关键字，那么
			  k1<=x.key1<=k2<=x.key2<=...<=x.keyn<=k_(n+1)
			4、每个叶节点具有相同的深度，即树的高度h
			5、每个节点锁包含的关键字个数有上界和下界。用一个被称为B树的最小度数(minimum degree)的固定整数t>=2来表示这些界：
				a、除了根节点以外的每个节点必须至少有t-1个关键字。因此，除了根节点以外的每个内部节点至少有t个孩子。如果树非空，根节点至少有一个关键字。
				b、每个节点至多可以包含2t-1个关键词。因此，一个内部节点至多可以有2t个孩子。当一个节点恰好有2t-1个关键词时，称该节点是满的。

 * @Author: wangchengdg@gmail.com
 * @Date: 2020-03-04 23:01:26
 * @LastEditTime: 2020-03-05 11:39:04
 * @LastEditors:
*/

package DataStruct

type BTree struct {
}

func NewBTree() *BTree {
	return &BTree{}
}

func (t *BTree) Search() {

}
func (t *BTree) Insert() {

}
