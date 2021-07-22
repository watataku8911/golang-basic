package main

import "fmt"

type User struct {
	email string
	Name string
	Age int
}

func newUser(name string, age int, email string) *User {
	u := new(User)
	u.Name = name
	u.Age = age
	u.email = email
	return u
}

func (u User) greet() string {
	return "hello," + u.Name;
}

func main() {
	var user *User = newUser("taro",30, "hoge@example.ac.jp")
	fmt.Println(user.Name)
	fmt.Println(user.Age)
	fmt.Println(user.email)
	fmt.Println(user.greet())
}