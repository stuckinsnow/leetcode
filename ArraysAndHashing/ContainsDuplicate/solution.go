package main

import "fmt"

func hasDuplicate(nums []int) bool {
	seen := make(map[int]struct{})
	for _, num := range nums {
		seen[num] = struct{}{}
	}
	return len(seen) < len(nums)
}

func main() {
	num := []int{1, 2, 3, 1}

	fmt.Println(hasDuplicate(num))
	fmt.Println(num)
}
