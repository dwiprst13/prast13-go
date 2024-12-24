package koneksi

import (
	"fmt"
	"net/http"
)

func Routes() {
	// Handler root
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<h1>Selamat Datang di Server</h1>")
	})

	// Handler untuk path /test
	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<p>Anaknya 12</p>")
	})

}
