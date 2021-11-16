package main

import "fmt"

func main() {
	var W, H, x, y, r int
	fmt.Scan(&W, &H, &x, &y, &r)

	if x < 0 || x+r > W {
		fmt.Println("No")
	} else if y < 0 || y+r > H {
		fmt.Println("No")
	} else {
		fmt.Println("Yes")
	}
}
