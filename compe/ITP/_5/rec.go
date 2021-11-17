package main

import "fmt"

func main() {
	var H, W int
	for {
		fmt.Scan(&H, &W)
		if H == 0 && W == 0 {
			break
		}
		for i := 0; i < H; i++ {
			for j := 0; j < W; j++ {
				fmt.Printf("#")
			}
			fmt.Println()
		}
		fmt.Println()
	}
}
