package p167_two_sum_sorted_input

func twoSum(numbers []int, target int) []int {
	var i = 0
	var j = len(numbers) - 1

	for i < j {
		sum := numbers[i] + numbers[j]
		if sum == target {
			return []int{i + 1, j + 1}
		}
		if sum > target {
			j--
		} else {
			i++
		}
	}

	return []int{}
}
