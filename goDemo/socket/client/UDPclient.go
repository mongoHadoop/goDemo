package main

import (
	"fmt"
	"net"
)

func main() {
	socker, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 8009,
	})
	if err != nil {
		panic(fmt.Sprintf("连接服务器失败,err:%v", err))
	}
	defer socker.Close()
	sendStr:="你好呀"
	_, err = socker.Write([]byte(sendStr))
	if err != nil {
		panic(fmt.Sprintf("数据发送失败,err:%v", err))
	}
}