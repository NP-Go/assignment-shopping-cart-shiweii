package main

import (
	"fmt"
	"sort"
	"strconv"
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

func (s ShoppingList) add() {
	var name, cat string
	var unit, catIdx int
	var cost float64
	var item Item

	fmt.Println("\n============")
	fmt.Println("Add New Item")
	fmt.Println("============")

	for {
		fmt.Println("\nWhat is the name of your item?")
		name = readInput()
		if len(name) > 0 {
			if _, exist := s.containsIgnoreCase(name); exist {
				fmt.Println("\nItem already exist in the Shopping List.")
			} else {
				break
			}
		} else {
			fmt.Println(noInput)
		}
	}

	for {
		fmt.Println("\nWhat category does it belong to?")
		cat = readInput()
		if len(name) > 0 {
			if i, exist := category.containsIgnoreCase(cat); exist {
				catIdx = i
				break
			} else {
				fmt.Println("\nCategory entered is not valid, please enter another category.")
			}
		} else {
			fmt.Println(noInput)
		}
	}

	for {
		fmt.Println("\nHow many units are these?")
		r, ok := readInputAsInt()
		if ok {
			unit = r
			break
		} else {
			fmt.Println("\nUnit entered is not valid, please re-enter unit.")
		}
	}

	for {
		fmt.Println("\nHow much does it cost per units?")
		r, ok := readInputAsFloat()
		if ok {
			cost = r
			break
		} else {
			fmt.Println("\nUnit entered is not valid, please re-enter cost.")
		}
	}

	item.category = catIdx
	item.quantity = unit
	item.cost = cost

	shoppingList[name] = item

	fmt.Println("\n[New item", name, "created]")

	mainMenu()
}

func (s ShoppingList) modify() {
	var item, itemNew Item
	var nameOld, nameNew string

	fmt.Println("\n===========")
	fmt.Println("Modify Item")
	fmt.Println("===========")

	if len(s) > 0 {
		for {
			fmt.Println("\nWhat item would you wish to modify?")
			ret := readInput()
			if (len(ret)) > 0 {
				if k, exist := s.containsIgnoreCase(ret); exist {
					nameOld = k
					item = shoppingList[nameOld]
					item.print(nameOld)
					break
				} else {
					fmt.Printf("\nItem [%s] does not exist in the Shopping List.", nameOld)
				}
			} else {
				fmt.Println(noInput)
			}
		}

		fmt.Println("\nEnter new name. Enter for no change.")
		nameNew = readInput()
		if (len(nameNew)) > 0 {
			delete(shoppingList, nameOld)
			shoppingList[nameNew] = item
		} else {
			nameNew = nameOld
		}

		for {
			fmt.Printf("\nEnter new Category. Enter for no change. [Current value: %s]\n", category[item.category])
			cat := readInput()
			if len(cat) > 0 {
				if v, exist := category.containsIgnoreCase(cat); exist {
					itemNew.category = v
					break
				} else {
					fmt.Println("\nCategory enter does not exist. Either enter a existing category or create a new Category")
				}
			} else {
				itemNew.category = item.category
				break
			}
		}

		for {
			fmt.Printf("\nEnter new Quantity. Enter for no change. [Current value: %d]\n", item.quantity)
			ret := readInput()
			if len(ret) > 0 {
				if v, err := strconv.Atoi(ret); err == nil {
					if v > 0 {
						itemNew.quantity = v
						break
					} else {
						fmt.Println("Quantity cannot be negative")
					}
				} else {
					fmt.Println("Please enter a valid Quantity.")
				}
			} else {
				itemNew.quantity = item.quantity
				break
			}
		}

		for {
			fmt.Printf("\nEnter new Cost. Enter for no change. [Current value: %.2f]\n", item.cost)
			ret := readInput()
			if len(ret) > 0 {
				if v, err := strconv.ParseFloat(ret, 64); err == nil {
					if v > 0 {
						itemNew.cost = v
						break
					} else {
						fmt.Println("Cost cannot be negative.")
					}
				} else {
					fmt.Println("Please enter a valid Cost.")
				}
			} else {
				itemNew.cost = item.cost
				break
			}
		}

		fmt.Println("")

		if i := strings.Compare(nameOld, nameNew); i == 0 {
			fmt.Println("No changes to item name made")
		}

		msg, diff := item.compare(itemNew)

		if diff {
			shoppingList[nameNew] = itemNew
		}

		if len(msg) > 0 {
			for _, r := range msg {
				fmt.Println(r)
			}
		}

		if i := strings.Compare(nameOld, nameNew); i == 0 && diff {
			fmt.Printf("\n[Item %s modifed]\n", nameNew)
		} else {
			fmt.Printf("\n[Item %s not modifed]\n", nameNew)
		}
	} else {
		fmt.Println(shpListEmpty)
	}

	mainMenu()
}

func (i ShoppingList) delete() {
	var name string
	fmt.Println("\n===========")
	fmt.Println("Delete Item")
	fmt.Println("===========")
	if len(i) > 0 {
		for {
			fmt.Println("\nEnter item name to delete:")
			name = readInput()
			if len(name) > 0 {
				if k, exist := i.containsIgnoreCase(name); exist {
					delete(i, k)
					fmt.Printf("\n[Item %s deleted from Shopping List]\n", k)
					break
				} else {
					fmt.Println("\nItem does not exist in the Shopping List.")
				}
			} else {
				fmt.Println(noInput)
			}
		}
	} else {
		fmt.Println(shpListEmpty)
	}
	mainMenu()
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
