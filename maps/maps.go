package main

import (
	"fmt"
	"maps"
)

func main() {
	//Creating a map
	m1 := make(map[string]string)

	//Setting up Value
	m1["name"] = "golang"
	m1["company"] = "Google"
	fmt.Println(m1)

	//Fetching values
	fmt.Println(m1["name"])
	fmt.Println(m1["data"])  //if key is not present and type is string it return empty string if value is int it return 0 is boolean it return false

	//Check the length of the map
	fmt.Println(len(m1))

	//Delete the Key
	delete(m1,"name") //map name and key

	//Clear Whole Map
	clear(m1)

	//Intializing map

	m2  := map[string]string{"name":"Anurag","email":"anurag@gmail.com"}
	fmt.Println(m2)

	//Multiple Return value
	// _,ok := m2["name"] //If no type
	
	name,ok := m2["name"] 
	fmt.Println(name)

	if ok {
		fmt.Println("User is There")
	} else{
		fmt.Println("User is not there")
	}

	//Check if to map are equal or not
	fmt.Println(maps.Equal(m1,m2))

}