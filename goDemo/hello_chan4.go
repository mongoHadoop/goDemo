package main

import (
	"fmt"
)

func producer(chnl chan int) {
	for i := 0; i < 10; i++ {
		chnl <- i
	}
	close(chnl)
}
func main() {
	ch := make(chan int)
	go producer(ch)
	for {
		v, ok := <-ch
		if ok == false {
			break
		}
		fmt.Println("Received ", v, ok)
	}
}
/**
关闭信道和使用 for range 遍历信道

数据发送方可以关闭信道，通知接收方这个信道不再有数据发送过来。
当从信道接收数据时，接收方可以多用一个变量来检查信道是否已经关闭。
v, ok := <- ch


在上述的程序中，producer 协程会从 0 到 9 写入信道 chn1，然后关闭该信道。
主函数有一个无限的 for 循环（第 16 行），使用变量 ok（第 18 行）检查信道是否已经关闭。
如果 ok 等于 false，说明信道已经关闭，于是退出 for 循环。
如果 ok 等于 true，会打印出接收到的值和 ok 的值。
***/