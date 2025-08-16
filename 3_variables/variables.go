package main

import "fmt"

func main() {
	//strings
	var name1 = "Hello! Golang"
	var name string = "Golang"

	fmt.Println(name1)
	fmt.Println(name)

	//int
	var age1 = 18
	var age int = 23

	fmt.Println(age1)
	fmt.Println(age)

	//bool
	var truthy bool = true
	fmt.Println(truthy)
	truthy = false
	fmt.Println(truthy)

	//shorthand syntax
	names := "Satyam"
	fmt.Println(names)

}
