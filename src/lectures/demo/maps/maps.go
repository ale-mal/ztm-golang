package main

import "fmt"

func main() {
	shoppingList := make(map[string]int)
	shoppingList["eggs"] = 11
	shoppingList["milk"] = 10
	shoppingList["bread"] += 1

	shoppingList["eggs"] += 1

	fmt.Println(shoppingList)

	delete(shoppingList, "milk")
	fmt.Println("Milk deleted, shopping list after:", shoppingList)

	fmt.Println("need", shoppingList["eggs"], "eggs")

	cereal, found := shoppingList["cereal"]
	if !found {
		fmt.Println("Cereal not found")
	} else {
		fmt.Println("Cereal found, count is", cereal)
	}

	totalItems := 0
	for _, amount := range shoppingList {
		totalItems += amount
	}
	fmt.Println("Total items is", totalItems)
}
