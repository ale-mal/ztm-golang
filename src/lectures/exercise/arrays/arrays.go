//--Summary:
//  Create a program that can store a shopping list and print
//  out information about the list.
//
//--Requirements:
//* Using an array, create a shopping list with enough room
//  for 4 products
//  - Products must include the price and the name
//* Insert 3 products into the array
//* Print to the terminal:
//  - The last item on the list
//  - The total number of items
//  - The total cost of the items
//* Add a fourth product to the list and print out the
//  information again

package main

import "fmt"

type Product struct {
	name  string
	price int
}

func printProducts(products [4]Product) {
	var totalNumber int
	var totalCost int
	var lastIndex int
	lastIndex = -1

	for i := 0; i < len(products); i++ {
		j := i

		if products[j].name == "" {
			continue
		}

		totalNumber++
		totalCost += products[j].price
		lastIndex = j
	}

	if lastIndex == -1 {
		fmt.Println("Products are empty!")
		return
	}

	fmt.Println("Total number is", totalNumber)
	fmt.Println("Total cost is", totalCost)
	fmt.Println("Last item name is", products[lastIndex])
}

func main() {
	var products [4]Product

	products[0] = Product{"Milk", 60}
	products[1] = Product{"Banana", 10}
	products[2] = Product{"Onigiri", 20}

	printProducts(products)

	products[3] = Product{"Whiskey", 900}

	printProducts(products)
}
