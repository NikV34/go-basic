package main

import "fmt"

func main() {
	var i, sq float64
	fmt.Scan(&i)
	sq = i * i
	switch {
	case i <= 0:
		fmt.Printf("число %2.2f не подходит",i)
	
	case sq > 10000:
		fmt.Printf("%e",i)

	default:
		fmt.Printf("%4f", sq)	
	}
}