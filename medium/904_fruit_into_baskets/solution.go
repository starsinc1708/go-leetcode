package main

import "fmt"

// two baskets
func totalFruit(fruits []int) int {
	begin := 0
	windowState := make(map[int]int)
	result := 0

	for end := 0; end < len(fruits); end++ {
		windowState[fruits[end]]++
		for len(windowState) > 2 {
			windowState[fruits[begin]]--
			if windowState[fruits[begin]] == 0 {
				delete(windowState, fruits[begin])
			}
			begin++
		}

		result = max(result, end-begin+1)
	}

	return result
}

func main() {
	fmt.Println(totalFruit([]int{1, 2, 3, 2, 2})) // 4
	fmt.Println(totalFruit([]int{0, 1, 2, 2}))    // 3
	fmt.Println(totalFruit([]int{1, 2, 1}))       // 3
}
