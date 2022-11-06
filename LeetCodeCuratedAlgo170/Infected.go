package main

import "fmt"

//garbage = ["MMM","PGM","GP"], travel = [3,10]
func main() {

	//["tars","rats","arts","star"]
	fmt.Println(numSimilarGroups([]string{"tars", "rats", "arts", "star"}))
}

func numSimilarGroups(strs []string) int {
	visited := make(map[string]bool)
	var result int
	for _, str := range strs {
		if !visited[str] {
			dfs(str, visited, strs)
			result++
		}
	}
	return result
}

func dfs(str string, visited map[string]bool, strs []string) {

	if visited[str] {
		return
	}
	visited[str] = true
	for _, str1 := range strs {
		if !visited[str1] && isSim(str, str1) {
			dfs(str1, visited, strs)
		}
	}
}

func isSim(str1, str2 string) bool {
	var count int
	if len(str1) != len(str2) {
		return false
	}

	for i := 0; i < len(str1); i++ {
		if str1[i] != str2[i] {
			count++
		}

		if count > 2 {
			return false
		}
	}
	return count <= 2
}
