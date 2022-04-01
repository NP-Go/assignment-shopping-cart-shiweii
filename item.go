package main

import "fmt"

type Item struct {
	category int
	quantity int
	cost     float64
	_        struct{}
}

func (i Item) print(name string) {
	fmt.Printf("Current item name is %s - Category is %s - Quantity is %d - Unit Cost is %.2f\n", name, category[i.category], i.quantity, i.cost)
}

func (i Item) totalCost() float64 {
	var tc float64
	tc = float64(i.quantity) * i.cost
	return float64(tc)
}

func (i Item) compare(in Item) ([]string, bool) {
	var msg []string
	var isDiff bool = false
	if i.category != in.category {
		isDiff = true
	} else {
		msg = append(msg, "No changes to category made.")
	}
	if i.quantity != in.quantity {
		isDiff = true
	} else {
		msg = append(msg, "No changes to quantity made.")
	}
	if i.cost != in.cost {
		isDiff = true
	} else {
		msg = append(msg, "No changes to cost made.")
	}
	return msg, isDiff
}
