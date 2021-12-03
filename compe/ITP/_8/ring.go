package main

import "fmt"

func main() {
	var p, s string
	var i, j, J, ret int
	fmt.Scan(&s, &p)

	for {
		if p[i] == s[j] {
			if i == len(p)-1 {
				fmt.Println("Yes")
				break
			} else if i == 0 {
				J = j
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
				j = J + 1
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
