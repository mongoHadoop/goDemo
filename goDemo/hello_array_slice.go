package main

/**
数组 切片
切片
切片是由数组建立的一种方便、灵活且功能强大的包装（Wrapper）。切片本身不拥有任何数据。它们只是对现有数组的引用。
 */


import (
"fmt"
)

func main() {
	a := [5]int{76, 77, 78, 79, 80}
	var b []int = a[1:4] // creates a slice from a[1] to a[3]
	fmt.Println(b)
	//让我们看看另一种创建切片的方法。

	c := []int{6, 7, 8} // creates and array and returns a slice reference
	/**
	在上述 c：= [] int {6，7，8} 创建一个有 3 个整型元素的数组，并返回一个存储在 c 中的切片引用。
	*/
	fmt.Println(c)
	/***
	切片的修改
	切片自己不拥有任何数据。它只是底层数组的一种表示。对切片所做的任何修改都会反映在底层数组中。
	*/
	darr := [...]int{57, 89, 90, 82, 100, 78, 67, 69, 59}
	dslice := darr[2:5]
	fmt.Println("array before", darr)
	for i := range dslice {
		dslice[i]++
	}
	fmt.Println("array after", darr)
	//当多个切片共用相同的底层数组时，每个切片所做的更改将反映在数组中。
	numa := [3]int{78, 79 ,80}
	nums1 := numa[:] // creates a slice which contains all elements of the array
	nums2 := numa[:]
	fmt.Println("array before change 1", numa)
	nums1[0] = 100
	fmt.Println("array after modification to slice nums1", numa)
	nums2[1] = 101
	fmt.Println("array after modification to slice nums2", numa)


	//切片的长度和容量 切片的长度是切片中的元素数。切片的容量是从创建切片索引开始的底层数组中元素数。
	fruitarray := [...]string{"apple", "orange", "grape", "mango", "water melon", "pine apple", "chikoo"}
	fruitslice := fruitarray[1:3]
	fmt.Printf("length of slice %d capacity %d", len(fruitslice), cap(fruitslice))
	fmt.Println()
	// length of is 2 and capacity is 6
	//切片可以重置其容量。任何超出这一点将导致程序运行时抛出错误。
	fruitslice = fruitslice[:cap(fruitslice)] // re-slicing furitslice till its capacity
	fmt.Println("After re-slicing length is",len(fruitslice), "and capacity is",cap(fruitslice))
	//使用 make 创建一个切片
	/**
	  使用 make 创建一个切片
	  func make（[]T，len，cap）[]T 通过传递类型，长度和容量来创建切片。容量是可选参数, 默认值为切片长度。make 函数创建一个数组，并返回引用该数组的切片。
	*/

	i := make([]int, 5, 5)
	fmt.Println(i)
	/**
	  追加切片元素
	  正如我们已经知道数组的长度是固定的，它的长度不能增加。切片是动态的，使用 append 可以将新元素追加到切片上。append 函数的定义是 func append（s[]T，x ... T）[]T。
	*/
	/**
	  有一个问题可能会困扰你。如果切片由数组支持，并且数组本身的长度是固定的，那么切片如何具有动态长度。以及内部发生了什么，
	  当新的元素被添加到切片时，会创建一个新的数组。现有数组的元素被复制到这个新数组中，并返回这个新数组的新切片引用。
	  现在新切片的容量是旧切片的两倍。很酷吧 :)。下面的程序会让你清晰理解。
	*/
	//切片追加
	cars := []string{"Ferrari", "Honda", "Ford"}
	fmt.Println("cars:", cars, "has old length", len(cars), "and capacity", cap(cars)) // capacity of cars is 3
	cars = append(cars, "Toyota")
	//现在新切片的容量是旧切片的两倍
	fmt.Println("cars:", cars, "has new length", len(cars), "and capacity", cap(cars)) // capacity of cars is doubled to 6
	/**
	切片类型的零值为 nil。一个 nil 切片的长度和容量为 0。可以使用 append 函数将值追加到 nil 切片。
	*/
	var names []string //zero value of a slice is nil
	if names == nil {
		fmt.Println("slice is nil going to append")
		names = append(names, "John", "Sebastian", "Vinay")
		fmt.Println("names contents:",names)
	}
	/**
	    也可以使用 ... 运算符将一个切片添加到另一个切片。你可以在可变参数函数[25]教程中了解有关此运算符的更多信息。
	**/
	veggies := []string{"potatoes", "tomatoes", "brinjal"}
	fruits := []string{"oranges", "apples"}
	food := append(veggies, fruits...)
	fmt.Println("food:",food)

	/**
	    切片的函数传递
	    我们可以认为，切片在内部可由一个结构体类型表示。这是它的表现形式，
	        type slice struct {
	            Length        int
	            Capacity      int
	            ZerothElement *byte
	        }
	**/
	nos := []int{8, 7, 6}
	fmt.Println("slice before function call", nos)
	//元素递减 2。在函数调用后打印切片时，这些更改是可见的
	subtactOne(nos)                               // function modifies the slice
	fmt.Println("slice after function call", nos) // modifications are visible outside

	//多维切片,类似于数组，切片可以有多个维度。
	pls := [][]string {
		{"C", "C++"},
		{"JavaScript"},
		{"Go", "Rust"},
	}
	for _, v1 := range pls {
		for _, v2 := range v1 {
			fmt.Printf("%s ", v2)
		}
		fmt.Printf("\n")
	}
	/**
	  内存优化
	  切片持有对底层数组的引用。只要切片在内存中，数组就不能被垃圾回收。在内存管理方面，这是需要注意的。让我们假设我们有一个非常大的数组，
	  我们只想处理它的一小部分。然后，我们由这个数组创建一个切片，并开始处理切片。这里需要重点注意的是，在切片引用时数组仍然存在内存中。
	  一种解决方法是使用 copy[30] 函数 func copy(dst，src[]T)int 来生成一个切片的副本。这样我们可以使用新的切片，原始数组可以被垃圾回收。
	  **/
	countriesNeeded := countries()
	fmt.Println(countriesNeeded)

}

func countries() []string {
	countries := []string{"USA", "Singapore", "Germany", "India", "Australia"}
	neededCountries := countries[:len(countries)-2]
	countriesCpy := make([]string, len(neededCountries))
	// 一种解决方法是使用 copy[30] 函数 func copy(dst，src[]T)int 来生成一个切片的副本。这样我们可以使用新的切片，原始数组可以被垃圾回收。
	copy(countriesCpy, neededCountries) //copies neededCountries to countriesCpy
	return countriesCpy
}
func subtactOne(numbers []int) {
	for i := range numbers {
		numbers[i] -= 2
	}
}


