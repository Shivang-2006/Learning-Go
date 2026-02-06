package main

import "fmt"

func main() {
	anime := make(map[string]string)
	anime["One piece"] = "Luffy"
	anime["Naruto"] = "Itachi"
	anime["Bleach"] = "Ichigo"

	fmt.Println("The charatcer form anime is:",anime["Naruto"])
	anime["Naruto"]="madara"
	fmt.Println("new char is:",anime["Naruto"])

	//delete(anime,"Naruto")
	//fmt.Println("new char is:",anime["Naruto"])
	
	name,exist:=anime["Humterxhunter"]
	fmt.Println("name of character:",name,"\nIts exixtance:",exist)

	for index,value:=range anime{

		fmt.Println("name of anime is:",index,"---- ","Name of char is:",value)
	}





}
