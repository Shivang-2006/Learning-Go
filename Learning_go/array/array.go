package main

import "fmt"

func main() {

	fmt.Println("Learning arrays")

	var subs[3]string

	subs[0]="Maths"

	subs[2]="science"
	subs[1]="English"
	fmt.Println("Name of subs are",subs)

	var mrks = [3]int{90,60,70}
	fmt.Println("Marks obtained:",mrks)

}