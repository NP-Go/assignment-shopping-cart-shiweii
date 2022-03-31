package main

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
