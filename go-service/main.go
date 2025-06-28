package main

import (
	"log"
	"net/http"
	"os"

	"go-service/handler"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PDF_SERVICE_PORT")
	if port == "" {
		port = "8081"
	}

	r := mux.NewRouter()

	r.HandleFunc("/api/v1/students/{id}/report", handler.GenerateStudentsReportHandler).Methods("GET")

	log.Println("🚀 PDF service running on http://localhost:" + port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
