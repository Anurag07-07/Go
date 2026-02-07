package main

import "fmt"

//There is only for keyword for loops
func main()  {
	i:=1
	for i<=3 {
		fmt.Println(i)
		i = i+1
	}


	//Infinite loop

	for{
		break		
	}

	//Classic For loop
	for i := 0; i < 5; i++ {
		if i==2 {
			continue
		}
		fmt.Println(i)
	}

	//Range
	for i:=range 3 {   //3 is Excluded
		fmt.Println(i)
	}
}