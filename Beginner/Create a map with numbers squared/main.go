package main

import (
	"fmt"
	"log"
)

func createMap(number int)(map[int]int, error){
	result := make(map[int]int)

	for i := 1; i <= number; i++ {
		result[i] = i*i
	}

	return result, nil
}

func main(){
	fmt.Print("Enter a positive number: ")
	var input int
	_, err := fmt.Scanln(&input)

	if err != nil {
		log.Fatal("Please enter a number")
		return
	}

	if input == 0 {
		log.Fatal("Please enter a number greater then 0")
		return
	}

	result, err := createMap(input)

	fmt.Println(result)

}