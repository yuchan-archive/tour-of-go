package main

import (
"fmt"
)

func fb(n int) {
	m := (n%3) ^ (n%5)
	fmt.Printf("%d",m)
	fmt.Println("Fizz")
	fmt.Println("Buzz")
	fmt.Println("FizzBuzz")
	fmt.Println("")
}

func main() {
	for i := 0; i < 20; i++ {
		fb(i)
	}
}
