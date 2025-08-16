package main

import "fmt"

type paymenter interface{
	pay(amount float32)
}

type payment struct{
	gateway paymenter
}

func (p payment) makePayment(amount float32)  {
	p.gateway.pay(amount)
}

type razorpay struct{}

func (r razorpay) pay(amount float32)  {
	fmt.Println("Making Payment using razorpay", amount)
}

// type stripe struct{}

// func (s stripe) pay(amount float32)  {
// 	fmt.Println("Making Payment using stripe", amount)
// }

type fakePayment struct{}

func (f fakePayment) pay(amount float32)  {
	fmt.Println("Making Payment using fake testing", amount)
}



func main()  {
	// stripePaymentGw := stripe{}
	// razorpayPaymentGw := razorpay{}
	fakePaymentGw := fakePayment{}
	newPayment := payment{
		gateway: fakePaymentGw,
	}

	newPayment.makePayment(100)
}