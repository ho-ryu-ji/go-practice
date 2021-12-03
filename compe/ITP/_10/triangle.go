package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, C float64
	fmt.Scan(&a, &b, &C)
	fmt.Println(a * b * math.Sin(C*math.Pi/180) / 2)
	fmt.Println(a + b + math.Sqrt(math.Pow(a, 2)+math.Pow(b, 2)-2*a*b*math.Cos(C*math.Pi/180)))
	fmt.Println(b * math.Sin(C*math.Pi/180))
}
