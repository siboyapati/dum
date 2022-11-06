package main

import "fmt"

func main() {
	grid := [][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}}
	fmt.Println(orangesRotting(grid))
}

func orangesRotting(grid [][]int) int {

	var fresh [][]int
	var rotten [][]int
	result := -1
	var direction = [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	var totalFresh int

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				fresh = append(fresh, []int{i, j})
			} else if grid[i][j] == 2 {
				rotten = append(rotten, []int{i, j})
			}
		}
	}


	if len(fresh) == 0{
		return 0
	}
	totalFresh = len(fresh)
	for len(rotten) > 0 {
		result++
		size := len(rotten)
		for i1 := 0; i1 < size; i1++ {
			i, j := rotten[0][0], rotten[0][1]
			rotten = rotten[1:]
			for _, d := range direction {
				x, y := i+d[0], j+d[1]
				if x >= 0 && y >= 0 && x < len(grid) && y < len(grid[0]) {
					if grid[x][y] == 1 {
						totalFresh--
						grid[x][y] = 2
						rotten = append(rotten, []int{x, y})
					}
				}
			}
		}
	}
	if totalFresh > 0 {
		return -1
	}
	return result
}
