package main

import "fmt"

func main() {
	var a, b int
	var op string

	for op != "?" {
		fmt.Scanf("%d %s %d", &a, &op, &b)
		if op == "+" {
			fmt.Println(a + b)
		} else if op == "-" {
			fmt.Println(a - b)
		} else if op == "*" {
			fmt.Println(a * b)
		} else if op == "/" {
			fmt.Println(a / b)
		}
	}
}
