package main

import "fmt"

func sum(nums ...int) int  {
	total := 0

	for _, num := range nums {
		total += num
	} 

	return  total
}

func main()  {
	//variadic funciton -> a function in which we could pass any number of parameter
	fmt.Println(1,2,3,4,"hello")

	fmt.Println(sum(1,2,3,4,4))

	nums := []int{1,2,3,4,4}
	result := sum(nums...)
	fmt.Println(result)
}