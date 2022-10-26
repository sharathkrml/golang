package main

import "fmt"

var outsideFunction = "outsideFunction"

// outsideFunctiontwo:= ":= be used outside function"
func main() {
	//  string
	var nameOne string = "hii"
	var nameTwo = "sharath" // type inference
	var nameThree string
	nameFour := "SHARATH"

	fmt.Println(nameOne+" "+nameTwo+" "+nameThree, nameFour)

	nameThree = "hiiiiiiii"
	fmt.Println(nameOne+" "+nameTwo+" "+nameThree, nameFour)

	fmt.Println(outsideFunction)
}
