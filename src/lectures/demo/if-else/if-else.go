package main

import "fmt"

func average(a, b, c int) float32 {
	// Convert the sum of the scores into a float32
	return float32(a+b+c) / 3
}

func main() {
	quiz1, quiz2, quiz3 := 9, 7, 8

	if quiz1 > quiz2 {
		fmt.Println("q1 is greater than q2")
	} else if quiz1 < quiz2 {
		fmt.Println("q1 is lower than q2")
	} else {
		fmt.Println("q1 and a2 are the same")
	}

	if average(quiz1, quiz2, quiz3) > 7 {
		fmt.Println("acceptable")
	} else {
		fmt.Println("not acceptable")
	}
}
