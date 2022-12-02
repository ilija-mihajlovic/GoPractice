package main

import (
	"fmt"
	"log"
)

func computeFactorials(num int) (uint, error) {
	var i int
	var factorial uint = 1

	if num < 0 {
		return 0, fmt.Errorf("factorial for negativ values is not allowed")
	}
	if num == 1 {
		return 1, nil
	}
	for i = num; i > 0; i-- {
		factorial *= uint(i)
	}
	return factorial, nil
}

func main() {
	fmt.Print("Enter the number: ")

	var number int
	
	_, err := fmt.Scanln(&number)

	if err != nil {
		log.Fatal("Please enter a number")
	}

	result, err := computeFactorials(number)

	fmt.Printf("\nFactorial of %d = %d", number, result)
}