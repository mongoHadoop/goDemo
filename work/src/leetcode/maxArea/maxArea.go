package main
//main包是一个特殊的包名，它表示当前是一个可执行程序，而不是一个库。
import "fmt"

func maxArea(height []int) int {
	max, start, end := 0, 0, len(height)-1
	for start < end {
		width := end - start
		high := 0
		if height[start] < height[end] {
			high = height[start]
			start++
		} else {
			high = height[end]
			end--
		}

		temp := width * high
		if temp > max {
			max = temp
		}
	}
	return max
}

func main() {
	scores := []int{5, 4, 23, 6, 9}
	fmt.Println("worst=", maxArea(scores))
}
