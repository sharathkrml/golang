package main

import "fmt"

func main() {
	menu := map[string]float64{
		"soup": 4.99,
		"pie":  5.0,
	}
	fmt.Println(menu)
	fmt.Println(menu["soup"])
	menu["cb"] = 100
	for k, v := range menu {
		fmt.Println(k, v)
	}

	menu["pie"] = 100
	for k, v := range menu {
		fmt.Println(k, v)
	}
}
