package main

import (
	_ "database/sql"
	"fmt"
	"log"
	"prast13/database"
	"prast13/proses"
)

func main() {
	// Membuka koneksi ke database
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal("Error opening database: ", err)
	}
	defer db.Close() // Menutup koneksi setelah selesai digunakan

	// Menyisipkan data
	id, err := proses.InsertData(db)
	if err != nil {
		log.Fatal("Error inserting data:", err)
	}

	// Tampilkan ID record baru
	fmt.Printf("Inserted record ID: %d\n", id)
}
