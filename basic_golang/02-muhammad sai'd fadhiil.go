package main

import (
	"fmt"
)

func main() {
	var c string

	fmt.Scan(&c)

	switch c {
	case "a":
		fmt.Println("vocal")
	case "i":
		fmt.Println("vocal")
	case "u":
		fmt.Println("vocal")
	case "e":
		fmt.Println("vocal")
	case "o":
		fmt.Println("vocal")
	default:
		fmt.Println("konsonan")
	}
}
