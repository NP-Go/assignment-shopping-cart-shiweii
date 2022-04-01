package main

import (
	"fmt"
	"sort"
	"strings"
)

type ShoppingList map[string]Item

func (s ShoppingList) list() {
	fmt.Println("\n=========================")
	fmt.Println("Displaying Shopping List")
	fmt.Println("=========================")
	if len(s) > 0 {
		for n, i := range s {
			fmt.Printf("Category: %s - Item: %s Quantity: %d Unit Coast: %.2f\n", category[i.category], n, i.quantity, i.cost)
		}
	} else {
		fmt.Println(noData)
	}
}

func (i ShoppingList) listSorted() {
	type tempItem = struct {
		name     string
		category int
		quantity int
		cost     float64
	}

	sortSlice := make([]tempItem, 0, len(i))

	for i, k := range i {
		sortSlice = append(sortSlice, tempItem{name: i, category: k.category, quantity: k.quantity, cost: k.cost})
	}

	sort.Slice(sortSlice, func(i, j int) bool { return sortSlice[i].category < sortSlice[j].category })

	for _, item := range sortSlice {
		fmt.Printf("Category: %s - Item: %s Quantity: %d Unit Coast: %.2f\n", category[item.category], item.name, item.quantity, item.cost)
	}
}

func (s ShoppingList) print() {
	fmt.Println("\n==================")
	fmt.Println("Print Current Data")
	fmt.Println("==================")
	if len(s) > 0 {
		for i, k := range s {
			fmt.Printf("%s - {%d %d %.2f}\n", i, k.category, k.quantity, k.cost)
		}
	} else {
		fmt.Println(noData)
	}
	mainMenu()
}

func (s ShoppingList) totalCost() {
	fmt.Println("\nTotal cost by Category")
	fmt.Println("----------------------")
	for cIdx, cName := range category {
		var totalCost float64 = 0
		for _, item := range s {
			if item.category == cIdx {
				totalCost = totalCost + item.totalCost()
			}
		}
		fmt.Printf("%s cost: %.2f\n", cName, totalCost)
	}
	mainMenu()
}

func (s ShoppingList) contains(v string) bool {
	if _, ok := s[v]; ok {
		return true
	}
	return false
}

// Method to check if value exist in Map ignoreing case
// prevent item of same name but with different case to be inserted into Shopping List
func (s ShoppingList) containsIgnoreCase(v string) (string, bool) {
	var d string
	v = strings.ToUpper(v)
	for k, _ := range s {
		if r := strings.Compare(strings.ToUpper(k), strings.ToUpper(v)); r == 0 {
			d = k
			return d, true
		}
	}
	return d, false
}

func (s ShoppingList) deleteByCategoryIdx(v int) int {
	var count int
	for p, r := range s {
		if r.category == v {
			count++
			delete(s, p)
		}
	}
	return count
}

func (s ShoppingList) updateByCategoryIdx(v int) int {
	var count int
	for p, r := range s {
		if r.category > v {
			count++
			r.category = r.category - 1
			s[p] = r
		}
	}
	return count
}
