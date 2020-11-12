package main

import (
	"fmt"
	"shopping"
	"shopping/models"
)

func main() {
	fmt.Println(models.NewItem().Price)
	fmt.Println(shopping.PriceCheck(4343))
}
