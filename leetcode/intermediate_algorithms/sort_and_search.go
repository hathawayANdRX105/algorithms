package intermediate_algorithms

import (
	"math/rand"
	"time"
)

// 以下是 intermediate_algorithms 关于 sort_and_search 的代码实现部分

// 1.sortColors 颜色分类
// 值0，1，2 分别代表一种颜色
func sortColors(nums []int) {

	// 利用三指针排序
	// nums[0, l] 维护 值0
	// nums[l+1, i-1] 维护 值1
	// nums[i, len(nums)-1] 维护 值2
	l, i, r := -1, 0, len(nums)

	for i < r {
		if nums[i] < 1 {
			l++
			nums[l], nums[i] = nums[i], nums[l]
		} else if nums[i] > 1 {
			r--
			nums[r], nums[i] = nums[i], nums[r]
			i-- // 抵消下面的偏移
		}

		i++
	}

}

// 2.topKFrequent 前k个高频元素
// 过程详解： O(n) 统计一次出现数字的个数， O(nlgn) 构建堆取 前k频繁出现的数
// 代码注解不做过多，参考 intro_algorithms 的 cp06 中 关于最大堆 swim, sink 代码实现
func topKFrequent(nums []int, k int) []int {

	unique := make(map[int]int, k)
	for _, v := range nums {
		if _, ok := unique[v]; !ok {
			unique[v] = 0
		}
		unique[v]++

	}

	// 添加并维护堆，swim
	var size int
	for k := range unique {
		nums[size] = k

		i := size
		for 0 < i && unique[nums[(i-1)/2]] < unique[nums[i]] {
			nums[(i-1)/2], nums[i] = nums[i], nums[(i-1)/2]
			i = (i - 1) / 2
		}
		size++

	}

	// 取k个最大堆的统计最大值，sink 并维护堆
	ans := make([]int, k)
	for i := 0; i < k; i++ {
		ans[i] = nums[0]

		size--
		nums[0] = nums[size]

		var j int
		for j*2+1 < size {
			j = j*2 + 1
			if j+1 < size && unique[nums[j]] < unique[nums[j+1]] {
				j++
			}

			// 如果当前节点已经是最大堆结构，则退出
			if unique[nums[j]] < unique[nums[(j-1)/2]] {
				break
			}

			// 否则交换根节点更大的计数，并且维护下一个子树最大堆性质
			nums[(j-1)/2], nums[j] = nums[j], nums[(j-1)/2]

		}

	}

	return ans
}

// swim ...
func swim(i int, nums []int) {
	for 0 < i && nums[(i-1)/2] < nums[i] {
		nums[(i-1)/2], nums[i] = nums[i], nums[(i-1)/2]
		i = (i - 1) / 2
	}
}

// sink ...
func sink(size int, nums []int) {
	var i int
	for i*2+1 < size {
		i = i*2 + 1
		if i+1 < size && nums[i] < nums[i+1] {
			i++
		}

		if nums[i] < nums[(i-1)/2] {
			break
		}

		nums[(i-1)/2], nums[i] = nums[i], nums[(i-1)/2]
	}
}

// 3.findKthLargest 寻找数组中第k大的元素
func findKthLargest(nums []int, k int) int {
	k = len(nums) - k + 1
	for i := 0; i < len(nums); i++ {

		// 维护 length - k + 1 的最大堆， 最大元素为候选第k大元素
		// ex: nums有9个元素，求第4大元素，则有 9-4=5，第4大数之前有5个比它小的元素
		if i < k {
			swim(i, nums)
		} else if nums[0] > nums[i] {
			// 如果出现比当前最大堆最大值 更小的值，说明 堆中最大值不是第k大值
			nums[0] = nums[i]
			sink(k, nums)
		}
	}

	return nums[0]
}

