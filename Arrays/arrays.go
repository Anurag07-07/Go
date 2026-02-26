package main
import "fmt"
func main() {

	var nums [4]int
	fmt.Println(len(nums)) // Output: 4
	nums[0] = 1
	fmt.Println(nums[0]) // Output: 1
	var vals [5]bool // By Default Value is false
	fmt.Println(vals) // Output: [false false false false false]
	var car [5]string
	car[0] = "maclaren" // assign "maclaren" to the first slot
	fmt.Println(car) // Output: [maclaren    ]

	numss := [3]int{1, 2, 3}
	fmt.Println(numss) // Output: [1 2 3]
	twod := [3][3]int{{1, 2, 3}, {1, 2, 3}, {1, 2, 3}}
	fmt.Println(twod) // Output: [[1 2 3] [1 2 3] [1 2 3]]
	
}
