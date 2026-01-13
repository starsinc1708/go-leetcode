package p283_move_zeros

func moveZeroes(nums []int) {
	p1, p2 := 0, 0
	for ; p2 < len(nums); p2++ {
		if nums[p2] != 0 {
			nums[p1], nums[p2] = nums[p2], nums[p1]
			p1++
		}
	}
}
