package main

import "fmt"

func main() {
	var b int = add(1, 3)
	fmt.Print(b)
}

func add(a int, b int) int {
	return a + b
}
