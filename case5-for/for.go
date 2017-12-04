package main

import "fmt"

func main() {
	var sum int = 0
	for index := 0; index < 10; index++ {
		sum += index
	}
	fmt.Println("加總為：", sum)
}
