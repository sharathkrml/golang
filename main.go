package main

import "fmt"

func main() {
	// age := 45
	// fmt.Println(age >= 50)
	// fmt.Println(age <= 50)
	// fmt.Println(age == 30)
	// fmt.Println(age != 40)

	// if-else
	// if age < 40 {
	// 	fmt.Println("age less than 40")
	// } else if age < 50 {
	// 	fmt.Println("age less than 50")
	// } else {
	// 	fmt.Println("damn age!!")
	// }

	names := []string{"a", "b", "c", "d", "e"}
	for index, value := range names {
		if index == 1 {
			fmt.Println("continuing at pos", index)
			continue
		}
		fmt.Printf("value at pos %v is %v \n", index, value)
	}

	for index, value := range names {
		if index == 1 {
			fmt.Println("breaking at pos", index)
			break
		}
		fmt.Printf("value at pos %v is %v \n", index, value)
	}
}
