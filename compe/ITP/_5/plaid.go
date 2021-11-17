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
				if i%2 == 0 && j%2 == 0 || i%2 != 0 && j%2 != 0 {
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
