package main

import "fmt"

func main() {
	var r, c int
	fmt.Scan(&r, &c)

	sheet := make([][]int, r+1)
	for i := 0; i < r+1; i++ {
		sheet[i] = make([]int, c+1)
	}

	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			fmt.Scan(&sheet[i][j])
			sheet[i][c] += sheet[i][j]
			sheet[r][j] += sheet[i][j]
		}
		sheet[r][c] += sheet[i][c]
	}

	for i := 0; i < r+1; i++ {
		for j := 0; j < c+1; j++ {
			if j != c {
				fmt.Printf("%d ", sheet[i][j])
			} else {
				fmt.Println(sheet[i][j])
			}
		}
	}
}
