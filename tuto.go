package main

import (
	"errors"
	"fmt"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// Hello Go
	fmt.Printf("Hello World\n")

	// imaginary number
	var com complex64 = 5 + 5i
	fmt.Printf("Value is: %v\n", com)

	// change a char in string
	s := "hello"
	c := []byte(s)
	c[0] = 'c'
	s2 := string(c)
	fmt.Printf("%s\n", s2)

	// err
	err := errors.New("emit macho dwarf: elf heder corrupted")
	if err != nil {
		fmt.Print(err)
	}

	// array n slise
	var arr = [10]int{43, 13, 22, 54, 23, 35, 16, 78, 91, 27}
	a := arr[2:5]
	b := arr[3:5]
	fmt.Printf("first %d\n", arr[0])
	fmt.Printf("last %d\n", arr[9])
	fmt.Printf("a = %d b = %d\n", a, b)

	// array in array
	doublearr := [2][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}}
	fmt.Printf("double array %d, %d\n", doublearr[0][1], doublearr[1][1])

	// map
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

	// range
	for k, v := range a {
		fmt.Println("map's key:", k)
		fmt.Println("map's val:", v)
	}

	x, y, z := 3, 4, 5
	fmt.Printf("max(%d, %d) = %d\n", x, y, max(x, y))
	fmt.Printf("max(%d, %d) = %d\n", x, z, max(x, z))

}