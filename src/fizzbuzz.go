package main

import "fmt"

func main() {
  for i := 1; i <= 100; i++ {
    var word = ""
    if i % 3 == 0 {
			word += "Fizz"
    } else if i % 5 == 0 {
			word += "Buzz"
		} else {
			word = fmt.Sprintf("%v",i)
		}
		fmt.Println(word)
		word = ""
  }
}