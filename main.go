package main

import "fmt"

var outsideFunction = "outsideFunction"

// outsideFunctiontwo:= ":= be used outside function"
func main() {
	//  string
	// var nameOne string = "hii"
	// var nameTwo = "sharath" // type inference
	// var nameThree string
	// nameFour := "SHARATH"

	// fmt.Println(nameOne+" "+nameTwo+" "+nameThree, nameFour)

	// nameThree = "hiiiiiiii"
	// fmt.Println(nameOne+" "+nameTwo+" "+nameThree, nameFour)

	// fmt.Println(outsideFunction)

	// ints
	var ageOne int = 20
	var ageTwo = 30
	ageThree := 40
	fmt.Println(ageOne, ageTwo, ageThree)

	// bits and memory

	// var numOne int8 = 25
	// var num3 uint8 = 10

	// float

	// var scoreOne float32 = -1.5
	// var scoreTwo float64 = 89398432.23423
	// scoreThree := 1.5 // infered to float64
}
