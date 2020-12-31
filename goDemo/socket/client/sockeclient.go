
package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

//客户端
func main() {
	conn, err := net.Dial("tcp", "10.1.174.39:8008")
	if err != nil {
		fmt.Println("连接服务器失败",err)
	}
	defer conn.Close()
	inputReader:=bufio.NewReader(os.Stdin)
	for{
		fmt.Println(":")
		input,_:=inputReader.ReadString('\n')
		_, err = conn.Write([]byte(input))
		if err != nil {
			fmt.Println("发送成功")
		}
	}
}
