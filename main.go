package main

import (
	"database/sql"
	"fmt"

	"br.com.simple_order_service/internal/infra/database"
	"br.com.simple_order_service/internal/usecase"

	// User underline to avoid errors on Go, and allow me to save it
	_ "github.com/mattn/go-sqlite3"
)

/* to create the database and the table
sqlite3 db.sqlite3
CREATE TABLE orders (id varchar(255) NOT NULL, price float NOT NULL, tax float NOT NULL, final_price float NOT NULL, PRIMARY KEY(id));
*/

func main() {
	db, err := sql.Open("sqlite3", "db.sqlite3")
	if err != nil {
		panic(err)
	}

	//* Defer wait everything works, to close at the end
	defer db.Close()

	orderDB := database.NewOrderRepository(db)
	uc := usecase.CalculateFinalPrice{
		OrderDBInterface: orderDB,
	}

	input := usecase.OrderInput{
		ID:    "123",
		Price: 12.0,
		Tax:   2.0,
	}

	output, err := uc.Execute(input)
	if err != nil {
		panic(err)
	}

	fmt.Println(output)
}
