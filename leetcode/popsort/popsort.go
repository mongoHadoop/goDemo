package popsort

import "fmt"

func   Popsort  (nums *[]int) *[]int {
	fmt.Println("交换前....")
	for i,x:= range *nums {
		fmt.Printf("第 %d 位 x 的值 = %d\n", i,x)
	}
	for i:=0; i<len(*nums)-1;i++{
		for j:=0;j<len(*nums)-1-i;j++{
			if((*nums)[j]>(*nums)[j+1]){
				temp1 := (*nums)[j]
				(*nums)[j]=(*nums)[j+1]
				(*nums)[j+1] = temp1
			}
		}
	}
	fmt.Println("交换后....")
	for i,x:= range (*nums) {
		fmt.Printf("第 %d 位 x 的值 = %d\n", i,x)
	}
	return nums
}