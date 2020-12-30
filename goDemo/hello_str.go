package main

import (
	"fmt"
	"unicode/utf8"
)

/****
Go 经典入门系列 14：字符串
由于字符串是一个字节切片,所以我们可以获取字符串的每一个字节。
//然后我们用了一个 for 循环以 16 进制的形式打印这些字节。
%x 格式限定符用于指定 16 进制编码,上面的程序输出
48 65 6c 6c 6f 20 57 6f 72 6c 64,
这些打印出来的字符是 "Hello World" 以 Unicode UTF-8 编码
 */


func printBytes(s string) {
	for i:= 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
}
func printChars(s string) {
	for i:= 0; i < len(s); i++ {
		fmt.Printf("%c ",s[i])
	}
}

func printChars2(s string) {
	runes := []rune(s)
	fmt.Println()
	fmt.Println("---------------------runes------------------")
	for i:= 0; i < len(runes); i++ {
		fmt.Printf("%c ",runes[i])
	}
}

func printCharsAndBytes(s string) {
	for index, rune := range s {
		fmt.Printf("%c starts at byte %d\n", rune, index)
	}
}
func main()  {
	name := "Hello World"
	printBytes(name)
	fmt.Printf("\n")
	printChars(name)
	fmt.Printf("\n")
	name = "Señor"
	printBytes(name)
	fmt.Printf("\n")
	printChars(name)


	/**
	  rune 是 Go 语言的内建类型，它也是 int32 的别称。
	  在 Go 语言中，rune 表示一个代码点。代码点无论占用多少个字节，
	  都可以用一个 rune 来表示。让我们修改一下上面的程序，用 rune 来打印字符。
	  **/

	/***
	  字符串的 for range 循环
	  **/
	name2 := "Señor"
	printChars2(name2)
	//用字节切片构造字符串
	byteSlice := []byte{0x43, 0x61, 0x66, 0xC3, 0xA9}
	str := string(byteSlice)
	fmt.Println(str)


	//用 rune 切片构造字符串
	runeSlice := []rune{0x0053, 0x0065, 0x00f1, 0x006f, 0x0072}
	str = string(runeSlice)
	fmt.Println(str)

	//字符串的长度
	word1 := "Señor"
	length(word1)
	word2 := "Pets"
	length(word2)

	//字符串是不可变的
	h := "hello"
	fmt.Println(mutate(h))
	//为了修改字符串，可以把字符串转化为一个 rune 切片。然后这个切片可以进行任何想要的改变，然后再转化为一个字符串
	fmt.Println(mutate2([]rune(h)))
}


func length(s string) {
	fmt.Printf("length of %s is %d\n", s, utf8.RuneCountInString(s))
}
/**
字符串是不可变的
 由于字符串是不可变的，因此这个操作是非法的。所以程序抛出了一个错误   cannot assign to s[0]。

**/
func mutate(s string)string {
	//s[0] = 'a'//any valid unicode character within single quote is a rune
	return s
}
/**
为了修改字符串，可以把字符串转化为一个 rune 切片。然后这个切片可以进行任何想要的改变，然后再转化为一个字符串

mutate 函数接收一个 rune 切片参数，它将切片的第一个元素修改为 'a'，然后将 rune 切片转化为字符串，并返回该字符串。
**/

func mutate2(s []rune) string {
	s[0] = 'a'
	return string(s)
}