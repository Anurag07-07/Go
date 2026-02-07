package main

import (
	"fmt"
	"time"
)

func main() {
	// age := 19
	// if age > 18 {
	// 	fmt.Println("Adult")
	// }


	temp := 15

	if temp<=15 {
		fmt.Println("Cold Day")
	}else if temp >=15 && temp<=25 {
		fmt.Println("Moderate Day")
	}else{
		fmt.Println("Hot Day")
	} 

	//Variable declare with condition
	if age:=15; age>=18   {
		fmt.Println("Adult")
	}else{
		fmt.Println("Teenager")
	}

	//Go does not have ternary Operator


	//Switch
	var i int = 1
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("four")
	}

	//Multiple Condition Switch
	switch time.Now().Weekday(){
	case time.Saturday,time.Sunday:
		fmt.Println("Weekend")
	default:
		fmt.Println("Work Days")
	}


	//Type switch
	whoAmI:=func (i interface{})  {
		switch t := i.(type){
		case int:
			fmt.Println("Its an Integer")
		case string:
			fmt.Println("Its an String")
		case bool:
			fmt.Println("Its an Boolean")
		default:
			fmt.Printf("Unknown type: %T\n", t)
		}
	}
	
	whoAmI("golang")
}