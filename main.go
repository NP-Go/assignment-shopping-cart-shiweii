package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var category Category
var shoppingList ShoppingList
var menuSelection map[int]string
var shoppingSlice []map[string]Item

func generateReport() {
	var reportSelection int
	fmt.Println("\n================")
	fmt.Println("Generate Report")
	fmt.Println("================")
	fmt.Println("1. Total Cost of each category.")
	fmt.Println("2. List of item by category.")
	fmt.Println("3. Main Menu.")
	fmt.Println("\nChoose your report:")
	fmt.Scanln(&reportSelection)
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

func mainMenu() {

	var optionSelected int

	fmt.Println("\nShopping List Application")
	fmt.Println("=========================")

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
		inputReader := bufio.NewReader(os.Stdin)
		selection, _ := inputReader.ReadString('\n')
		selection = strings.TrimSpace(selection)
		// Check that user input valid selection
		value, _ := strconv.Atoi(selection)
		_, inMenuSelectionMap := menuSelection[value]
		if inMenuSelectionMap {
			fmt.Printf("Selected [%s]\n", menuSelection[value])
			optionSelected = value
			break
		} else {
			fmt.Println("==========================================")
			fmt.Println("Invalid choice, please select your choice.")
			fmt.Println("==========================================")
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

	// Init Shopping List Slice
	shoppingSlice = append(shoppingSlice, shoppingList)

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
		10: "Exit",
	}
}

func main() {
	mainMenu()
}
