package graphs

// Pattern: multi-source BFS from all zero cells.
// Invariant: the first distance assigned to a cell is its nearest-zero distance.
// Complexity: O(rows * cols) time, O(rows * cols) space.
// Interview line: start from all sources at once instead of running BFS from every one cell.
func UpdateMatrix(mat [][]int) [][]int {
	rows := len(mat)
	if rows == 0 {
		return mat
	}
	cols := len(mat[0])

	type point struct{ r, c int }
	var q []point

	res := make([][]int, rows)
	for r := 0; r < rows; r++ {
		res[r] = make([]int, cols)
		for c := 0; c < cols; c++ {
			if mat[r][c] == 0 {
				res[r][c] = 0
				q = append(q, point{r, c})
			} else {
				res[r][c] = 1e9 // unvisited sentinel
			}
		}
	}

	dirs := []point{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	for len(q) > 0 {
		curr := q[0]
		q = q[1:]

		for _, d := range dirs {
			nr, nc := curr.r+d.r, curr.c+d.c
			if nr >= 0 && nr < rows && nc >= 0 && nc < cols {
				if res[curr.r][curr.c]+1 < res[nr][nc] {
					res[nr][nc] = res[curr.r][curr.c] + 1
					q = append(q, point{nr, nc})
				}
			}
		}
	}

	return res
}

// Pattern: multi-source BFS with elapsed time.
// Invariant: queue spreads rot one minute at a time from all initially rotten oranges.
// Complexity: O(rows * cols) time, O(rows * cols) space.
// Interview line: count fresh oranges up front so the final answer is easy to validate.
func OrangesRotting(grid [][]int) int {
	rows := len(grid)
	if rows == 0 {
		return 0
	}
	cols := len(grid[0])

	type point struct{ r, c int }
	var q []point
	fresh := 0

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == 2 {
				q = append(q, point{r, c})
			} else if grid[r][c] == 1 {
				fresh++
			}
		}
	}

	if fresh == 0 {
		return 0
	}

	minutes := 0
	dirs := []point{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	for len(q) > 0 && fresh > 0 {
		size := len(q)
		for i := 0; i < size; i++ {
			curr := q[0]
			q = q[1:]

			for _, d := range dirs {
				nr, nc := curr.r+d.r, curr.c+d.c
				if nr >= 0 && nr < rows && nc >= 0 && nc < cols && grid[nr][nc] == 1 {
					grid[nr][nc] = 2
					fresh--
					q = append(q, point{nr, nc})
				}
			}
		}
		minutes++
	}

	if fresh == 0 {
		return minutes
	}
	return -1
}
