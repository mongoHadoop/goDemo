package main

import (
"bufio"
"fmt"
"io"
"net"
)

func process1(conn net.Conn) {

	defer conn.Close()
	reader := bufio.NewReader(conn)
	buf := make([]byte,1024)
	for {
		n,err := reader.Read(buf)
	//读完了
	if err == io.EOF {
		fmt.Println("读完了")
		break
	}
	//读错了
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
		return
	}
	defer lister.Close()
	for {
		fmt.Println("等待建立连接，此时会阻塞住")
		conn,err := lister.Accept() //等待建立连接
		fmt.Println("连接建立成功，继续")
		if err != nil {
			fmt.Println("建立连接失败",err)
			//继续监听下次链接
			continue
		}
		go process1(conn)
	}
}