package main

import (
	"fmt"

	//go "golang.org/x/text/number"
)

func main() {

	for i := 0; i < 6; i++ {
		fmt.Println("numbers are:",i)

	}
	counter := 0
	for {
		fmt.Println("the infinite loop is :",counter)
		counter++
		 if counter == 3{
			break
		 }
		 
	}


	number:=[]int{1,2,3,4,5,6,7,8,9}

	for index,value:=range number{
		fmt.Printf("index of numver is %d:,Value of humber is %d\n",index,value)



	}
}