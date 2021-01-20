package trap

import "math"

//https://leetcode-cn.com/problems/volume-of-histogram-lcci/
func trap(height []int) int {
	if len(height) <= 2 {
		return 0
	}
	result := 0
	//双指针
	leftMax, rightMax := 0, 0
	i, j := 0, len(height)-1

	//指针分别从两端向中间移动
	//结束条件为两个指针相遇
	for i < j {
		leftMax = int(math.Max(float64(leftMax), float64(height[i])))
		rightMax = int(math.Max(float64(rightMax), float64(height[j])))

		//每次移动都将当前移动过的柱子能容纳的水量加到结果中
		if leftMax < rightMax {
			result += leftMax - height[i]
			//左指针移动
			i++
		} else {
			result += rightMax - height[j]
			//右指针移动
			j--
		}
	}
	return result
}
