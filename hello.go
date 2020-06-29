package main

import "fmt"

// type myType int

// func (value myType) println() {
// 	fmt.Println(value)
// }

// func main() {
// 	var z myType = 12345

// 	z.println()
// }

type Person struct {
	name string
	age  int
}

func (person Person) greet() {
	fmt.Println("Hello" + person.name)
}

func main() {
	person := Person{
		name: "tyankatsu",
		age:  26,
	}

	person.greet()
}
