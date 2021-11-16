package main

import "fmt"

func main() {
	numbers := make(map[string]int)
	numbers["one"] = 1
	numbers["ten"] = 10
	numbers["three"] = 3
	fmt.Println("3: ", numbers["three"])

	rating := map[string]float32{"C": 5, "Go": 4.5, "Pthon": 4.5, "C++": 2}
	csharpRating, ok := rating["C#"]
	cRating, OK := rating["C"]
	if ok {
		fmt.Println("C# is in the map and its rating is ", csharpRating)
	} else {
		fmt.Println("We have no rating associated with C# inthe map")
	}
	if OK {
		fmt.Println("C is in the map and its rating is ", cRating)
	}
}
