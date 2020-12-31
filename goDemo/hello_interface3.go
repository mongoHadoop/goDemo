package main

import (
	"fmt"
)

type Test interface {
	Tester()
}

type MyFloat float64

func (m MyFloat) Tester() {
	fmt.Println(m)
}

func describe(t Test) {
	fmt.Printf("Interface type %T value %v\n", t, t)
}

func main() {
	var t Test
	f := MyFloat(89.7)
	t = f
	describe(t)
	t.Tester()
}
/**
接口的内部表示
我们可以把接口看作内部的一个元组 (type, value)。type 是接口底层的具体类型（Concrete Type），而 value 是具体类型的值。

Test 接口只有一个方法 Tester()，而 MyFloat 类型实现了该接口。
在第 24 行，我们把变量 f（MyFloat 类型）赋值给了 t（Test 类型）.
现在 t 的具体类型为 MyFloat，而 t 的值为 89.7。第 17 行的 describe 函数打印出了接口的具体类型和值。该程序输出：
***/