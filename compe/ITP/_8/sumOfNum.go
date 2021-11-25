package main

import (
	"fmt"
)

func main() {
	for {
		var s string
		var sum int
		fmt.Scan(&s)

		if s == "0" {
			break
		}

		c := []byte(s)

		for _, v := range c {
			sum += int(v - '0')
		}
		fmt.Println(sum)
	}
}
