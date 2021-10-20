package main

import (
	"fmt"

	"rsc.io/quote"
)

var a int
var b *int

func main() {
	a = 10
	b = &a
	fmt.Println(a)  //10
	fmt.Println(&a) //addr of a
	fmt.Println(b)  //addr of a
	fmt.Println(&b) //addr of b
	fmt.Println(*b) //10

	fmt.Println(quote.Go())

}
