package main

import (
	"fmt"
	"time"
)

func main() {
	i := 1

	switch i {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	case 4:
		fmt.Println("Four")
	case 5:
		fmt.Println("Five")
	default:
		fmt.Println("Hello Go!")
	}

	//multiple condition switch

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's a Weekend!")
	default:
		fmt.Println("It's a working day!")
	}
}
