package main

import "fmt"

func printSlice[T any](items []T) {
	for _, val := range items {
		fmt.Println(val)
	}
}
//Or
func printSlice1[T comparable](items []T) {
	for _, val := range items {
		fmt.Println(val)
	}
}

type stack[T any] struct{
	elements []T
}

func main() {
	nums:=[]int{1,2,3,4,5}
	printSlice(nums)

	myStack:=stack[string]{
		elements: []string{"golang"},
	}

	fmt.Println(myStack)
}