package longestPalindrome

//https://leetcode-cn.com/problems/longest-palindromic-substring/
func longestPalindrome(s string) string {
	index := 0
	size := len(s)
	maxLen := 0
	result := ""

	for index < size {
		pointerL := index
		pointerR := index

		//中心扩散，先找出和中心不相等的左、右索引，这样可以忽略奇、偶长度对称性的差异
		for pointerL >= 0 && s[pointerL] == s[index] {
			pointerL--
		}
		for pointerR < size && s[pointerR] == s[index] {
			pointerR++
		}

		//下一次搜索直接将索引移动到上一次的初始右索引，节省重复计算
		nextPoint := pointerR
		for pointerL >= 0 && pointerR < size && s[pointerR] == s[pointerL] {
			pointerL--
			pointerR++
		}

		tempMaxLenth := pointerR - pointerL - 1
		//如果此次for index < size{}循环中最大长度比上一次大，则替换上一次存储的结果字符串且更新最大长度
		if tempMaxLenth > maxLen {
			result = s[pointerL+1 : pointerR]
			maxLen = tempMaxLenth
		}
		index = nextPoint
	}

	return result
}
