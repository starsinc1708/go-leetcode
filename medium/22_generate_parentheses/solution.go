package p22_generate_parentheses

// https://leetcode.com/problems/generate-parentheses
// Time complexity: O(4^N/Sqrt(N))
func generateParenthesis(n int) []string {
	var ans []string
	backTrack(&ans, "", 0, 0, n)
	return ans
}

func backTrack(result *[]string, currStr string, openCount int, closeCount int, n int) {
	if len(currStr) == n*2 {
		*result = append(*result, currStr)
		return
	}

	if openCount < n {
		backTrack(result, currStr+"(", openCount+1, closeCount, n)
	}
	if closeCount < openCount {
		backTrack(result, currStr+")", openCount, closeCount+1, n)
	}

}
