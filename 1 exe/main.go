package main

import (
	
	"fmt"
)

func main()  {

	slice:= []int{1,2,3,4,5,6,7,8,9}

	acc := 0

	for _, i := range slice {
		acc += i
	}
fmt.Println(acc)
}

