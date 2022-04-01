package main

import (
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
