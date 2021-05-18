package main

import "fmt"

func multiplesOfThree(n int) bool {
	return n%3 == 0
}

func multiplesOfFive(n int) bool {
	return n%5 == 0
}

func multiplesOfThreeAndFive(i int) bool {
	return (i%3 == 0) && (i%5 == 0)
}

// FizzBuzz prints fizz (3), buzz (5), fizzbuzz(3&5) or interger
func FizzBuzz(i int) interface{} {
	if multiplesOfThreeAndFive(i) {
		return "fizzbuzz"
	} else if multiplesOfFive(i) {
		return "buzz"
	} else if multiplesOfThree(i) {
		return "fizz"
	}
	return i
}

func main() {
	fmt.Println("Higher-order functions")
	for i := 1; i <= 100; i++ {
		fmt.Printf("%v\n", FizzBuzz(i))
	}
}
