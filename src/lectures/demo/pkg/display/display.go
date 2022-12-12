package display

import "fmt"

// capital letter on function name means it'll be exported (i.e. public)
func Display(msg string) {
	fmt.Println(msg)
}

func hello(msg string) {
	fmt.Println(msg)
}
