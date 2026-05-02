package main

import (
	"fmt"
	"log"

	// Ganti "service_CICD" dengan nama module yang ada di file go.mod kamu
	"order-tariff-service/internal/repository/postgres"
)

func main() {
	fmt.Println("Order & Tariff Service sedang berjalan....")

	// 1. Coba hubungkan ke PostgreSQL
	db, err := postgres.InitDB()
	if err != nil {
		// Jika gagal konek, aplikasi akan berhenti dan memberi tahu error-nya
		log.Fatalf("Gagal inisialisasi database: %v", err)
	}
	// Pastikan koneksi ditutup saat aplikasi selesai berjalan
	defer db.Close()

	log.Println("Database terkoneksi dengan sukses! Siap memproses data.")
}
