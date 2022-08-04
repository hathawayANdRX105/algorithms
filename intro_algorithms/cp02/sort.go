package cp02

// InsertSort sustain a sequence array from start to end
func InsertSort(arr []int, cmp func(o1, o2 int) bool) []int {
	length := len(arr)
	if length < 2 {
		return arr
	}

	for j := 1; j < length; j++ {
		key := arr[j]

		i := j - 1
		for ; i > -1 && !cmp(arr[i], key); i-- {
			arr[i+1] = arr[i]
		}

		arr[i+1] = key

	}

	return arr
}

// SelectSort min/max num from sub array to insert into squence order
func SelectSort(arr []int, cmp func(o1, o2 int) bool) []int {
	length := len(arr)
	if length < 2 {
		return arr
	}

	for i := 0; i < length-1; i++ {
		ex_index := i
		for j := i + 1; j < length; j++ {
			if cmp(arr[j], arr[ex_index]) {
				ex_index = j
			}
		}
		Exchange(arr, i, ex_index)
	}

	return arr
}

// use tow arr to merge sort
func MergeSort(arr []int, cmp func(o1, o2 int) bool) []int {

	length := len(arr)
	arr2 := make([]int, length)
	copy(arr2, arr)

	partion(arr, arr2, 0, length-1, cmp)
	return arr
}

// iterative operartion to split sub-question until simple unit's 1 of index
func partion(sortArr, mergeArr []int, prev, rear int, cmp func(o1, o2 int) bool) {
	if rear <= prev {
		return
	}

	mid := prev + (rear-prev)/2

	partion(mergeArr, sortArr, prev, mid, cmp)
	partion(mergeArr, sortArr, mid+1, rear, cmp)

	merge(sortArr, mergeArr, prev, mid, rear, cmp)
}

// use mergeArr to sort order
// use sortArr to keep order
func merge(sortArr, mergeArr []int, prev, mid, rear int, cmp func(o1, o2 int) bool) {
	// rear - prev = 1
	if cmp(mergeArr[mid], mergeArr[mid+1]) {
		for i := prev; i <= rear; i++ {
			sortArr[i] = mergeArr[i]
		}
		return
	}

	i, s, m := prev, prev, mid+1
	for ; i <= rear; i++ {

		if cmp(mergeArr[s], mergeArr[m]) {
			sortArr[i] = mergeArr[s]
			s++
			if mid < s {
				s = rear
			}
		} else {
			sortArr[i] = mergeArr[m]
			m++
			if rear < m {
				m = mid
			}
		}
	}

}
