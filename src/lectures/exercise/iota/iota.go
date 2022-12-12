//--Summary:
//  Create a calculator that can perform basic mathematical operations.
//
//--Requirements:
//* Mathematical operations must be defined as constants using iota
//* Write a receiver function that performs the mathematical operation
//  on two operands
//* Operations required:
//  - Add, Subtract, Multiply, Divide
//* The existing function calls in main() represent the API and cannot be changed
//
//--Notes:
//* Your program is complete when it compiles and prints the correct results

package main

import "fmt"

type Operation int

const (
	add Operation = iota
	sub
	mul
	div
)

func (op Operation) calculate(lhs, rhs int) int {
	switch op {
	case add:
		return lhs + rhs
	case sub:
		return lhs - rhs
	case mul:
		return lhs * rhs
	case div:
		if rhs == 0 {
			panic("can't divide by zero")
		}
		return lhs / rhs
	default:
		panic("unknown operation")
	}
}

func main() {
	fmt.Println(add.calculate(2, 2)) // = 4

	fmt.Println(sub.calculate(10, 3)) // = 7

	fmt.Println(mul.calculate(3, 3)) // = 9

	fmt.Println(div.calculate(100, 2)) // = 50
}
