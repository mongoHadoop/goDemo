package main

import "fmt"
/**
类型选择（Type Switch）
类型选择用于将接口的具体类型与很多 case 语句所指定的类型进行比较。
它与一般的 switch 语句类似。唯一的区别在于类型选择指定的是类型，而一般的 switch 指定的是值。

类型选择的语法类似于类型断言。类型断言的语法是 i.(T)，
而对于类型选择，类型 T 由关键字 type 代替。下面看看程序是如何工作的。

*/
func findType(i interface{}) {
	switch i.(type) {
	case string:
		fmt.Printf("I am a string and my value is %s\n", i.(string))
	case int:
		fmt.Printf("I am an int and my value is %d\n", i.(int))
	default:
		fmt.Printf("Unknown type\n")
	}
}
func main() {
	findType("Naveen")
	findType(77)
	findType(89.98)
}
