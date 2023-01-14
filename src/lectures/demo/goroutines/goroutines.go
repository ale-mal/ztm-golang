package main

import (
	"fmt"
	"time"
	"unicode"
)

func main() {
	data := []rune{'a', 'b', 'c', 'd'}
	capitalized := data

	for i := 0; i < len(data); i++ {
		// use channel to prevent data races
		ch := make(chan rune)
		go func() {
			ch <- unicode.ToUpper(data[i])
		}()
		capitalized[i] = <-ch
	}
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("Capitalized: %c\n", capitalized)
}
