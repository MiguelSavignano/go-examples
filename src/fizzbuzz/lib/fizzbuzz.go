package fizzbuzz

import "fmt"

func Fizzbuzz(size int, number_clousure func(int) string) {
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
