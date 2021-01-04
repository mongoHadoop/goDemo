package main

import "fmt"

/**
什么是信道？
信道可以想像成 Go 协程之间通信的管道。如同管道中的水会从一端流到另一端，
通过使用信道，数据也可以从一端发送，在另一端接收。
所有信道都关联了一个类型。信道只能运输这种类型的数据，
而运输其他类型的数据都是非法的。
chan T 表示 T 类型的信道。
*/

func main() {
	var a chan int
	if a == nil {
		fmt.Println("channel a is nil, going to define it")
		a = make(chan int)
		fmt.Printf("Type of a is %T", a)
	}
}
/***
通过信道进行发送和接收

如下所示，该语法通过信道发送和接收数据
data := <- a // 读取信道 a
a <- data // 写入信道 a

信道旁的箭头方向指定了是发送数据还是接收数据。
在第一行，箭头对于 a 来说是向外指的，因此我们读取了信道 a 的值，并把该值存储到变量 data。
在第二行，箭头指向了 a，因此我们在把数据写入信道 a。

发送与接收默认是阻塞的

发送与接收默认是阻塞的。这是什么意思？当把数据发送到信道时，程序控制会在发送数据的语句处发生阻塞，直到有其它 Go 协程从信道读取到数据，才会解除阻塞。与此类似，当读取信道的数据时，如果没有其它的协程把数据写入到这个信道，那么读取过程就会一直阻塞着。
信道的这种特性能够帮助 Go 协程之间进行高效的通信，不需要用到其他编程语言常见的显式锁或条件变量。

*/