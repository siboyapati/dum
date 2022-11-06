package main

import (
	"fmt"
	"sort"
)

func main() {
	t := [][]string{{"MUC", "LHR"}, {"JFK", "MUC"}, {"SFO", "SJC"}, {"LHR", "SFO"}}
	fmt.Println(findItinerary(t))
}

//Output: ["JFK","MUC","LHR","SFO","SJC"]
//Input: tickets = [["MUC","LHR"],["JFK","MUC"],["SFO","SJC"],["LHR","SFO"]]
func findItinerary(tickets [][]string) []string {

	adjMap := make(map[string][]string)
	//var result []string
	var q [][]string
	for _, ticket := range tickets {
		adjMap[ticket[0]] = append(adjMap[ticket[0]], ticket[1])
	}
	for _, val := range adjMap {
		sort.Slice(val, func(i, j int) bool {
			return val[i] < val[j]
		})
	}
	fmt.Println(adjMap)
	q = append(q, []string{"JFK"})

	var temp []string
	for len(q) != 0 {

		temp = q[0]
		q = q[1:]
		lastIndex := temp[len(temp)-1]

		if val, ok := adjMap[lastIndex]; ok {
			if len(val) == 0 {
				break
			}
			temp = append(temp, val[0])
			adjMap[lastIndex] = val[1:]
		} else {
			break
		}
		q = append(q, temp)
	}
	fmt.Println(temp)
	return temp
}
