package main

import "fmt"

func calcSquares(number int, squareop chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit
		number /= 10
	}
	squareop <- sum
}

func calcCubes(number int, cubeop chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit * digit
		number /= 10
	}
	cubeop <- sum
}

func main() {
	number := 589
	sqrch := make(chan int)
	cubech := make(chan int)
	go calcSquares(number, sqrch)
	go calcCubes(number, cubech)
	squares, cubes := <-sqrch, <-cubech
	fmt.Println("Final output", squares + cubes)
}
/**
信道的另一个示例
我们再编写一个程序来更好地理解信道。
该程序会计算一个数中每一位的平方和与立方和，
然后把平方和与立方和相加并打印出来。
我们会这样去构建程序：在一个单独的 Go 协程计算平方和，
而在另一个协程计算立方和，最后在 Go 主协程把平方和与立方和相加。
***/