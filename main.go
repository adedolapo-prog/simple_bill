package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInputs (r *bufio.Reader, prompt string) (string, error) {
	fmt.Print(prompt)

	input, err := r.ReadString('\n')
	return strings.TrimSpace(input), err
}

func inputBill() bill {
	reader := bufio.NewReader(os.Stdin)

	name, _ := getInputs(reader, "Create a new bill name: ")

	b := createBill(name)
	fmt.Println("created bill - ", b.name)
	return b
} 

func promptOptions (b bill) {
	reader := bufio.NewReader(os.Stdin)

	opt, _ := getInputs(reader, "Choose options: (a: add item, s: save bill, t: add tip): ")
	
	switch opt {
	case "a":
		name, _ := getInputs(reader, "Item Name: ")
		price, _ := getInputs(reader, "Item Price: ")

		p, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("The price must be a number")
			promptOptions(b)
		}

		b.addItem(name, p)

		fmt.Println("Item Added -", name, price)
		promptOptions(b)
	
	case "t":
		tip, _ := getInputs(reader, "Enter tip amount ($): ")

		t, err := strconv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Println("The tip must be a number")
			promptOptions(b)
		}

		b.updateTip(t)

		fmt.Println("Your tip:", tip)
		promptOptions(b)
	
	case "s":
		b.saveBill()
		fmt.Println("You saved the file -", b.name)

	default:
		fmt.Println("That was not a valid option")
		promptOptions(b)
	}
}

func main() {
	myBill := inputBill()
	promptOptions(myBill)
}

