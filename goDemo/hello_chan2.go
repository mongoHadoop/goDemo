package main

import "fmt"

func sendData(sendch chan<- int) {
	sendch <- 10
}

func main() {
	sendch := make(chan<- int)//chan<- int 定义了唯送信道
	go sendData(sendch)
	fmt.Println(<-sendch)//因为箭头指向了 chan。在第 12 行，我们试图通过唯送信道接收数据，于是编译器报错：
}
/**

单向信道
我们目前讨论的信道都是双向信道，即通过信道既能发送数据，
又能接收数据。其实也可以创建单向信道，这种信道只能发送或者接收数据。

上面程序的第 10 行，我们创建了唯送（Send Only）信道 sendch。chan<- int 定义了唯送信道，
因为箭头指向了 chan。在第 12 行，
我们试图通过唯送信道接收数据，于是编译器报错：
**/