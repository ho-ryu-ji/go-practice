package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	stdin := bufio.NewScanner(os.Stdin)

	stdin.Scan()
	str := stdin.Text()

	for _, c := range str {
		if c >= 'A' && c <= 'Z' {
			fmt.Printf("%c", c+'a'-'A')
		} else if c >= 'a' && c <= 'z' {
			fmt.Printf("%c", c-'a'+'A')
		} else {
			fmt.Printf("%c", c)
		}
	}
	fmt.Println()
}
