package main

import (
"fmt"
)

type SalaryCalculator_2 interface {
	DisplaySalary()
}

type LeaveCalculator_2 interface {
	CalculateLeavesLeft() int
}
//接口的嵌套
type EmployeeOperations interface {
	SalaryCalculator_2
	LeaveCalculator_2
}

type Employee_3 struct {
	firstName string
	lastName string
	basicPay int
	pf int
	totalLeaves int
	leavesTaken int
}

func (e Employee_3) DisplaySalary() {
	fmt.Printf("%s %s has salary $%d", e.firstName, e.lastName, (e.basicPay + e.pf))
}

func (e Employee_3) CalculateLeavesLeft() int {
	return e.totalLeaves - e.leavesTaken
}

func main() {
	e := Employee_3 {
		firstName: "Naveen",
		lastName: "Ramanathan",
		basicPay: 5000,
		pf: 200,
		totalLeaves: 30,
		leavesTaken: 5,
	}
	var empOp EmployeeOperations = e
	empOp.DisplaySalary()
	fmt.Println("\nLeaves left =", empOp.CalculateLeavesLeft())
}

/**

接口的嵌套
尽管 Go 语言没有提供继承机制,但可以通过嵌套其他的接口,创建一个新接口。

在上述程序的第 15 行,我们创建了一个新的接口 EmployeeOperations,它嵌套了两个接口：SalaryCalculator 和 LeaveCalculator。
如果一个类型定义了 SalaryCalculator 和 LeaveCalculator 接口里包含的方法,我们就称该类型实现了 EmployeeOperations 接口。
在第 29 行和第 33 行,由于 Employee 结构体定义了 DisplaySalary 和 CalculateLeavesLeft 方法,因此它实现了接口 EmployeeOperations。
在 46 行,empOp 的类型是 EmployeeOperations,e 的类型是 Employee,我们把 empOp 赋值为 e。接下来的两行,empOp 调用了 DisplaySalary() 和 CalculateLeavesLeft() 方法。
*/