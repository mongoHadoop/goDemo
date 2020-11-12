package main


import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("start")
	//go协程类似一个线程,但是go协程是由go自己调度,而不是系统.在协程中的代码可以和其他代码并发执行.让我们看一个例子：
	go process()
	/***
	但是最重要的是了解我们是如何启动一个go协程.
	我们只是简单的将go关键字附在我们想要执行的函数前面即可
	一个go协程的开销和系统线程比起来相对很低（一般都是几KB）.在现代的硬件上,有可能拥有成千上万个go协程
	**/
	time.Sleep(time.Millisecond * 10) // this is bad, don't do this!
	fmt.Println("done")
}

func process() {
	fmt.Println("processing")
}