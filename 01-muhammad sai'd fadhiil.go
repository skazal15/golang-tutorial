package main

import (
	"fmt"
)

func main() {

	var c, reverse string
	fmt.Scan(&c)

	a := func(c string) bool {
		for i := len(c) - 1; i >= 0; i-- {
			reverse += string(c[i])
		}

		for i := range c {
			if c[i] != reverse[i] {
				return false
			}
		}
		return true
	}
	fmt.Println(a(c))

}
