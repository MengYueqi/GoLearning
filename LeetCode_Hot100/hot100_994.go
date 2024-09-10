package main

type point struct {
	x, y int
}

func orangesRotting(grid [][]int) int {
	// 没有新鲜橘子的情况
	flagNew := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				flagNew = 1
			}
		}
	}
	if flagNew == 0 {
		return 0
	}

	queue := make([]point, 0)
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 2 {
				queue = append(queue, point{x: i, y: j})
			}
		}
	}
	cnt := 0
	for len(queue) != 0 {
		cnt += 1
		length := len(queue)
		for i := 0; i < length; i++ {
			x, y := queue[0].x, queue[0].y
			queue = queue[1:]
			if x-1 >= 0 && grid[x-1][y] == 1 {
				grid[x-1][y] = 2
				queue = append(queue, point{x: x - 1, y: y})
			}
			if y+1 < len(grid[0]) && grid[x][y+1] == 1 {
				grid[x][y+1] = 2
				queue = append(queue, point{x: x, y: y + 1})
			}
			if y-1 >= 0 && grid[x][y-1] == 1 {
				grid[x][y-1] = 2
				queue = append(queue, point{x: x, y: y - 1})
			}
			if x+1 < len(grid) && grid[x+1][y] == 1 {
				grid[x+1][y] = 2
				queue = append(queue, point{x: x + 1, y: y})
			}
		}
	}

	flagNew = 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				flagNew = 1
			}
		}
	}
	if flagNew == 0 {
		return cnt - 1
	} else {
		return -1
	}
}

func main() {
	orangesRotting([][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}})
}
