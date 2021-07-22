package main

import (
	"./user" //自作パッケージ
    "fmt"
)
func main() {
	name := user.Name("hoge")
	fmt.Println(name)

	email := user.Email("hoge@xx.xx.jp")
	fmt.Println(email)
}