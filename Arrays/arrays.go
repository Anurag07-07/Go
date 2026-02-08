package main

import "fmt"

func main() {
	var nums [4]int

	fmt.Println(len((nums)))

	nums[0] = 1

	fmt.Println(nums[0])

	var vals [5]bool //By Default Value is False

	fmt.Println(vals)

	var car [5]string
	car[0] = "maclaren"
	fmt.Println(car)

	// int->0 bool->false string->"" By Default Value

	//To Declare it in single line
	numss:=[3]int{1,2,3}
	fmt.Println(numss)

	//2d Array in single line
	twod:=[3][3]int{{1,2,3},{1,2,3},{1,2,3}}
	fmt.Println(twod)

	
}