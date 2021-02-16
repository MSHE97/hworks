package types

// Money - денежная суммы в минимальных еденицах (дирамы, копейки, центы и т.д.)
type Money int64

// Category представляет собой категорию, в которой был совершен платёж (авто, аптеки, рестораны и т.д.).
type Category string

// Payment представлет информацию о платеже.
type Payment struct {
	ID int
	Amount Money
	Category Category
}