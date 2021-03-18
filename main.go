package main

import "fmt"

func main() {
	// var i, m int
	// fmt.Scan(&i)
	// counter := 0
	// for ;i!=0; {
	// 	switch {
	// 	case i>m:
	// 		m = i
	// 		counter = 1
	// 		case i==m: 
	// 		counter++
	// 	}
	// 	fmt.Scan(&i)
	// }
	// fmt.Println(counter)

	// var n int
	// for fmt.Scan(&n); n<=100; {
	// 	if n > 9 {
	// 		fmt.Println(n)
	// 	}
	// 	fmt.Scan(&n)
	// }

	// var x, p, y int
	// counter :=0
	// fmt.Scan(&x, &p, &y)
	// for ; x < y; {
	// 	x = x * (100+p) / 100
	// 	fmt.Println(x)
	// 	counter++
	// }
	// fmt.Println(counter)

	var x,y,copyX, copyY int
	fmt.Scan(&x,&y)
	reverseX :=0
	copyX = x
	copyY = y
	for ;copyX > 0; copyX=copyX/10{
		reverseX = reverseX *10 + copyX%10
	}
	x= reverseX
	for ;x>0; x=x/10{
		r:=x%10
		if r > 0 {
			for copyY = y;copyY>0;copyY=copyY/10 {
				if copyY%10 == r {
					fmt.Print(r, " ")
					break
				}
			}
		}
	}
}