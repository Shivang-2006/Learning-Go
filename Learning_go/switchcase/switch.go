package main

import "fmt"

func main() {
	fmt.Println("Learning switch case")

	//basic one
	day := 4

	switch day {

	case 1:
		fmt.Println("monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
    default:
		fmt.Println("Unknown Day")
	} 



    month:= "January"

	switch month{
	case "January","Feburary","March":
		fmt.Println("sardiii")
	case "April","May","June":	
	    fmt.Println("summer")
	default:
		fmt.Println("Other season")
	}


	temp:=-25

	switch{
	case temp>0:
		fmt.Println("thandaaa")
	case temp>=25:
		fmt.Println("Garammm")
	case temp<0:
		fmt.Println("Jm gye guru")
	}




}
