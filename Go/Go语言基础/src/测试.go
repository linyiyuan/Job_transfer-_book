package main

import "fmt"

func main() {
	var a [3]int
	fmt.Println(a[0])
	fmt.Println(a[len(a)-1])

	for k, v := range a {
		fmt.Printf("键为：%v, 值为: %v", k, v)
	}

}
