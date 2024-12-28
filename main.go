package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	chiMid "github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"github.com/kasyap1234/expense-tracker/config"
	appMid "github.com/kasyap1234/expense-tracker/middleware"
	"github.com/kasyap1234/expense-tracker/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	DB_URL := os.Getenv("DB_URL")
	if DB_URL == "" {
		log.Fatal("DB_URL is not set")
	}
	fmt.Println("DB_URL:", DB_URL)
	config.InitDB(DB_URL)

	r := chi.NewRouter()

	r.Use(chiMid.Logger)
	r.Use(chiMid.Recoverer)
	r.Get("/auth/google/login", routes.GoogleLogin)
	r.Get("/auth/google/callback", routes.OauthCallback)

	// Protected routes
	r.Route("/expenses", func(r chi.Router) {
		r.Use(appMid.JWTMiddleware)
		r.Post("/", routes.CreateExpense)
		r.Get("/", routes.GetExpenses)
		r.Get("/{expenseID}", routes.GetExpense)
		r.Put("/{expenseID}", routes.UpdateExpense)
		r.Delete("/{expenseID}", routes.DeleteExpense)
	})
	ssl := http.ListenAndServeTLS(":"+port, "cert.pem", "key.pem", r)
	if ssl != nil {
		log.Fatal(err)
	}

}
