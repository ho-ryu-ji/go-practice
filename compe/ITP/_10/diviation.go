package main

import (
	"fmt"
	"math"
)

func sum(arr []float64) float64 {
	a := 0.0
	for _, v := range arr {
		a += v
	}
	return a
}

func meanSD(arr []float64) float64 {
	n := float64(len(arr))
	m := sum(arr) / n
	s := 0.0
	for _, x := range arr {
		s += (x - m) * (x - m)
	}
	return math.Sqrt(s / n)
}

func main() {
	for {
		var n int
		fmt.Scan(&n)

		if n == 0 {
			break
		}

		arr := make([]float64, n)
		for i := 0; i < n; i++ {
			fmt.Scan(&arr[i])
		}

		fmt.Println(meanSD(arr))
	}
}
