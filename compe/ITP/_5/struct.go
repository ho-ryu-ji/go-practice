package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		var x int
		x = i
		if x%3 == 0 {
			fmt.Printf(" %d", i)
		} else {
			for {
				if x%10 == 3 {
					fmt.Printf(" %d", i)
					break
				}
				x /= 10
				if x == 0 {
					break
				}
			}
		}
	}
	fmt.Println()
}
