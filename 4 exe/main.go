package main

import "fmt"

func main()  {
	

nums:= []int{5, -2, 7, -3, 0}
negatives:= []int{}

for _, x:= range nums {
    if x < 0 {
        negatives = append(negatives, x)
    }
}

if len(negatives) > 0 {
    fmt.Println(negatives)
} else {
    fmt.Println("отрицательных чисел нет")
}

}