package main

import (
"fmt"
)

type SalaryCalculator interface {
	CalculateSalary() int
}

type Permanent struct {
	empId    int
	basicpay int
	pf       int
}

type Contract struct {
	empId  int
	basicpay int
}

//salary of permanent employee is sum of basic pay and pf
func (p Permanent) CalculateSalary() int {
	return p.basicpay + p.pf
}

//salary of contract employee is the basic pay alone
func (c Contract) CalculateSalary() int {
	return c.basicpay
}

/*
接口的实际用途
我们现在讨论一下接口的实际应用场景。
我们编写一个简单程序，根据公司员工的个人薪资，计算公司的总支出
total expense is calculated by iterating though the SalaryCalculator slice and summing
the salaries of the individual employees
接口编程,利用接口为参数
*/
func totalExpense(s []SalaryCalculator) {
	expense := 0
	for _, v := range s {
		expense = expense + v.CalculateSalary()
	}
	fmt.Printf("Total Expense Per Month $%d", expense)
}

func main() {
	pemp1 := Permanent{1, 5000, 20}
	pemp2 := Permanent{2, 6000, 30}
	cemp1 := Contract{3, 3000}
	employees := []SalaryCalculator{pemp1, pemp2, cemp1}
	totalExpense(employees)

}
/**
上面程序的第 7 行声明了一个 SalaryCalculator 接口类型，它只有一个方法 CalculateSalary() int。
在公司里，我们有两类员工，即第 11 行和第 17 行定义的结构体
：Permanent 和 Contract。长期员工（Permanent）的薪资是 basicpay 与 pf 相加之和，
而合同员工（Contract）只有基本工资 basicpay。在第 23 行和第 28 行中，
方法 CalculateSalary 分别实现了以上关系。由于 Permanent 和 Contract 都声明了该方法，因此它们都实现了 SalaryCalculator 接口。

第 36 行声明的 totalExpense 方法体现出了接口的妙用。该方法接收一个
SalaryCalculator 接口的切片（[]SalaryCalculator）作为参数。在第 49 行，我们向 totalExpense 方法传递了一个包含 Permanent 和 Contact 类型的切片。
在第 39 行中，通过调用不同类型对应的 CalculateSalary 方法，totalExpense 可以计算得到支出。

这样做最大的优点是：totalExpense 可以扩展新的员工类型，而不需要修改任何代码。
假如公司增加了一种新的员工类型 Freelancer，它有着不同的薪资结构。Freelancer只需传递到 totalExpense 的切片参数中，
无需 totalExpense 方法本身进行修改。只要 Freelancer 也实现了 SalaryCalculator 接口，totalExpense 就能够实现其功能。
该程序输出 Total Expense Per Month $14050。
*/