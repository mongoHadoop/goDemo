package shopping

/***
大多数人容易认为，导入shopping/db有点特殊，
因为我们已经在shopping文件夹里面了。
实际上，我们是导入$GOPATH/src/shopping/db，
这就意味如果你的工作空间src/test文件下有一个db包，这也可以通过test/db导入这个db包。
*/
import (
	"shopping/db"
)

func PriceCheck(itemId int) (float64, bool) {
	var item = db.LoadItem(itemId)
	if item == nil {
		return 0, false
	}
	return item.Price, true
}
