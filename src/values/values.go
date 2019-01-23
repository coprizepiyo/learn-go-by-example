package main

import "fmt"

func main() {
	// 字符串可以用 + 号连接
	fmt.Println("go" + "lang")

	// 整型和浮点型
	fmt.Println("1 + 1 = ", 1 + 1)
	fmt.Println("7.0 / 3.0 = ", 7.0 / 3.0)

	// 布尔值和布尔运算符
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}