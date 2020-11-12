package main
import (
  "fmt"
)

type phone interface{
  call()
}

type NokiaPhone struct {

}

type IPhone struct {

}

func (noika NokiaPhone) call(){
 fmt.Println("iam nokia i call you")
}

func (iphone IPhone) call(){
 fmt.Println("i am iphone i call you")
}

func main(){
	var phone phone
	phone = new (NokiaPhone)
	phone.call()

	phone = new(IPhone)
    	phone.call()
}
