package main

import "fmt"

func main() {
	// var可以声明一个变量
	var a = "initial"
	fmt.Println(a)

	// var也可以声明多个变量
	var b, c, d int = 1, 2, 3
	fmt.Println(b, c, d)

	// 不用指明类型，通过初始化值来推导
	var e = true
	fmt.Println(e)

	// 没有初始化的声明默认都是零值
	var f string
	var g int
	var h bool
	fmt.Println("sring => " + f, "\nint => ", g, "\nbool => ", h)

	// 还有一种定义变量的方式
	i := "short" //相当于 var i string = "short"
	fmt.Println(i)
}