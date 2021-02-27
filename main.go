package main

import (
	"log"
	"net/http"
	"os"

	"github.com/sql-mvc/services"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	log.Println("Server started on port: " + port)
	http.HandleFunc("/", services.Index)
	http.HandleFunc("/show", services.Show)
	http.HandleFunc("/new", services.New)
	http.HandleFunc("/edit", services.Edit)
	http.HandleFunc("/insert", services.Insert)
	http.HandleFunc("/update", services.Update)
	http.HandleFunc("/delete", services.Delete)
	http.ListenAndServe(":"+port, nil)
}
