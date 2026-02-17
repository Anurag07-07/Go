package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

// Multiple Return type 
func getLanguage() (string,string,string){
	return "golang","java","C++"
}

//Returning function
func processIt() func(a int) int{
	return  func(a int) int {
		return 4
	}
}  

func main() {
	res := add(45, 56)
	fmt.Println(res)

	lang1,lang2,lang3 := getLanguage();
	// lang1,lang2,_ := getLanguage(); if we dont want to get some value
	fmt.Println(lang1)
	fmt.Println(lang2)
	fmt.Println(lang3)


	processIt()
}