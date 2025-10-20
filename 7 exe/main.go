package main

import"fmt"

func main() {
	nums:= []int{0, 5, 0, 3, 7, 0, 2}
noZeros:= []int{}

for _, x:= range nums {
    if x!= 0 {
        noZeros = append(noZeros, x)
    }
}

fmt.Println(noZeros) 

}