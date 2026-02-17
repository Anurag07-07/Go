package main

import "fmt"

func main() {

	sum:= 0

	nums := []int{1, 3, 2, 4, 5}
	for _,num := range nums{
		sum = sum + num
	}

	fmt.Println(sum)

	user :=map[string]string{"A":"Something","B":"Something"}

	for k,v := range user{
		fmt.Println(k,v)
	}

	for i,c := range "golang"{
		fmt.Println(i,c)  //Unicode code point code
		fmt.Println(string(c))
	}
	
}