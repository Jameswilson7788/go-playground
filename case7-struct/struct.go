package main

import "fmt"

type user struct {
	name    string
	age     int
	address string
}

func Older(p1, p2 user) (user, int) {
	if p1.age > p2.age {
		return p1, p1.age - p2.age
	}
	return p2, p2.age - p1.age
}

func main() {

	Tom := user{
		name:    "Tom",
		age:     20,
		address: "我家",
	}

	var James user
	James.name = "James"
	James.age = 25
	James.address = "我家家"

	tbolder, tbdiff := Older(Tom, James)
	fmt.Printf("%s , %d", tbolder.name, tbdiff)
}
