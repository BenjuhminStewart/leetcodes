package leetcodes

func HouseRobber(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	if n == 1 {
		return nums[0]
	}

	dp := make([]int, n+1)
	dp[1] = nums[0]
	dp[2] = max(nums[0], nums[1])

	for i := 3; i <= n; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i-1])
	}

	return dp[n]
}
