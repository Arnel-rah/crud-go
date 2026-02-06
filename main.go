package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"go-crud-psql/internal/config" 
	"go-crud-psql/internal/handlers"
	"go-crud-psql/internal/repositories"
	"go-crud-psql/internal/services"
)

const port = 8080

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Échec du chargement de la configuration : %v", err)
	}

	db, err := config.ConnectDB(cfg) 
	if err != nil {
		log.Fatalf("Échec de la connexion à la base de données : %v", err)
	}

	router := mux.NewRouter()

	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	router.HandleFunc("/users", userHandler.CreateUser).Methods(http.MethodPost)
	router.HandleFunc("/users", userHandler.GetAllUsers).Methods(http.MethodGet)
	router.HandleFunc("/users/{id}", userHandler.GetUser).Methods(http.MethodGet) 
	router.HandleFunc("/users/{id}", userHandler.UpdateUser).Methods(http.MethodPut)
	router.HandleFunc("/users/{id}", userHandler.DeleteUser).Methods(http.MethodDelete)

	fmt.Printf("Serveur démarré sur http://localhost:%d\n", port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}