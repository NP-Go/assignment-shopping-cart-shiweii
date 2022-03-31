package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var category = Category{
	"Household",
	"Food",
	"Drinks",
}

var inventory = Inventory{
	"Fork":   {category: 0, quantity: 4, cost: 3},
	"Plates": {category: 0, quantity: 4, cost: 3},
	"Cups":   {category: 0, quantity: 5, cost: 3},
	"Bread":  {category: 1, quantity: 2, cost: 2},
	"Cake":   {category: 1, quantity: 3, cost: 1},
	"Coke":   {category: 2, quantity: 5, cost: 2},
	"Sprite": {category: 2, quantity: 5, cost: 2},
}

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
		computeTotalCost()
	case 2:
		fmt.Println("\nList by Category")
		fmt.Println("----------------")
		inventory.listSorted()
	case 3:
		mainMenu()
	}
}

func computeTotalCost() {
	fmt.Println("\nTotal cost by Category")
	fmt.Println("----------------------")
	for cIdx, cName := range category {
		var totalCost float64 = 0
		for _, item := range inventory {
			if item.category == cIdx {
				totalCost = totalCost + item.totalCost()
			}
		}
		fmt.Printf("%s cost: %.2f\n", cName, totalCost)
	}
}

func addItem() {
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

	inventory[name] = item

	mainMenu()
}

func modifyItem() {
	var nameInput, catInput, qtyInput, costInput string
	var nameCur, nameNew, catNew string
	var qtyNew int
	var costNew float64
	var itemNew Item
	//var item Item

	fmt.Println("\nWhat item would you wish to modify?")
	fmt.Scanln(&nameCur)

	item, found := inventory[nameCur]

	if found {
		item.print(nameCur)
	}

	fmt.Println("\nEnter new name. Enter for no change.")
	fmt.Scanln(&nameInput)
	nameNew = strings.TrimSpace(nameInput)

	if (len(nameNew)) > 0 {
		for i, v := range inventory {
			if i == nameCur {
				delete(inventory, nameCur)
				inventory[nameNew] = v
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

	fmt.Println("nameCur", nameCur)
	fmt.Println("nameNew", nameNew)

	res := item.isDifferent(itemNew)

	if len(res) > 0 {
		for _, r := range res {
			fmt.Println(r)
		}
		inventory[nameNew] = itemNew
	}

	mainMenu()
}

func deleteItem() {
	var input string

	fmt.Println("\n===========")
	fmt.Println("Delete Item")
	fmt.Println("===========")

	fmt.Println("\nEnter item name to delete:")
	fmt.Scanln(&input)

	if inventory.delete(input) {
		fmt.Println("Deleted", input)
	} else {
		fmt.Println("Item not found. Noting to delete!")
	}
	mainMenu()
}

func printData() {
	fmt.Println("\n==================")
	fmt.Println("Print Current Data")
	fmt.Println("==================")
	inventory.print()
	mainMenu()
}

func mainMenu() {

	var optionSelected int

	var menuSelection = map[int]string{
		1: "View entire shopping list.",
		2: "Generate Shopping List Report.",
		3: "Add Items.",
		4: "Modify Items.",
		5: "Delete Items.",
		6: "Print Current Data.",
		7: "Add New Category Name",
	}

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
		fmt.Println("\n=========================")
		fmt.Println("Displaying Shopping List")
		fmt.Println("=========================")
		inventory.list()
		mainMenu()
	case 2:
		generateReport()
	case 3:
		addItem()
	case 4:
		modifyItem()
	case 5:
		deleteItem()
	case 6:
		printData()
	}
}

func main() {

	mainMenu()

}
