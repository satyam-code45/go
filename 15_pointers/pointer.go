package main

import "fmt"


//pass by value
// func changeNum(num int)  {
// 	num = 5;
// 	fmt.Println("InChangeNum: ", num)
// }

//pass by reference

func changeNum(num * int)  {
	*num = 5;
	fmt.Println("InChangeNum: ", *num)
}

func main(){
	num  := 1

	fmt.Println(num)

	changeNum(&num)

	fmt.Println(num)
}