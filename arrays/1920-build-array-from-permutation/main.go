package main

import (
	"fmt"
)

// Given a zero-based permutation nums (0-indexed), build an array ans of the same length where ans[i] = nums[nums[i]] for each 0 <= i < nums.length and return it.
// A zero-based permutation nums is an array of distinct integers from 0 to nums.length - 1 (inclusive).

// constraints:
// 1 <= nums.length <= 1000
// 0 <= nums[i] < nums.length

// initial solution

func buildArray(nums []int) []int {
    // define intermediary and final slice
    var idm []int
    var ans []int

    for idx, _ := range nums {
        idm = append(idm, nums[idx])
        ans = append(ans, nums[idm[idx]])
    }
    
    return ans
}

func main() {
	
	input := []int{0,2,1,5,3,4}

	ans := buildArray(input)
	fmt.Println(ans)

}

// solution by shukhratutaboev with O(n) time and O(1) space
// func buildArray2(nums []int) []int {
//	for i, n := range nums {
//		nums[i] += (nums[n] % 1000) * 1000 // take modulo of 1000 since max length is 1000
//	}
//
//	for i, _ := range nums {
//		nums[i] = nums[i] / 1000
//	}
//
//	return nums
//}