// 4.findPeakElement 寻找峰值O(lgn)
func findPeakElement(nums []int) int {
	if len(nums) < 2 || nums[0] > nums[1] {
		// case 1: nums 只有一个元素
		// case 2: nums 有两个元素以上，并且 第一个元素为峰值
		return 0
	} else if nums[len(nums)-1] > nums[len(nums)-2] {
		// case 3: nums 有两个元素以上，并且 最后一个元素为峰值
		return len(nums) - 1
	}

	l, r := 0, len(nums)-1
	for l < r {
		m := l + (r-l)>>1

		if nums[m] < nums[m+1] {
			// 左边界 优先关注更大元素
			l = m + 1
		} else {
			// 右边界 次优先关注
			r = m
		}

	}

	return l
}

// 5.searchRange1 使用两次二分查询分别寻找左右边界
func searchRange1(nums []int, target int) []int {
	// 第一次二分查询，寻找左边界
	var m int
	l1, r1 := 0, len(nums)-1
	for l1 < r1 {
		// 中间值靠左计算
		m = l1 + (r1-l1)>>1

		// mid < t
		// 左边界逼近，右边界保留结果
		if nums[m] < target {
			l1 = m + 1
		} else {
			r1 = m
		}
	}

	// 判断当前数组大小如果为0 以及 数组不存在目标元素，返回 未找到结果
	if len(nums) < 1 || nums[l1] != target {
		return []int{-1, -1}
	}

	// 第二次二分查询，寻找右边界
	l2 := l1
	r1 = len(nums) - 1
	for l2 < r1 {
		// 中间值靠右计算
		m = l2 + (r1-l2+1)>>1

		// t < mid
		// 左边界保留可能结果，右边界逼近
		if target < nums[m] {
			r1 = m - 1
		} else {
			l2 = m
		}
	}

	return []int{l1, l2}
}

// searchRange2 先用二分查询搜索目标值，找到后根据 中心拓展，或者二次二分查询探索左右边界
func searchRange2(nums []int, target int) []int {
	var m int
	ans := []int{-1, -1}
	l, r := 0, len(nums)-1

	for l <= r {
		m = l + (r-l)>>1

		if nums[m] == target {
			break
		} else if target < nums[m] {
			r = m - 1
		} else {
			l = m + 1
		}
	}

	// 1.中心扩增法
	// 寻找到target后, 以中点拓展边界
	// ans[0], ans[1] = m, m
	// for i := m - 1; l <= i && nums[i] == target; i-- {
	// 	ans[0]--
	// }

	// for j := m + 1; j <= r && nums[j] == target; j++ {
	// 	ans[1]++
	// }

	// 2.做两次二分查询
	// 中值寻找到target，划分两个二分查询左右边界
	p1, p2 := m, m
	// 与searchRange1 探索左边界相似
	for l < p1 {
		m = l + (p1-l)>>1
		if nums[m] == target {
			p1 = m
			continue
		}

		l = m + 1
	}

	// 与searchRange1 探索右边界相似
	for p2 < r {
		m = p2 + (r-p2+1)>>1
		if nums[m] == target {
			p2 = m
			continue
		}

		r = m - 1
	}
	ans[0], ans[1] = l, r

	return ans
}

// partition according the start value in range
func partitionForMerge(arr [][]int, lo, hi int) int {
	// 随机抽取其中某值作为划分值，期望partition算法O(lgn)
	i := rand.Int()%(hi-lo+1) + lo
	arr[i], arr[hi] = arr[hi], arr[i]

	i = lo - 1
	for lo < hi {
		if arr[lo][0] <= arr[hi][0] {
			i++

			arr[i], arr[lo] = arr[lo], arr[i]
		}

		lo++
	}

	//exchange cmpValue with value of index j
	arr[i+1], arr[hi] = arr[hi], arr[i+1]

	return i + 1
}

func quickSortForMerge(intervals [][]int, lo, hi int) {

	// 使用尾递归 快速排序
	for lo < hi {
		q := partitionForMerge(intervals, lo, hi)
		quickSortForMerge(intervals, lo, q-1)

		lo = q + 1
	}

	// quickSortForMerge(intervals, q+1, hi)
}

