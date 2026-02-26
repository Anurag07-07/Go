package main

import (
	"fmt"
	"time"
)

//Sending
func processChan(numChan chan int){
	for num:=range numChan{
		fmt.Println("processing number",num)
		time.Sleep(time.Second*1)
	}
	// fmt.Println("processing number",<-numChan)
}

func task(done chan bool){
	defer func ()  {
		done<-true
	}()

	fmt.Println("Processing...")
}

func sum(result chan int,num1 int ,num2 int){
	numresult := num1+num2
	result <- numresult
}

func emailSender(emailChan chan string,done chan bool){
	defer func ()  {
		done<- true
	}()

	for email := range emailChan{
		fmt.Println("sending email to",email)
		time.Sleep(time.Second)
	}
}

func emailrecieve(emailChan <-chan string,done chan bool){
	defer func ()  {
		done<- true
	}()

	for email := range emailChan{
		fmt.Println("sending email to",email)
		time.Sleep(time.Second)
	}
}

func emailonlysend(emailChan chan string,done chan<- bool){
	defer func ()  {
		done<- true
	}()

	for email := range emailChan{
		fmt.Println("sending email to",email)
		time.Sleep(time.Second)
	}
}


func main() {

	chan1:=make(chan int)
	chan2:=make(chan string)


	go func() {
		chan1<-10
	}()

	go func() {
		chan2<-"pong"
	}()

	for i := 0; i < 2; i++ {
		select{
		case chan1Val:=<-chan1:
		fmt.Println(chan1Val)
		case chan2Val:=<-chan2:
		fmt.Println(chan2Val)
		}
	}

	// emailChan := make(chan string,100)
	// done := make(chan bool)

	// go emailSender(emailChan,done)
	// for i := 0; i < 100; i++ {
	// 	emailChan<-fmt.Sprintf("%d@gmail.com",i)
	// }


	// fmt.Println("done sending")
	// close(emailChan)
	// <-done



	// done:=make(chan bool)
	// go task(done)

	// <-done //block

	// numChan := make(chan int)

	// go processChan(numChan)


	// for {
	// 	numChan <- rand.Intn(100)
	// }

	// result:= make(chan int)
	// go sum(result,4,5)
	// res:=<-result
	// fmt.Println(res)

	// numChan<- 5
	// time.Sleep(time.Second*2)

	// message := make(chan string)

	// message <- "ping"

	// msg := <-message

	// fmt.Println(msg)
}