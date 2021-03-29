package main

import "fmt"

func main() {
	var workArray [10]uint8
	var n uint8
    var i int
	for i=0; i<10; i++ {
		fmt.Scan(&n)
		workArray[i] = n
	} 
	
	var a,b int
	for i=1; i<4; i++ {
		fmt.Scan(&a)
		fmt.Scan(&b)
		workArray[a],workArray[b] = workArray[b],workArray[a]
	}

	for i=0; i<10; i++ {
		fmt.Print(workArray[i]," ")
	}  
}