package array_and_string

//ReverseString 利用双指针思想进行遍历一个byte数组
// 更节省内存的实现可以尝试单指针遍历半个数组，通过中心对称计算另一边的位置
func ReverseString(s []byte) {
	if len(s) == 0 {
		return
	}

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}

}
