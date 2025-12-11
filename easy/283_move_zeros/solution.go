package main

import "fmt"

func moveZeroes(nums []int) {
	p1, p2 := 0, 0
	for ; p2 < len(nums); p2++ {
		if nums[p2] != 0 {
			nums[p1], nums[p2] = nums[p2], nums[p1]
			p1++
		}
	}
}

func main() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	fmt.Println(nums)
}
