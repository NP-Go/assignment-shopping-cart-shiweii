package main

import (
	"fmt"
	"os"
	"sort"
)

var category Category
var shoppingList ShoppingList
var shoppingList2 ShoppingList
var menuSelection map[int]string
var shoppingListSlice []map[string]Item

// Default to first shopping list
var selectedList int = 0

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
		fmt.Println("\nList by Category")
		fmt.Println("----------------")
		shoppingList.listSorted()
	case 3:
		mainMenu()
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
	mainMenu()
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
	mainMenu()
}

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
		shoppingList.add()
	case 4:
		shoppingList.modify()
	case 5:
		shoppingList.delete()
	case 6:
		shoppingList.print()
	case 7:
		category.add()
	case 8:
		category.modify()
	case 9:
		category.delete()
	case 10:
		createShoppingList()
	case 11:
		setShoppingList()
	case 12:
		os.Exit(0)
	}
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
	shoppingList2 = map[string]Item{
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
