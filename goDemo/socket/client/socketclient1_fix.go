package main

import (
	"fmt"
	"goDemo/goDemo/socket/socker_stick"
	"net"
)

func main() {
	conn,err := net.Dial("tcp", "10.1.174.39:8008")
	if err != nil {
		fmt.Println("连接服务器失败",err)
	}
	defer conn.Close()
	for i := 0; i < 10; i++ {
		sendStr := "hello world ads asdf asd fads fadsf ads ads asd asd ads "
		data,err := socker_stick.Encode(sendStr)
		if err != nil {
			fmt.Println("编码失败",err)
			return
		}
		conn.Write(data)
		//time.Sleep(time.Second)
	}
}