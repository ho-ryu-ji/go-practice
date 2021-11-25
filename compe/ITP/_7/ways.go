package main

import "fmt"

func main() {
	for {
		var n, x int
		fmt.Scan(&n, &x)
		var cnt int

		if n == 0 && x == 0 {
			break
		}

		for i := 1; i <= n; i++ {
			for j := i + 1; j <= n; j++ {
				for k := j + 1; k <= n; k++ {
					if i+j+k == x {
						cnt++
					}
				}
			}
		}
		fmt.Println(cnt)
	}
}
