package main

/**
GO 基础赋值
 */
import (
	"fmt"
)

func main() {
	num := 10
	/**
	  if  condition {
	  }
	*/
	if num % 2 == 0 { //checks if number is even
		fmt.Println("the number is even")
	}  else {
		fmt.Println("the number is odd")
	}
/**
  if statement; condition {
  }
 */
	if num := 10; num % 2 == 0 { //checks if number is even
		fmt.Println(num,"is even")
	}  else {
		fmt.Println(num,"is odd")
	}
	/***
	  在上面的程序中，
	  num 在 if 语句中进行初始化，num 只能从 if 和 else 中访问。
		也就是说 num 的范围仅限于 if else 代码块。
		如果我们试图从其他外部的 if 或者 else 访问 num,编译器会不通过。
	*/
}
