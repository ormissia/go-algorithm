package twoSum

//https://leetcode-cn.com/problems/two-sum/
func twoSum(nums []int, target int) []int {
	for i, v := range nums {
		temp := target - v
		for j, k := range nums {
			if temp == k && i != j {
				return []int{i, j}
			}
		}
	}
	return []int{}
}
