package main

import (
"fmt"
)
//空接口表示为 interface{}。由于空接口没有方法，因此所有类型都实现了空接口
func describe2(i interface{}) {
	fmt.Printf("Type = %T, value = %v\n", i, i)
}

func main() {
	s := "Hello World"
	describe2(s)
	i := 55
	describe2(i)
	strt := struct {
		name string
	}{
		name: "Naveen R",
	}
	describe2(strt)
}
/***
空接口
没有包含方法的接口称为空接口。
空接口表示为 interface{}。由于空接口没有方法，因此所有类型都实现了空接口
在上面的程序的第 7 行，describe(i interface{}) 函数接收空接口作为参数，因此，可以给这个函数传递任何类型。

在第 13 行、第 15 行和第 21 行，我们分别给 describe 函数传递了 string、int 和 struct。该程序打印：
*/