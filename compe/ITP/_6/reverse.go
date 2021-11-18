package main

import "fmt"

func Reverse(list []int) {
	n := len(list)
	for i := 0; i < (n / 2); i++ {
		list[i], list[n-i-1] = list[n-i-1], list[i]
	}
}

func main() {
	var n int
	fmt.Scan(&n)

	list := make([]int, n)

	for i := range list {
		fmt.Scan(&list[i])
	}

	for i := 0; i < (n / 2); i++ {
		list[i], list[n-i-1] = list[n-i-1], list[i]
	}

	for i := 0; i < n; i++ {
		if i != n-1 {
			fmt.Printf("%d ", list[i])
		} else {
			fmt.Println(list[i])
		}
	}
}
