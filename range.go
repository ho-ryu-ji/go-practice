package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	a := arr[2:5]

	for k, v := range a {
		fmt.Println("map's key:", k)
		fmt.Println("map's val:", v)
	}
}
