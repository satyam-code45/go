package main

import (
	"fmt"
	"time"
)

func task(id int) {
	fmt.Println("doing task: ", id)
}

func main() {
	for i := 0; i <= 10; i++ {
		go task(i)
	}

	fmt.Println("Hello!")
	time.Sleep(time.Second * 1)

}
