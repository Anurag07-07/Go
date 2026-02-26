package main

import "fmt"

type MyType string

// type OrderStatus int

// const (
	// 	Recieved OrderStatus = iota
	// 	Confirmed
	// 	Prepared
// 	Delivered
// )

type OrderStatus string

const (
	Recieved OrderStatus = "recieved"
	Confirmed = "Confirmed"
	Prepared = "prepared"
	Delivered = "delivered"
)

func changeOrderStatus(status OrderStatus) {
	fmt.Println("Changing order status to ",status)
}

func main() {	
	changeOrderStatus(Recieved)
}