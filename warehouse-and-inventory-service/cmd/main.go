package main

import (
	"fmt"
	"log"
	"net/http"
	
	"warehouse-inventory-service/internal/handler"
	"warehouse-inventory-service/internal/repository"
	"warehouse-inventory-service/internal/service"
)

func main() {
	fmt.Println("Memulai Warehouse & Inventory Service...")

	// 1. Setup Layer Repository (Gunakan Dummy sementara untuk Tahap 2)
	repo := repository.NewDummyInboundRepo()

	// 2. Setup Layer Service (Inject Repository ke dalam Service)
	svc := service.NewInboundService(repo)

	// 3. Setup Layer Handler (Inject Service ke dalam Handler)
	inboundHandler := handler.NewInboundHandler(svc)

	// 4. Mendaftarkan Rute HTTP (Routing/Endpoint)
	http.HandleFunc("/api/v1/inbound", inboundHandler.HandleProcessInbound)

	// 5. Menjalankan Server HTTP
	port := ":8080"
	fmt.Printf("Server berjalan di http://localhost%s\n", port)
	
	// Server akan standby dan mendengarkan request masuk
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalf("Gagal menjalankan server: %v", err)
	}
}
