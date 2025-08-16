package main

import (
	"fmt"
)

//instead of int | string -> comparable
func printSlice[T comparable](items []T) {
	for _, item := range items {
		fmt.Println(item)
	}
}

func main() {
	nums := []int{1, 2, 3, 4, 5}

	names := []string{"c", "golang", "c++", "pyhton"}
	printSlice(nums)
	printSlice(names)
}
