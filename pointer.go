package main

import "fmt"


// func main() {
//     var hoge int = 10 //int型の変数　
//     var bow *int = &hoge //int型へのポイント型

//     fmt.Println(hoge) //=> 10
//     fmt.Println(bow) //=> 0x10414020   
// }


func main() {
    var hoge int = 10
    var bow *int = &hoge
    fmt.Println(*bow) //=> 10
}


