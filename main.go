package main

import "fmt"

func main() {
	age := 30
	fmt.Print("hello, \n") //PrintLn adds \n by default
	fmt.Print("world , ")

	// println
	fmt.Println("my age is", 30)

	// formatted strings
	fmt.Printf("My Age is %v and my name is %v  \n", "30", "Sharath")
	fmt.Printf("My Age is %q and my name is %q  \n", "30", "Sharath") // add "quotes"
	fmt.Printf("type of age is %T \n", age)                           //Type of variable
	fmt.Printf("float %0.2f", 2.555)                                  // float upto 2 decimal points
}
