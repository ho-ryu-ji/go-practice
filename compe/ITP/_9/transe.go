package main

import (
	"fmt"
)

func print(a int, b int, s []byte) {
	for i := a; i <= b; i++ {
		fmt.Printf("%c", s[i])
	}
	fmt.Println()
}

func reverse(a int, b int, s []byte) {
	for i, j := a, b; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func replace(a int, b int, p string, s []byte) {
	var j int
	for i := a; i <= b; i++ {
		s[i] = p[j]
		j++
	}
}

func main() {
	var q int
	var str string
	s := make([]byte, 0)

	fmt.Scan(&str, &q)
	for i := 0; i < len(str); i++ {
		s = append(s, str[i])
	}

	for i := 0; i < q; i++ {
		var a, b int
		var p, req string
		fmt.Scan(&req)
		if req == "print" {
			fmt.Scan(&a, &b)
			print(a, b, s)
		} else if req == "reverse" {
			fmt.Scan(&a, &b)
			reverse(a, b, s)
		} else {
			fmt.Scan(&a, &b, &p)
			replace(a, b, p, s)
		}
	}
}
