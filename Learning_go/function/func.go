package main

import "fmt"

func add(a,b int)(int){
	return a+b
}

func multiply(c,d int)(int){
	return c*d
}

func main() {
	fmt.Println("Learning function")
    add:=add(3,4)
	fmt.Println("sum of two no is:",add)
	multiply:=multiply(5,8)
	fmt.Println("Product of two No is:",multiply)

}
