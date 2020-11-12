package main

/**
  结构的声明和初始化
使用了&操作符去获得我们值的地址（&叫取地址符
*/
import "fmt"

type Saiyan0 struct {
	Name  string
	Power int
}

func add(s Saiyan0) {
	s.Power += 10000
}

//现在我们传递了一个地址类型*Saiyan，这里的*X表示一个指向类型X的一个指针
func add2(s *Saiyan0) {
	s.Power += 10000
}

func main() {
	//
	goku := Saiyan0{"Goku", 9000}
	add(goku)

	fmt.Println(goku.Power)
	//使用了&操作符去获得我们值的地址（&叫取地址符
	goku2 := &Saiyan0{"Goku", 9000}
	add2(goku2)
	fmt.Println(goku2.Power)
}
