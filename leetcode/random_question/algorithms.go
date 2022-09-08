package random_question

import (
	"fmt"
	"sort"
)

// 题库

// 1403.minSubsequence 非递增顺序的最小子序列
// Done at <2022-08-04 周四>
func minSubsequence(nums []int) []int {
	// 倒叙排序
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))

	// 求和
	var sum int
	for _, v := range nums {
		sum += v
	}

	// 找到 遇到 subtract > sum 第一次的情况
	var i, subtract int
	for ; i < len(nums) && sum-nums[i] >= subtract+nums[i]; i++ {
		sum, subtract = sum-nums[i], subtract+nums[i]
	}

	return nums[:i+1]
}

// 899.orderlyQueue 有序队列
// Done at <2022-08-04 周四>
func orderlyQueue(s string, k int) string {
	if 1 < k {
		b := []byte(s)
		sort.Slice(b, func(i, j int) bool {
			return b[i] < b[j]
		})
		return string(b)
	}

	m := len(s)
	minI := m - 1

	for i := 0; i < m-1; i++ {
		if s[i] < s[minI] {
			minI = i
		} else if s[i] == s[minI] {
			p, q := i+1, minI+1

			for p-i < m {
				if s[p%m] == s[q%m] {
					p++
					q++
					continue
				} else if s[p%m] < s[q%m] {
					minI = i
				}

				break
			}
		}
	}

	// b = append(b, b[:minI]...)

	// return string(b[minI:])
	return s[minI:] + s[:minI]
}

// 592.fractionAddition 分数加减运算
// Done at <2022-08-05 周五>

// gcd 针对 rune 类型的 辗转求余法，求出 m, n 最大公倍数
func gcd(m, n rune) rune {
	var rem rune
	for n > 0 {
		rem = m % n
		m = n
		n = rem
	}

	return m
}

func fractionAddition(expression string) string {
	// 添加终结符
	expression = fmt.Sprintf("%v+", expression)

	// 初始化 '+'，提前处理
	if expression[0] != '-' {
		expression = fmt.Sprintf("+%v", expression)
	}

	// ex: b/a +- d/c
	// 分子： b * c +- d * a
	// 分母： a * c
	// mole,deno 为总分母，分子合并
	// 利用 preDeno 做每次数字拼接，遇上 ‘/’ 时，说明当前数字拼接完时分子，更新分子，重置数字拼接条件，为分母数字拼接准备
	var mole, deno rune = 0, 1
	var preMole, preDeno rune = 0, 1

	// multiple 做数字拼接使用， singal 记录上一次符号， ‘-’ ascii 为 45， ‘+’ ascii 为 43 ， 利用中位数 44 - singal 做正负判断
	var multiple, singal rune = 1, '+'

	for _, v := range expression {

		if v == '+' || v == '-' {
			// 计算当前分数
			// fmt.Printf("%v, %v/%v , %v/%v  \n", string(singal), mole, deno, preMole, preDeno)
			mole = preDeno*mole + deno*preMole*(44-singal)
			deno *= preDeno

			// 记录当前 符号，为下一次使用
			singal = v

		} else if v == '/' {
			// 将当前拼接的数字更新到 分子上
			preMole = preDeno

		} else {
			// 数字拼接
			preDeno = preDeno*multiple + (v - '0')
			multiple *= 10
			// 跳过下面重置
			continue
		}

		// case 1: 遇上 '+' or '-' 字符，数字拼接重置, 为新的分子做准备
		// case 2: 遇上 '/' 字符，数字拼接重置条件，为分母做准备
		preDeno = 0
		multiple = 1
	}

	// 特殊处理
	if mole == 0 {
		return "0/1"
	}

	// 求余之前对符号进行处理
	singal = 1
	if mole < 0 {
		singal = -1
		mole *= singal
	}

	// 求最大公倍数，并且除去，求最简分子/分母
	commonFactor := gcd(mole, deno)
	mole, deno = mole/commonFactor, deno/commonFactor

	// 按形式返回
	return fmt.Sprintf("%v/%v", singal*mole, deno)
}

// 4.addOneRow
// Done at <2022-08-05 周五>
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func addSpecifedDepthNode(node *TreeNode, val int, depth int) *TreeNode {
	if node == nil {
		return node
	}

	// 当深度为2时，只需要添加相应子树方向的节点，存放值为val
	if depth == 2 {
		node.Left = &TreeNode{Val: val, Left: node.Left}
		node.Right = &TreeNode{Val: val, Right: node.Right}

		return node
	}

	node.Left = addSpecifedDepthNode(node.Left, val, depth-1)
	node.Right = addSpecifedDepthNode(node.Right, val, depth-1)

	return node

}

func addOneRow(root *TreeNode, val int, depth int) *TreeNode {
	if depth == 1 {
		// 当深度为1时，需要将当前根节点存放在值为val的根节点的左子树中
		return &TreeNode{Val: val, Left: root}
	}

	return addSpecifedDepthNode(root, val, depth)
}

// 1470.shuffle 重新排列数组
// Dont at <2022-08-29 周一>
func shuffle(nums []int, n int) []int {
	ans := make([]int, len(nums))

	for i := 0; i < n; i++ {
		ans[i*2], ans[i*2+1] = nums[i], nums[n+i]
	}

	return ans
}

// 1475.finalPrices商品折扣后的最终价格
func finalPrices(prices []int) []int {
	st := []int{0}
	var p int
	ans := make([]int, len(prices))

	for i := len(prices) - 1; -1 < i; i-- {

		// 利用栈记录可能小于当前值的元素，并且是倒叙压栈，下标为最小下标
		for p > -1 {
			if st[p] <= prices[i] {
				// 找到最近下标小于当前值的元素，则进行折扣计算，并且退出
				ans[i] = prices[i] - st[p]
				break
			}
			st = st[:p]
			p--
		}

		// 压栈，每个可能的元素
		st = append(st, prices[i])
		p++
	}

	return ans
}

// 687.longestUnivaluePath 最长同路径
// Done at <2022-09-02 周五>
func longestUnivaluePath(root *TreeNode) int {

	// 通过dfs ， 自底向上寻找可能的最长单边， l + r 比较最长边可以同样作用于单边最长
	var maxLen int
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		var l, r int
		sl := dfs(node.Left)
		sr := dfs(node.Right)

		if node.Left != nil && node.Left.Val == node.Val {
			l = sl + 1
		}

		if node.Right != nil && node.Right.Val == node.Val {
			r = sr + 1
		}

		// 左右子树边连接情况 以及 单边最长
		if l+r > maxLen {
			maxLen = l + r
		}

		if r > l {
			l = r
		}

		return l
	}

	dfs(root)
	return maxLen
}

// 646.findLongestChain 最长数对链
// Done at <2022-09-03 周六>
// (a, b) < (c, d)
func findLongestChain(pairs [][]int) int {
	if len(pairs) < 2 {
		return 1
	}

	// 选择 后一位进行排序
	// 要挑选最长数对链的第一个数对时，最优的选择是挑选第二个数字最小的，这样能给挑选后续的数对留下更多的空间。
	sort.Slice(pairs, func(i, j int) bool { return pairs[i][1] < pairs[j][1] })

	// 跳过第一个
	count := 1
	end := pairs[0][1]
	for i := 1; i < len(pairs); i++ {
		if end < pairs[i][0] {
			end = pairs[i][1]
			count++
		}
	}
	return count
}
