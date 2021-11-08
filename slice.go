package main

import "fmt"

func main() {
	var arr = [10]int{43, 13, 22, 54, 23, 35, 16, 78, 91, 27}
	a := arr[2:5]
	b := arr[3:5]
	fmt.Printf("first %d\n", arr[0])
	fmt.Printf("last %d\n", arr[9])
	fmt.Printf("a = %d b = %d\n", a, b)
}
