package main

import (
	"fmt"
	"goDemo/goDemo/structs/computer"
)
type name struct {
	firstName string
	lastName string
}
//如果结构体包含不可比较的字段，则结构体变量也不可比较。
type image struct {
	data map[int]int
}

func main() {
	var spec computer.Spec
	spec.Maker = "apple"
	spec.Price = 50000
	//如果我们试图访问未导出的字段 model，编译器会报错。将 main.go 的内容替换为下面的代码。
	//  spec.model = "Mac Mini"
	fmt.Println("Spec:", spec)
	/****

	结构体相等性（Structs Equality）
	结构体是值类型。如果它的每一个字段都是可比较的，
	则该结构体也是可比较的。如果两个结构体变量的对应字段相等，则这两个变量也是相等的。
	 */
	name1 := name{"Steve", "Jobs"}
	name2 := name{"Steve", "Jobs"}
	if name1 == name2 {
		fmt.Println("name1 and name2 are equal")
	} else {
		fmt.Println("name1 and name2 are not equal")
	}

	name3 := name{firstName:"Steve", lastName:"Jobs"}
	name4 := name{}
	name4.firstName = "Steve"
	if name3 == name4 {
		fmt.Println("name3 and name4 are equal")
	} else {
		fmt.Println("name3 and name4 are not equal")
	}

/*	image1 := image{data: map[int]int{
		0: 155,
	}}
	image2 := image{data: map[int]int{
		0: 155,
	}}
	if image1 == image2 {
		fmt.Println("image1 and image2 are equal")
	}*/

}

