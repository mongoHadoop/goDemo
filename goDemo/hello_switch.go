package main
import (
	"fmt"
)

func main() {
	finger := 4
	switch finger {
	case 1:
		fmt.Println("Thumb")
	case 2:
		fmt.Println("Index")
	case 3:
		fmt.Println("Middle")
	case 4:
		fmt.Println("Ring")
	case 5:
		fmt.Println("Pinky")

	}
	letter := "i"
	switch letter {
	case "a", "e", "i", "o", "u": // 一个选项多个表达式
		fmt.Println("vowel")
	default:
		fmt.Println("not a vowel")
	}
	//无表达式的 switch
	/***
	在 switch 语句中，表达式是可选的，可以被省略。如果省略表达式，则表示这个 switch 语句等同于 switch true，并且每个 case 表达式都被认定为有效，相应的代码块也会被执行。
	 */
	num := 75
	switch { // 表达式被省略了
	case num >= 0 && num <= 50:
	fmt.Println("num is greater than 0 and less than 50")
	case num >= 51 && num <= 100:
	fmt.Println("num is greater than 51 and less than 100")
	case num >= 101:
	fmt.Println("num is greater than 100")
	}

}