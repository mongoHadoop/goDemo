package main

import (
	"fmt"
	"goDemo/goDemo/geometry/rectangle" // 导入自定义包
	"goDemo/goDemo/geometry/rectangle2"
	//使用空白标识符（Blank Identifier）
	/***
	有时候我们导入一个包，只是为了确保它进行了初始化，而无需使用包中的任何函数或变量。
	例如，我们或许需要确保调用了 rectangle 包的 init 函数，
	而不需要在代码中使用它。这种情况也可以使用空白标识符，如下所示.*/
	_"goDemo/goDemo/geometry/rectangle2" // 导入自定义包
	"log"
)
/***
导入了包，却不在代码中使用它,这在 Go 中是非法的,当这么做时,编译器是会报错的.
其原因是为了避免导入过多未使用的包,从而导致编译时间显著增加,将 geometry.go 中的代码替换为如下代码：
 */
var _ = rectangle2.Area2 // 错误屏蔽器
/***
var _ = rectangle.Area
这一行屏蔽了错误。我们应该了解这些错误屏蔽器（Error Silencer）的动态，
在程序开发结束时就移除它们，包括那些还没有使用过的包。由此建议在 import 语句下面的包级别范围中写上错误屏蔽器。
 */

/*
 * 1. 包级别变量
 */
var rectLen, rectWidth float64 = 6, 7

/*
*2. init 函数会检查长和宽是否大于0
 */
func init() {
	println("main package initialized")
	if rectLen < 0 {
		log.Fatal("length is less than zero")
	}
	if rectWidth < 0 {
		log.Fatal("width is less than zero")
	}
}
/**
main 包的初始化顺序为：

首先初始化被导入的包。因此，首先初始化了 rectangle 包。
接着初始化了包级别的变量 rectLen 和 rectWidth。
调用 init 函数。
最后调用 main 函数。
 */

func main() {
	var rectLen, rectWidth float64 = 6, 7
	fmt.Println("Geometrical shape properties")
	/*Area function of rectangle package used*/
	//Printf 内的格式说明符 %.2f 会将浮点数截断到小数点两位。应用程序的输出为：
	fmt.Printf("area of rectangle %.2f\n", rectangle.Area(rectLen, rectWidth))
	/*Diagonal function of rectangle package used*/
	fmt.Printf("diagonal of the rectangle %.2f ", rectangle.Diagonal(rectLen, rectWidth))
	//导出名字（Exported Names）
	/**
	我们将 rectangle 包中的函数 Area 和 Diagonal 首字母大写.在 Go 中这具有特殊意义.
	在 Go 中,任何以大写字母开头的变量或者函数都是被导出的名字.
	其它包只能访问被导出的函数和变量.在这里,我们需要在 main 包中访问 Area 和 Diagonal 函数,因此会将它们的首字母大写
	在 rectprops.go 中，如果函数名从 Area(len, wid float64) 变为 area(len, wid float64)，并且在 geometry.go 中，
	rectangle.Area(rectLen, rectWidth) 变为 rectangle.area(rectLen, rectWidth)，
	则该程序运行时，编译器会抛出错误 geometry.go:11: cannot refer to unexported name rectangle.area。因为如果想在包外访问一个函数，它应该首字母大写。
	**/


	/***
	所有包都可以包含一个 init 函数。init 函数不应该有任何返回值类型和参数，在我们的代码中也不能显式地调用它。init 函数的形式如下：
	func init() {
	}
	 */

}