package main

import "fmt"

func main() {
	a := make([]int, 26)

	for {
		var c byte
		fmt.Scan(&c)
		if c == nil {
			break
		}

		if 'a' <= c && c <= 'z' {
			a[c-'a']++
		}
		if 'A' <= c && c <= 'Z' {
			a[c-'A']++
		}
	}

	for i := 0; i < 26; i++ {
		fmt.Printf("%c : %d\n", i+'a', a[i])
	}
}
