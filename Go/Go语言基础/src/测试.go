package main

import "fmt"

func main(){
	//var e float64 = 0.5468465
	//var a float64 = 1123
	//
	//fmt.Printf("%T, %.3f\n", a-e, a-e)
	//
	//
	//
	//i := true
	//
	//if i {
	//	fmt.Printf("%T, %.3f\n", a-e, a-e)
	//}
	//
	//var str string = "测"
	//var str1 string = "试"
	//
	//
	//println(str + str1)

	// 准备一个字符串类型
	var house = "Malibu Point 10880, 90265"
	// 对字符串取地址, ptr类型为*string
	ptr := &house
	// 打印ptr的类型
	fmt.Printf("ptr type: %T\n", ptr)
	// 打印ptr的指针地址
	fmt.Printf("address: %p\n", ptr)

	value := *ptr
	fmt.Printf("value type: %T\n", value)
	fmt.Printf("value : %s\n", value)
}
