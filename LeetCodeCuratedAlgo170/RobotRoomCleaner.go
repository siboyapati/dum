package main

type Set map[[2]int]bool

//Input: maze = [[0,0,1,0,0],[0,0,0,0,0],[0,0,0,1,0],[1,1,0,1,1],[0,0,0,0,0]], start = [0,4], destination = [3,2]
//Output: false
//Explanation: There is no way for the ball to stop at the destination. Notice that you can pass through the destination but you cannot stop there.
func hasPath(maze [][]int, start []int, destination []int) bool {

	visited := make(Set)

	dir := [][]int{[]int{0, 1}, []int{0, -1}, []int{1, 0}, []int{-1, 0}}
	return helper12(maze, start, destination, dir, 0, visited)
}

func helper12(maze [][]int, start []int, destination []int, dir [][]int, d int, visited map[[2]int]bool) bool {

	if visited[[2]int{start[0], start[1]}] {
		return false
	}
	visited[[2]int{start[0], start[1]}] = true
	if start[0] == destination[0] && start[1] == destination[1] {
		x, y := start[0]+dir[d][0], start[1]+dir[d][1]
		if maze[x][y] == 1 {
			return true

		}
	}

	for i := 0; i < len(dir); i++ {
		x, y := start[0]+dir[d][0], start[1]+dir[d][1]
		if visited[[2]int{x, y}] {
			continue
		}
		if x < 0 || y < 0 || x >= len(maze) || y >= len(maze[0]) || maze[x][y] == 1 {
			d = (d + 1) % 4
		}
		visited[[2]int{x, y}] = true
		if helper12(maze, []int{x, y}, destination, dir, d, visited) {
			return true
		}
		visited[[2]int{x, y}] = false
	}
	return false
}
