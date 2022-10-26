package main

import "fmt"

var score = 9.99 // inside package's scope - if commented,it cannot be used from greetings.go

func main() {
	sayHello("sharath")

	// var score = 9.99 //outside package's scope

	for _, v := range points {
		fmt.Println(v)
	}

	showScore()
}

//go run main.go greetings.go
