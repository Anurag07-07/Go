package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	fmt.Printf("Hello my name is %s and I am %d year Old","Anurag",22)

	var name1 string = "Anurag";
	var name2 = "Anurag"
	name3:="Anurag"
	var name4 string
	name4 = "Anurag"
	msg:= fmt.Sprintf("Hello This is Multiple type of Decelaration %s %s %s %s",name1,name2,name3,name4)
	fmt.Println(msg)

	var(
		x float32 = 45.56
		y int = 456
		z string = "anurag"
	)

	fmt.Println(x,y,z)
}