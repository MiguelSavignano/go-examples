package main

import "strings"
import "fmt"
import (
	L "./fizzbuzz/lib"
)
func main() {
	boom_word_clousure := func(number int) string {
		var o_repeats = strings.Repeat("o", number)
		return fmt.Sprintf("B%sm!!", o_repeats)
	}

	L.Fizzbuzz(100, boom_word_clousure)
}