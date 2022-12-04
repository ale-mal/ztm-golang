package main

import "fmt"

func fibonacci(num uint) uint {
	if num <= 1 {
		return num
	}
	return fibonacci(num-1) + fibonacci(num-2)
}

func main() {
	for i := uint(0); i <= 10; i++ {
		fmt.Println("i:", i, ", fibonacci is:", fibonacci(i))
	}
}
