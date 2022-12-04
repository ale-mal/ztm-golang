//--Summary:
//  Implement the classic "FizzBuzz" problem using a `for` loop.
//
//--Requirements:
//* Print integers 1 to 50, except:
//  - Print "Fizz" if the integer is divisible by 3
//  - Print "Buzz" if the integer is divisible by 5
//  - Print "FizzBuzz" if the integer is divisible by both 3 and 5
//
//--Notes:
//* The remainder operator (%) can be used to determine divisibility

package main

import "fmt"

func main() {
	for i := 1; i <= 50; i++ {
		div3, div5 := i%3 == 0, i%5 == 0
		if div3 && div5 {
			fmt.Println("fizzbuzz")
		} else if div3 {
			fmt.Println("fizz")
		} else if div5 {
			fmt.Println("buzz")
		} else {
			fmt.Println(i)
		}
	}
}
