package main

import "fmt"

func main() {
	var x, y int
	for {
		fmt.Scan(&x, &y)

		if x == 0 && y == 0 {
			break
		} else {
			if x > y {
				var tmp int
				tmp = x
				x = y
				y = tmp
			}
			fmt.Println(x, y)
		}
	}
}
