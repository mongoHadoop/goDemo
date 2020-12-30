package main

import "fmt"

/**
数组是同一类型元素的集合。
例如，整数集合 5,8,9,79,76 形成一个数组。
Go 语言中不允许混合不同类型的元素，
例如包含字符串和整数的数组。（译者注：当然，如果是 interface{} 类型数组，可以包含任意类型）

数组的声明
一个数组的表示形式为 [n]T,n 表示数组中元素的数量,T 代表每个元素的类型。元素的数量 n 也是该类型的一部分（稍后我们将详细讨论这一点）

 */
func main() {
	var a [3]int
	a1 := [3]int{12, 78, 50} // short hand declaration to create array
	fmt.Println(a)
	fmt.Println(a1)
	//在简略声明中，不需要将数组中所有的元素赋值。
	a2:= [3]int{12}
	fmt.Println(a2)

	//你甚至可以忽略声明数组的长度，并用 ... 代
	a3 := [...]int{12, 78, 50} // ... makes the compiler determine the length
	fmt.Println(a3)
	//数组是值类型 Go中的数组不是引用类型
	//数组作为参数传递给函数时，它们是按值传递，而原始数组保持不变。
	a4:=[...]string{"USA","CHINA","INDIA","Gemany","France"}
	b:=a4
	b[0]="Sia"
	fmt.Println("a4 is ",a4)
	fmt.Println("b is ",b)
//同样，当数组作为参数传递给函数时，它们是按值传递，而原始数组保持不变。
	num := [...]int{5, 6, 7, 8, 8}
	fmt.Println("before passing to function ", num)
	changeLocal(num) //num is passed by value
	fmt.Println("after passing to function ", num)


	//使用 range 迭代数组
	a6 := [...]float64{67.7, 89.8, 21, 78}
	sum := float64(0)
	for i, v := range a6 {//range returns both the index and value
		fmt.Printf("%d the element of a is %.2f\n", i, v)
		sum += v
	}
	fmt.Println("\nsum of all elements of a",sum)
	/***
	如果你只需要值并希望忽略索引，则可以通过用 _ 空白标识符替换索引来执行。
	*/

	a7 := [3][2]string{
		{"lion", "tiger"},
		{"cat", "dog"},
		{"pigeon", "peacock"}, // this comma is necessary. The compiler will complain if you omit this comma
	}
	printarray(a7)
	var b1[3][2]string
	b1[0][0] = "apple"
	b1[0][1] = "samsung"
	b1[1][0] = "microsoft"
	b1[1][1] = "google"
	b1[2][0] = "AT&T"
	b1[2][1] = "T-Mobile"
	fmt.Printf("\n")
	printarray(b1)
}
/**
同样，当数组作为参数传递给函数时，它们是按值传递，而原始数组保持不变。
 */
func changeLocal(num [5]int) {
	num[0] = 55
	fmt.Println("inside function ", num)
}
/*多维数组
到目前为止我们创建的数组都是一维的，Go 语言可以创建多维数组。
*/

func printarray(a [3][2]string) {
	for _, v1 := range a {
		for _, v2 := range v1 {
			fmt.Printf("%s ", v2)
		}
		fmt.Printf("\n")
	}
}