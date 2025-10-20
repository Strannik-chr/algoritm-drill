package main

import "fmt"

func main()  {
	numbers:= []int{1, 5, 8, 12, 20}
a, b:= 5, 15
inRange:= []int{}

for _, x:= range numbers {
    if a <= x && x <= b {
        inRange = append(inRange, x)
    }
}

fmt.Println(inRange) 
}