package koneksi

import (
	"fmt"
	"net/http"
)

// KoneksiServer menjalankan server HTTP
func KoneksiServer() {

    Routes()
	// Menjalankan server HTTP
	fmt.Println("Server berjalan di http://localhost:8081")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error menjalankan server:", err)
	}
}
