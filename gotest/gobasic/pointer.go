package main

import (
	"fmt"
)

type ByteSlice []byte

func (slice ByteSlice) Append1(data []byte) []byte {
	
}

func (p *ByteSlice) Write(data []byte) (n int, err error) {
	slice := *p

	*p = slice
	return len(data),nil
}