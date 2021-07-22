package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var name string = "綿野拓也"

	//実際の文字数
	fmt.Println(utf8.RuneCountInString(name))

	age := 22
	if age >= 20 {
		fmt.Println(name, "は成人ですのでお酒が飲めます")
	} else {
		fmt.Println(name, "は未成年なのでお酒が飲めません")
	}
}