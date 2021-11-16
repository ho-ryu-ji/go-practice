package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	x, y, z := 3, 4, 5
	fmt.Printf("max(%d, %d) = %d\n", x, y, max(x, y))
	fmt.Printf("max(%d, %d) = %d\n", x, z, max(x, z))

}
