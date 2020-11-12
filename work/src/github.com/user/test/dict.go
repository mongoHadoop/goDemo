package main

import "fmt"

func main() {
	mapkv := make(map[string]int)
	mapkv["age"] = 19
	mapkv["sex"] = 1
	mapkv["phone"] = 13438045673
	mapkv["name"] = 111
	age, err := mapkv["name"]
	fmt.Println(age, err)
	//使用len可以获得映射中键的个数。使用delete可以删除映射中的一个键值对。
	// 返回 1
	total := len(mapkv)
	fmt.Println(total)
	// 没有返回值, 可以调用一个不存在的键
	delete(mapkv, "goku")
}
