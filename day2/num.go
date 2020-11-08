func numDecodings(s string) int {
	dp := make([]int, len(s)+1)
	dp[0] = 1
	for i := 1; i <= len(s); i++ {
		if i > 1 {
			num := 10*int(s[i-2]-'0') + int(s[i-1]-'0')
			if num >= 10 && num <= 26 {
				dp[i] += dp[i-2]
			}
		}

		if int(s[i-1]-'0') != 0 {
			dp[i] += dp[i-1]
		}

		if dp[i] == 0 {
			return 0
		}
	}

	return dp[len(s)]
}
