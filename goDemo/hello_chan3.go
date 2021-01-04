package main

import "fmt"
//函数 sendData 里的参数 sendch chan<- int 把 cha1 转换为一个唯送信道 于是该信道在 sendData 协程里是一个唯送信道
func sendData2(sendch chan<- int) {
	sendch <- 10
}

func main() {
	cha1 := make(chan int)
	go sendData2(cha1)
	fmt.Println(<-cha1)
}
/**
在上述程序的第 10 行，我们创建了一个双向信道 cha1。
在第 11 行 cha1 作为参数传递给了 sendData 协程。在第 5 行，
函数 sendData 里的参数 sendch chan<- int
把 cha1 转换为一个唯送信道。于是该信道在 sendData 协程里是一个唯送信道，
而在 Go 主协程里是一个双向信道。该程序最终打印输出 10。

一切都很顺利，只不过一个不能读取数据的唯送信道究竟有什么意义呢？
这就需要用到信道转换（Channel Conversion）了。把一个双向信道转换成唯送信道或者唯收（Receive Only）信道都是行得通的，但是反过来就不行。
*/