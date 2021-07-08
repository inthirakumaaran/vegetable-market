package common

type Vegetable struct {
	Name string
	Price int
	Quantity int
}

type UpdateVegetable struct {
	Name string
	Value int
}

const MARKET_DB_PATH = "../market.tmp"