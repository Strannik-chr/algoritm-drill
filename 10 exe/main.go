package main

import "fmt"

func main()  {
	nums := []int {1,-2,3,4,-5,6,-7,8,9}
	minV := nums[0]
	maxV := nums[0]
	for _, i := range nums {
		if i < minV {
			minV = i
		}else  if i > maxV{
			maxV = i
		}
	} 
	fmt.Println(maxV,minV)
}