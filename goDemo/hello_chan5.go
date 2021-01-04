package main

import (
	"fmt"
)

func producer2(chnl chan int) {
	for i := 0; i < 10; i++ {
		chnl <- i
	}
	close(chnl)
}
func main() {
	ch := make(chan int)
	go producer2(ch)
	for v := range ch {
		fmt.Println("Received ",v)
	}
}
/***
接下来我们使用 for range 循环重写上面的代码。

在第 16 行，for range 循环从信道 ch 接收数据，
直到该信道关闭。一旦关闭了 ch，循环会自动结束。该程序会输出：
我们可以使用 for range 循环，重写信道的另一个示例[14]这一节里面的代码，提高代码的可重用性。
**/