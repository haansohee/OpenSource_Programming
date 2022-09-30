package main

import "fmt"

func main() {
	a := 4 // var int a = 4
	pa := &a

	fmt.Println(*pa)
	fmt.Println(pa)
	fmt.Println(&a)
}
