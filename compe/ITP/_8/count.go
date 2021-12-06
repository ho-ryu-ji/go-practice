package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	a := make([]int, 26)
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		s := scanner.Text()

		for i := 0; i < len(s); i++ {
			if 'a' <= s[i] && s[i] <= 'z' {
				a[s[i]-'a']++
			}
			if 'A' <= s[i] && s[i] <= 'Z' {
				a[s[i]-'A']++
			}
		}
	}

	for i := 0; i < 26; i++ {
		fmt.Printf("%c : %d\n", i+'a', a[i])
	}
}
