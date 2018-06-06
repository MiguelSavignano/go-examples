package main

import "fmt"
import (
	L "./fizzbuzz/lib"
)
func main() {
	number_word_clousure := func(number int) string {
	  return fmt.Sprintf("%v", number)
	}

	L.Fizzbuzz(100, number_word_clousure)
}