package main

import (
	"fmt"
	"time"
	// "math/rand"
	// "time"
)

//channels used to communicate between two go routines.

//send
// func processNum(numChan chan int) {

// 	for num := range numChan{
// 		fmt.Println("Processing Number: ", num)
// 		time.Sleep(time.Second)
// 	}

// }

// receive
// func sum(result chan int, num1 int, num2 int) {
// 	numResult := num1 + num2

// 	result <- numResult
// }

// func task(done chan bool) {
// 	defer func() {
// 		done <- true
// 	}()

// 	fmt.Println("Processing......")

// }

func emailSender(emailChan chan string, done chan bool)  {
	defer func ()  {
		done <- true
	}()
	for email := range emailChan{
		fmt.Println("Sending Email to: ", email)
		time.Sleep(time.Second)
	}
}


func main() {

	emailChan := make(chan string, 100)
	done := make(chan bool)

	go emailSender(emailChan,done) 

	for i := 0; i < 10; i++ {
		emailChan <- fmt.Sprintf("%d@gamil.com",i)
	}

	fmt.Println("Done Sending!")

	close(emailChan)
	<-done










	// done := make(chan bool)
	// go task(done)

	// <-done //block

	// result := make(chan int)

	// go sum(result, 4, 5)

	// res := <-result

	// fmt.Println(res)

	//send
	// numChan := make(chan int)

	// go processNum(numChan)

	// for {
	// 	numChan <- rand.Intn(100)
	// }

}
