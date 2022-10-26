package main

import "fmt"

func main() {
	// array in go is fixed length
	// var ages [3]int = [3]int{1, 2, 3}
	var ages = [3]int{1, 2, 3}
	names := [3]string{"hi", "hello", "gm"}
	ages[0] = 100 //changing value by index
	fmt.Println(ages, len(ages))
	fmt.Println(names, len(names))

	// slices - array without length
	// they use arrays under the hood

	var scores = []int{100, 200, 300, 400}
	scores[2] = 20
	scores = append(scores, 9999) // append creates a new slice,which is assigned to scores manually
	fmt.Println(scores, len(scores))

	// slice ranges
	rangeOne := names[1:3] //index 1,2
	rangeOne = names[0:2]  //index 0,1
	fmt.Println(rangeOne)

	rangeTwo := scores[1:3]
	fmt.Println(rangeTwo)
	letsseeslice := []string{"first", "second", "third", "fourth"}
	rangeThree := letsseeslice[1:] //everything starting from 1st index

	fmt.Println(rangeThree)
	rangeAppended := append(rangeOne, "wohooo")
	fmt.Println(rangeAppended)

}
