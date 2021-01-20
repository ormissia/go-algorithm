package getWinner

//https://leetcode-cn.com/problems/find-the-winner-of-an-array-game
func getWinner(arr []int, k int) int {
	N := len(arr)

	//如果k比数组最大长度大，找出数组中最大值即可
	if k >= N {
		return max(arr)
	}
	cnt := 0
	for i := 1; i < N; i++ {
		if arr[0] > arr[i] {
			cnt++
		} else {
			arr[i], arr[0] = arr[0], arr[i]
			cnt = 1
		}
		if cnt == k {
			return arr[0]
		}
	}
	return arr[0]
}

func max(a []int) int {
	ret := 0
	for i := 0; i < len(a); i++ {
		if ret < a[i] {
			ret = a[i]
		}
	}
	return ret
}
