package intermediate_algorithms

import (
	"bytes"
	"math/rand"
	"strconv"
	"strings"
)

// 以下是 intermediate_algorithms 关于 design_problem 的代码实现

// 1.二叉树的序列化与反序列化
/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var buffer bytes.Buffer

type Codec struct {
	index int      // 遍历反序列使用
	data  []string // 反序列字符串数组
}

func Constructor() Codec {
	return Codec{index: 0, data: nil}
}

func (this *Codec) preorderTraversalSerialize(node *TreeNode) {
	if node == nil {
		buffer.WriteByte('e')
		buffer.WriteByte(',')
		return
	}

	buffer.WriteString(strconv.Itoa(node.Val))
	buffer.WriteByte(',')

	this.preorderTraversalSerialize(node.Left)
	this.preorderTraversalSerialize(node.Right)
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}

	this.preorderTraversalSerialize(root)

	defer buffer.Reset()
	return buffer.String()[:buffer.Len()-1]
}

// preorderTraversalDeserialize ...
func (this *Codec) preorderTraversalDeserialize() *TreeNode {
	if this.index >= len(this.data) || this.data[this.index] == "e" {
		this.index++
		return nil
	}

	// 使用atoi 加快速度
	val, _ := strconv.Atoi(this.data[this.index])
	node := &TreeNode{Val: val}
	this.index++

	node.Left = this.preorderTraversalDeserialize()
	node.Right = this.preorderTraversalDeserialize()

	return node
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	// 没有元素 [ ], 则返回 nil
	if len(data) < 3 {
		return nil
	}

	// split data format: [ 1 2 3 null null 4 5 ]
	// 包括 [ ]
	splitData := strings.Split(data, ",")

	this.data = splitData

	return this.preorderTraversalDeserialize()
}

// 2.常数时间插入，删除和获取随机元素
/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */

type RandomizedSet struct {
	hash map[int]int // unique set
	data []int       // random set
}

func Constructor2() RandomizedSet {

	return RandomizedSet{hash: make(map[int]int, 2e5), data: make([]int, 0, 2e5)}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.hash[val]; ok {
		return false
	}

	this.data = append(this.data, val)
	this.hash[val] = len(this.data) - 1

	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	pos, ok := this.hash[val]
	if !ok {
		return false
	}

	// 将删除的元素与最后元素交换
	this.data[pos] = this.data[len(this.data)-1]

	// 更新 最后元素的索引
	this.hash[this.data[pos]] = pos

	// 删除记录，以及调整切片长度，使其可以append
	delete(this.hash, val)
	// 调整需要在更新索引之后，防止删除的元素就是第一个的情况，切片提前收缩，导致不能让hashmap删除记录
	this.data = this.data[:len(this.data)-1]

	return true
}

// 题提示会在至少有一个元素的前提下获取随机元素，不用判断
func (this *RandomizedSet) GetRandom() int {
	randI := rand.Int() % len(this.data)
	return this.data[randI]
}
