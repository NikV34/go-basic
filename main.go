package main

import "fmt"

func main()  {
	var N int
	fmt.Scan(&N)
	array := []int{}
	for ; N > 0; {
		array = append(array,N%10)
		N=N/10
	}
	for _,value := range array {
		fmt.Print(value)
	}
}