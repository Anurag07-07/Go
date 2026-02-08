package main

import (
	"fmt"
	"slices"
)

//Slices are like dynmaic Array
//most used contruct in go
//useful methods
func main(){
	var nums []string
	fmt.Println(nums==nil)
	fmt.Println(len(nums))

	//Declaration and size 
	var nums1 = make([]int, 5,50)  //Five here is size of array generally we put 0 and 50 is capacity 
	fmt.Println(nums1)


	//Append the value inside the Array
	nums1 = append(nums1, 1,2,3,4,5)  //It will append from last 

	fmt.Println(cap(nums1))  //Maximum number of element can fit 
	fmt.Println(nums1)
	
	nums1[0] = 45
	
	fmt.Println(nums1)

	//Slice Copy
	var nums2 = make([]int, len(nums1))

	copy(nums2,nums1) //copy(duplicate array,og array)

	//Slice Operator

	var nums3 = []int{4,5,6};
	fmt.Println(nums3[0:1]) //Start from 0 and take one element
	fmt.Println(nums3[:1]) //Start from 0 and take one element
	fmt.Println(nums3[1:]) //Start from 0 and take one element

	fmt.Println(nums1,nums2)

	//Compare two slices
	fmt.Println(slices.Equal(nums1,nums2))

	//Two D
	var nums5 = [][]int{{1,2,3},{4,6,5}}
	fmt.Println(nums5)
}