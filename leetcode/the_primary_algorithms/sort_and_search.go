package the_primary_algorithms

// 以下是 the_primary_algorithms 关于 搜索与排序 的代码实现部分

// 1.Merge 将nums2 归并到nums1 中，并且保持顺序
func Merge(nums1 []int, m int, nums2 []int, n int) {

	// 三指针遍历
	// 倒叙迭代，nums1 关于m索引之前都是有序， nums2 关于n索引之前都是有序
	// 主要是针对nums2合并到nums1中，因此可以判断n是否全部迭代完
	var endIndex int

	for m, n, endIndex = m-1, n-1, m+n-1; -1 < n; endIndex-- {
		if m < 0 || nums1[m] < nums2[n] {
			nums1[endIndex] = nums2[n]
			n--
		} else {
			nums1[endIndex] = nums1[m]
			m--
		}
	}
}

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */
func isBadVersion(version int) bool { return false }

// 2.firstBadVersion 寻找第一个错误版本，已经提供检查检查版本是否出问题的api
func firstBadVersion(n int) int {
	var l, r int = 1, n
	for l < r {
		if isBadVersion(l + (r-l)>>1) {
			r = l + (r-l)>>1
		} else {
			l = l + 1 + (r-l)>>1
		}
	}

	return l
}
