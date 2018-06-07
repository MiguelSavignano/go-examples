package main

import "fmt"

func fizzbuzz(i int, c chan string) {
	var word string = ""
	if i % 3 == 0 {
		word += "Fizz"
	} else if i % 5 == 0 {
		word += "Buzz"
	} else {
		word = fmt.Sprintf("%v", i)
	}
	c <- fmt.Sprintf("%v => %s", i, word)
}

func printer(c chan string) {
	for {
		msg := <- c
		fmt.Println(msg)
	}
}

func main() {
	var print_channel chan string = make(chan string)
  for i := 1; i <= 100; i++ {
		go fizzbuzz(i, print_channel)
		go printer(print_channel)
  }
  var input string
  fmt.Scanln(&input)
}