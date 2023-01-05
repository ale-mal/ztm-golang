package main

import "fmt"

type Preparer interface {
	PrepareDish()
}

type Chicken string
type Salad string

func (c Chicken) PrepareDish() {
	fmt.Println("cook chicken")
}

func (s Salad) PrepareDish() {
	fmt.Println("chop salad")
	fmt.Println("add dressing")
}

func prepareDishes(dishes []Preparer) {
	fmt.Println("Prepare dishes:")
	for i := 0; i < len(dishes); i++ {
		fmt.Printf("--Dish: %v--\n", dishes[i])
		dishes[i].PrepareDish()
	}
	fmt.Println()
}

func main() {
	dishes := []Preparer{
		Chicken("Chicken wings"),
		Salad("Caesar salad"),
	}
	prepareDishes(dishes)
}
