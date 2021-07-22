package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	//fmt.Println("HelloWold");
	exe, err := os.Executable()
	if err!= nil {
		fmt.Println(err)
	}
	fmt.Println(exe)
	fmt.Println(filepath.Dir(exe))

	p, _ := os.Getwd()
	fmt.Println(p)
}