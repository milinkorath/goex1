package main

import (
	"fmt"
	"strings"
)

func main() {
	a := "11 21 31"
	b := strings.Fields(a)

	for _, value := range b {
		fmt.Println(value)
	}
}
