package proses

import (
	"database/sql"
	"fmt"
)

// InsertData memasukkan data ke dalam tabel users
func InsertData(db *sql.DB) (int, error) {
	sqlStatement := `
	INSERT INTO users (age, email, first_name, last_name)
	VALUES ($1, $2, $3, $4)
	RETURNING id`

	var id int
	err := db.QueryRow(sqlStatement, 24, "yunianisa@gmail.com", "Yuni", "Anisa").Scan(&id)
	if err != nil {
		return 0, err
	}

	fmt.Println("New record ID is:", id)
	return id, nil
}
