package database

import (
	"database/sql"

	"br.com.simple_order_service/internal/entity"
)

type OrderDB struct {
	Db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderDB {
	return &OrderDB{
		Db: db,
	}
}

func (r *OrderDB) Save(order *entity.Order) error {
	//Execute this SQL statement
	_, err := r.Db.Exec("INSERT INTO orders (id, price, tax, final_price) VALUES (?,?,?,?)",
		order.ID, order.Price, order.Tax, order.FinalPrice)
	if err != nil {
		return err
	}
	return nil
}

func (r *OrderDB) GetTotalTransactions() (int, error) {
	var total int
	// Get the results and put the value on total
	err := r.Db.QueryRow("SELECT COUNT(*) FROM orders").Scan(&total)
	if err != nil {
		return 0, err
	}
	return total, nil
}
