package main

import "fmt"

// Greeter comment 入れないとエラーになる
type Greeter interface {
	Greet(name string) string
}

// Male 男性
type Male struct {
	name string
}

// Female 女性
type Female struct {
	name string
}

func (sex Male) Greet(name string) string {
	return "こんにちは！！" + name + "くん！！"
}
func (sex Female) Greet(name string) string {
	return "こんにちは！！" + name + "ちゃん！！"
}

func main() {
	var grt Greeter

	var male Male
	var female Female

	grt = male

	fmt.Println(grt.Greet("tyankatsu"))

	grt = female

	fmt.Println(grt.Greet("tyankatsu"))
}
