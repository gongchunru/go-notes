package main

func main() {
	println(longestPalindrome("babad"))
}
func longestPalindrome(s string) string {

	strLen := len(s)
	dp := make([][]bool, strLen)
	maxStart, maxEnd := 0, 0

	for i := 0; i < strLen; i++ {
		dp[i] = make([]bool, strLen)
		dp[i][i] = true
	}

	for r := 1; r < strLen; r++ {
		for l := 0; l < r; l++ {
			if s[l] == s[r] && (r-l <= 2 || dp[l+1][r-1]) {
				dp[l][r] = true
				if r-l > maxEnd-maxStart {
					maxStart = l
					maxEnd = r
				}
			}
		}
	}
	return s[maxStart : maxEnd+1]
}
