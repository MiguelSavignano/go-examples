package main

import "fmt"

func fizzbuzz(size int, number_clousure func(int) string) {
	for i := 1; i <= size; i++ {
		var word = ""
		if i % 3 == 0 {
			word += "Fizz"
		} else if i % 5 == 0 {
			word += "Buzz"
		} else {
			word = number_clousure(i)
		}
		fmt.Println(word)
		word = ""
	}
}

func main() {
	number_word_clousure := func(number int) string {
		return fmt.Sprintf("%v",number)
	}

	fizzbuzz(100, number_word_clousure)
}