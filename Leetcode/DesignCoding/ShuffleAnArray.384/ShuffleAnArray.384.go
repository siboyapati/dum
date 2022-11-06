package ShuffleAnArray_384

import (
	"math/rand"
)

//Leetcode
//384. Shuffle an Array
//Given an integer array nums, design an algorithm to randomly shuffle the array. All permutations of the array should be equally likely as a result of the shuffling.
//
//Implement the Solution class:
//
//Solution(int[] nums) Initializes the object with the integer array nums.
//int[] reset() Resets the array to its original configuration and returns it.
//int[] shuffle() Returns a random shuffling of the array.
//
//Input
//["Solution", "shuffle", "reset", "shuffle"]
//[[[1, 2, 3]], [], [], []]
//Output
//[null, [3, 1, 2], [1, 2, 3], [1, 3, 2]]
//
//Explanation
//Solution solution = new Solution([1, 2, 3]);
//solution.shuffle();    // Shuffle the array [1,2,3] and return its result.
//// Any permutation of [1,2,3] must be equally likely to be returned.
//// Example: return [3, 1, 2]
//solution.reset();      // Resets the array back to its original configuration [1,2,3]. Return [1, 2, 3]
//solution.shuffle();    // Returns the random shuffling of array [1,2,3]. Example: return [1, 3, 2]

type Solution struct {
	arr    []int
	newArr []int
}

func Constructor(nums []int) Solution {
	newArr := append([]int{}, nums...)
	return Solution{
		arr:    nums,
		newArr: newArr,
	}
}

func (this *Solution) Reset() []int {
	return this.arr
}

func (this *Solution) Shuffle() []int {
	//tmp := make([]int, len(this.arr))
	//copy(tmp, this.arr)
	tmp := append([]int{}, this.arr...)
	//rand.Seed(time.Now().UnixNano())
	//rand.Shuffle(len(tmp), func(i, j int) {
	//	tmp[i], tmp[j] = tmp[j], tmp[i]
	//})

	// This is much faster than using rand.Shuffle;
	for i := 0; i < len(this.arr); i++ {
		r := rand.Intn(len(this.arr))
		tmp[i], tmp[r] = tmp[r], tmp[i]
	}
	return tmp
}