// 6.merge 合并区间
// 区间乱序, 先排序，后根据条件合并
// array_and_string array.go 中已有实现，这里对快排进行优化，以及使用快慢指针节省空间，速度慢了
func merge(intervals [][]int) [][]int {
	size := len(intervals)
	if size < 2 {
		return intervals
	}

	rand.Seed(time.Now().UnixNano())
	quickSortForMerge(intervals, 0, size-1)

	// 1.利用快慢指针
	// （2）.可以直接创建数组直接存，参考 array_and_string 中 array 的 merge 实现
	l, r := 0, 1
	for ; r < size; r++ {
		// case 3: 两个区间不存在交界，更新候选合并区间
		if intervals[l][1] < intervals[r][0] {
			l++
			intervals[l] = intervals[r]
			continue
		}

		// 需要合并
		if intervals[l][1] < intervals[r][1] && (intervals[l][0] == intervals[r][0] || intervals[l][1] >= intervals[r][0]) {
			// case 1: 同左边界，取最大右边界
			// case 2: 存在交界，合并最大区间
			intervals[l][1] = intervals[r][1]
		}
	}

	return intervals[:l+1]
}

// 7.search 搜索旋转排序数组
func search(nums []int, target int) int {
	l, r := 0, len(nums)-1

	for l <= r {
		m := l + (r-l)>>1

		// fmt.Printf("before %v, %v, %v \n",l, r, m)
		if nums[m] == target {
			// case 1: 寻找中值为目标值
			return m
		} else if nums[0] <= target && nums[m] < nums[0] {
			// case 2: 如果 target 在左边子数组，并且中值落在 右边子数组，右边界靠拢
			r = m - 1
		} else if nums[0] > target && nums[m] >= nums[0] {
			// case 3: 如果 target 在右边子数组，并且中值落在 左边子数组，左边界靠拢
			l = m + 1
		} else if nums[m] < target {
			// case 4: nums[l, r] 已落在升序数组范围内，二分查找
			//   subcase 1： mid < target 说明 target 在右边，左边界靠拢
			l = m + 1
		} else {
			//   subcase 2： 否则 target 在左边，右边界考虑
			r = m - 1
		}

		// fmt.Printf("after %v, %v, %v \n", l, r, m)
	}

	// case 5: 未找到target
	return -1
}

// 8.searchMatrix1 搜索二维矩阵 II
// 假定matrix 为 m行，n列,则 searchMatrix1 算法复杂度为 O(lgm + mlgn)
func searchMatrix1(matrix [][]int, target int) bool {
	var m int
	l, r := 0, len(matrix)-1

	// 二分查询寻找 target 求下底的行值
	for l <= r {
		m = l + (r-l)>>1

		if matrix[m][0] == target {
			return true
		} else if matrix[m][0] < target {
			l = m + 1
			continue
		}

		r = m - 1
	}

	// 二分查询可能包含 target 的每个子数组
	row := r
	for -1 < row {
		l, r = 0, len(matrix[row])-1
		if matrix[row][r] < target {
			break
		}

		for l < r {
			m = l + (r-l)>>1
			if matrix[row][m] == target {
				return true
			} else if matrix[row][m] < target {
				l = m + 1
				continue
			}

			r = m - 1
		}

		row--
	}

	return false
}

// searchMatrix2
// 由于 矩阵 每行的元素从左到右升序排列，且每列的元素从上到下升序排列
// 如果从左上角，右下角开始，有升降的趋势
// 左上角取值开始， 向下升序，向左降序
func searchMatrix2(matrix [][]int, target int) bool {
	x, y := 0, len(matrix[0])-1

	for x < len(matrix) && -1 < y {
		if matrix[x][y] == target {
			// case 1: matrix[x][y]为target
			return true
		} else if matrix[x][y] < target {
			// case 2: matrix[x][y] < target, 说明 target 在 [x+1:row] 行的子数组中
			x++
		} else {
			// case 3: matrix[x][y] > target, 说明 target 在 [0: y-1] 列的子数组中
			y--
		}
	}

	return false
}
