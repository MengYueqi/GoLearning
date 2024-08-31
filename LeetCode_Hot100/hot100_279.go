package main

import (
	"fmt"
	"math"
)

func numSquares(n int) int {
	cnt := make([]int, n)
	sqrt := int(math.Sqrt(float64(n)))
	if n == 1 || n == sqrt*sqrt {
		return 1
	} else {
		cnt[0] = 1
		for i := 2; i <= n; i++ {
			cntNow := 1000000
			sqrt = int(math.Sqrt(float64(i)))
			if i == sqrt*sqrt {
				cnt[i-1] = 1
				continue
			}
			for j := 1; j <= sqrt; j++ {
				cntNow = min(cntNow, cnt[i-j*j-1]+1)
			}
			cnt[i-1] = cntNow
		}
	}
	return cnt[n-1]
}

func main() {
	fmt.Println(numSquares(5))
}
