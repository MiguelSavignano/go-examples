package main

// import "fmt"
import (
	L "./fizzbuzz/lib"
)
func main() {
	boom_word_clousure := func(number int) string {
		return "Boom"
	}

	L.Fizzbuzz(100, boom_word_clousure)
}