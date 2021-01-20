package transpose

//https://leetcode-cn.com/problems/transpose-matrix/
func transpose(A [][]int) [][]int {
	if len(A) == 0 {
		return A
	}

	//原矩阵的高
	height := len(A)
	//原矩阵的宽
	width := len(A[0])
	//返回结果矩阵
	result := make([][]int, width)

	//为返回结果的每一行创建切片
	for t := 0; t < width; t++ {
		result[t] = make([]int, height)
	}

	//进行转置
	for i, tempList := range A {
		for j, tempValue := range tempList {
			result[j][i] = tempValue
		}
	}
	return result
}
