package main

import "fmt"

/**
for 是 Go 语言唯一的循环语句。Go 语言中并没有其他语言比如 C 语言中的 while 和 do while 循环。
for initialisation; condition; post {
}
初始化语句只执行一次。循环初始化后，将检查循环条件。如果条件的计算结果为 true ，
则 {} 内的循环体将执行，接着执行 post 语句。post 语句将在每次成功循环迭代后执行。
在执行 post 语句后，条件将被再次检查。如果为 true, 则循环将继续执行, 否则 for 循环将终止。
 */
func main() {
	for i := 1; i <= 10; i++ {
		fmt.Printf(" %d",i)
	}
	for i := 1; i <= 10; i++ {
		if i > 5 {
			break //loop is terminated if i > 5
		}
		fmt.Printf("%d ", i)
	}
	fmt.Printf("\nline after for loop")
	i:=0
	/// for 循环的三部分,初始化语句,条件语句,POST语句都是可选
	for ;i<=10;{
		fmt.Printf("%d ",i)
		i+=2
	}
	//无限循环
	for {
		fmt.Println("HEllo")
	}
}
