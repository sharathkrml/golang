package main

import "fmt"

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// make new bill

func newBill(name string) bill {
	b := bill{name: name, items: map[string]float64{"pie": 5.99, "cake": 3.79}, tip: 0}
	return b
}

// format the bill

// receiver function,only can be called by bill (like methods inside object)
func (b *bill) format() string {
	fs := "bill breakdown \n"
	var total float64 = 0

	for k, v := range b.items {
		total += v
		fs += fmt.Sprintf("%-15v...$%v \n", k+":", v)
	}
	total += b.tip
	fs += fmt.Sprintf("%-15v...$%0.2f \n", "tip:", b.tip)
	fs += fmt.Sprintf("%-15v...$%0.2f", "total:", total)
	return fs
}

//update the tip

func (b *bill) updateTip(tip float64) {
	b.tip = tip
}

func (b *bill) addItem(name string, price float64) {
	b.items[name] = price
}
