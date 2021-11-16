package main

import (
	"fmt"
	"sort"
)

func main() {
	var arr []int

	for i := 0; i < 3; i++ {
		var a int
		fmt.Scan(&a)
		arr = append(arr, a)
	}

	sort.Ints(arr)
	for i := 0; i < 3; i++ {
		fmt.Printf("%d", arr[i])
		if i != 2 {
			fmt.Printf(" ")
		} else {
			fmt.Println()
		}
	}
}
