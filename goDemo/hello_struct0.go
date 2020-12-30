package main

import "fmt"

/**
什么是结构体？
结构体是用户定义的类型，表示若干个字段（Field）的集合。
有时应该把数据整合在一起，而不是让这些数据没有联系。这种情况下可以使用结构体。
例如，一个职员有 firstName、lastName 和 age 三个属性，
而把这些属性组合在一个结构体 employee 中就很合理。
结构体的声明
type Employee struct {
    firstName string
    lastName  string
    age       int
}

通过把相同类型的字段声明在同一行，
结构体可以变得更加紧凑。在上面的结构体中，
firstName 和 lastName 属于相同的 string 类型，于是这个结构体可以重写为
type Employee struct {
    firstName, lastName string
    age, salary         int
}
Employee 称为 命名的结构体（Named Structure）。
我们创建了名为 Employee 的新类型，而它可以用于创建 Employee 类型的结构体变量。

匿名结构体（Anonymous Structure）
声明结构体时也可以不用声明一个新类型，这样的结构体类型称为 匿名结构体（Anonymous Structure）。
var employee struct {
    firstName, lastName string
    a
***/

type Employee struct {
	firstName, lastName string
	age, salary         int
}

func main()  {
	//creating structure using field names
	/***
	  我们创建了一个命名的结构体 Employee，通过指定每个字段名的值，我们定义了结构体变量 emp1。字段名的顺序不一定要与声明结构体类型时的顺序相同。在这里，我们改变了 lastName 的位置，将其移到了末尾。这样做也不会有任何的问题
	  **/
	emp1 := Employee{
		firstName: "Sam",
		age:       25,
		salary:    500,
		lastName:  "Anderson",
	}
	//creating structure without using field names
	emp2 := Employee{"Thomas", "Paul", 29, 800}

	fmt.Println("Employee 1", emp1)
	fmt.Println("Employee 2", emp2)

	//创建匿名结构体
	/**
	  我们定义了一个匿名结构体变量 emp3。上面我们已经提到，之所以称这种结构体是匿名的，是因为它只是创建一个新的结构体变量 em3，而没有定义任何结构体类型。
	  **/
	emp3 := struct {
		firstName, lastName string
		age, salary         int
	}{
		firstName: "Andreah",
		lastName:  "Nikola",
		age:       31,
		salary:    5000,
	}

	fmt.Println("Employee 3", emp3)

	//结构体的零值（Zero Value）
	var emp4 Employee //zero valued structure
	fmt.Println("Employee 4", emp4)
	//当然还可以为某些字段指定初始值，而忽略其他字段。这样，忽略的字段名会赋值为零值。


	emp5 := Employee{
		firstName: "John",
		lastName:  "Paul",
	}
	fmt.Println("Employee 5", emp5)

	//访问结构体的字段   点号操作符 . 用于访问结构体的字段。
	emp6 := Employee{"Sam", "Anderson", 55, 6000}
	fmt.Println("First Name:", emp6.firstName)
	fmt.Println("Last Name:", emp6.lastName)
	fmt.Println("Age:", emp6.age)
	fmt.Printf("Salary: $%d", emp6.salary)
	//还可以创建零值的 struct，以后再给各个字段赋值。
	var emp7 Employee
	emp7.firstName = "Jack"
	emp7.lastName = "Adams"
	fmt.Println("Employee 7:", emp7)

	//结构体的指针

	emp8 := &Employee{"Sam", "Anderson", 55, 6000}
	fmt.Println("First Name:", (*emp8).firstName)
	fmt.Println("Age:", (*emp8).age)
	//等价
	/***
	emp8 是一个指向结构体 Employee 的指针。(*emp8).firstName
	表示访问结构体 emp8 的 firstName 字段。该程序会输出：
	Go 语言允许我们在访问 firstName 字段时，
	可以使用 emp8.firstName 来代替显式的解引用 (*emp8).firstName。
	    ***/
	emp8_ := &Employee{"Sam", "Anderson", 55, 6000}
	fmt.Println("First Name:", emp8_.firstName)
	fmt.Println("Age:", emp8_.age)

	//匿名字段 当我们创建结构体时，字段可以只有类型，而没有字段名。这样的字段称为匿名字段（Anonymous Field）。
	p := Person{"Naveen", 50}
	fmt.Println(p)
	/**
	  虽然匿名字段没有名称，但其实匿名字段的名称就默认为它的类型。
	比如在上面的 Person 结构体里，虽说字段是匿名的，
	但 Go 默认这些字段名是它们各自的类型。所以 Person 结构体有两个名为 string 和 int 的字段。
	  我们访问了 Person 结构体的匿名字段，我们把字段类型作为字段名，分别为 "string" 和 "int"
	  ***/
	var p1 Person
	p1.string = "naveen"
	p1.int = 50
	fmt.Println(p1)


	//嵌套结构体（Nested Structs）
	var p2 Person2
	p2.name = "Naveen"
	p2.age = 50
	p2.address = Address {
		city: "Chicago",
		state: "Illinois",
	}
	fmt.Println("Name:", p2.name)
	fmt.Println("Age:",p2.age)
	fmt.Println("City:",p2.address.city)
	fmt.Println("State:",p2.address.state)


	//提升字段（Promoted Fields）
	//如果是结构体中有匿名的结构体类型字段，则该匿名结构体里的字段就称为提升字段

	var p3 Person3
	p3.name = "Naveen"
	p3.age = 50
	p3.Address = Address{
		city:  "Chicago",
		state: "Illinois",
	}
	fmt.Println("Name:", p3.name)
	fmt.Println("Age:", p3.age)
	/**
	我们使用了语法 p3.city 和 p3.state,
	访问提升字段 city 和 state 就像它们是在结构体 p 中声明的一样。该程序会输出：
	***/

	fmt.Println("City:", p3.city) //city is promoted field
	fmt.Println("State:", p3.state) //state is promoted field
	/**
	在上面的代码片段中，Person 结构体有一个匿名字段 Address，而
	Address 是一个结构体。现在结构体 Address 有 city 和 state 两个字段，
	访问这两个字段就像在 Person3 里直接声明的一样，因此我们称之为提升字段。
	**/

}

//匿名字段
type Person struct {
	string
	int
}

type Address struct {
	city, state string
}
type Person2 struct {
	name string
	age int
	address Address
}

type Person3 struct {
	name string
	age  int
	Address
}