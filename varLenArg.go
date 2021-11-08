package main

import "fmt"

func varLenArg(arg ...int) {
	for _, n := range arg {
		fmt.Printf("And the number is: %d\n", n)
	}
}

func main() {
	varLenArg(1, 2, 3, 100)
}
