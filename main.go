package main

import (
	"fmt"
	"strings"
)

func main() {
	greeting := "hello there!!"
	fmt.Println(strings.Contains(greeting, "hello"))
	fmt.Println(strings.ReplaceAll(greeting, "hello", "hii"))
	fmt.Println(strings.ToUpper((greeting)))

	fmt.Println(strings.Index(greeting, " "))
	fmt.Println(strings.Split(greeting, " "))
	fmt.Println(greeting)
}
