//Summary:
//  Print basic information to the terminal using various variable
//  creation techniques. The information may be printed using any
//  formatting you like.
//
//Requirements:
//* Store your favorite color in a variable using the `var` keyword
//* Store your birth year and age (in years) in two variables using
//  compound assignment
//* Store your first & last initials in two variables using block assignment
//* Declare (but don't assign!) a variable for your age (in days),
//  then assign it on the next line by multiplying 365 with the age
// 	variable created earlier
//
//Notes:
//* Use fmt.Println() to print out information
//* Basic math operations are:
//    Subtraction    -
// 	  Addition       +
// 	  Multiplication *
// 	  Division       /

package main

import "fmt"

func main() {
	var favColor string = "Blue"
	fmt.Println("fav color is", favColor)

	birthYear, ageInYears := 1993, 29
	fmt.Println("birth year is", birthYear, "age in years is", ageInYears)

	var (
		firstInitial = "A"
		lastInitial  = "M"
	)
	fmt.Println("Initials are", firstInitial, lastInitial)

	var ageInDays int
	ageInDays = ageInYears * 365
	fmt.Println("age in days is", ageInDays)
}
