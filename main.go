package main

import "fmt"

func main()  {
	var N int
	fmt.Scan(&N)
	fmt.Printf("It is %d hours %d minutes.", N/3600, N%3600/60)
}