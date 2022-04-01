package main

type Item struct {
	category int
	quantity int
	cost     float64
	_        struct{}
}

func (i Item) totalCost() float64 {
	var tc float64
	tc = float64(i.quantity) * i.cost
	return float64(tc)
}
