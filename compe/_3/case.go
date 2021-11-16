package main

import "fmt"

func main() {
	var x, cnt int

	for {
		fmt.Scan(&x)
		cnt++
		if x != 0 {
			fmt.Printf("Case %d: %d\n", cnt, x)
		} else {
			break
		}
	}
}
