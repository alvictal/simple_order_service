package entity

type OrderDBInterface interface {
	Save(order *Order) error
	GetTotalTransactions() (int, error)
}
