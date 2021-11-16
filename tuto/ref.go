package main

import "fmt"

func add1(a *int) int {
	*a = *a + 1
	return *a
}

func main() {
	d := 3
	fmt.Println("d = ", d)
	d1 := add1(&d)
	fmt.Println("d+1 = ", d1)
	fmt.Println("d = ", d)
}
