package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')
	return strings.TrimSpace(input), err
}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)
	name, _ := getInput("Created a new bill name ", reader)

	b := newbill(name)
	fmt.Println("Created the bill - ", b.name)

	return b
}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)
	opt, _ := getInput("Choose Option (a - add Item, s - save a bill, t - add a tip) ", reader)
	switch opt {
	case "a":
		name, _ := getInput("Item name ", reader)
		price, _ := getInput("Item price ", reader)

		p, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("the price must be a number")
			promptOptions(b)
		}
		b.addItem(name, p)
		fmt.Println("item added --:", name, price)
		promptOptions(b)
	case "t":
		tip, _ := getInput("Enter tip amount ($) ", reader)
		t, err := strconv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Println("the tip must be a number")
			promptOptions(b)
		}
		b.updateTip(t)
		fmt.Println("Tipe added --:", tip)
		promptOptions(b)
	case "s":
		b.save()
		fmt.Println("you choose to saved the file -", b.name)
	default:
		fmt.Println("that was a wrong option.....!")
		promptOptions(b)
	}
}

func main() {
	mybill := createBill()
	promptOptions(mybill)

}
