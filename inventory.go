package main

import (
	"fmt"
	"sort"
)

type Inventory map[string]Item

func (i Inventory) delete(v string) bool {
	_, ok := i[v]
	if ok {
		delete(i, v)
		return true
	}
	return false
}

func (i Inventory) list() {
	if len(i) > 0 {
		for n, i := range i {
			fmt.Printf("Category: %s - Item: %s Quantity: %d Unit Coast: %.2f\n", category[i.category], n, i.quantity, i.cost)
		}
	} else {
		fmt.Println("No data found!")
	}
}

func (i Inventory) listSorted() {
	type tempItem = struct {
		name     string
		category int
		quantity int
		cost     float64
	}

	sortSlice := make([]tempItem, 0, len(inventory))

	for i, k := range inventory {
		sortSlice = append(sortSlice, tempItem{name: i, category: k.category, quantity: k.quantity, cost: k.cost})
	}

	sort.Slice(sortSlice, func(i, j int) bool { return sortSlice[i].category < sortSlice[j].category })

	for _, item := range sortSlice {
		fmt.Printf("Category: %s - Item: %s Quantity: %d Unit Coast: %.2f\n", category[item.category], item.name, item.quantity, item.cost)
	}
}

func (i Inventory) print() {
	if len(i) > 0 {
		for i, k := range i {
			fmt.Printf("%s - {%d %d %.2f}\n", i, k.category, k.quantity, k.cost)
		}
	} else {
		fmt.Println("No data found!")
	}
}
