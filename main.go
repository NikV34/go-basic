package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)
	var arr []int
	for i:=0;i<=N-1;i++ {
		var n int
		fmt.Scan(&n)
		arr = append(arr, n)
	}
	fmt.Print(arr[3])
}