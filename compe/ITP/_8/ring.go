package main

import "fmt"

func main() {
	var p, s string
	var i, j, ret int
	fmt.Scan(&s, &p)

	for {
		if p[i] == s[j] {
			if i == len(p)-1 {
				fmt.Println("Yes")
				break
			}
			i++
			j++
		} else {
			if ret == 1 {
				fmt.Println("No")
				break
			}
			if i != 0 {
				i = 0
			} else {
				j++
			}
		}

		if j == len(s) {
			j = 0
			ret = 1
		}
	}
}
