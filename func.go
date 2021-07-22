package main
import "fmt"

func main() {
	hello()
	n := plus(1, 2)
	fmt.Println(n)
}

func hello() {
	fmt.Println("HelloWould\n")
}

func plus (i, j int) int {
	n := i + j
	return n
}