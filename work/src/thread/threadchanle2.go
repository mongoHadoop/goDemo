package main

import (
	"fmt"
	"math/rand"
	"time"
)
type Worker2 struct {
	id int
}
func main() {
	c := make(chan int)
	for i := 0; i < 5; i++ {
		worker := &Worker2{id: i}
		go worker.process2(c)
	}
	/**
	我们每秒往通道中发送20个信息，但是我们的程序，每秒只能处理10个信息；因此，有一半的信息被丢弃。
	 */
	for {
		select {
		case c <- rand.Int():
			//可选的代码
		default:
			//这里可以留下空行以丢弃数据
			fmt.Println("dropped")
		}
		time.Sleep(time.Millisecond * 50)
	}
}


/**
我们需要知道从一个通道接收或者发送数据时会阻塞.
当我们从一个通道接收数据时,直到数据可用,
go协程才会继续执行.类似的,往一个通道发送数据时，在数据被接收之前, go协程也不会继续执行.
 **/
func (w *Worker2) process2(c chan int) {
	for {
		data := <-c
		data++
		fmt.Printf("worker %d got %d\n", w.id, data)
		c<-data
		time.Sleep(time.Millisecond * 500)
	}
}