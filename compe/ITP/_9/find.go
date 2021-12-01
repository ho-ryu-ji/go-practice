package main

import (
	"fmt"
	"strings"
)

func main() {
	var cnt int
	var W, s string
	fmt.Scan(&W)
	W = strings.ToLower(W)

	for {
		fmt.Scan(&s)

		if s == "END_OF_TEXT" {
			fmt.Println(cnt)
			break
		}

		s = strings.ToLower(s)

		if s == W {
			cnt++
		}
	}
}
