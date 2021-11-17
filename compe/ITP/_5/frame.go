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
				if i == 0 || i == H-1 || j == 0 || j == W-1 {
					fmt.Printf("#")
				} else {
					fmt.Printf(".")
				}
			}
			fmt.Println()
		}
		fmt.Println()
	}
}
