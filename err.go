package main

import (
	"errors"
	"fmt"
)

func main() {
	err := errors.New("emit macho dwarf: elf heder corrupted")
	if err != nil {
		fmt.Println(err)
	}
}
