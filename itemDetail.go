package main

import "fmt"

type itemDetail struct {
	name string
	item Item
}

func (i itemDetail) print() {
	fmt.Printf("Current item name is %s - Category is %s - Quantity is %d - Unit Cost is %.2f\n", i.name, category[i.item.category], i.item.quantity, i.item.cost)
}
