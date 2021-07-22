package main
import "fmt"

func main() {
	arr := [3]string{"watataku", "minami", "yoda"}
	fmt.Println("配列：", arr[1] ,"\n")
	fmt.Println("-----------スライス------------")
	var srice[]string
	srice = append(srice, "a")
	srice = append(srice, "b")
	srice = append(srice, "c")
	for i,s := range srice {
		fmt.Println(i, ":", s)
	}
	fmt.Println("------------マップ------------")
	maps := map[int]string{
		1 : "takuya",
		2 : "mai",
		3 : "nanase",
	}
	for key, value := range maps {
		fmt.Println(key, ":", value)
	}

	delete(maps, 1)
	fmt.Println("キー１の値を消しました")
	for key, value := range maps {
		fmt.Println(key, ":",value)
	}
}

