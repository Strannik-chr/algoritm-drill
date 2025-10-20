package main

import "fmt"

func main()  {
	slice := []int {1,2,3,4,5,6,7,8,9}
	NonZero := []int {}
	zeroCount := 0
	for _, i := range slice{
		if i == 0 {
			zeroCount++
		}else {
		NonZero = append(NonZero, i)
	}
	} 

	for i:= 0; i < zeroCount; i++ {
    NonZero = append(NonZero, 0)
}

fmt.Println(NonZero) 
}