package convert

import "strings"

//https://leetcode-cn.com/problems/zigzag-conversion/
func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	//转成rune数组
	sRune := []rune(s)
	//数组长度
	runeLength := len(sRune)

	//每个周期的字符数
	cycle := numRows*2 - 2

	//定义一个numRows行的string数组
	res := make([]string, numRows)

	for i := 0; i < runeLength; i++ {
		var mod = i % cycle
		if mod < numRows {
			res[mod] += string(sRune[i])
		} else {
			res[cycle-mod] += string(sRune[i])
		}
	}
	return strings.Join(res, "")
}
