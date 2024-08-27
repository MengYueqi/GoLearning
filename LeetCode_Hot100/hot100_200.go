package main

import "fmt"

func numIslands(grid [][]byte) int {
	num := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' {
				num++
				dfs(grid, i, j)
			}
		}
	}
	return num
}

func dfs(grid [][]byte, x, y int) {
	if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[0]) {
		return
	}
	grid[x][y] = '0'
	if x-1 >= 0 && grid[x-1][y] == '1' {
		dfs(grid, x-1, y)
	}
	if x+1 < len(grid) && grid[x+1][y] == '1' {
		dfs(grid, x+1, y)
	}
	if y-1 >= 0 && grid[x][y-1] == '1' {
		dfs(grid, x, y-1)
	}
	if y+1 < len(grid[x]) && grid[x][y+1] == '1' {
		dfs(grid, x, y+1)
	}
}

func main() {
	// Test case 1: Simple grid with one island
	grid1 := [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}
	fmt.Println("Test Case 1:", numIslands(grid1)) // Expected: 3

	// Test case 2: All water
	grid2 := [][]byte{
		{'0', '0', '0'},
		{'0', '0', '0'},
		{'0', '0', '0'},
	}
	fmt.Println("Test Case 2:", numIslands(grid2)) // Expected: 0

	// Test case 3: All land
	grid3 := [][]byte{
		{'1', '1', '1'},
		{'1', '1', '1'},
		{'1', '1', '1'},
	}
	fmt.Println("Test Case 3:", numIslands(grid3)) // Expected: 1

	// Test case 4: Mixed grid
	grid4 := [][]byte{
		{'1', '0', '0', '1'},
		{'0', '1', '1', '0'},
		{'0', '0', '0', '1'},
		{'1', '0', '1', '0'},
	}
	fmt.Println("Test Case 4:", numIslands(grid4)) // Expected: 5

	// Test case 5: Larger grid with multiple islands
	grid5 := [][]byte{
		{'1', '1', '0', '0', '0', '1', '1'},
		{'1', '0', '0', '0', '1', '0', '0'},
		{'0', '0', '1', '0', '0', '1', '0'},
		{'0', '1', '0', '1', '0', '0', '1'},
		{'1', '0', '1', '0', '1', '1', '0'},
	}
	fmt.Println("Test Case 5:", numIslands(grid5)) // Expected: 9
}
