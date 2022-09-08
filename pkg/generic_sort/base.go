package generic_sort

/*
Comparator:
   将用来约束数据结构，数据的比较大小，比较情况，以及交换
   方法的自定义为了数据结构中更多的自定义，将放大了操作不同数据元素的粒度，不只限于builtin的数据类型
   ex.自定义结构类型的比较大小，比较情况，交换都可以提前在调用数据结构初始化之前定义好处理情况
func info:
   IsLess 逻辑大小比较，正序比较为 i < j，反序比较为 i > j，根据需要还能定义 i <= j情况
   IsEqual 逻辑比较， 自定义两个数据元素是否相等
   Exchange 逻辑交换，交换的数据元素能够被自定义，可以操作数据元素的内部的变量等。
*/

type Comparator[T any] interface {
	IsLess(i, j *T) bool // ex:if i < j , then return true, else return false.
	IsEqual(i, j T) bool // ex:if i == j, then return true, else return false.
	Swap(i, j *T) (T, T) // ex: return *j, *i
}

type Number interface {
	int | int8 | int16 | int32 | int64 | float32 | float64 | uint | uint8 | uint16 | uint32 | uint64 | string
}

type NumberComparator[T Number] struct {
	orderFlag bool // if orderFlag is true , then use *i > *j, else *i < *j
}

func (nc *NumberComparator[T]) SetOrder(flag bool) {
	nc.orderFlag = flag
}

func (nc *NumberComparator[T]) IsLess(i, j *T) bool {
	if nc.orderFlag {
		return *i > *j
	}

	return *i < *j
}

func (nc *NumberComparator[T]) IsEqual(i, j T) bool {
	return i == j
}

func (nc *NumberComparator[T]) Swap(i, j *T) (T, T) {
	return *j, *i
}
