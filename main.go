package main

import (
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
	godotenv.Load()

	config.InitDB()

	r := chi.NewRouter()
	r.Use(chiMid.Logger)
	r.Use(chiMid.Recoverer)

	r.Get("/auth/google/login", routes.GoogleLogin)
	r.Get("/auth/google/callback", routes.OauthCallback)

	r.Route("/expenses", func(r chi.Router) {
		r.Use(appMid.JWTMiddleware)
		r.Post("/", routes.CreateExpense)
		r.Get("/", routes.GetExpenses)
		r.Get("/{expenseID}", routes.GetExpense)
		r.Put("/{expenseID}", routes.UpdateExpense)
		r.Delete("/{expenseID}", routes.DeleteExpense)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	log.Printf("Server starting on port %s with HTTPS", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal(err)
	}
	log.Default().Print("Server started on port ",port)
	
}
