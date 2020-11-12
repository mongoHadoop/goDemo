package main

import "fmt"

var scores1 [10]int

func main() {
	scores1[2] = 9
	scores333 := [4]int{9001, 9333, 212, 33}
	//和声明数组不同的是，声明切片不需要在方括号中指定其大小
	scores := []int{1, 4, 293, 4, 9}
	//下面使用另外一种方式创建切片，这里使用make：
	//我们使用make，没有使用new，是因为创建一个切片不仅仅是分配一段内存（new只能分配一段内存）
	//在使用make创建切片时，我们可以分别的指定长度和容量大小：
	scores_ := make([]int, 10)
	//创建了一个长度为0但是容量为10的切片
	scores_ = append(scores_, 5)
	fmt.Println(scores)
	fmt.Println(scores333)
	fmt.Println(scores_)
	scores_1 := make([]int, 0, 10)
	//切片的第6个元素赋值
	scores_1 = scores_1[0:6]
	//事实证明，append比较特殊。
	// 如果底层的数组已经达到上限，append会重新创建一个更大的数组，并将所有的值复制过去（这就是动态数组的工作原理，例如php、python、ruby和javascript等等）
	scores_1[5] = 9033
	fmt.Println(scores_1)
	//最后，这里提供了4种常用的方式去初始化一个切片：
	/*names := []string{"leto", "jessica", "paul"}
	checks := make([]bool, 10)
	var names []string
	scores_v1 := make([]int, 0, 20)*/
	scores_v2 := []int{1, 2, 3, 4, 5}
	slice := scores_v2[2:]
	slice[0] = 999
	fmt.Println(scores_v2)




}
