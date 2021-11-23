package main

import "fmt"

func main() {
	var n, m, l int
	fmt.Scan(&n, &m, &l)

	A := make([][]int, n)
	for i := 0; i < n; i++ {
		A[i] = make([]int, m)
	}
	B := make([][]int, m)
	for i := 0; i < m; i++ {
		B[i] = make([]int, l)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Scan(&A[i][j])
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < l; j++ {
			fmt.Scan(&B[i][j])
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < l; j++ {
			var c int
			for k := 0; k < m; k++ {
				c += A[i][k] * B[k][j]
			}
			if j == l-1 {
				fmt.Println(c)
			} else {
				fmt.Printf("%d ", c)
			}
		}
	}
}
