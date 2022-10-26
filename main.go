package main

import "fmt"

func main() {
	myBill := newBill("mario's bill")

	myBill.updateTip(100)
	myBill.addItem("biriyani", 150)
	fmt.Println(myBill)
	fmt.Println(myBill.format())
}
