package main

import "fmt"

func main()  {
	slice := []int{1,2,3,4,5,6,7,8,9}
	min := 1
	for i := 2; i >= len(slice); i++ {
		if i < min {
			i += min
		}
	}
	fmt.Println(min)
}