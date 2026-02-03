package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main(){
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")


	psqlconn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
	dbHost, dbPort, dbUser, dbPassword, dbName)

	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}

	defer db.Close()
	err = db.Ping()
	CheckError(err)
	fmt.Println("Connected to the database successfully!")

	rows, err := db.Query(`SELECT "id", "name", "email" FROM "users"`)
	CheckError(err)
	for rows.Next(){
		var id int
		var name string
		var email string

		err = rows.Scan(&id, &name, &email)
		CheckError(err)

		fmt.Println(id, name, email)
	}

}	

func CheckError(err error)  {
		if err != nil {
			panic(err)
	}
}
