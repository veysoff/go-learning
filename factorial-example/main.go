package main

import (
	"fmt"
)

func main() {
	var num int
	fmt.Print("Enter a number: ")
	_, err := fmt.Scan(&num)

	if err != nil {
		fmt.Println("Invalid input.")
		return
	}

	if num < 0 {
		fmt.Println("Invalid negative number.")
		return
	}

	fmt.Printf("Factorial of %d is %d\n", num, factorial(num))
}

// factorial calculates the factorial of a given number using recursion
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}
