package main

import (
	"crud-app/internal/repository/psql"
	"crud-app/internal/service"
	"crud-app/internal/transport/rest"
	"crud-app/pkg/database"
	"log"
	"net/http"
	"time"

	_ "github.com/lib/pq"
)

func main() {
	// Сделать .env
	db, err := database.NewPostgresConnection(database.ConnectionInfo{
		Host:     "localhost",
		Port:     5432,
		Username: "postgres",
		DBName:   "postgres",
		SSLMode:  "disable",
		Password: "qwerty123",
	})
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// init deps
	booksRepo := psql.NewBooks(db)
	booksService := service.NewBooks(booksRepo)
	handler := rest.NewHandler(booksService)

	// init & run server
	srv := &http.Server{
		Addr:    ":8080",
		Handler: handler.InitRouter(),
	}

	log.Println("SERVER STARTED AT", time.Now().Format(time.RFC3339))

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
