package main

import "fmt"

func main()  {
	var N, a int
	var counter = 0
	fmt.Scan(&N)
	array := []int{}
	for i:=0; i < N; i++{
		fmt.Scan(&a)
		array = append(array,a)
		if array[i] > 0 {
			counter++
		}
	}
	fmt.Print(counter)
}