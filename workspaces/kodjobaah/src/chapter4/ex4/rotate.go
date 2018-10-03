package main

import "fmt"

func rotate(nums []int, k int) []int {

	toChange := make([]int, len(nums), cap(nums))
	lright := len(nums[k:])
	lleft := len(nums) - lright
	mid := 0
	if lleft > lright {
		mid = lleft
	} else {
		mid = lright
	}
	for i := 0; i < mid; i++ {
		if i < lright {
			toChange[i] = nums[lleft+i]
		}
		if i < lleft {
			toChange[lright+i] = nums[i]
		}
	}
	return toChange

}

func rotate2(nums []int, k int) []int {
	return append(nums[k:], nums[:k]...)
}

func main() {
	s := [...]int{0, 1, 2, 3, 4, 5}
	fmt.Println(s)
	fmt.Println(rotate(s[:], 4))

}
