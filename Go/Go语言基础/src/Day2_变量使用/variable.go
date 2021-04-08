package main

func main(){
	//带有类型以及值的变量声明
	var b string = "World"
	//不带类型的变量声明
	var a = "Hello"
	//不带值变量声明
	var c bool
	//只能在函数体里面使用的声明
	d := 10
	//简单变量声明
	println(a, b, c, d)


	//多变量声明
	//第一种方式 (以逗号分隔，声明与赋值分开，若不赋值，存在默认值)
	var name1, name2, name3 string
	name1, name2, name3 = "test1", "test2", "test3"
	println(name1, name2, name3)

	//第二种方式(直接赋值，下面的变量类型可以是不同的类型)
	var name4, name5, name6  = "1", "2", "3"
	println(name4, name5, name6)

	//第三种方式 （集合类型）
	var (
		name7 string
		name8 int
		name9 bool
	)
	name1 = "string"
	println(name7, name8, name9)

	//在同一个作用域中，已存在同名的变量，则之后的声明初始化，则退化为赋值操作。但这个前提是，最少要有一个新的变量被定义，
	//且在同一作用域，例如，下面的y就是新定义的变量
	x := 100
	println(x)
	x, y := 200, "string2"
	println(x, y)
}