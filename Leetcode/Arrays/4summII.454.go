package Arrays

//Given four integer arrays nums1, nums2, nums3, and nums4 all of length n, return the number of tuples (i, j, k, l) such that:
//
//0 <= i, j, k, l < n
//nums1[i] + nums2[j] + nums3[k] + nums4[l] == 0

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	var result int
	//m1 := make(map[int]int)
	var m1 map[int]int
	for _, n1 := range nums1 {
		for _, n2 := range nums2 {
			m1[n1+n2]++
		}
	}
	for _, n3 := range nums3 {
		for _, n4 := range nums4 {
			sum := n3+n4
			if val,ok := m1[-sum]; ok {
				result += val
			}

		}
	}
	return result
}
