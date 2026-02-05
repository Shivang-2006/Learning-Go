package main

import (
	"fmt"
	//"golang.org/x/text/number"
)

func main() {

	fmt.Println("Learning slice")

	numbers := make([]int,3,5)
	numbers[0]=2
	numbers[1]=3
	numbers[2]=4
	numbers=append(numbers, 3,5,6,6)

	fmt.Println("The numbers are ", numbers)
	fmt.Println("The Length of Number is", len(numbers))
	fmt.Println("The capacity of number is  ", cap(numbers))

}
