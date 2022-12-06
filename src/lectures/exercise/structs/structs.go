//--Summary:
//  Create a program to calculate the area and perimeter
//  of a rectangle.
//
//--Requirements:
//* Create a rectangle structure containing its coordinates
//* Using functions, calculate the area and perimeter of a rectangle,
//  - Print the results to the terminal
//  - The functions must use the rectangle structure as the function parameter
//* After performing the above requirements, double the size
//  of the existing rectangle and repeat the calculations
//  - Print the new results to the terminal
//
//--Notes:
//* The area of a rectangle is length*width
//* The perimeter of a rectangle is the sum of the lengths of all sides

package main

import "fmt"

type Vector2 struct {
	x int
	y int
}

/**
 *     |
 *     |
 * ----+--->x
 *     |
 *     V y
 */

type Rectangle struct {
	a Vector2 // left upper
	b Vector2 // right bottom
}

func calculateArea(rect Rectangle) int {
	return (rect.b.x - rect.a.x) * (rect.b.y - rect.a.y)
}

func calculatePerimeter(rect Rectangle) int {
	return (rect.b.x-rect.a.x)*2 + (rect.b.y-rect.a.y)*2
}

func printInfo(rect Rectangle) {
	fmt.Println("Rect is", rect)
	fmt.Println("Area of rect is", calculateArea(rect))
	fmt.Println("Perimeter of rect is", calculatePerimeter(rect))
}

func main() {
	rect := Rectangle{
		a: Vector2{0, 0},
		b: Vector2{2, 2},
	}

	printInfo(rect)

	rect.b.x *= 2
	rect.b.y *= 2

	printInfo(rect)
}
