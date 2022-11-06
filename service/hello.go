package service

func Add(a, b int) int {
	return a + b
}

func twoSum(nums []int, target int) []int {

	//m := make(map[int]int)

	var m = map[int]int{}
	//var m []int
	for i, v := range nums {
		tar := target - v
		if index, ok := m[tar]; ok {
			return []int{index, i}
		}
		m[v] = i
	}
	return []int{-1, -1}
}
