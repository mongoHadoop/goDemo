package main
/**
GO 变量类型 demo
有符号整型
int8：表示 8 位有符号整型大小：8 位范围：-128～127
int16：表示 16 位有符号整型大小：16 位范围：-32768～32767
int32：表示 32 位有符号整型大小：32 位范围：-2147483648～2147483647
int64：表示 64 位有符号整型大小：64 位范围：-9223372036854775808～
int：根据不同的底层平台（Underlying Platform），
表示 32 或 64 位整型。
除非对整型的大小有特定的需求，否则你通常应该使用 int 表示整型。
大小：在 32 位系统下是 32 位，
而在 64 位系统下是 64 位。
范围：在 32 位系统下是 -2147483648～2147483647，
而在 64 位系统是 -9223372036854775808～9223372036854775807。
 */

/**
无符号整型
uint8：表示 8 位无符号整型大小：8 位范围：0～255
uint16：表示 16 位无符号整型大小：16 位范围：0～65535
uint32：表示 32 位无符号整型大小：32 位范围：0～4294967295
uint64：表示 64 位无符号整型大小：64 位范围：0～18446744073709551615
uint：根据不同的底层平台，表示 32 或 64 位无符号整型。大小：
在 32 位系统下是 32 位，而在 64 位系统下是 64 位。范围：在 32 位系统下是 0～4294967295，而在 64 位系统是 0～18446744073709551615。
*/

/**
浮点型
float32：32 位浮点数float64：64 位浮点数


*/
/**
复数类型
complex64：实部和虚部都是 float32 类型的的复数。complex128：实部和虚部都是 float64 类型的的复数。
*/

import (
	"fmt"
	"unsafe"
)

func main() {
	a := true
	b := false
	fmt.Println("a:", a, "b:", b)
	c := a && b
	fmt.Println("c:", c)
	d := a || b
	fmt.Println("d:", d)
	/**
	在 Printf 方法中，使用 %T 格式说明符（Format Specifier），可以打印出变量的类型。Go 的
	unsafe[3] 包提供了一个 Sizeof[4] 函数，该函数接收变量并返回它的字节大小。
	unsafe 包应该小心使用，因为使用 unsafe 包可能会带来可移植性问题。不过出于本教程的目的，我们是可以使用的。
	 */
	fmt.Printf("type of a is %T, size of a is %d", a, unsafe.Sizeof(a)) // a 的类型和大小
	fmt.Printf("\ntype of b is %T, size of b is %d", b, unsafe.Sizeof(b)) // b 的类型和大小
	fmt.Println("")
	fmt.Println("---------------------------------------------------")
	a1, b1 := 5.67, 8.97
	fmt.Printf("type of a %T b %T\n", a1, b1)
	sum := a1 + b1
	diff := a1 - b1
	fmt.Println("sum", sum, "diff", diff)

	no1, no2 := 56, 89
	fmt.Println("sum", no1+no2, "diff", no1-no2)
	fmt.Println("")
	fmt.Println("---------------------------------------------------")
	/**
	复数类型
	complex64：实部和虚部都是 float32 类型的的复数。complex128：实部和虚部都是 float64 类型的的复数。
	内建函数 **complex**[7] 用于创建一个包含实部和虚部的复数。complex 函数的定义如下：
	*/
	c1 := complex(5, 7)
	//还可以使用简短语法来创建复数：
	c2 := 8 + 27i
	cadd := c1 + c2
	fmt.Println("sum:", cadd)
	cmul := c1 * c2
	fmt.Println("product:", cmul)

	/**
	  其他数字类型 byte 是 uint8 的别名。rune 是 int32 的别名。
	 */
	fmt.Println("")
	fmt.Println("---------------------------------------------------")
	/**
	string 类型 在 Golang 中，字符串是字节的集合
	 */
	first := "Naveen"
	last := "Ramanathan"
	name := first +" "+ last
	fmt.Println("My name is",name)
	fmt.Println("")
	fmt.Println("---------------------------------------------------")
	/**
	类型转换
	Go 有着非常严格的强类型特征。Go 没有自动类型提升或类型转换
	 */
	i1 := 55      //int
	j1 := 67.8    //float64
	//sum1 := i1 + j1 //不允许 int + float64 此是错误的
	sum1 := i1 + int(j1) //j is converted to int 此才是正确的
	fmt.Println(sum1)
	i2 := 10
	var j2 float64 = float64(i2) // 若没有显式转换，该语句会报错
	fmt.Println("j2", j2)
}