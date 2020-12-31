package main
import (
"bufio"
"fmt"
	"goDemo/goDemo/socket/socker_stick"
	"io"
"net"
)

func process2(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)

	for {
		msg,err := socker_stick.Decode(reader)
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

		fmt.Println("客户端发送过来的值:",msg)
	}

}
//对了，只有TCP才有粘包
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
		go process2(conn)
	}
}