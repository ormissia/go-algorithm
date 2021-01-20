package reverse

//https://leetcode-cn.com/problems/reverse-integer
func reverse(x int) int {
	var result int

	for x != 0 {
		if temp := int32(result); (temp*10)/10 != temp {
			return 0
		}
		//输入123
		//循环x  - result
		//    3  -   12
		//    32 -   1
		//    321-   0
		result = result*10 + x%10
		x = x / 10
	}

	return result
}
