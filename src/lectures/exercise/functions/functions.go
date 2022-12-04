//--Summary:
//  Use functions to perform basic operations and print some
//  information to the terminal.
//
//--Requirements:
//* Add three numbers together using any combination of the existing functions.
//  * Print the result
//* Call every function at least once

package main

import "fmt"

//* Write a function that accepts a person's name as a function
//  parameter and displays a greeting to that person.
func greet(name string) {
	fmt.Println("greetins,", name, "!")
}

//* Write a function that returns any message, and call it from within
//  fmt.Println()
func getAnyMessage() string {
	return "any message"
}

//* Write a function to add 3 numbers together, supplied as arguments, and
//  return the answer
func tripleSum(a, b, c int) int {
	return a + b + c
}

//* Write a function that returns any number
func getAnyNumber() int {
	return 8
}

//* Write a function that returns any two numbers
func getAnyTwoNumbers() (int, int) {
	return 1, 2
}

func main() {
	greet("han yolo")

	fmt.Println(getAnyMessage())

	one, two, three := 1, 2, 3
	fmt.Println("sum of ", one, two, three, "is", tripleSum(one, two, three))

	one, two = getAnyTwoNumbers()
	three = getAnyNumber()
	fmt.Println("sum of ", one, two, three, "is", tripleSum(one, two, three))
}
