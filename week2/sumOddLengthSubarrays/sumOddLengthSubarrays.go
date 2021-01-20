package sumOddLengthSubarrays

import "fmt"

//https://leetcode-cn.com/problems/sum-of-all-odd-length-subarrays/
// func sumOddLengthSubarrays(arr []int) int {
// 	result := 0
// 	//TODO 尝试使用goroutine将不同的任务分发下去，
// 	for i := 1; i <= len(arr); i += 2 {
// 		for j := 0; j+i <= len(arr); j++ {
// 			result += ArraySum(arr[j : j+i])
// 		}
// 	}
// 	return result
// }

//goroutine版
func sumOddLengthSubarrays(arr []int) int {
	result := 0

	inNum := make(chan int, 100)
	defer close(inNum)
	outNum := make(chan int, 100)
	defer close(outNum)

	for i := 0; i < len(arr); i++ {
		go worker(inNum, outNum, arr)
	}

	go func() {
		for i := 1; i <= len(arr); i += 2 {
			inNum <- i
		}
	}()

	//for i := 1; i <= len(arr); i += 2 {
	for i := 1; i <= len(arr); i += 2 {
		res := <-outNum
		//fmt.Println(res)
		result += res
	}

	return result
}

func worker(inNum <-chan int, outNum chan<- int, arr []int) {
	for iInNum := range inNum {
		result := 0
		for j := 0; j+iInNum <= len(arr); j++ {
			result += ArraySum(arr[j : j+iInNum])
		}
		fmt.Println(iInNum, result)
		outNum <- result
	}
}

//计算切片的和
func ArraySum(array []int) (sum int) {
	for _, v := range array {
		sum += v
	}
	return
}
