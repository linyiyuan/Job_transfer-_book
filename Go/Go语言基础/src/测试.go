package main

import "fmt"

func main() {
	var a = []int{1, 2, 3, 4, 5}

	a = append(a[:2], append([]int{6, 7, 8}, a[2:]...)...)

	fmt.Println(a)
}
