package main

import "fmt"

func main()  {
	slice:= []int{1,2,3,4,5,6,7,8}
AccSum := 0
for _,i := range slice{
	AccSum += i
}	

n:= len(slice)
average:= float64(AccSum) / float64(n)

fmt.Println(average)
}

