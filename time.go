package main

import (
	"fmt"
	"time"
)

func main()  {
	day := time.Now()
	//"2006年1月2日15時04分05秒"以外の時刻は正しいフォーマット文字列としては認識されない。
	const layout2 = "2006-01-02 15:04:05"
	fmt.Println(day.Format(layout2))
}