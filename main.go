package main

import (
	"fmt"
	"math"
)

func sayGreetings(n string) {
	fmt.Printf("good morning %v \n", n)
}

func sayBye(n string) {
	fmt.Printf("good bye %v \n", n)
}

func cycleNames(n []string, f func(string)) {
	for _, value := range n {
		f(value)
	}
}

func circleArea(r float64) float64 {
	return math.Pi * r * r
}
func main() {
	names := []string{"sharath", "hgihjfe", "jhnoief"}
	// function which takes in function as argument
	cycleNames(names, sayBye)
	// function that returns a value
	a1 := circleArea(10)
	a2 := circleArea(2.5)
	fmt.Printf("are of cirlce a1 is %0.3f \n", a1)
	fmt.Printf("are of cirlce a2 is %0.3f \n", a2)

}
