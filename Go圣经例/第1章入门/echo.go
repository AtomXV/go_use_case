package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		sep = " "
		s += sep + arg
		fmt.Println("index:")
	}
}
