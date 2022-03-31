package main

import "fmt"

type Category []string

func (c Category) contains(v string) bool {
	for _, r := range c {
		if v == r {
			return true
		}
	}
	return false
}

func (c Category) getIndexByName(v string) int {
	var idx int
	for i, r := range c {
		if v == r {
			idx = i
		}
	}
	return idx
}

func (c *Category) add() {
	var input string

	fmt.Println("\n=====================")
	fmt.Println("Add New Category Name")
	fmt.Println("=====================")

	fmt.Println("\nWhat is the New Category Name to add?")
	fmt.Scanln(&input)

	if (len(input)) > 0 {
		if c.contains(input) {
			fmt.Printf("Category: %s already exist at index %d!\n", input, c.getIndexByName(input))

		} else {
			*c = append(*c, input)
			fmt.Printf("New Category: %s added at %d!\n", input, c.getIndexByName(input))
		}
	} else {
		fmt.Println("No Input Found!")
	}
	mainMenu()
}

func (c *Category) modify() {
	var input, inputNew string

	fmt.Println("\n====================")
	fmt.Println("Modify Category name")
	fmt.Println("====================")

	fmt.Println("\nWhich Category name to modify?")
	fmt.Scanln(&input)

	if (len(input)) > 0 {
		if c.contains(input) {
			fmt.Println("\nPlease enter new Category name:")
			fmt.Scanln(&inputNew)
			if c.contains(inputNew) {
				fmt.Printf("Category: %s already exist, please enter a new name!\n", inputNew)
			} else if input == inputNew {
				fmt.Printf("Same Category name is entered")
			} else {
				(*c)[c.getIndexByName(input)] = inputNew
			}
		} else {
			fmt.Printf("Category: %s does not exist!\n", input)
		}
	} else {
		fmt.Println("No Input Found!")
	}
	mainMenu()
}

func (c *Category) delete() {
	var input string

	fmt.Println("\n===============")
	fmt.Println("Delete Category")
	fmt.Println("===============")

	fmt.Println("\nWhich Category to delete?")
	fmt.Scanln(&input)

	if (len(input)) > 0 {
		if c.contains(input) {
			idx := c.getIndexByName(input)
			*c = append((*c)[:idx], (*c)[idx+1:]...)
			shoppingList.deleteByCategoryIdx(idx)
			shoppingList.updateByCategoryIdx(idx)
		} else {
			fmt.Printf("Category: %s does not exist!\n", input)
		}
	} else {
		fmt.Println("No Input Found!")
	}
	mainMenu()
}
