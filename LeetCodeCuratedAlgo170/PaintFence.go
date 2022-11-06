package main

import "fmt"

//Input: n = 1, k = 1
//Output: 1

//Input: n = 7, k = 2
//Output: 42
func numWays(n int, k int) int {

	dp := make([]int, n+1)
	dp[1] = k

	for i := 2; i <= n; i++ {

		if i%3 == 0 {
			dp[i] = dp[i-1] * (k - 1)
		} else {
			dp[i] = dp[i-1] * k

		}
	}
	return dp[n]
}

func main() {
	fmt.Println(numWays2(3, 2))
	//fmt.Println(numWays(7, 2))
}

func numWays2(n int, k int) int {
	if n == 1 {
		return k
	}
	if n == 2 {
		return k * k
	}

	totalWays := make([]int, n+1)
	totalWays[1] = k
	totalWays[2] = k * k
	for i := 3; i <= n; i++ {
		totalWays[i] = (k - 1) * (totalWays[i-1] + totalWays[i-2])
	}

	return totalWays[n]
}
