package main

import "fmt"

type MyArray struct {
	//exported struct
	Size  int   //exported field
	Array *[]int //unexported field
}
var myArray MyArray
func insert(element int,index int)  {
	arry1:=*myArray.Array
	if(index<0||index>len(arry1)){
		fmt.Println("超出数组实际元 素范围！")
		return
	}
	if(myArray.Size> len(*myArray.Array)){
		resize(arry1)
	}
	//
	//if myArray.Size==0{
	//	arry1[0]=element
	//	myArray.Size++
	//	return
	//}
	//从右向左循环，将元素逐个向右挪1位
	for i:=myArray.Size;i>index;i--{
		arry1[i+1]=arry1[i]
	}
	//腾出的位置放入新元素
	arry1[index] = element
	myArray.Size++

}
func resize(array1 []int )  {
	copy_slice := make([]int, len(array1))
	copy(copy_slice, array1)
	myArray.Array=&copy_slice

}
func outprint()  {
	array1:=*myArray.Array
	for i:=0;i<myArray.Size;i++{
		fmt.Println(array1[i])
	}
}

func main()  {

	myArray = MyArray{
		Size: 0,
		Array: &[]int{},
	}
	v1:=make([]int, 10)
	myArray.Array= &v1
	insert(3,0)
	 insert(7,1)
	 insert(9,2)
	 insert(5,3)
	insert(6,4)
	 outprint()
}