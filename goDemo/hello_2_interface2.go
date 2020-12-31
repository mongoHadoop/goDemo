package main

import (
	"fmt"
)

type SalaryCalculator_ interface {
	DisplaySalary()
}

type LeaveCalculator_ interface {
	CalculateLeavesLeft() int
}

type Employee_ struct {
	firstName string
	lastName string
	basicPay int
	pf int
	totalLeaves int
	leavesTaken int
}

func (e Employee_) DisplaySalary() {
	fmt.Printf("%s %s has salary $%d", e.firstName, e.lastName, (e.basicPay + e.pf))
}

func (e Employee_) CalculateLeavesLeft() int {
	return e.totalLeaves - e.leavesTaken
}

func main() {
	e := Employee_ {
		firstName: "Naveen",
		lastName: "Ramanathan",
		basicPay: 5000,
		pf: 200,
		totalLeaves: 30,
		leavesTaken: 5,
	}
	var s SalaryCalculator_ = e
	s.DisplaySalary()
	var l LeaveCalculator_ = e
	fmt.Println("\nLeaves left =", l.CalculateLeavesLeft())
}
/**
上述程序在第 7 行和第 11 行分别声明了两个接口：SalaryCalculator 和 LeaveCalculator。
第 15 行定义了结构体 Employee，它在第 24 行实现了 SalaryCalculator 接口的 DisplaySalary 方法，
接着在第 28 行又实现了 LeaveCalculator 接口里的 CalculateLeavesLeft 方法。于是 Employee 就实现了 SalaryCalculator 和 LeaveCalculator 两个接口。

第 41 行，我们把 e 赋值给了 SalaryCalculator 类型的接口变量 ，
而在 43 行，我们同样把 e 赋值给 LeaveCalculator 类型的接口变量 。
由于 e 的类型 Employee 实现了 SalaryCalculator 和 LeaveCalculator 两个接口，因此这是合法的。
***/