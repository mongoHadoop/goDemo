package main

import (
	"fmt"
)

func digits(number int, dchnl chan int) {
	for number != 0 {
		digit := number % 10
		dchnl <- digit
		number /= 10
	}
	close(dchnl)
}
func calcSquares2(number int, squareop chan int) {
	sum := 0
	dch := make(chan int)
	go digits(number, dch)
	for digit := range dch {
		sum += digit * digit
	}
	squareop <- sum
}

func calcCubes2(number int, cubeop chan int) {
	sum := 0
	dch := make(chan int)
	go digits(number, dch)
	for digit := range dch {
		sum += digit * digit * digit
	}
	cubeop <- sum
}

func main() {
	number := 589
	sqrch := make(chan int)
	cubech := make(chan int)
	go calcSquares2(number, sqrch)
	go calcCubes2(number, cubech)
	squares, cubes := <-sqrch, <-cubech
	fmt.Println("Final output", squares+cubes)
}
/***

如果你仔细观察这段代码，
会发现获得一个数里的每位数的代码在 calcSquares 和 calcCubes 两个函数内重复了
我们将把这段代码抽离出来，放在一个单独的函数里，然后并发地调用它.

上述程序里的 digits 函数，包含了获取一个数的每位数的逻辑，
并且 calcSquares 和 calcCubes 两个函数并发地调用了 digits。当计算完数字里面的每一位数时，
第 13 行就会关闭信道。calcSquares 和 calcCubes
两个协程使用 for range 循环分别监听了它们的信道，直到该信道关闭。程序的其他地方不变，该程序同样会输出：

上述程序里的 digits 函数，包含了获取一个数的每位数的逻辑，并且 calcSquares 和 calcCubes 两个函数并发地调用了 digits。当计算完数字里面的每一位数时，
第 13 行就会关闭信道。calcSquares 和 calcCubes 两个协程使用 for range 循环分别监听了它们的信道，直到该信道关闭。程序的其他地方不变，该程序同样会输出：

*/