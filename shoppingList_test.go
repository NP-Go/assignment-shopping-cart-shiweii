package main

import (
	"testing"
)

func containTest(t *testing.T) {
	got := shoppingList.contains("Coke")
	res := true
	if got != res {
		t.Errorf("shoppingList.contains(Coke) = %t; want %t got %t", got, res, got)
	}

	got = shoppingList.contains("Pepsi")
	res = false
	if got != res {
		t.Errorf("shoppingList.contains(Pepsi) = %t; want %t got %t", got, res, got)
	}

	_, got = shoppingList.containsIgnoreCase("coke")
	res = true
	if got != res {
		t.Errorf("shoppingList.contains(coke) = %t; want %t got %t", got, res, got)
	}

	_, got = shoppingList.containsIgnoreCase("COKE")
	res = true
	if got != res {
		t.Errorf("shoppingList.contains(COKE) = %t; want %t got %t", got, res, got)
	}

	_, got = shoppingList.containsIgnoreCase("COkE")
	res = true
	if got != res {
		t.Errorf("shoppingList.contains(COkE) = %t; want %t got %t", got, res, got)
	}

	_, got = shoppingList.containsIgnoreCase("pepsi")
	res = false
	if got != res {
		t.Errorf("shoppingList.contains(pepsi) = %t; want %t got %t", got, res, got)
	}
}
