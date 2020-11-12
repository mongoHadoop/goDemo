package main

/**
  结构体上的函数
*/
import "fmt"

type Saiyan struct {
	Name  string
	Power int
}

func (s *Saiyan) add() {
	s.Power += 10000
}

func main() {
	goku := &Saiyan{"Goku", 9001}
	goku.add()
	fmt.Println("sss=", goku.Power)
	//尽管没有构造函数，go有一个内置的函数new，可以用来分配一个类型需要的内存。new(X)和&X{}是等效的：
	goku2 := new(Saiyan)
	goku2.Name = "ssv"
	goku2.Power = 900
	goku2.add()
	fmt.Println("sss=", goku2.Power)
}
