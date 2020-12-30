package main

import "fmt"

/**
Go 经典入门系列 12：可变参数函数
语法
如果函数最后一个参数被记作 ...T，这时函数可以接受任意个 T 类型参数作为最后一个参数。
请注意只有函数的最后一个参数才允许是可变的。
append 函数可以接受不同数量的参数。
func append(slice []Type, elems ...Type) []Type
上面是 append 函数的定义。在定义中 elems 是可变参数。这样 append 函数可以接受可变化的参数。
 */

func find(num int, nums ...int) {
	fmt.Printf("type of nums is %T\n", nums)
	found := false
	for i, v := range nums {
		if v == num {
			fmt.Println(num, "found at index", i, "in", nums)
			found = true
		}
	}
	if !found {
		fmt.Println(num, "not found in ", nums)
	}
	fmt.Printf("\n")
}
/**
在上面程序中 func find(num int, nums ...int)
中的 nums 可接受任意数量的参数。在 find 函数中，参数 nums 相当于一个整型切片。
find 函数中的可变参数是 89，90，95 。find 函数接受一个 int 类型的可变参数。
因此这三个参数被编译器转换为一个 int 类型切片 int []int{89, 90, 95} 然后被传入 find 函数。

**/

func main() {
	find(89, 89, 90, 95)
	find(45, 56, 67, 45, 90, 109)
	find(78, 38, 56, 98)
	find(87)
	//给可变参数函数传入切片
	nums := []int{89, 90, 95}
	//find(89, nums)//这种情况下无法通过编译，编译器报出错误  cannot use nums (type []int) as type int in argument to find 。
	/**
	为什么无法工作呢？原因很直接，find 函数的说明如下，由可变参数函数的定义可知，nums ...int 意味它可以接受 int 类型的可变参数。
	nums 作为可变参数传入 find 函数。前面我们知道，这些可变参数参数会被转换为 int 类型切片然后在传入 find 函数中。
	但是在这里 nums 已经是一个 int 类型切片，编译器试图在 nums 基础上再创建一个切片，像下面这样
	find(89, []int{nums})
	这里之所以会失败是因为 nums 是一个 []int类型 而不是 int类型。
	那么有没有办法给可变参数函数传入切片参数呢？答案是肯定的。
	有一个可以直接将切片传入可变参数函数的语法糖，你可以在在切片后加上 ... 后缀。如果这样做，切片将直接传入函数，不再创建新的切片
	***/
	find(89, nums...)
	/**
	我们使用了语法糖 ... 并且将切片作为可变参数传入 change 函数。
	正如前面我们所讨论的，如果使用了 ... ，welcome 切片本身会作为参数直接传入，不需要再创建一个新的切片。这样参数 welcome 将作为参数传入 change 函数

	*/
	welcome := []string{"hello", "world"}
	change(welcome...)
	fmt.Println(welcome)
	change2(welcome...)
	fmt.Println(welcome)
}

func change(s ...string) {
	s[0] = "Go"
}
func change2(s ...string) {
	s[0] = "Go"
	s = append(s, "playground")
	fmt.Println(s)
}