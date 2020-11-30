package main

import (
	"fmt"
	"goDemo/leetcode/quicksort"
	//"goDemo/leetcode/popsort"
)


func main()  {
	scores := []int{5, 4, 1,2,23, 6, 9,6,5}
	//ss:=duplicatarry.RemoveDuplicates(scores)
	fmt.Println("信息")
	//popsort.Popsort(&scores)
	quicksort.QuickSort(&scores)
}