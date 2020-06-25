package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var isTrue bool = true

	fmt.Println(isTrue)

	var i int = 12345
	var int64 int64 = int64(i)

	fmt.Println(int64)

	var str string = "あいうえ"
	str += "お"

	fmt.Println(str)

	var en string = "golang"
	var ja string = "Go言語"

	fmt.Println(en, " length: ", len(en))
	fmt.Println(ja, " length: ", utf8.RuneCountInString(ja))
}
