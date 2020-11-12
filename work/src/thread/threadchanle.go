package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	for i := 0; i < 5; i++ {
		worker := &Worker{id: i}
		go worker.process(c)
	}
	//我们不知道哪个worker将获得数据。我们所知道的是，go语言确保，往一个通道发送数据时，仅有一个单独的接收器可以接收。
	/**
	注意：通道是唯一共享的状态,通过通道,可以安全的,并发发送和接收数据.
	通道提供了我们需要的所有同步代码，并且也确保，在任意的特定时刻，只有一个go协程，可以访问数据的特定部分
	 */
		c <- 0
		time.Sleep(time.Millisecond * 50)
}

type Worker struct {
	id int
}
/**
我们需要知道从一个通道接收或者发送数据时会阻塞.
当我们从一个通道接收数据时,直到数据可用,
go协程才会继续执行.类似的,往一个通道发送数据时，在数据被接收之前, go协程也不会继续执行.
 **/
func (w *Worker) process(c chan int) {
	for {
		data := <-c
		data++
		fmt.Printf("worker %d got %d\n", w.id, data)
		c<-data
		time.Sleep(time.Millisecond * 500)
	}
}