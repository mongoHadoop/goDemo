package quicksort

import "fmt"
func QuickSort(arr *[]int)   {
	quickSort2(arr,0,len(*arr)-1)
	for _,value :=range (*arr){
		fmt.Printf("value %d\n", value)
	}
}

func quickSort2( arr *[]int, leftIndex int, rightIndex int)  {
	if(leftIndex>=rightIndex){
		return
	}
	//left,right   := leftIndex,rightIndex
	var left,right  int = leftIndex,rightIndex
	var key   int = (*arr)[left]
	for left<right {
		for right>left && (*arr)[right]>=key{
			right--
			//fmt.Printf("right=%d  left = %d\n", right,left)
		}
		//找到这种元素将arr[right]放入arr[left]中
		(*arr)[left] =(*arr)[right]
		for left<right   && (*arr)[right]<=key{
			//从左往右扫描，找到第一个比基准值大的元素
			left++
		}
		//找到这种元素将arr[left]放入arr[right]中
		(*arr)[right] = (*arr)[left]

	}
	//基准值归位
	(*arr)[left] = key;
	quickSort2(arr,leftIndex,left-1);
	quickSort2(arr,right+1,rightIndex);
}