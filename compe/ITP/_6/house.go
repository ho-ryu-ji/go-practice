package main

import "fmt"

func main() {
	var n, b, f, r, v int
	var p [4][3][10]int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&b, &f, &r, &v)
		p[b-1][f-1][r-1] += v
	}
	for i := 0; i < 4; i++ {
		for j := 0; j < 3; j++ {
			for k := 0; k < 10; k++ {
				if k == 9 {
					fmt.Printf(" %d\n", p[i][j][k])
				} else {
					fmt.Printf(" %d", p[i][j][k])
				}
			}
		}
		if i == 3 {
			break
		} else {
			fmt.Printf("####################\n")
		}
	}
}
