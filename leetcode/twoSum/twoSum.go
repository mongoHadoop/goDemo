package main
//main包是一个特殊的包名，它表示当前是一个可执行程序，而不是一个库。
import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		another := target - nums[i]
		if _, ok := m[another]; ok {
			return []int{m[another], i}
		}
		m[nums[i]] = i
		fmt.Println(m)
	}
	return nil
}
func main() {
	scores := []int{5, 4, 23, 6, 9}
	fmt.Println("worst=", twoSum(scores,29))
}