package main

import "fmt"

func main() {
	for {
		var n, x int
		fmt.Scan(&n, &x)
		var cnt int

		I := make([]int, n-1)
		J := make([]int, n-2)
		K := make([]int, n-3)

		if n == 0 && x == 0 {
			break
		}
		for i := 1; i <= n; i++ {
			for j := 2; j <= n; j++ {
				for k := 3; k <= n; k++ {
					if i == j || i == k || j == k {
						continue
					} else if i+j+k == x {
						for i := 1; i <= n; i++ {
							if i != I[i] || i != J[i] || i != K[i] {
								cnt++
							}
						}
					}
				}
			}
		}
		fmt.Println(cnt)
	}
}
