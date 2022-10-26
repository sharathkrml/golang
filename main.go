package main

import "fmt"

func updateName(x *string) {
	// takes value from pointer & updates it
	*x = "wedge"
}

func updateNumber(n *int) {
	*n = 1000
}

func main() {
	name := "tifa"
	pointer := &name

	fmt.Println("memory address of name:", pointer)
	fmt.Println("value at address of pointer:", *pointer)

	updateName(pointer)

	fmt.Println(name)

	num := 1
	updateNumber(&num)
	fmt.Println("number is:", num)
}
