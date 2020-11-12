package main

import "fmt"

type Person struct {
	Name string
}

func (p *Person) Introduce() {
	fmt.Printf("Hi, I'm %s\n", p.Name)
}

//go支持组合，即一种结构体包含另外一个结构体。在一些语言中，这叫混入类或者特性
//结构体Saiyan有一个字段时*Persion类型。
// 因此我们没有明确的给它一个字段名，我们可以间接的使用这个组合类型的字段和方法。然而，go编译器给会给该字段一个名字，认为这是完全有效的。
type Saiyan2 struct {
	*Person
	Power int
}

func (s *Saiyan2) Introduce() {
	fmt.Printf("Hi, I'm %s. Ya!\n", s.Name)
}

/***
组合优于继承吗？很多人都认为组合是一种更健壮的共享代码的方式。当你使用继承，你的类将和你的超类紧耦合，并且你最终更关注继承，而不是行为
*/
func main() {
	// 使用:
	goku := &Saiyan2{
		Person: &Person{"Goku"},
		Power:  9001,
	}
	goku.Person.Introduce()
}
