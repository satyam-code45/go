package main

import "fmt"

func main() {
	//creating map
	// m := make(map[string]string)

	// m["name"] = "Satyam"

	m := map[string]string{"name": "Satyam"}

	fmt.Println(m)
	fmt.Println(m["name"])
	fmt.Println(len(m))

	m["age"] = "21"

	fmt.Println(m["age"])
	fmt.Println(len(m))
	fmt.Println(m)

	delete(m, "age")

	fmt.Println(m["age"])
	fmt.Println(len(m))
	fmt.Println(m)

	clear(m)
	fmt.Println(m)
}
