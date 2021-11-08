package main

import "fmt"

func main() {
	doublearr := [2][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}}
	fmt.Printf("double array %d, %d\n", doublearr[0][1], doublearr[1][1])
}
