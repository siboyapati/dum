package main

import "fmt"

type SnakeGame struct {
	food   [][]int
	snake  [][]int
	width  int
	height int
}

func Constructor(width int, height int, food [][]int) SnakeGame {
	return SnakeGame{
		width:  width - 1,
		height: height - 1,
		food:   food,
		snake:  [][]int{[]int{0, 0}},
	}
}

func (this *SnakeGame) Move(direction string) int {

	path := map[string][][]int{
		"U": [][]int{[]int{-1, 0}},
		"D": [][]int{[]int{1, 0}},
		"L": [][]int{[]int{0, -1}},
		"R": [][]int{[]int{0, 1}},
	}

	head := this.snake[len(this.snake)-1]
	nextX := head[0] + path[direction][0][0]
	nextY := head[1] + path[direction][0][1]
	fmt.Println("nextx,y", nextX, nextY)

	if nextX < 0 || nextX > this.width || nextY < 0 || nextY > this.height {
		return -1
	}
	foodCord := this.food[0]

	if foodCord[0] == nextX && foodCord[1] == nextY {
		this.food = this.food[1:]
		this.snake = append(this.snake, []int{nextX, nextY})
	} else {
		this.snake = this.snake[1:]
		this.snake = append(this.snake, []int{nextX, nextY})

		for i := 0; i < len(this.snake)-1; i++ {
			cur := this.snake[i]
			if cur[0] == nextX && cur[1] == nextY {
				return -1
			}
		}
	}
	return len(this.snake) - 1
}
