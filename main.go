package main

import "fmt"

func updateName(x string) string {
	x = "wedge"
	return x
}

func updateMenu(y map[string]float64) {
	y["coffee"] = 9999
}

func main() {
	// non-pointer values -> strings,int,bools,float,arrays,structs
	name := "tifa"

	// updateName(name) // returns wedge

	name = updateName(name) // returns tifu & saves it to variable
	fmt.Println(name)

	// pointer wrapped values -> slices,maps,functions
	menu := map[string]float64{
		"lime":   45,
		"coffee": 10,
	}
	updateMenu(menu)
	fmt.Println(menu)
}
