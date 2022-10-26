package main

import (
	"fmt"
	"sort"
)

func main() {
	ages := []int{135, 45, 90}
	sort.Ints(ages) //alters orignal slice
	fmt.Println(ages)
	index := sort.SearchInts(ages, 80)
	fmt.Println(index)
	names := []string{"b", "a", "g", "z", "f"}
	sort.Strings(names)
	fmt.Println(names)
	fmt.Println(sort.SearchStrings(names, "g"))
}
