package main

import "fmt"

func main() {
	var s string
	var m, h int

	for {
		q := make([]byte, 0)
		fmt.Scan(&s, &m)
		if s == "-" {
			break
		}

		for i := 0; i < len(s); i++ {
			q = append(q, s[i])
		}

		for i := 0; i < m; i++ {
			fmt.Scan(&h)
			for i := 0; i < h; i++ {
				q = append(q, q[i])
			}
			q = q[h:]
		}

		for i := 0; i < len(q); i++ {
			fmt.Printf("%c", q[i])
		}
		fmt.Println()

	}
}
