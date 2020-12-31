package main

import "fmt"

type Describer interface {
	Describe()
}
type Person_ struct {
	name string
	age  int
}

func (p Person_) Describe() {
	fmt.Printf("%s is %d years old", p.name, p.age)
}
//还可以将一个类型和接口相比较。如果一个类型实现了接口，那么该类型与其实现的接口就可以互相比较。
func findType2(i interface{}) {
	switch v := i.(type) {
	case Describer:
		v.Describe()
	default:
		fmt.Printf("unknown type\n")
	}
}

func main() {
	findType2("Naveen")
	p := Person_{
		name: "Naveen R",
		age:  25,
	}
	findType2(p)
}
/***
在上面程序中，结构体 Person 实现了 Describer 接口。
在第 19 行的 case 语句中，
v 与接口类型 Describer 进行了比较。p 实现了 Describer，
因此满足了该 case 语句，于是当程序运行到第 32 行的 findType(p) 时，程序调用了 Describe() 方法。
**/