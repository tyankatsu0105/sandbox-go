package main

import "fmt"

func main() {
	type myInteger int

	var i myInteger = 12345

	fmt.Println(i)

	type myStruct struct {
		a int
		b int
	}

	var integer int = 12345
	var str string = string(integer)

	fmt.Println(str)
}
