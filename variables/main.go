package main

import "fmt"

func main() {
	//Variable with type
	var name string = "golang"
	fmt.Println(name)

	//Variable without type 
	// It can automatically Infer the value
	var isActive = true
	fmt.Println(isActive)

	//Shortand Syntax
	first_name:="Anurag"
	fmt.Println(first_name)

 //When we only define the variable we must have to tell the type
	var company string
	company = "Google"
	fmt.Println(company)

	//Assign float
	var price float32 = 45.23
	fmt.Println(price)
}