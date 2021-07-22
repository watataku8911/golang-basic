package main
import (
	"fmt"
	"os"
	"log"

)

func main() {
	file, err := os.OpenFile("test.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND , 0666)
	if err != nil {
		log.Fatal("エラー：ファイルがありません")
	}
	defer file.Close()
	/*
	*書き込み
	*/
	_, err = file.WriteString("ばか\n")
	if err != nil {
		panic(err)
	} else {
		fmt.Println("書き込みました")
	}
	
}