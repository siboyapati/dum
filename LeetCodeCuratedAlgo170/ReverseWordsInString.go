package main

import "fmt"

func main() {
	fmt.Println(getFactors(12))
}

func getFactors(n int) [][]int {

	var result [][]int
	var temp []int
	if n == 1 {
		return result
	}

	product := 1
	index := 2
	helper1(&result, temp, n, product, index)

	return result

}

func helper1(result *[][]int, temp []int, n, product, index int) {

	if product > n {
		return
	}

	if product == n {
		newTemp := make([]int, len(temp))
		copy(newTemp, temp)
		*result = append(*result, newTemp)
		return
	}

	for i := index; i <= n/product; i++ {

		if n%i == 0 && i != n {
			temp = append(temp, i)
			helper1(result, temp, n, product*i, i)
			temp = temp[:len(temp)-1]
		}
	}
}
