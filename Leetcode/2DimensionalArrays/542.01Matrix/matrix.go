package main

/*542. 01 Matrix
Given an m x n binary matrix mat, return the distance of the nearest 0 for each cell.
The distance between two adjacent cells is 1.

Input: mat = [[0,0,0],[0,1,0],[0,0,0]]
Output: [[0,0,0],[0,1,0],[0,0,0]]*/

/*Input: mat = [[0,0,0],[0,1,0],[1,1,1]]
Output: [[0,0,0],[0,1,0],[1,2,1]]*/

func main() {

}

func updateMatrix(mat [][]int) [][]int {
	var q [][]int
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[0]); j++ {
			if mat[i][j] == 0 {
				q = append(q, []int{i, j})
			} else {
				mat[i][j] = -1
			}
		}
	}
	for len(q) > 0 {
		var direction = [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
		x, y := q[0][0], q[0][1]
		q = q[1:]
		for _, d := range direction {
			x1 := x + d[0]
			y1 := y + d[1]
			if x1 >= 0 && y1 >= 0 && x1 < len(mat) && y1 < len(mat[0]) {
				if mat[x1][y1] == -1 {
					mat[x1][y1] = mat[x][y] + 1
					q = append(q, []int{x1, y1})
				}
			}
		}
	}
	return mat
}
