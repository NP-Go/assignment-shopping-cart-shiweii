package main

import (
	"fmt"
	"strings"
)

type Category []string

func (c Category) contains(v string) bool {
	for _, r := range c {
		if v == r {
			return true
		}
	}
	return false
}

// Method to check if value exist in Slice ignoreing case
// prevent item of same name but with different case to be inserted into Shopping List
func (c Category) containsIgnoreCase(v string) (int, bool) {
	v = strings.ToUpper(v)
	for i, k := range c {
		if r := strings.Compare(strings.ToUpper(k), strings.ToUpper(v)); r == 0 {
			return i, true
		}
	}
	return 0, false
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
	var cat string

	fmt.Println("\n=====================")
	fmt.Println("Add New Category Name")
	fmt.Println("=====================")

	for {
		fmt.Println("\nWhat is the New Category Name to add?")
		cat = readInput()
		if (len(cat)) > 0 {
			if v, exist := c.containsIgnoreCase(cat); exist {
				fmt.Printf("Category [%s] exists!\n", category[v])
			} else {
				*c = append(*c, cat)
				fmt.Printf("New Category: %s added at index: %d!\n", cat, c.getIndexByName(cat))
				break
			}
		} else {
			fmt.Println(noInput)
		}
	}
	mainMenu()
}

func (c *Category) modify() {
	//var input, inputNew string
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
				if v, exist := c.containsIgnoreCase(cat); exist {
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
				if exist := c.contains(catNew); exist {
					fmt.Printf("Same Category name is entered\n")
				} else {
					(*c)[catIdx] = catNew
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
	mainMenu()
}

func (c *Category) delete() {
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
				if v, exist := c.containsIgnoreCase(cat); exist {
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
		*c = append((*c)[:idx], (*c)[idx+1:]...)
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
	mainMenu()
}
