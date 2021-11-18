package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	var cards [4][13]int

	for i := 0; i < n; i++ {
		var s string
		var r int
		fmt.Scan(&s)
		fmt.Scan(&r)
		if s == "S" {
			cards[0][r-1] = 1
		} else if s == "H" {
			cards[1][r-1] = 1
		} else if s == "C" {
			cards[2][r-1] = 1
		} else {
			cards[3][r-1] = 1
		}
	}

	for i := 0; i < 4; i++ {
		for j := 0; j < 13; j++ {
			if cards[i][j] != 1 {
				if i == 0 {
					fmt.Println("S", j+1)
				} else if i == 1 {
					fmt.Println("H", j+1)
				} else if i == 2 {
					fmt.Println("C", j+1)
				} else {
					fmt.Println("D", j+1)
				}
			}
		}
	}
}
