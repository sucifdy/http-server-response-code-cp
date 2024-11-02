package main

import (
	"net/http"
)

// Data students
var students = []string{
	"Aditira",
	"Dito",
	"Afis",
	"Eddy",
}

// Fungsi untuk mengecek apakah nama ada di dalam slice students
func IsNameExists(name string) bool {
	for _, n := range students {
		if n == name {
			return true
		}
	}
	return false
}

// Handler untuk mengecek nama siswa berdasarkan kondisi yang diberikan
func CheckStudentName() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Cek apakah metode yang digunakan adalah GET
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Header().Set("Content-Type", "text/plain; charset=utf-8")
			w.Write([]byte("Method is not allowed"))
			return
		}

		// Ambil nilai parameter `name` dari URL
		name := r.URL.Query().Get("name")

		// Jika parameter `name` tidak ada atau kosong
		if name == "" {
			w.WriteHeader(http.StatusNotFound)
			w.Header().Set("Content-Type", "text/plain; charset=utf-8")
			w.Write([]byte("Data not found"))
			return
		}

		// Cek apakah nama ada di dalam slice students
		if IsNameExists(name) {
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "text/plain; charset=utf-8")
			w.Write([]byte("Name is exists"))
		} else {
			w.WriteHeader(http.StatusNotFound)
			w.Header().Set("Content-Type", "text/plain; charset=utf-8")
			w.Write([]byte("Data not found"))
		}
	}
}

// Fungsi untuk membuat mux dan menetapkan routing
func GetMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/students", CheckStudentName())
	return mux
}

func main() {
	http.ListenAndServe("localhost:8080", GetMux())
}
