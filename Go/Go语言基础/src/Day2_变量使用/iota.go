package main

import "fmt"

func main() {
	const (
		a = iota
		b = iota
		c = iota
	)

	fmt.Println(c)

	const (
		d = 1
		e
		f
		g = "aa"
		h
		i = 100
		j = iota
	)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(i)
	fmt.Println(j)
}
