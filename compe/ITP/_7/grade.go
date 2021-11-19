package main

import "fmt"

func main() {
	for {
		var m, f, r int
		fmt.Scan(&m, &f, &r)

		if m == -1 && f == -1 && r == -1 {
			break
		} else if m == -1 || f == -1 || m+f < 30 {
			fmt.Println("F")
		} else if m+f >= 30 && m+f < 50 {
			if r >= 50 {
				fmt.Println("C")
			} else {
				fmt.Println("D")
			}
		} else if m+f >= 50 && m+f < 65 {
			fmt.Println("C")
		} else if m+f >= 65 && m+f < 80 {
			fmt.Println("B")
		} else if m+f >= 80 {
			fmt.Println("A")
		}
	}
}
