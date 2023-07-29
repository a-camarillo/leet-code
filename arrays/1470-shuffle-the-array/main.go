package main

import "fmt"

// Given the array nums consisting of 2n elements in the form [x1,x2,...,xn,y1,y2,...,yn].

// Return the array in the form [x1,y1,x2,y2,...,xn,yn].


// initial solution
func shuffle(nums []int, n int) []int {

	shuf := []int{}

	// create two pointers for going through array
	p1 := 0
	p2 := n

	for range nums {
		if p2 < 2*n {
			x := nums[p1]
			y := nums[p2]

			shuf = append(shuf, x)
			shuf = append(shuf, y)	

			p1 += 1
			p2 += 1
		}
	}

	return shuf	
}

func main() {
	shuf := shuffle([]int{2,5,1,3,4,7},3)

	fmt.Println(shuf)
}
