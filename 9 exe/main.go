package main

import "fmt"

func main() {
	nums := []int{1, 3, 5, 7, 9}
	sorted := true

	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			sorted = false
			break
		}
	}

	fmt.Println(sorted)
}
