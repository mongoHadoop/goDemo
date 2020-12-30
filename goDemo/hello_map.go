package main

import "fmt"

/**
Go 经典入门系列 12：Maps
map 是在 Go 中将值（value）与键（key）关联的内置类型。通过相应的键可以获取到值。

如何创建 map ？
通过向 make 函数传入键和值的类型，可以创建 map make(map[type of key]type of value) 是创建 map 的语法。

通过向 make 函数传入键和值的类型，可以创建 map。
map 的零值是 nil。如果你想添加元素到 nil map 中，会触发运行时 panic。因此 map 必须使用 make 函数初始化。
personSalary := make(map[string]int)
*/
func main()  {
	var personSalary map[string]int
	if personSalary == nil {
		fmt.Println("map is nil. Going to make one.")
		personSalary = make(map[string]int)
	}

	//给 map 添加元素
	personSalary1 := make(map[string]int)
	personSalary1["steve"] = 12000
	personSalary1["jamie"] = 15000
	personSalary1["mike"] = 9000
	fmt.Println("personSalary1 map contents:", personSalary1)


	//你也可以在声明的时候初始化 map。
	personSalary2 := map[string]int {
		"steve": 12000,
		"jamie": 15000,
	}

	personSalary2["mike"] = 9000
	fmt.Println("personSalary map contents:", personSalary2)
	//键不一定只能是 string 类型。所有可比较的类型，如 boolean，interger，float，complex，string 等，都可以作为键。关于可比较的类型，如果你想了解更多，请访问http://golang.org/ref/spec#Comparison_operators


	//获取 map 中的元素
	personSalary2["mike"] = 9000
	employee := "jamie"
	fmt.Println("Salary of", employee, "is", personSalary2[employee])

	//如果获取一个不存在的元素，会发生什么呢？map 会返回该元素类型的零值。在 personSalary 这个 map 里，如果我们获取一个不存在的元素，会返回 int 类型的零值 0。

	//如果我们想知道 map 中到底是不是存在这个 key，该怎么做：
	//value, ok := map[key]
	//上面就是获取 map 中某个 key 是否存在的语法。如果 ok 是 true，表示 key 存在，key 对应的值就是 value ，反之表示 key 不存在。
	newEmp := "joe"
	value, ok := personSalary2[newEmp]
	if ok == true {
		fmt.Println("Salary of", newEmp, "is", value)
	} else {
		fmt.Println("key",newEmp,"not found")
	}



	//遍历 map 中所有的元素需要用 for range 循环。
	//有一点很重要，当使用 for range 遍历 map 时，不保证每次执行程序获取的元素顺序相同。

	fmt.Println("All items of a map")
	for key, value := range personSalary2 {
		fmt.Printf("personSalary[%s] = %d\n", key, value)
	}
	//删除 map 中的元素

	fmt.Println("map before deletion", personSalary2)
	delete(personSalary2, "steve")
	fmt.Println("map after deletion", personSalary2)

	//获取 map 的长度
	fmt.Println("length is", len(personSalary2))
	/***
	Map 是引用类型
	和 slices 类似，map 也是引用类型。当 map 被赋值为一个新变量的时候,
	它们指向同一个内部数据结构。因此，改变其中一个变量，就会影响到另一变量。
	 **/
	fmt.Println("Original person salary", personSalary2)
	newPersonSalary := personSalary2
	newPersonSalary["mike"] = 1000000
	fmt.Println("Person salary changed", personSalary2)



}