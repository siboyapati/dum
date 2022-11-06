package GraphTheory

func shortestPathBinaryMatrix(grid [][]int) int {
	if grid == nil || grid[0][0] == 1 {
		return -1
	}
	q := [][]int{{0, 0, 1}}
	dirs := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}, {1, 1}, {1, -1}, {-1, 1}, {-1, -1}}
	for len(q) != 0 {
		size := len(q)
		for i := 0; i < size; i++ {
			temp := q[0]
			q = q[1:]
			if temp[0] == len(grid)-1 && temp[1] == len(grid[0])-1 {
				return temp[2]
			}
			for _, dir := range dirs {
				x, y := temp[0]+dir[0], temp[1]+dir[1]
				if x < 0 || y < 0 || x >= len(grid) || y >= len(grid[0]) {
					continue
				}
				if grid[x][y] == 1 {
					continue
				}
				q = append(q, []int{x, y, temp[2] + 1})
				grid[x][y] = 1
			}
		}
	}
	return -1
}
