package main

import "fmt"

type itemDetail struct {
	name     string
	category int
	quantity int
	cost     float64
}

func (i itemDetail) print() {
	fmt.Printf("Current item name is %s - Category is %s - Quantity is %d - Unit Cost is %.2f\n", i.name, category[i.category], i.quantity, i.cost)
}

func (i itemDetail) mapItem() Item {
	var item Item
	item.category = i.category
	item.quantity = i.quantity
	item.cost = i.cost
	return item
}
