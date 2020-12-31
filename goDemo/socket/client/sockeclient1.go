package main

import (
	"fmt"
	"net"
)

//客户端
func main() {
	conn, err := net.Dial("tcp", "10.1.174.39:8008")
	if err != nil {
		fmt.Println("连接服务器失败",err)
	}
	defer conn.Close()
	for i := 0; i < 10; i++ {
		sendStr := "hello world ads asdf asd fads fadsf ads ads asd asd ads "
		conn.Write([]byte(sendStr))
		//time.Sleep(time.Second)
		/***
		主要原因是因为我们是应用层软件，是跑在操作系统之上的软件，
		当我们向服务器发送一个数据时，是调用操作系统的相关接口发送的，操作系统再经过各种复杂的操作，发送到对方机器
		但是操作系统有一个发送数据缓冲区，默认情况如果缓冲区是有大小的，如果缓冲区没满，是不会发送数据的，
		所以上述客户端在发送数据时，系统的缓冲区都没满，一直压在了操作系统的缓冲区中，最后发现没数据了，才一次都发送到服务端
		但是为什么sleep(1)又管用了呢?这是因为缓冲区不止一个程序在用，1s的时间足够其他程序将缓冲区打满，
		然后各自发各自的数据，这也是为什么第一次操作没问题，第二次有问题，因为第二次全部都是我们客户端打满的
		**/
	}
}