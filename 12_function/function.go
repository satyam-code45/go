package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func getLanguages() (string,string,string){
	return "Go Lang", "c", "c++" 
}

func main() {
	result := add(3, 4)

	fmt.Println(result)

	fmt.Println(getLanguages())
}
