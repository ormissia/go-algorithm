package sortedSquares

//https://leetcode-cn.com/problems/squares-of-a-sorted-array/
func sortedSquares(a []int) []int {
	n := len(a)
	lastNegIndex := -1
	for i := 0; i < n && a[i] < 0; i++ {
		lastNegIndex = i
	}

	ans := make([]int, 0, n)
	for i, j := lastNegIndex, lastNegIndex+1; i >= 0 || j < n; {
		if i < 0 {
			ans = append(ans, a[j]*a[j])
			j++
		} else if j == n {
			ans = append(ans, a[i]*a[i])
			i--
		} else if a[i]*a[i] < a[j]*a[j] {
			ans = append(ans, a[i]*a[i])
			i--
		} else {
			ans = append(ans, a[j]*a[j])
			j++
		}
	}

	return ans
}

// func sortedSquares(a []int) []int {
//     ans := make([]int, len(a))
//     for i, v := range a {
//         ans[i] = v * v
//     }
//     sort.Ints(ans)
//     return ans
// }
