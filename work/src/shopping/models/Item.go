package models

type Item struct {
	Price float64
}

/**
可以通过models.NewItem()调用这个函数。但是如果这个函数名为newItem，那么我们从其他的包是不能调用这个函数的。
例如，如果你将结构体Item的Price字段改成price，你会得到一个错误。
即当　属性字段或者函数,public暴露出去就，大写
如果要private 则小写
*/
func NewItem() *Item {
	item := new(Item)
	item.Price = 200
	return item
}
