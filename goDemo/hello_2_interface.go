package main

import "fmt"

type Describer_ interface {
	Describe()
}
type Person_1 struct {
	name string
	age  int
}
//结构体 Person 使用值接受者，实现了 Describer 接口。
func (p Person_1) Describe() { // 使用值接受者实现
	fmt.Printf("%s is %d years old\n", p.name, p.age)
}

type Address_ struct {
	state   string
	country string
}
//在 22 行，结构体 Address 使用指针接受者实现了 Describer 接口。
func (a *Address_) Describe() { // 使用指针接受者实现
	fmt.Printf("State %s Country %s", a.state, a.country)
}

func main() {
	var d1 Describer_
	p1 := Person_1{"Sam", 25}
	d1 = p1
	d1.Describe()
	p2 := Person_1{"James", 32}
	d1 = &p2
	d1.Describe()

	var d2 Describer_
	a := Address_{"Washington", "USA"}


	/* 如果下面一行取消注释会导致编译错误：
	   cannot use a (type Address) as type Describer
	   in assignment: Address does not implement
	   Describer (Describe method has pointer
	   receiver)这是因为在第 22 行，我们使用 Address 类型的指针接受者实现了接口 Describer，
	你应该会很惊讶，因为我们曾经学习过，使用指针接受者的方法，无论指针还是值都可以调用它。那么为什么第 45 行的代码就不管用呢？ */
	//d2 = a // 我们试图用 a 来赋值 d2,然而 a 属于值类型，它并没有实现 Describer 接口,其原因是：对于使用指针接受者的方法，用一个指针或者一个可取得地址的值来调用都是合法的。,
//第 47 行就可以成功运行，因为我们将 a 的地址 &a 赋值给了 d2。
	d2 = &a // 这是合法的
	// 因为在第 22 行，Address 类型的指针实现了 Describer 接口
	d2.Describe()

}
/***

在上面程序中的第 13 行，结构体 Person 使用值接受者，实现了 Describer 接口。
我们在讨论方法的时候就已经提到过，使用值接受者声明的方法，既可以用值来调用，也能用指针调用。不管是一个值，还是一个可以解引用的指针，调用这样的方法都是合法的。
p1 的类型是 Person，在第 29 行，p1 赋值给了 d1。由于 Person 实现了接口变量 d1，因此在第 30 行，会打印 Sam is 25 years old。
接下来在第 32 行，d1 又赋值为 &p2，在第 33 行同样打印输出了 James is 32 years old。棒棒哒。:)

在 22 行，结构体 Address 使用指针接受者实现了 Describer 接口。
在上面程序里，如果去掉第 45 行的注释，
我们会得到编译错误：
main.go:42: cannot use a (type Address) as type Describer in assignment:
Address does not implement Describer (Describe method has pointer receiver)。
这是因为在第 22 行，我们使用 Address 类型的指针接受者实现了接口 Describer，
而接下来我们试图用 a 来赋值 d2。
然而 a 属于值类型，它并没有实现 Describer 接口。
你应该会很惊讶，因为我们曾经学习过，使用指针接受者的方法，无论指针还是值都可以调用它。那么为什么第 45 行的代码就不管用呢？

其原因是：对于使用指针接受者的方法，用一个指针或者一个可取得地址的值来调用都是合法的。但接口中存储的具体值（Concrete Value）
并不能取到地址，因此在第 45 行，对于编译器无法自动获取 a 的地址，于是程序报错。
第 47 行就可以成功运行，因为我们将 a 的地址 &a 赋值给了 d2。
*/
