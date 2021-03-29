package main

import "fmt"

func main() {
	var workArray [10]int
	var n,i int
	for i=0; i<10; i++ {
		fmt.Scan(&n)
		workArray[i] = n
	} 
	
	var a,b,c int
	for i=1; i<4; i++ {
		fmt.Scan(&a)
		fmt.Scan(&b)
		c = workArray[a]
		workArray[a] = workArray[b]
		workArray[b] = c
	}

	for i=0; i<10; i++ {
		fmt.Print(workArray[i]," ")
	}  
}