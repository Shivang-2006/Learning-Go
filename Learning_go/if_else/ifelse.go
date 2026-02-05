package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Learning if else condition")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter a number: ")
	input, _ := reader.ReadString('\n')

	input = strings.TrimSpace(input)

	a, _ := strconv.Atoi(input) 

	if a > 5 {
		fmt.Println("a is greater than 5")
	} else {
		fmt.Println("a is smaller than or equal to 5")
	}
}
