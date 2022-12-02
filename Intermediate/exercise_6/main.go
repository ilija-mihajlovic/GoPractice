package main

/*
	Write a program that calculates and prints the value according to the given formula:
	Q = Square root of [(2 * C * D)/H]

	Following are the fixed values of C and H:
	C is 50. H is 30.

	D is the variable whose values should be input to your program in a comma-separated sequence.
*/

import (
	"encoding/json"
	"fmt"
	"math"
	"strconv"
	"strings"
)

const C = 50
const H = 30

func main(){
	fmt.Print("Enter comma-separeted numbers: ")
	
	var input string

	fmt.Scanf("%v", &input)

	numbers := strings.Split(input, ",")

	num := make([]float64, len(numbers))

	for index, value := range numbers {
		num[index], _ = strconv.ParseFloat(value, 8)
	}

	result := make([]int, len(num))

	for index, value := range num {
		result[index] = int(math.Round(math.Sqrt((2*C*value)/H)))
	}

	s, _ := json.Marshal(result)
	fmt.Println(strings.Trim(string(s), "[]"))

}