package main

import "fmt"

// ーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーー
// ーーーーーーーーーーーーーーーーーーーーー①変数定義後にフィールドを設定する方法ーーーー
// ーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーー

// type Person struct {
//    firstName string 
//    age int
// }

// func main(){
//   var mike Person
//   mike.firstName = "Mike"
//   mike.age = 20
//   fmt.Println(mike.firstName, mike.age) //=> Mike 20
// }

// ーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーー
// ーーーーーーーーーーーーーーーーーーーーー②{}で順番にフィールドの値を渡す方法ーーーー
// ーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーー

// type Person struct {
//    firstName string 
//    age int
// }

// func main(){
//   bob := Person{"Bob", 30}
//   fmt.Println(bob.firstName, bob.age) //=>Bob 30
// }

// ーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーー
// ーーーーーーーーーーーーーーーーーーーーー③フィールド名を ： で指定する方法ーーーー
// ーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーー

// type Person struct {
//    firstName string 
//    age int
// }

// func main(){
//   sam := Person{age: 15, firstName: "Sam"}
//   fmt.Println(sam.firstName, sam.age) //=>Sam 15
// }

// ーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーー
// ーーーーーーーーーーーーーーーーーーーーー④初期化関数を作成することで初期化をするーーーー
// ーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーー

type Person struct {
   firstName string 
   age int
}

func newPerson(firstName string, age int) *Person{
     person := new(Person)
     person.firstName = firstName
     person.age = age
     return person
}

func main(){
  var jen *Person = newPerson("Jennifer", 40)
  fmt.Println(jen.firstName, jen.age) //=>Jennifer 40
}
