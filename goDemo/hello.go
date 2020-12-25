package main

/**
GO 基础赋值
 */
import "fmt"

func main() {
	/**
	  声明变量
	**/
	// 声明单个变量
	var age int // 变量声明
	fmt.Println("my age is", age)
	//类型推断（Type Inference）
	var age2 = 29 // 可以推断类型

	fmt.Println("my age is", age2)

	fmt.Printf("hello, world\n")
	//Go 能够通过一条语句声明多个变量。
	//声明多个变量的语法是 var name1, name2 type = initialvalue1, initialvalue2。

	var width, heigh int = 100, 50 // 声明多个变量
	fmt.Println("width is", width, "heigh is", heigh)
	/***
	在有些情况下，我们可能会想要在一个语句中声明不同类型的变量。其语法如下：
	*/
	var (
		name   = "naveen"
		age3    = 29
		height int
	)
	fmt.Println("my name is", name, ", age is", age3, "and height is", height)

	//简短声明
	name1, age4 := "naveen", 29 // 简短声明

	fmt.Println("my name is", name1, "age is", age4)
	/**
	简短声明的语法要求 := 操作符的左边至少有一个变量是尚未声明的。考虑下面的程序：
	*/
	a, b := 20, 30 // 声明变量a和b
	fmt.Println("a is", a, "b is", b)
	b, c := 40, 50 // b已经声明，但c尚未声明
	fmt.Println("b is", b, "c is", c)
	b, c = 80, 90 // 给已经声明的变量b和c赋新值
	fmt.Println("changed b is", b, "c is", c)
}
