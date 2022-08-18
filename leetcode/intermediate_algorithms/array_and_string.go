package intermediate_algorithms

import (
	"math"
	"sort"
)

// 以下是 intermediate_algorithms 关于 array_and_string 的代码实现部分

// 1.threeSum 三数之和
func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}

	sort.Ints(nums)
	var ans [][]int

	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			return ans
		}
		if 0 < i && nums[i-1] == nums[i] {
			continue
		}

		l, r := i+1, len(nums)-1

		for l < r {
			if -nums[i] == nums[l]+nums[r] {
				ans = append(ans, []int{nums[i], nums[l], nums[r]})

				for l < r && nums[l+1] == nums[l] {
					l++
				}

				for l < r && nums[r-1] == nums[r] {
					r--
				}

				l++
				r--
			} else if -nums[i] < nums[l]+nums[r] {
				r--
			} else {
				l++
			}
		}

	}

	return ans
}

// 2.setZeroes 要求常数空间解决, 进行矩阵置零
// 将第一行，第一列延迟置零处理，作为记录空间进行检查每个元素是否为零，如果为零，则置该行该列的第一项为零
func setZeroes(matrix [][]int) {
	var zeroMarker byte

	for i := 0; i < len(matrix); i++ {
		if matrix[i][0] == 0 {
			zeroMarker = 1 << 1
			break
		}
	}

	for j := 0; j < len(matrix[0]); j++ {
		if matrix[0][j] == 0 {
			zeroMarker++
			break
		}
	}

	for r := 1; r < len(matrix); r++ {
		for c := 1; c < len(matrix[0]); c++ {
			if matrix[r][c] == 0 {
				matrix[r][0] = 0
				matrix[0][c] = 0
			}
		}
	}

	for r := 1; r < len(matrix); r++ {
		for c := 1; c < len(matrix[0]); c++ {
			if matrix[r][0] == 0 || matrix[0][c] == 0 {
				matrix[r][c] = 0
			}
		}
	}

	if zeroMarker&1 != 0 {
		for j := 0; j < len(matrix[0]); j++ {
			matrix[0][j] = 0
		}
	}

	if (zeroMarker>>1)&1 != 0 {
		for i := 0; i < len(matrix); i++ {
			matrix[i][0] = 0
		}
	}

}

// 3.groupAnagrams 字母异位词分组, 将相同字母组合的单词分组
func groupAnagrams(strs []string) [][]string {
	find := map[string]int{}

	var ans [][]string
	var lastI int

	for _, word := range strs {
		// 词分量进行 map 查询
		letterVector := make([]byte, 26)

		for _, cv := range word {
			letterVector[cv-'a']++
		}

		if pos, ok := find[string(letterVector)]; ok {
			ans[pos] = append(ans[pos], word)
			continue
		}

		ans = append(ans, []string{word})
		find[string(letterVector)] = lastI
		lastI++
	}

	return ans
}

// 4.lengthOfLongestSubstring 无重复字符的最长字串
func lengthOfLongestSubstring(s string) int {
	if len(s) < 1 {
		return 0
	}

	find := map[byte]int{}

	// l,r :指向 s[l, r] 不含重复字符的子串左边界和右边界
	var l, r, maxLen int
	for ; r < len(s); r++ {
		// 如果s[r] 出现重复，收缩边界到刚好不含当前字符的边界
		// 可能出现的重复字符的上一个索引小于当前 l，则不更新
		if pos, ok := find[s[r]]; ok && pos+1 > l {
			// 需要跳跃时再检查，减少检查次数
			if r-l > maxLen {
				maxLen = r - l
			}

			l = pos + 1
		}

		// 记录新的无重复字符索引
		find[s[r]] = r

	}

	// 可能最后的无重复字符子串没有检查
	if r-l > maxLen {
		maxLen = r - l
	}

	return maxLen
}

// 5.longestPalindrome 在 array_and_string 的 string 中实现

// "abc" -> "^#a#b#c#&"
func preProcess(s string) []byte {
	T := make([]byte, (len(s)+1)*2+1)

	for i := 1; i < len(s)+1; i++ {
		T[i*2] = s[i-1]
		T[i*2+1] = '#'
	}

	T[0] = '^'
	T[1] = '#'
	T[len(T)-1] = '$'

	return T
}

func longestPalindrome(s string) string {
	if len(s) < 1 {
		return ""
	}

	T := preProcess(s)
	radius := make([]int, len(T))

	var C, R int
	var startCenter, maxLen int
	// "^#a#b#a#&" 从 T[1,len(T)-1] 检查回文
	for i := 1; i < len(T)-1; i++ {
		// 利用已有回文加速
		if R > i {
			// 对称点 iMirror = C - (i - C)
			iMirror := 2*C - i

			if R-i > radius[iMirror] {
				radius[i] = radius[iMirror]
			} else {
				radius[i] = R - i
			}
		} else {
			radius[i] = 0
		}

		// 以 i 为中心， radius[i]为半径进行中心拓展匹配回文
		for T[i-radius[i]-1] == T[i+radius[i]+1] {
			radius[i]++
		}

		// 记录最长回文中心以及半径长度
		if radius[i] > maxLen {
			startCenter = i
			maxLen = radius[i]
		}

		// 更新最远右边界回文
		if R < i+radius[i] {
			C = i
			R = i + radius[i]
		}

	}

	start := (startCenter - maxLen) / 2
	return s[start : start+maxLen]
}

// 6.increasingTriplet 寻找递增的三元子序列
func increasingTriplet(nums []int) bool {
	if len(nums) < 3 {
		return false
	}

	// nums[0] 为最小值， mid 为第二小值， 作为减少空间浪费
	mid := math.MaxInt

	// 从num[1:]开始遍历
	for _, v := range nums[1:] {

		// *贪心思路*，让返回条件提前出现，能提速
		if v > mid {
			// 比第二小值大，说明以及找到递增的 三元子序列数组
			return true
		} else if v > nums[0] {
			// 如果等于或小于第二小值，换言之大于最小值， 则更新第二小值
			mid = v
		} else {
			// 小于最小值时，更新最小值
			nums[0] = v
		}

	}

	return false
}
