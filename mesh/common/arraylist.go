/*
 * @Description: 实现动态数组
 * @Author: wangchengdg@gmail.com
 * @Date: 2020-03-08 22:20:02
 * @LastEditTime: 2020-03-10 12:02:10
 * @LastEditors:
 */
package Common

type ArrayList struct {
	_datas    []interface{}   //实际存储的数据
	_capacity int             //容量多少
	_size     int             //当前有多少数据
	_compare  NodeCompareFunc //元素的比较函数
}

/**
 * @description:新建动态数组
 * @param n:初始容量
 * @param compare:比较函数
 * @return:
 */
func NewArrayList(n int, compare NodeCompareFunc) *ArrayList {
	n = MaxInt(n, 8)
	datas := make([]interface{}, n)
	return &ArrayList{_datas: datas, _capacity: n, _size: 0}
}

func (l *ArrayList) Datas() []interface{} {
	return l._datas

}
func (l *ArrayList) Size() int {
	return l._size
}
func (l *ArrayList) Capacity() int {
	return l._capacity
}

/**
 * @description: 分为自动调整容量和手动调整容量两种情况
	1、自动调整：容量已满，或者当前占有量不到容量的1/4的时候
	2、手动调整：需要输入cnts参数，表明要调整到多少
 * @param cnts ：可选，表明需要手动调整到多大的容量
 * @return: void
*/
func (l *ArrayList) checkResize(cnts ...int) {
	cnt := 0
	if len(cnts) > 0 {
		cnt = cnts[0]
	}
	//到8就不能再缩了
	if l._capacity > 8 || cnt >= 8 {
		new_capacity := cnt
		if cnt <= 0 {
			if l._size == l._capacity {
				new_capacity = l._capacity * 2
			} else if l._size < l._capacity/4 {
				new_capacity = l._capacity / 2
			}
		}

		if new_capacity > 0 {

			datas := make([]interface{}, new_capacity)
			for i := 0; i < l._size; i++ {
				datas[i] = l._datas[i]
			}
			l._datas = datas
			l._capacity = new_capacity
		}
	}
}

func (l *ArrayList) Append(item interface{}) {
	l.checkResize()

	l._datas[l._size] = item
	l._size++
}

/**
 * @description: Clear只是清理数据，并没有释放多余的空间
 * @return: void
 */
func (l *ArrayList) Clear() {
	for i := 0; i < l._size; i++ {
		l._datas[i] = nil
	}
	l._size = 0
}

/**
 * @description: 插入到某个位置
 * @param {type}
 * @return:
 */
func (l *ArrayList) Insert(item interface{}, index int) bool {
	l.checkResize()
	//等于也是可以的，表示最佳
	if index >= 0 && index <= l._size {
		//整体后移，将位置index空出来
		for i := l._size; i > index; i-- {
			l._datas[i] = l._datas[i-1]
		}
		l._datas[index] = item
		l._size++
		return true
	}
	return false
}

func (l *ArrayList) DeleteAt(index int) (bool, interface{}) {
	if index >= 0 && index < l._size {
		ret := l._datas[index]
		//整体前移，最后一个位置空出来，然后置为nil
		for i := index; i < l._size-1; i++ {
			l._datas[i] = l._datas[i+1]
		}
		l._datas[l._size-1] = nil
		l._size--
		return true, ret
	}
	return false, nil
}

func (l *ArrayList) IndexOf(item interface{}) int {
	for i := 0; i < l._size; i++ {
		if l._datas[i] == item {
			return i
		}
	}
	return -1
}

func (l *ArrayList) ItemAt(index int) interface{} {
	if index >= 0 && index < l._size {
		return l._datas[index]
	}
	return nil
}

func (l *ArrayList) Delete(item interface{}) bool {
	move := false
	for i := 0; i < l._size; i++ {
		if !move {
			if l._datas[i] == item {
				move = true
			}
		} else {
			l._datas[i-1] = l._datas[i]
		}
	}
	if move {
		l._datas[l._size-1] = nil
		l._size--
		l.checkResize()
		return true
	}
	return false
}

func (l *ArrayList) Concate(listPtr *ArrayList) {
	if listPtr != nil {
		size1 := l.Size()
		size2 := (*listPtr).Size()
		needSize := size1 + size2
		if needSize > l.Capacity() {
			new_capacity := l.Capacity()
			for i := 0; i < 32; i++ {
				new_capacity = new_capacity * 2
				if new_capacity > needSize || new_capacity >= INT_MAX/2 {
					break
				}
			}

			l.checkResize(new_capacity)

			datas := (*listPtr).Datas()
			for i := 0; i < size2; i++ {
				l.Append(datas[i])
			}
		}
	}
}

func (l *ArrayList) Clone() *ArrayList {
	list := NewArrayList(l._capacity, l._compare)
	for i := 0; i < l._size; i++ {
		list._datas[i] = l._datas[i]
	}
	list._size = l._size
	return list
}

func (l *ArrayList) Empty() bool {
	return l._size <= 0
}
