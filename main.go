package main

import "fmt"

func main()  {
	array := []int{}
	var N, a int
	fmt.Scan(&N)
	for i:=0; i < N; i++{
		fmt.Scan(&a)
		if i%2==0 {
			array = append(array, a)
		}
	}
	for _,value := range array {
		fmt.Print(value)
	}
}