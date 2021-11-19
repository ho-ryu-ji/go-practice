package main

import "fmt"

func main() {
	for {
		var n, x int
		fmt.Scan(&n, &x)
		var cnt int

		// cnt := make([][][]int, n)
		// for i := 0; i < n; i++ {
		// 	cnt[i] = make([][]int, n)
		// }
		// for i := 0; i < n; i++ {
		// 	for j := 0; j < n; j++ {
		// 		cnt[i][j] = make([]int, n)
		// 	}
		// }

		if n == 0 && x == 0 {
			break
		}
		for i := 1; i <= n; i++ {
			for j := 2; j <= n; j++ {
				for k := 3; k <= n; k++ {
					if i == j || i == k || j == k {
						continue
					} else if i+j+k == x {
						if 
						cnt++
					}
				}
			}
		}
		fmt.Println(cnt)
	}
}
