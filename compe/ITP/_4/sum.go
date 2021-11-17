package main

import "fmt"

func main() {
	var n, tmp int
	fmt.Scan(&n)
	min, max, sum := 10000000, -1000000, 0

	for i := 0; i < n; i++ {
		fmt.Scan(&tmp)

		if min > tmp {
			min = tmp
		}
		if max < tmp {
			max = tmp
		}
		sum += tmp
	}
	fmt.Printf("%d %d %d\n", min, max, sum)
}
