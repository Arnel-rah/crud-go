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
var port = 8080

func main() {
	cfg, err := LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	db, err := ConnectDB(cfg)
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	router := mux.NewRouter()

	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)


	router.HandleFunc("/users", userHandler.CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", userHandler.GetUser).Methods("GET")
	router.HandleFunc("/users/{id}", userHandler.GetUserById).Methods("GET")
	router.HandleFunc("/users/{id}", userHandler.UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", userHandler.DeleteUser).Methods("DELETE")

	fmt.Println(`server run at localhost:`, port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}
