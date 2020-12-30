package main

import "fmt"

/**
指针的声明
指针变量的类型为 *T，该指针指向一个 T 类型的变量。
*/

func main()  {
	b := 255
	var a *int = &b
	fmt.Printf("Type of a is %T\n", a)
	fmt.Println("address of b is", a)

	/**
	  & 操作符用于获取变量的地址。上面程序的第 9 行我们把 b 的地址赋值给 *int 类型的 a。我们称 a 指向了 b。当我们打印 a 的值时，会打印出 b 的地址。程序将输出：
	  ***/
	//  指针的解引用
	a2 := &b
	fmt.Println("address of b is", a2)
	//值访问
	fmt.Println("value of b is", *a2)

	fmt.Println("value of a before function call is",b)
	b1 := &b
	change12(b1)
	fmt.Println("value of a after function call is", b)

	/***
	  不要向函数传递数组的指针，而应该使用切片
	  ***/
	a1 := [3]int{89, 90, 91}
	modify(&a1)
	fmt.Println(a1)
	modify2(&a1)
	fmt.Println(a1)
	//所以别再传递数组指针了，而是使用切片吧
	modify3(a1[:])
	fmt.Println(a1)

	//Go 不支持指针运算
	//b2 := [...]int{109, 110, 111}
	//p := &b2
	//p++ //Go 不支持指针运算

}
//向函数传递指针参数
func change12(val *int) {
	*val = 55
}
/***
假如我们想要在函数内修改一个数组，并希望调用函数的地方也能得到修改后的数组，
一种解决方案是把一个指向数组的指针传递给这个函数。
不要向函数传递数组的指针，而应该使用切片
**/
func modify(arr *[3]int) {
	(*arr)[0] = 90
}
/**
a[x] 是 (*a)[x] 的简写形式，因此上面代码中的 (*arr)[0]
可以替换为 arr[0]。下面我们用简写形式重写以上代码。
这种方式向函数传递一个数组指针参数，并在函数内修改数组。
尽管它是有效的，但却不是 Go 语言惯用的实现方式。我们最好使用切片[8]来处理。
**/
func modify2(arr *[3]int) {
	arr[0] = 910
}



func modify3(sls []int) {
	sls[0] = 1001
}