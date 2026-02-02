package main

import "fmt"

func main() {

	fmt.Println("hey what is your name?")
	var name string;
	fmt.Scan(&name)
	fmt.Println("Hello Mr",name)
	fmt.Println("What is your age?")
	var age int;
	fmt.Scan(&age)
	fmt.Println("so yor are",age,"years old")


}