package main

import (
	"fmt"
	"math"
)

func minko(n int, p float64, x []float64, y []float64) float64 {
	var d float64
	for i := 0; i < n; i++ {
		a := x[i] - y[i]
		if a < 0 {
			a *= -1
		}
		d += math.Pow(a, p)
	}

	return math.Pow(d, 1/p)
}

func cheby(n int, x []float64, y []float64) float64 {
	var d float64
	for i := 0; i < n; i++ {
		a := x[i] - y[i]
		if a < 0 {
			a *= -1
		}
		d = math.Max(d, a)
	}

	return d
}

func main() {
	var n int
	fmt.Scan(&n)
	var p float64 = 3

	x := make([]float64, n)
	y := make([]float64, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&x[i])
	}
	for i := 0; i < n; i++ {
		fmt.Scan(&y[i])
	}

	for i := 1; i <= int(p); i++ {
		fmt.Println(minko(n, float64(i), x, y))
	}
	fmt.Println(cheby(n, x, y))
}
