package main

import "fmt"

func main() {
	var n, m int

	fmt.Scan(&n, &m)

	A := make([][]int, n)
	for i := 0; i < n; i++ {
		A[i] = make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Scan(&A[i][j])
		}
	}

	b := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Scan(&b[i])
	}

	for i := 0; i < n; i++ {
		var c int
		for j := 0; j < m; j++ {
			c += A[i][j] * b[j]
		}
		fmt.Println(c)
	}
}
