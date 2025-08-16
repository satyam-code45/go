package main

import "fmt"

func main() {
	nums := []int{6, 7, 8}

	// for i := 0; i < len(nums); i++ {
	// 	fmt.Println(nums[i])
	// }

	sum := 0

	for i, num := range nums {
		fmt.Println(num, i)
		sum += num
	}

	fmt.Println(sum)

	m := map[string]string{"name ": "satyam", "age" : "21"}

	for i, value := range m{
		fmt.Println(i,value)
	}
}
