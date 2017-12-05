package main

import "fmt"

type Rectangle struct {
	width, height float64
}

func main() {
	fmt.Println("你媽媽")
}

func area(r Rectangle) float64 {
	return r.width * r.height
}
