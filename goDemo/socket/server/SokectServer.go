
package main

import (
	"bufio"
	"fmt"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()
	for {
		reader := bufio.NewReader(conn)
		buf := make([]byte,128)
		n,err := reader.Read(buf)
		if err != nil {
			fmt.Println("数据读取失败",err)
			return
		}
		recvStr := string(buf[:n])
		fmt.Println("客户端发送过来的值:",recvStr)
	}

}
func main() {
	lister,err := net.Listen("tcp","0.0.0.0:8008")
	if err != nil {
		fmt.Println("连接失败",err)
	}
	for {
		fmt.Println("等待建立连接，此时会阻塞住")
		conn,err := lister.Accept() //等待建立连接
		fmt.Println("连接建立成功，继续")
		if err != nil {
			fmt.Println("建立连接失败",err)
			//继续监听下次链接
			continue
		}
		go process(conn)
	}
}