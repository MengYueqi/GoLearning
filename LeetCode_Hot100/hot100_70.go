package main

func climbStairs(n int) int {
	if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	} else {
		cnt := make([]int, n)
		cnt[0], cnt[1] = 1, 2
		for i := 2; i < n; i++ {
			cnt[i] = cnt[i-1] + cnt[i-2]
		}
		return cnt[n-1]
	}
}
