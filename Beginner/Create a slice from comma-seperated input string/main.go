package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main(){

	fmt.Print("Enter comma-separeted numbers: ")
	
	var input string

	fmt.Scanf("%v", &input)

	numbers := strings.Split(input, ",")

	var num = make([]int, len(numbers))

	for index, value := range numbers {
		num[index], _ = strconv.Atoi(value)
	}

	fmt.Println(num)


}