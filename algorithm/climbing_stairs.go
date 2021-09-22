package algorithm

// 1. recursion
// 2. memory recursion
// 3. fibonacci
// 4. dp
func climbStairs(n int) int {
	var mem = make([]int, n+1)
	return climb_stairs(mem, n)
}

func climb_stairs(mem []int, n int) int {
	// terminator
	if n <= 1 {
		return 1
	}
	// process current level logic
	if mem[n] == 0 {
		// drill down
		mem[n] = climb_stairs(mem, n-1) + climb_stairs(mem, n-2)
	}

	return mem[n]

	// reverse status if needed

}

// dp
func climbStairs2(n int) int {
	var dp = make([]int, n+1)
	dp[0] = 1
	dp[1] = 2

	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n-1]
}

func climbStairs3(n int) int {
	if n == 1 {
		return 1
	}
	var f1 = 1
	var f2 = 2

	for i := 2; i < n; i++ {
		var f3 = f1 + f2
		f1 = f2
		f2 = f3
	}
	return f2
}
