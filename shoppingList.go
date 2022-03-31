package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type ShoppingList map[string]Item

func (i ShoppingList) delete() {

	var input string

	fmt.Println("\n===========")
	fmt.Println("Delete Item")
	fmt.Println("===========")

	fmt.Println("\nEnter item name to delete:")
	fmt.Scanln(&input)

	_, ok := i[input]
	if ok {
		delete(i, input)
		fmt.Println("Deleted", input)
	} else {
		fmt.Println("Item not found. Noting to delete!")
	}

	mainMenu()
}

func (s ShoppingList) list() {
	fmt.Println("\n=========================")
	fmt.Println("Displaying Shopping List")
	fmt.Println("=========================")
	if len(s) > 0 {
		for n, i := range s {
			fmt.Printf("Category: %s - Item: %s Quantity: %d Unit Coast: %.2f\n", category[i.category], n, i.quantity, i.cost)
		}
	} else {
		fmt.Println("No data found!")
	}
	mainMenu()
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
	mainMenu()
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
		fmt.Println("No data found!")
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

func (s ShoppingList) add() {
	var name, cat string
	var unit, catIdx int
	var cost float64
	var item Item

	fmt.Println("\nWhat is the name of your item?")
	fmt.Scanln(&name)
	fmt.Println("What category does it belong to?")
	fmt.Scanln(&cat)
	fmt.Println("How many units are these?")
	fmt.Scanln(&unit)
	fmt.Println("How much does it cost per units?")
	fmt.Scanln(&cost)

	for i, name := range category {
		if strings.ToUpper(cat) == strings.ToUpper(name) {
			catIdx = i
		}
	}

	item.category = catIdx
	item.quantity = unit
	item.cost = cost

	shoppingList[name] = item

	mainMenu()
}

func (s ShoppingList) modify() {
	var nameInput, catInput, qtyInput, costInput string
	var nameCur, nameNew, catNew string
	var qtyNew int
	var costNew float64
	var itemNew Item
	//var item Item

	fmt.Println("\nWhat item would you wish to modify?")
	fmt.Scanln(&nameCur)

	item, found := shoppingList[nameCur]

	if found {
		item.print(nameCur)
	}

	fmt.Println("\nEnter new name. Enter for no change.")
	fmt.Scanln(&nameInput)
	nameNew = strings.TrimSpace(nameInput)

	if (len(nameNew)) > 0 {
		for i, v := range shoppingList {
			if i == nameCur {
				delete(shoppingList, nameCur)
				shoppingList[nameNew] = v
			}
		}
	} else {
		nameNew = nameCur
	}

	fmt.Println("\nEnter new Category. Enter for no change.")
	fmt.Scanln(&catInput)
	catNew = strings.TrimSpace(catInput)
	if (len(catNew)) > 0 {
		if category.contains(catNew) {
			idx := category.getIndexByName(catNew)
			itemNew.category = idx
		} else {
			fmt.Println("Category enter does not exist. Either enter a existing category or create a new Category")
		}
	} else {
		itemNew.category = item.category
	}

	fmt.Println("\nEnter new Quantity. Enter for no change.")
	fmt.Scanln(&qtyInput)
	qtyInput = strings.TrimSpace(qtyInput)
	if (len(qtyInput)) > 0 {
		qtyNew, _ = strconv.Atoi(qtyInput)
		if qtyNew > 0 {
			itemNew.quantity = qtyNew
		} else {
			fmt.Println("Quantity cannot be negative")
		}
	} else {
		itemNew.quantity = item.quantity
	}

	fmt.Println("\nEnter new Cost. Enter for no change.")
	fmt.Scanln(&costInput)
	costInput = strings.TrimSpace(costInput)
	if (len(costInput)) > 0 {
		costNew, _ = strconv.ParseFloat(costInput, 64)
		if costNew > 0 {
			itemNew.cost = costNew
		} else {
			fmt.Println("Cost cannot be negative")
		}
	} else {
		itemNew.cost = item.cost
	}

	if nameCur == nameNew {
		fmt.Println("No changes to item name made")
	}

	res := item.isDifferent(itemNew)

	if len(res) > 0 {
		for _, r := range res {
			fmt.Println(r)
		}
		shoppingList[nameNew] = itemNew
	}

	mainMenu()
}

func (s ShoppingList) deleteByCategoryIdx(v int) {
	for p, r := range s {
		if r.category == v {
			delete(s, p)
		}
	}
}

func (s ShoppingList) updateByCategoryIdx(v int) {
	for p, r := range s {
		if r.category > v {
			r.category = r.category - 1
			s[p] = r
		}
	}
}
