package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func ReadWrite() bool {
	file, err := os.Open("file.txt")

	if err != nil {
		// error
		fmt.Printf("error")
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if scanner.Err() != nil {
		log.Fatal(err)
	}
	return true
}

func main() {
	ReadWrite()
}
