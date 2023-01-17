package main

import (
	"fmt"
	"os"
)

type bill struct {
	name string
	items map[string]float64
	tip float64
}

func createBill(name string) bill {
	b := bill {
		name: name,
		items: map[string]float64 { },
		tip: 0,
	}

	return b
}

func (b *bill) format () string {
	fs := "Bill Breakdown: \n"

	var total float64 = 0

	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v ...$%v \n", k+":", v)
		total += v
	}

	//tip
	fs += fmt.Sprintf("%-25v ...$%0.2f \n", "tip:", b.tip)

	//total
	fs += fmt.Sprintf("%-25v ...$%0.2f", "total:", total+b.tip)
	return fs
}

func (b *bill) updateTip (tip float64) {
	(*b).tip = tip
}

func (b *bill) addItem(name string, price float64) {
	b.items[name] = price
}

//save bill
func (b *bill) saveBill () {
	data := []byte(b.format())

	err := os.WriteFile("bills/"+b.name+".txt", data, 0644)
	if err != nil {
		panic(err)
	}

	fmt.Println("Bill was saved to file")
}