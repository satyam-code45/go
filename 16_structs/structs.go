package main

import (
	"fmt"
	"time"
)

type order struct{
	id string
	amount float32
	status string
	createdAt time.Time
}

func main()  {
	myOrder := order{
		id: "1",
		amount: 45.00,
		status: "received",
	}

	myOrder.createdAt = time.Now()

	fmt.Println(myOrder)

	fmt.Println(myOrder.amount)
}