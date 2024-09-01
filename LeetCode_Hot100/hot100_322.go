package main

import "fmt"

func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	cnt := make([]int, amount+1)
	for i := 0; i < len(coins); i++ {
		if coins[i] <= amount {
			cnt[coins[i]] = 1
		}
	}
	for i := 1; i <= amount; i++ {
		if cnt[i] == 1 {
			continue
		} else {
			for j := 0; j < len(coins); j++ {
				if coins[j] < i {
					if cnt[i] == 0 {
						// 保证子数据是可以完成钱数组合的
						if cnt[i-coins[j]] != 0 {
							cnt[i] = cnt[i-coins[j]] + 1
						}
					} else {
						// 保证子数据是可以完成钱数组合的
						if cnt[i-coins[j]] != 0 {
							cnt[i] = min(cnt[i-coins[j]]+1, cnt[i])
						}
					}
				}
			}
		}
	}
	if cnt[amount] == 0 {
		return -1
	} else {
		return cnt[amount]
	}
}

func main() {
	fmt.Println(coinChange([]int{186, 419, 83, 408}, 6249))
}
