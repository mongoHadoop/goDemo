package main

import (
"fmt"
)

func assert(i interface{}) {
s := i.(int) //get the underlying int value from i
fmt.Println(s)
}
func assert2(i interface{}) {
	v, ok := i.(int)
	fmt.Println(v, ok)
}
func main() {
	var s interface{} = 56
 	var s2 interface{} = "Steven Paul"
	assert(s)
 	//该程序会报错：panic: interface conversion: interface {} is string, not int.。
 	//assert(s2)
 	/**
 	要解决该问题，我们可以使用以下语法：
	v, ok := i.(T)
	如果 i 的具体类型是 T，那么 v 赋值为 i 的底层值，而 ok 赋值为 true。
	如果 i 的具体类型不是 T，那么 ok 赋值为 false，v 赋值为 T 类型的零值，此时程序不会报错。
 	*/
	assert2(s2)

}
/***
类型断言
类型断言用于提取接口的底层值（Underlying Value）。
在语法 i.(T) 中，接口 i 的具体类型是 T，该语法用于获得接口的底层值。

在第 12 行，s 的具体类型是 int。在第 8 行，我们使用了语法 i.(int) 来提取 i 的底层 int 值。该程序会打印 56。
在上面程序中，如果具体类型不是 int，会发生什么呢？接下来看看。
 */
