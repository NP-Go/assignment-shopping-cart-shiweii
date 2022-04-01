package main

import (
	"fmt"
	"os"
	"sort"
)

var category Category
var shoppingList ShoppingList
var menuSelection map[int]string
var shoppingListSlice []map[string]Item

// Default to first shopping list
var selectedList int = 0

func mainMenu() {

	var optionSelected int

	fmt.Println("\nShopping List Application")
	fmt.Println("=========================")
	fmt.Println("Currently in Shopping list:", selectedList, "out of", len(shoppingListSlice)-1)
	fmt.Println("-------------------------------")
	// To store the keys in slice in sorted order
	menu := make([]int, len(menuSelection))
	i := 0
	for k := range menuSelection {
		menu[i] = k
		i++
	}
	sort.Ints(menu)

	for {
		for _, k := range menu {
			fmt.Printf("%d. %s\n", k, menuSelection[k])
		}
		// Display main menu
		fmt.Println("\nSelect your choice:")
		// Gather user input.
		r, ok := readInputAsInt()

		if ok {
			_, inMenuSelectionMap := menuSelection[r]
			if inMenuSelectionMap {
				fmt.Printf("\nSelected [%s]\n", menuSelection[r])
				optionSelected = r
				break
			} else {
				fmt.Println("\nInvalid choice, please select your choice.")
				fmt.Println("==========================================")
				fmt.Println("Currently in Shopping list:", selectedList, "out of", len(shoppingListSlice)-1)
			}
		} else {
			fmt.Println("\nInvalid choice, please select your choice.")
			fmt.Println("==========================================")
			fmt.Println("Currently in Shopping list:", selectedList, "out of", len(shoppingListSlice)-1)
		}
	}

	switch optionSelected {
	case 1:
		shoppingList.list()
	case 2:
		generateReport()
	case 3:
		addItem()
	case 4:
		modifyItem()
	case 5:
		deleteItem()
	case 6:
		shoppingList.print()
	case 7:
		addCategory()
	case 8:
		modifyCategory()
	case 9:
		deleteCategory()
	case 10:
		createShoppingList()
	case 11:
		setShoppingList()
	case 12:
		os.Exit(0)
	}
	mainMenu()
}

func init() {
	// Init Category Data
	category = append(category, "Household")
	category = append(category, "Food")
	category = append(category, "Drinks")

	// Init Shopping List Data
	shoppingList = map[string]Item{
		"Fork":   {category: 0, quantity: 4, cost: 3},
		"Plates": {category: 0, quantity: 4, cost: 3},
		"Cups":   {category: 0, quantity: 5, cost: 3},
		"Bread":  {category: 1, quantity: 2, cost: 2},
		"Cake":   {category: 1, quantity: 3, cost: 1},
		"Coke":   {category: 2, quantity: 5, cost: 2},
		"Sprite": {category: 2, quantity: 5, cost: 2},
	}

	// Init Shopping List 2 Data
	shoppingList2 := map[string]Item{
		"Fork":   {category: 0, quantity: 4, cost: 3},
		"Plates": {category: 0, quantity: 4, cost: 3},
	}

	// Init Shopping List Slice
	shoppingListSlice = append(shoppingListSlice, shoppingList)
	shoppingListSlice = append(shoppingListSlice, shoppingList2)
	shoppingList = shoppingListSlice[selectedList]

	// Init Menu
	menuSelection = map[int]string{
		1:  "View entire shopping list.",
		2:  "Generate Shopping List Report.",
		3:  "Add Items.",
		4:  "Modify Items.",
		5:  "Delete Items.",
		6:  "Print Current Data.",
		7:  "Add New Category Name",
		8:  "Modify Category Name",
		9:  "Delete Category",
		10: "Create new Shopping List",
		11: "Retrieve Previous Shopping List",
		12: "Exit",
	}
}

func main() {

	mainMenu()
}
