package main

import (
"fmt"
"net"
)

func main() {
	listen,err := net.ListenUDP("udp",&net.UDPAddr{
		IP:   net.IPv4(0,0,0,0),
		Port: 8009,
	})
	if err != nil {
		panic(fmt.Sprintf("udp启动失败,err:%v",err))
	}
	defer listen.Close()
	for{
		var data = make([]byte,1024)
		n,addr,err := listen.ReadFromUDP(data)
		if err != nil {
			panic(fmt.Sprintf("读取数据失败,err:%v",err))
		}
		fmt.Println(string(data[:n]),addr,n)
	}
}