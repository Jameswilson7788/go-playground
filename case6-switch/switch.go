package main

import "fmt"

func main() {
	i := 5
	switch i {
	case 1:
		fmt.Println("這是1")
	case 2, 3, 4, 5:
		fmt.Println("這是2,3,4,5")
	case 10:
		fmt.Println("這是10")
	}
}
