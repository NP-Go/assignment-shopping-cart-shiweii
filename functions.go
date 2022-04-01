package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func generateReport() {
	var reportSelection int

	for {
		fmt.Println("\n================")
		fmt.Println("Generate Report")
		fmt.Println("================")
		fmt.Println("1. Total Cost of each category.")
		fmt.Println("2. List of item by category.")
		fmt.Println("3. Main Menu.")
		fmt.Println("\nChoose your report:")
		v, b := readInputAsInt()
		if b {
			reportSelection = v
			break
		} else {
			fmt.Println("\nInvalid input, please select a valid option.")
		}
	}

	switch reportSelection {
	case 1:
		shoppingList.totalCost()
	case 2:
		shoppingList.listSorted()
	case 3:
		mainMenu()
	}
}

func addItem() {
	var detail itemDetail

	fmt.Println("\n============")
	fmt.Println("Add New Item")
	fmt.Println("============")

	for {
		fmt.Println("\nWhat is the name of your item?")
		name := readInput()
		if len(name) > 0 {
			if _, exist := shoppingList.containsIgnoreCase(name); exist {
				fmt.Println("\nItem already exist in the Shopping List.")
			} else {
				detail.name = name
				break
			}
		} else {
			fmt.Println(noInput)
		}
	}

	for {
		fmt.Println("\nWhat category does it belong to?")
		cat := readInput()
		if len(cat) > 0 {
			if i, exist := category.containsIgnoreCase(cat); exist {
				detail.item.category = i
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
			detail.item.quantity = r
			break
		} else {
			fmt.Println("\nUnit entered is not valid, please re-enter unit.")
		}
	}

	for {
		fmt.Println("\nHow much does it cost per units?")
		r, ok := readInputAsFloat()
		if ok {
			detail.item.cost = r
			break
		} else {
			fmt.Println("\nUnit entered is not valid, please re-enter cost.")
		}
	}

	shoppingList[detail.name] = detail.item

	fmt.Println("\n[New item", detail.name, "created]")
}

func modifyItem() {
	var nameOld, nameNew string
	var detail, detailNew itemDetail
	var msg []string

	fmt.Println("\n===========")
	fmt.Println("Modify Item")
	fmt.Println("===========")

	if len(shoppingList) > 0 {
		for {
			fmt.Println("\nWhat item would you wish to modify?")
			name := readInput()
			if (len(name)) > 0 {
				if k, exist := shoppingList.containsIgnoreCase(name); exist {
					detail = itemDetail{
						name: k,
						item: shoppingList[k],
					}
					detailNew = detail
					detail.print()
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
			detailNew.name = nameNew
		} else {
			msg = append(msg, "No changes to item name made.")
		}

		for {
			fmt.Printf("\nEnter new Category. Enter for no change. [Current value: %s]\n", category[detail.item.category])
			cat := readInput()
			if len(cat) > 0 {
				if ret, exist := category.containsIgnoreCase(cat); exist {
					detailNew.item.category = ret
					break
				} else {
					fmt.Println("\nCategory enter does not exist. Either enter a existing category or create a new Category")
				}
			} else {
				msg = append(msg, "No changes to category made.")
				break
			}
		}

		for {
			fmt.Printf("\nEnter new Quantity. Enter for no change. [Current value: %d]\n", detail.item.quantity)
			qty := readInput()
			if len(qty) > 0 {
				if ret, err := strconv.Atoi(qty); err == nil {
					if ret > 0 {
						detailNew.item.quantity = ret
						break
					} else {
						fmt.Println("Quantity cannot be negative")
					}
				} else {
					fmt.Println("Please enter a valid Quantity.")
				}
			} else {
				msg = append(msg, "No changes to quantity made.")
				break
			}
		}

		for {
			fmt.Printf("\nEnter new Cost. Enter for no change. [Current value: %.2f]\n", detail.item.cost)
			cost := readInput()
			if len(cost) > 0 {
				if ret, err := strconv.ParseFloat(cost, 64); err == nil {
					if ret > 0 {
						detailNew.item.cost = ret
						break
					} else {
						fmt.Println("Cost cannot be negative.")
					}
				} else {
					fmt.Println("Please enter a valid Cost.")
				}
			} else {
				msg = append(msg, "No changes to cost made.")
				break
			}
		}

		if len(msg) > 0 {
			fmt.Printf("")
			for _, r := range msg {
				fmt.Println(r)
			}
		}

		if !reflect.DeepEqual(detail, detailNew) {
			delete(shoppingList, detail.name)
			shoppingList[detailNew.name] = detailNew.item
			fmt.Printf("\n[Item %s modifed]\n", detailNew.name)
		} else {
			fmt.Printf("\n[Item %s not modifed]\n", detailNew.name)
		}
	} else {
		fmt.Println(shpListEmpty)
	}
}

func deleteItem() {
	var name string
	fmt.Println("\n===========")
	fmt.Println("Delete Item")
	fmt.Println("===========")
	if len(shoppingList) > 0 {
		for {
			fmt.Println("\nEnter item name to delete:")
			name = readInput()
			if len(name) > 0 {
				if k, exist := shoppingList.containsIgnoreCase(name); exist {
					delete(shoppingList, k)
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
		fmt.Println("No item in Shopping List to delete!")
	}
}

func addCategory() {
	var cat string

	fmt.Println("\n=====================")
	fmt.Println("Add New Category Name")
	fmt.Println("=====================")

	for {
		fmt.Println("\nWhat is the New Category Name to add?")
		cat = readInput()
		if (len(cat)) > 0 {
			if v, exist := category.containsIgnoreCase(cat); exist {
				fmt.Printf("Category [%s] exists!\n", category[v])
			} else {
				category = append(category, cat)
				fmt.Printf("New Category: %s added at index: %d!\n", cat, category.getIndexByName(cat))
				break
			}
		} else {
			fmt.Println(noInput)
		}
	}
	mainMenu()
}

func modifyCategory() {
	var cat, catNew string
	var catIdx int

	fmt.Println("\n====================")
	fmt.Println("Modify Category name")
	fmt.Println("====================")

	if (len(category)) > 0 {
		for {
			fmt.Println("\nWhich Category name to modify?")
			cat = readInput()
			if (len(cat)) > 0 {
				if v, exist := category.containsIgnoreCase(cat); exist {
					catIdx = v
					break
				} else {
					fmt.Printf(catNotFound, cat)
				}
			} else {
				fmt.Println(noInput)
			}
		}
		for {
			fmt.Println("\nPlease enter new Category name:")
			catNew = readInput()
			if (len(catNew)) > 0 {
				if exist := category.contains(catNew); exist {
					fmt.Printf("Same Category name is entered\n")
				} else {
					category[catIdx] = catNew
					fmt.Printf("\n[Category %s modified to %s]\n", cat, catNew)
					break
				}
			} else {
				fmt.Println(noInput)
			}
		}
	} else {
		fmt.Println(catEmpty)
	}
}

func deleteCategory() {
	var idx int
	var cat, catOld string

	fmt.Println("\n===============")
	fmt.Println("Delete Category")
	fmt.Println("===============")

	if (len(category)) > 0 {
		for {
			fmt.Println("\nWhich Category to delete?")
			cat = readInput()
			if (len(cat)) > 0 {
				if v, exist := category.containsIgnoreCase(cat); exist {
					idx = v
					catOld = category[idx]
					break
				} else {
					fmt.Printf(catNotFound, cat)
				}
			} else {
				fmt.Println(noInput)
			}
		}
		// Delete Category
		category = append(category[:idx], category[idx+1:]...)
		fmt.Printf("\n[Category %s deleted]\n", catOld)
		// Delete all items in this Category
		r := shoppingList.deleteByCategoryIdx(idx)
		fmt.Printf("[%d items belonging to %s deleted]\n", r, catOld)
		// Update Category of other items
		r = shoppingList.updateByCategoryIdx(idx)
		fmt.Printf("[Category of %d items updated]\n", r)
	} else {
		fmt.Println(catEmpty)
	}
}

func createShoppingList() {
	newShoppingList := make(map[string]Item)
	shoppingListSlice = append(shoppingListSlice, newShoppingList)
	shoppingList = newShoppingList
	selectedList = len(shoppingListSlice) - 1
	fmt.Println("\nNew Shopping List created at:", selectedList)
	fmt.Println("\nUsing new Shopping List", selectedList)
	shoppingList = shoppingListSlice[selectedList]
}

func setShoppingList() {
	var input int
	for {
		fmt.Println("\nSelect Shopping List by index:")
		v, b := readInputAsInt()
		if b {
			input = v
			break
		} else {
			fmt.Println("\nInvalid input, please select the shoping list by index.")
		}
	}
	if input > len(shoppingListSlice)-1 {
		fmt.Println("\nNo Shopping List at index", input, "please create a new shopping list.")
	} else {
		selectedList = input
		shoppingList = shoppingListSlice[selectedList]
	}
}
